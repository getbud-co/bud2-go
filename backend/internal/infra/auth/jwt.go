package auth

import (
	"time"

	"github.com/dsbraz/bud2/backend/internal/domain"
	"github.com/golang-jwt/jwt/v5"
)

type TokenIssuer struct {
	secret string
}

func NewTokenIssuer(secret string) *TokenIssuer {
	return &TokenIssuer{secret: secret}
}

func (ti *TokenIssuer) IssueToken(claims domain.UserClaims, ttl time.Duration) (string, error) {
	now := time.Now().UTC()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   claims.UserID.String(),
		"tenant_id": claims.TenantID.String(),
		"role":      claims.Role,
		"iat":       now.Unix(),
		"exp":       now.Add(ttl).Unix(),
	})

	return token.SignedString([]byte(ti.secret))
}
