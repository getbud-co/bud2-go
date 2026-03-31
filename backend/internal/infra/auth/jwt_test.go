package auth

import (
	"testing"
	"time"

	"github.com/dsbraz/bud2/backend/internal/domain"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTokenIssuer_IssueToken(t *testing.T) {
	issuer := NewTokenIssuer("test-secret")

	claims := domain.UserClaims{
		UserID:   domain.UserID(uuid.New()),
		TenantID: domain.TenantID(uuid.New()),
		Role:     "admin",
	}

	tokenString, err := issuer.IssueToken(claims, 24*time.Hour)
	require.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	// Parse and verify token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("test-secret"), nil
	})
	require.NoError(t, err)
	assert.True(t, token.Valid)

	// Verify claims
	mapClaims, ok := token.Claims.(jwt.MapClaims)
	require.True(t, ok)
	assert.Equal(t, claims.UserID.String(), mapClaims["user_id"])
	assert.Equal(t, claims.TenantID.String(), mapClaims["tenant_id"])
	assert.Equal(t, claims.Role, mapClaims["role"])
	assert.NotNil(t, mapClaims["iat"])
	assert.NotNil(t, mapClaims["exp"])
}

func TestTokenIssuer_IssueToken_DifferentSecrets(t *testing.T) {
	issuer1 := NewTokenIssuer("secret-1")
	issuer2 := NewTokenIssuer("secret-2")

	claims := domain.UserClaims{
		UserID:   domain.UserID(uuid.New()),
		TenantID: domain.TenantID(uuid.New()),
		Role:     "admin",
	}

	token1, err := issuer1.IssueToken(claims, 24*time.Hour)
	require.NoError(t, err)

	token2, err := issuer2.IssueToken(claims, 24*time.Hour)
	require.NoError(t, err)

	// Tokens should be different due to different signatures
	assert.NotEqual(t, token1, token2)
}

func TestTokenIssuer_IssueToken_VerifyWithWrongSecret(t *testing.T) {
	issuer := NewTokenIssuer("correct-secret")

	claims := domain.UserClaims{
		UserID:   domain.UserID(uuid.New()),
		TenantID: domain.TenantID(uuid.New()),
		Role:     "admin",
	}

	tokenString, err := issuer.IssueToken(claims, 24*time.Hour)
	require.NoError(t, err)

	// Try to verify with wrong secret
	_, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("wrong-secret"), nil
	})
	assert.Error(t, err)
}

func TestTokenIssuer_IssueToken_TTL(t *testing.T) {
	issuer := NewTokenIssuer("test-secret")

	claims := domain.UserClaims{
		UserID:   domain.UserID(uuid.New()),
		TenantID: domain.TenantID(uuid.New()),
		Role:     "admin",
	}

	// Issue token with 1 hour TTL
	tokenString, err := issuer.IssueToken(claims, time.Hour)
	require.NoError(t, err)

	// Parse and check expiration
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("test-secret"), nil
	})
	require.NoError(t, err)

	mapClaims := token.Claims.(jwt.MapClaims)
	exp := int64(mapClaims["exp"].(float64))
	iat := int64(mapClaims["iat"].(float64))

	// Expiration should be approximately 1 hour after issued at
	ttl := time.Duration(exp-iat) * time.Second
	assert.InDelta(t, time.Hour.Seconds(), ttl.Seconds(), 5) // Allow 5 seconds delta
}
