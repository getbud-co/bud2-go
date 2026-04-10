package user

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

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
	repo   usr.Repository
	logger *slog.Logger
}

func NewCreateUseCase(repo usr.Repository, logger *slog.Logger) *CreateUseCase {
	return &CreateUseCase{repo: repo, logger: logger}
}

func (uc *CreateUseCase) Execute(ctx context.Context, cmd CreateCommand) (*usr.User, error) {
	uc.logger.Debug("creating user", "email", cmd.Email, "tenant_id", cmd.TenantID)

	passwordHash, err := auth.HashPassword(cmd.Password)
	if err != nil {
		uc.logger.Error("failed to hash password", "error", err)
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
		uc.logger.Warn("user validation failed", "error", err, "email", cmd.Email)
		return nil, err
	}

	if _, err := uc.repo.GetByEmail(ctx, cmd.TenantID, cmd.Email); err == nil {
		uc.logger.Warn("email conflict", "email", cmd.Email, "tenant_id", cmd.TenantID)
		return nil, usr.ErrEmailExists
	} else if !errors.Is(err, usr.ErrNotFound) {
		uc.logger.Error("failed to check email uniqueness", "error", err, "email", cmd.Email)
		return nil, err
	}

	result, err := uc.repo.Create(ctx, u)
	if err != nil {
		uc.logger.Error("failed to create user", "error", err, "email", cmd.Email)
		return nil, err
	}

	uc.logger.Info("user created", "user_id", result.ID, "email", result.Email, "tenant_id", result.TenantID)
	return result, nil
}
