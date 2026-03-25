package shared

import (
	"net/http"

	"github.com/dsbraz/bud2/backend/internal/domain"
	"github.com/google/uuid"
)

// TenantMiddleware extracts tenant_id from X-Tenant-ID header.
// Placeholder until JWT authentication is implemented (ADR-012).
func TenantMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		raw := r.Header.Get("X-Tenant-ID")
		if raw == "" {
			WriteProblem(w, http.StatusUnauthorized, "Unauthorized", "X-Tenant-ID header is required")
			return
		}
		id, err := uuid.Parse(raw)
		if err != nil {
			WriteProblem(w, http.StatusUnauthorized, "Unauthorized", "invalid X-Tenant-ID format")
			return
		}
		ctx := domain.TenantIDToContext(r.Context(), domain.TenantID(id))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
