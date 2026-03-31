package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dsbraz/bud2/backend/internal/domain"
	"github.com/dsbraz/bud2/backend/internal/test/fixtures"
	"github.com/dsbraz/bud2/backend/internal/test/mocks"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware_ValidToken(t *testing.T) {
	secret := "test-secret"
	middleware := AuthMiddleware(AuthMiddlewareConfig{JWTSecret: secret})

	// Create valid token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   uuid.New().String(),
		"tenant_id": uuid.New().String(),
		"role":      "admin",
		"exp":       time.Now().Add(time.Hour).Unix(),
	})
	tokenString, _ := token.SignedString([]byte(secret))

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify context has claims
		claims, err := domain.ClaimsFromContext(r.Context())
		assert.NoError(t, err)
		assert.Equal(t, "admin", claims.Role)

		// Verify tenant ID is also in context
		tenantID, err := domain.TenantIDFromContext(r.Context())
		assert.NoError(t, err)
		assert.NotEqual(t, domain.TenantID{}, tenantID)

		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)
	rr := httptest.NewRecorder()

	middleware(handler).ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestAuthMiddleware_MissingHeader(t *testing.T) {
	middleware := AuthMiddleware(AuthMiddlewareConfig{JWTSecret: "secret"})

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	middleware(handler).ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	assert.Contains(t, rr.Body.String(), "Authorization header is required")
}

func TestAuthMiddleware_InvalidFormat(t *testing.T) {
	middleware := AuthMiddleware(AuthMiddlewareConfig{JWTSecret: "secret"})

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "InvalidToken")
	rr := httptest.NewRecorder()

	middleware(handler).ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	assert.Contains(t, rr.Body.String(), "Invalid authorization header format")
}

func TestAuthMiddleware_WrongScheme(t *testing.T) {
	middleware := AuthMiddleware(AuthMiddlewareConfig{JWTSecret: "secret"})

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Basic dXNlcjpwYXNz")
	rr := httptest.NewRecorder()

	middleware(handler).ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	assert.Contains(t, rr.Body.String(), "Invalid authorization header format")
}

func TestAuthMiddleware_InvalidToken(t *testing.T) {
	middleware := AuthMiddleware(AuthMiddlewareConfig{JWTSecret: "secret"})

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer invalid-token")
	rr := httptest.NewRecorder()

	middleware(handler).ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	assert.Contains(t, rr.Body.String(), "Invalid or expired token")
}

func TestAuthMiddleware_ExpiredToken(t *testing.T) {
	secret := "test-secret"
	middleware := AuthMiddleware(AuthMiddlewareConfig{JWTSecret: secret})

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   uuid.New().String(),
		"tenant_id": uuid.New().String(),
		"role":      "admin",
		"exp":       time.Now().Add(-time.Hour).Unix(), // Expired
	})
	tokenString, _ := token.SignedString([]byte(secret))

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)
	rr := httptest.NewRecorder()

	middleware(handler).ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	assert.Contains(t, rr.Body.String(), "Invalid or expired token")
}

func TestAuthMiddleware_WrongSignature(t *testing.T) {
	middleware := AuthMiddleware(AuthMiddlewareConfig{JWTSecret: "correct-secret"})

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   uuid.New().String(),
		"tenant_id": uuid.New().String(),
		"role":      "admin",
		"exp":       time.Now().Add(time.Hour).Unix(),
	})
	tokenString, _ := token.SignedString([]byte("wrong-secret"))

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)
	rr := httptest.NewRecorder()

	middleware(handler).ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
}

func TestAuthMiddleware_MissingUserID(t *testing.T) {
	secret := "test-secret"
	middleware := AuthMiddleware(AuthMiddlewareConfig{JWTSecret: secret})

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"tenant_id": uuid.New().String(),
		"role":      "admin",
		"exp":       time.Now().Add(time.Hour).Unix(),
	})
	tokenString, _ := token.SignedString([]byte(secret))

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)
	rr := httptest.NewRecorder()

	middleware(handler).ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	assert.Contains(t, rr.Body.String(), "Missing user_id in token")
}

func TestAuthMiddleware_InvalidUserIDFormat(t *testing.T) {
	secret := "test-secret"
	middleware := AuthMiddleware(AuthMiddlewareConfig{JWTSecret: secret})

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   "not-a-valid-uuid",
		"tenant_id": uuid.New().String(),
		"role":      "admin",
		"exp":       time.Now().Add(time.Hour).Unix(),
	})
	tokenString, _ := token.SignedString([]byte(secret))

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)
	rr := httptest.NewRecorder()

	middleware(handler).ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	assert.Contains(t, rr.Body.String(), "Invalid user_id format")
}

func TestAuthMiddleware_MissingTenantID(t *testing.T) {
	secret := "test-secret"
	middleware := AuthMiddleware(AuthMiddlewareConfig{JWTSecret: secret})

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": uuid.New().String(),
		"role":    "admin",
		"exp":     time.Now().Add(time.Hour).Unix(),
	})
	tokenString, _ := token.SignedString([]byte(secret))

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)
	rr := httptest.NewRecorder()

	middleware(handler).ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	assert.Contains(t, rr.Body.String(), "Missing tenant_id in token")
}

func TestAuthMiddleware_InvalidTenantIDFormat(t *testing.T) {
	secret := "test-secret"
	middleware := AuthMiddleware(AuthMiddlewareConfig{JWTSecret: secret})

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   uuid.New().String(),
		"tenant_id": "not-a-valid-uuid",
		"role":      "admin",
		"exp":       time.Now().Add(time.Hour).Unix(),
	})
	tokenString, _ := token.SignedString([]byte(secret))

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)
	rr := httptest.NewRecorder()

	middleware(handler).ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	assert.Contains(t, rr.Body.String(), "Invalid tenant_id format")
}

func TestAuthMiddleware_MissingRole(t *testing.T) {
	secret := "test-secret"
	middleware := AuthMiddleware(AuthMiddlewareConfig{JWTSecret: secret})

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   uuid.New().String(),
		"tenant_id": uuid.New().String(),
		"exp":       time.Now().Add(time.Hour).Unix(),
	})
	tokenString, _ := token.SignedString([]byte(secret))

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)
	rr := httptest.NewRecorder()

	middleware(handler).ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	assert.Contains(t, rr.Body.String(), "Missing role in token")
}

func TestAuthMiddleware_CaseInsensitiveBearer(t *testing.T) {
	secret := "test-secret"
	middleware := AuthMiddleware(AuthMiddlewareConfig{JWTSecret: secret})

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   uuid.New().String(),
		"tenant_id": uuid.New().String(),
		"role":      "admin",
		"exp":       time.Now().Add(time.Hour).Unix(),
	})
	tokenString, _ := token.SignedString([]byte(secret))

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Test with lowercase "bearer"
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "bearer "+tokenString)
	rr := httptest.NewRecorder()

	middleware(handler).ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestRequirePermission_Success(t *testing.T) {
	mockChecker := new(mocks.PermissionChecker)
	middleware := RequirePermission(mockChecker, "org", "write")

	mockChecker.On("Enforce", "admin", "org", "write").Return(true, nil)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	ctx := fixtures.NewContextWithAdminUser()
	req = req.WithContext(ctx)
	rr := httptest.NewRecorder()

	middleware(handler).ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockChecker.AssertExpectations(t)
}

func TestRequirePermission_Forbidden(t *testing.T) {
	mockChecker := new(mocks.PermissionChecker)
	middleware := RequirePermission(mockChecker, "org", "write")

	mockChecker.On("Enforce", "viewer", "org", "write").Return(false, nil)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	claims := fixtures.NewTestUserClaims()
	claims.Role = "viewer"
	ctx := fixtures.NewContextWithUserClaims(claims)
	req = req.WithContext(ctx)
	rr := httptest.NewRecorder()

	middleware(handler).ServeHTTP(rr, req)

	assert.Equal(t, http.StatusForbidden, rr.Code)
	assert.Contains(t, rr.Body.String(), "Insufficient permissions")
	mockChecker.AssertExpectations(t)
}

func TestRequirePermission_NoClaims(t *testing.T) {
	mockChecker := new(mocks.PermissionChecker)
	middleware := RequirePermission(mockChecker, "org", "write")

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	// No claims in context
	rr := httptest.NewRecorder()

	middleware(handler).ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	assert.Contains(t, rr.Body.String(), "Authentication required")
}

func TestRequirePermission_EnforcerError(t *testing.T) {
	mockChecker := new(mocks.PermissionChecker)
	middleware := RequirePermission(mockChecker, "org", "write")

	mockChecker.On("Enforce", "admin", "org", "write").Return(false, assert.AnError)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	ctx := fixtures.NewContextWithAdminUser()
	req = req.WithContext(ctx)
	rr := httptest.NewRecorder()

	middleware(handler).ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Contains(t, rr.Body.String(), "Authorization check failed")
	mockChecker.AssertExpectations(t)
}
