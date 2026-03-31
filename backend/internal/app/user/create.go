package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/dsbraz/bud2/backend/internal/domain"
	usr "github.com/dsbraz/bud2/backend/internal/domain/user"
	"github.com/dsbraz/bud2/backend/internal/infra/auth"
	"github.com/google/uuid"
)

type CreateCommand struct {
	TenantID domain.TenantID
	Name     string
	Email    string
	Password string
	Role     string
}

type CreateUseCase struct {
	repo usr.Repository
}

func NewCreateUseCase(repo usr.Repository) *CreateUseCase {
	return &CreateUseCase{repo: repo}
}

func (uc *CreateUseCase) Execute(ctx context.Context, cmd CreateCommand) (*usr.User, error) {
	// Hash password
	passwordHash, err := auth.HashPassword(cmd.Password)
	if err != nil {
		return nil, fmt.Errorf("%w: invalid password", domain.ErrValidation)
	}

	u := &usr.User{
		ID:           uuid.New(),
		TenantID:     cmd.TenantID,
		Name:         cmd.Name,
		Email:        cmd.Email,
		PasswordHash: passwordHash,
		Role:         usr.Role(cmd.Role),
		Status:       usr.StatusActive,
	}

	if err := u.Validate(); err != nil {
		return nil, err
	}

	if _, err := uc.repo.GetByEmail(ctx, cmd.TenantID, cmd.Email); err == nil {
		return nil, usr.ErrEmailExists
	} else if !errors.Is(err, usr.ErrNotFound) {
		return nil, err
	}

	return uc.repo.Create(ctx, u)
}
