package mocks

import (
	"time"

	"github.com/dsbraz/bud2/backend/internal/domain"
	"github.com/stretchr/testify/mock"
)

type TokenIssuer struct {
	mock.Mock
}

func (m *TokenIssuer) IssueToken(claims domain.UserClaims, ttl time.Duration) (string, error) {
	args := m.Called(claims, ttl)
	return args.String(0), args.Error(1)
}
