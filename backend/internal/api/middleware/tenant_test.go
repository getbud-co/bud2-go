package middleware

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dsbraz/bud2/backend/internal/test/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestTenantMiddleware_Success(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	ctx := fixtures.NewContextWithTenant(fixtures.NewTestTenantID())
	req = req.WithContext(ctx)
	rr := httptest.NewRecorder()

	TenantMiddleware(handler).ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestTenantMiddleware_MissingTenant(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	// No tenant in context
	rr := httptest.NewRecorder()

	TenantMiddleware(handler).ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	assert.Contains(t, rr.Body.String(), "tenant_id is required")
}

func TestTenantMiddleware_InvalidTenantType(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	// Add wrong type to context
	ctx := req.Context()
	ctx = context.WithValue(ctx, "tenant_id", "not-a-tenant-id")
	req = req.WithContext(ctx)
	rr := httptest.NewRecorder()

	TenantMiddleware(handler).ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	assert.Contains(t, rr.Body.String(), "tenant_id is required")
}
