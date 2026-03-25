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

type CreateUseCase struct {
	repo Repository
}

func NewCreateUseCase(repo Repository) *CreateUseCase {
	return &CreateUseCase{repo: repo}
}

func (uc *CreateUseCase) Execute(ctx context.Context, cmd CreateCommand) (*User, error) {
	user := &User{
		ID:       uuid.New(),
		TenantID: cmd.TenantID,
		Name:     cmd.Name,
		Email:    cmd.Email,
		Role:     Role(cmd.Role),
		Status:   StatusActive,
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if _, err := uc.repo.GetByEmail(ctx, cmd.TenantID, cmd.Email); err == nil {
		return nil, ErrEmailExists
	} else if !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	return uc.repo.Create(ctx, user)
}
