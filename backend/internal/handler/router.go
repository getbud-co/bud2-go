package handler

import (
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type RouterConfig struct {
	Env            string
	AllowedOrigins []string
	OpenAPISpec    []byte
}

func NewRouter(orgHandler *OrganizationHandler, cfg RouterConfig) *chi.Mux {
	r := chi.NewRouter()

	allowedOrigins := cfg.AllowedOrigins
	if len(allowedOrigins) == 0 {
		allowedOrigins = []string{"http://localhost:3000"}
	}
	// trim whitespace from each origin
	for i, o := range allowedOrigins {
		allowedOrigins[i] = strings.TrimSpace(o)
	}

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
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

	return r
}
