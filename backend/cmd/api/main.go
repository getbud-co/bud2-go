package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/exaring/otelpgx"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"

	apispec "github.com/dsbraz/bud2/backend/api"
	"github.com/dsbraz/bud2/backend/internal/api"
	apiauth "github.com/dsbraz/bud2/backend/internal/api/auth"
	apibootstrap "github.com/dsbraz/bud2/backend/internal/api/bootstrap"
	apiorg "github.com/dsbraz/bud2/backend/internal/api/organization"
	apiuser "github.com/dsbraz/bud2/backend/internal/api/user"
	appauth "github.com/dsbraz/bud2/backend/internal/app/auth"
	appbootstrap "github.com/dsbraz/bud2/backend/internal/app/bootstrap"
	apporg "github.com/dsbraz/bud2/backend/internal/app/organization"
	appuser "github.com/dsbraz/bud2/backend/internal/app/user"
	"github.com/dsbraz/bud2/backend/internal/config"
	infraauth "github.com/dsbraz/bud2/backend/internal/infra/auth"
	"github.com/dsbraz/bud2/backend/internal/infra/otel"
	"github.com/dsbraz/bud2/backend/internal/infra/postgres"
	"github.com/dsbraz/bud2/backend/internal/infra/postgres/sqlc"
	"github.com/dsbraz/bud2/backend/internal/infra/rbac"
)

func main() {
	cfg := config.Load()

	logger := initLogger(cfg.Env, cfg.LogLevel)
	slog.SetDefault(logger)

	if cfg.DatabaseURL == "" {
		slog.Error("DATABASE_URL is required")
		os.Exit(1)
	}

	if cfg.JWTSecret == "" {
		slog.Error("JWT_SECRET is required")
		os.Exit(1)
	}

	// Initialize Casbin enforcer
	if err := rbac.InitEnforcer(cfg.PolicyModel, cfg.PolicyFile); err != nil {
		slog.Error("failed to initialize authorization", "error", err)
		os.Exit(1)
	}

	// Initialize OpenTelemetry
	otelProvider, err := otel.NewProvider(otel.Config{
		Endpoint:    cfg.OTelEndpoint,
		ServiceName: cfg.OTelServiceName,
		Environment: cfg.OTelEnvironment,
	})
	if err != nil {
		slog.Error("failed to initialize OpenTelemetry", "error", err)
		os.Exit(1)
	}
	defer func() {
		if err := otelProvider.Shutdown(context.Background()); err != nil {
			slog.Error("failed to shutdown OpenTelemetry", "error", err)
		}
	}()

	// Run migrations
	runMigrations(cfg.DatabaseURL)

	// Database connection pool with OpenTelemetry instrumentation
	ctx := context.Background()
	pool, err := initDBPool(ctx, cfg.DatabaseURL)
	if err != nil {
		slog.Error("failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		slog.Error("database ping failed", "error", err)
		os.Exit(1)
	}

	// Infra
	queries := sqlc.New(pool)
	orgRepo := postgres.NewOrgRepository(queries)
	userRepo := postgres.NewUserRepository(queries)
	membershipRepo := postgres.NewMembershipRepository(queries)
	txManager := postgres.NewTxManager(pool)
	tokenIssuer := infraauth.NewTokenIssuer(cfg.JWTSecret)
	passwordHasher := infraauth.NewDefaultBcryptPasswordHasher()

	// Use cases
	createOrg := apporg.NewCreateUseCase(orgRepo, logger)
	getOrg := apporg.NewGetUseCase(orgRepo, logger)
	listOrg := apporg.NewListUseCase(orgRepo, logger)
	updateOrg := apporg.NewUpdateUseCase(orgRepo, logger)

	createUser := appuser.NewCreateUseCase(userRepo, membershipRepo, orgRepo, txManager, passwordHasher, logger)
	getUser := appuser.NewGetUseCase(userRepo, membershipRepo, logger)
	listUser := appuser.NewListUseCase(userRepo, membershipRepo, logger)
	updateUser := appuser.NewUpdateUseCase(userRepo, membershipRepo, txManager, logger)

	bootstrapUC := appbootstrap.NewUseCase(orgRepo, txManager, tokenIssuer, passwordHasher, logger)
	loginUC := appauth.NewLoginUseCase(userRepo, membershipRepo, orgRepo, tokenIssuer, passwordHasher, logger)
	getSessionUC := appauth.NewGetSessionUseCase(userRepo, membershipRepo, orgRepo, tokenIssuer, passwordHasher, logger)
	switchOrganizationUC := appauth.NewSwitchOrganizationUseCase(userRepo, membershipRepo, orgRepo, tokenIssuer, passwordHasher, logger)

	// Handlers + Router
	bootstrapHandler := apibootstrap.NewHandler(bootstrapUC)
	authHandler := apiauth.NewHandler(loginUC, getSessionUC, switchOrganizationUC)
	orgHandler := apiorg.NewHandler(createOrg, getOrg, listOrg, updateOrg)
	userHandler := apiuser.NewHandler(createUser, getUser, listUser, updateUser)
	router := api.NewRouter(bootstrapHandler, authHandler, orgHandler, userHandler, api.RouterConfig{
		Env:            cfg.Env,
		AllowedOrigins: strings.Split(cfg.AllowedOrigins, ","),
		OpenAPISpec:    apispec.Spec,
		JWTSecret:      cfg.JWTSecret,
		Enforcer:       rbac.Enforcer(),
		Pool:           pool,
	})

	slog.Info("starting server", "port", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		slog.Error("server error", "error", err)
		os.Exit(1)
	}
}

func initLogger(env, levelStr string) *slog.Logger {
	var level slog.Level
	switch strings.ToLower(levelStr) {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{Level: level}

	var handler slog.Handler
	if env == "production" {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	return slog.New(handler)
}

func initDBPool(ctx context.Context, databaseURL string) (*pgxpool.Pool, error) {
	pgxCfg, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return nil, err
	}

	pgxCfg.ConnConfig.Tracer = otelpgx.NewTracer()

	return pgxpool.NewWithConfig(ctx, pgxCfg)
}

func runMigrations(databaseURL string) {
	// golang-migrate pgx/v5 driver registers as "pgx5://" scheme
	migrationURL := strings.Replace(databaseURL, "postgres://", "pgx5://", 1)
	m, err := migrate.New("file://migrations", migrationURL)
	if err != nil {
		slog.Error("failed to initialize migrations", "error", err)
		os.Exit(1)
	}
	defer m.Close()

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		slog.Error("failed to run migrations", "error", err)
		os.Exit(1)
	}

	slog.Info("migrations applied")
}
