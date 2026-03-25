package shared

import (
	"strings"

	"github.com/dsbraz/bud2/backend/internal/organization"
	"github.com/dsbraz/bud2/backend/internal/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type RouterConfig struct {
	Env            string
	AllowedOrigins []string
	OpenAPISpec    []byte
}

func NewRouter(orgHandler *organization.Handler, userHandler *user.Handler, cfg RouterConfig) *chi.Mux {
	r := chi.NewRouter()

	allowedOrigins := cfg.AllowedOrigins
	if len(allowedOrigins) == 0 {
		allowedOrigins = []string{"http://localhost:3000"}
	}
	for i, o := range allowedOrigins {
		allowedOrigins[i] = strings.TrimSpace(o)
	}

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-Tenant-ID"},
		AllowCredentials: true,
	}))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	r.Get("/health", Health)

	if cfg.Env != "production" {
		r.Get("/swagger/", swaggerUIHandler)
		r.Get("/swagger/openapi.yml", openapiSpecHandler(cfg.OpenAPISpec))
	}

	r.Route("/organizations", func(r chi.Router) {
		r.Post("/", orgHandler.Create)
		r.Get("/", orgHandler.List)
		r.Get("/{id}", orgHandler.Get)
		r.Put("/{id}", orgHandler.Update)
	})

	r.Route("/users", func(r chi.Router) {
		r.Use(TenantMiddleware)
		r.Post("/", userHandler.Create)
		r.Get("/", userHandler.List)
		r.Get("/{id}", userHandler.Get)
		r.Put("/{id}", userHandler.Update)
	})

	return r
}
