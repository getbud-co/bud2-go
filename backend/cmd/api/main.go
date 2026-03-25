package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"strings"

	apispec "github.com/dsbraz/bud2/backend/api"
	"github.com/dsbraz/bud2/backend/internal/config"
	"github.com/dsbraz/bud2/backend/internal/handler"
	"github.com/dsbraz/bud2/backend/internal/infra/postgres"
	orguc "github.com/dsbraz/bud2/backend/internal/usecase/organization"
	useruc "github.com/dsbraz/bud2/backend/internal/usecase/user"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg := config.Load()

	if cfg.DatabaseURL == "" {
		slog.Error("DATABASE_URL is required")
		os.Exit(1)
	}

	// Structured logging
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	// Run migrations
	runMigrations(cfg.DatabaseURL)

	// Database connection pool
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, cfg.DatabaseURL)
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
	queries := postgres.New(pool)
	orgRepo := postgres.NewOrganizationRepository(queries)
	userRepo := postgres.NewUserRepository(queries)

	// Use cases
	createOrg := orguc.NewCreateUseCase(orgRepo)
	getOrg := orguc.NewGetUseCase(orgRepo)
	listOrg := orguc.NewListUseCase(orgRepo)
	updateOrg := orguc.NewUpdateUseCase(orgRepo)

	createUser := useruc.NewCreateUseCase(userRepo)
	getUser := useruc.NewGetUseCase(userRepo)
	listUser := useruc.NewListUseCase(userRepo)
	updateUser := useruc.NewUpdateUseCase(userRepo)

	// Handler + Router
	orgHandler := handler.NewOrganizationHandler(createOrg, getOrg, listOrg, updateOrg)
	userHandler := handler.NewUserHandler(createUser, getUser, listUser, updateUser)
	router := handler.NewRouter(orgHandler, userHandler, handler.RouterConfig{
		Env:            cfg.Env,
		AllowedOrigins: strings.Split(cfg.AllowedOrigins, ","),
		OpenAPISpec:    apispec.Spec,
	})

	slog.Info("starting server", "port", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		slog.Error("server error", "error", err)
		os.Exit(1)
	}
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
