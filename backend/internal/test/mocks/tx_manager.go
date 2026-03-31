package mocks

import (
	"context"

	"github.com/dsbraz/bud2/backend/internal/domain/organization"
	"github.com/dsbraz/bud2/backend/internal/domain/user"
	"github.com/stretchr/testify/mock"
)

type TxManager struct {
	mock.Mock
}

func (m *TxManager) WithTx(ctx context.Context, fn func(orgRepo organization.Repository, userRepo user.Repository) error) error {
	args := m.Called(ctx, fn)
	return args.Error(0)
}
