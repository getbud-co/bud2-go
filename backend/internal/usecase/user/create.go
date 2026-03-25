package user

import (
	"context"
	"errors"

	"github.com/dsbraz/bud2/backend/internal/domain"
	"github.com/google/uuid"
)

type CreateCommand struct {
	TenantID domain.TenantID
	Name     string
	Email    string
	Role     string
}

type CreateUseCase struct{ repo domain.UserRepository }

func NewCreateUseCase(repo domain.UserRepository) *CreateUseCase {
	return &CreateUseCase{repo: repo}
}

func (uc *CreateUseCase) Execute(ctx context.Context, cmd CreateCommand) (*domain.User, error) {
	user := &domain.User{
		ID:       uuid.New(),
		TenantID: cmd.TenantID,
		Name:     cmd.Name,
		Email:    cmd.Email,
		Role:     domain.UserRole(cmd.Role),
		Status:   domain.UserStatusActive,
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if _, err := uc.repo.GetByEmail(ctx, cmd.TenantID, cmd.Email); err == nil {
		return nil, domain.ErrUserEmailExists
	} else if !errors.Is(err, domain.ErrUserNotFound) {
		return nil, err
	}

	return uc.repo.Create(ctx, user)
}
