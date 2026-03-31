package middleware

import (
	"net/http"
	"strings"

	"github.com/dsbraz/bud2/backend/internal/api/httputil"
	"github.com/dsbraz/bud2/backend/internal/domain"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthMiddlewareConfig struct {
	JWTSecret string
}

func AuthMiddleware(cfg AuthMiddlewareConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				httputil.WriteProblem(w, http.StatusUnauthorized, "Unauthorized", "Authorization header is required")
				return
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				httputil.WriteProblem(w, http.StatusUnauthorized, "Unauthorized", "Invalid authorization header format")
				return
			}

			tokenString := parts[1]
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, jwt.ErrSignatureInvalid
				}
				return []byte(cfg.JWTSecret), nil
			})

			if err != nil || !token.Valid {
				httputil.WriteProblem(w, http.StatusUnauthorized, "Unauthorized", "Invalid or expired token")
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				httputil.WriteProblem(w, http.StatusUnauthorized, "Unauthorized", "Invalid token claims")
				return
			}

			userIDStr, ok := claims["user_id"].(string)
			if !ok {
				httputil.WriteProblem(w, http.StatusUnauthorized, "Unauthorized", "Missing user_id in token")
				return
			}
			userID, err := uuid.Parse(userIDStr)
			if err != nil {
				httputil.WriteProblem(w, http.StatusUnauthorized, "Unauthorized", "Invalid user_id format")
				return
			}

			tenantIDStr, ok := claims["tenant_id"].(string)
			if !ok {
				httputil.WriteProblem(w, http.StatusUnauthorized, "Unauthorized", "Missing tenant_id in token")
				return
			}
			tenantID, err := uuid.Parse(tenantIDStr)
			if err != nil {
				httputil.WriteProblem(w, http.StatusUnauthorized, "Unauthorized", "Invalid tenant_id format")
				return
			}

			role, ok := claims["role"].(string)
			if !ok {
				httputil.WriteProblem(w, http.StatusUnauthorized, "Unauthorized", "Missing role in token")
				return
			}

			userClaims := domain.UserClaims{
				UserID:   domain.UserID(userID),
				TenantID: domain.TenantID(tenantID),
				Role:     role,
			}

			ctx := domain.ClaimsToContext(r.Context(), userClaims)
			ctx = domain.TenantIDToContext(ctx, userClaims.TenantID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

type PermissionChecker interface {
	Enforce(rvals ...interface{}) (bool, error)
}

func RequirePermission(checker PermissionChecker, obj, act string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims, err := domain.ClaimsFromContext(r.Context())
			if err != nil {
				httputil.WriteProblem(w, http.StatusUnauthorized, "Unauthorized", "Authentication required")
				return
			}

			allowed, err := checker.Enforce(claims.Role, obj, act)
			if err != nil {
				httputil.WriteProblem(w, http.StatusInternalServerError, "Internal Server Error", "Authorization check failed")
				return
			}

			if !allowed {
				httputil.WriteProblem(w, http.StatusForbidden, "Forbidden", "Insufficient permissions")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
