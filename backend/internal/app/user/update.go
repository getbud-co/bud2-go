package user

import (
	"context"
	"errors"
	"log/slog"

	"github.com/dsbraz/bud2/backend/internal/domain"
	usr "github.com/dsbraz/bud2/backend/internal/domain/user"
	"github.com/google/uuid"
)

type UpdateCommand struct {
	TenantID domain.TenantID
	ID       uuid.UUID
	Name     string
	Email    string
	Role     string
	Status   string
}

type UpdateUseCase struct {
	repo   usr.Repository
	logger *slog.Logger
}

func NewUpdateUseCase(repo usr.Repository, logger *slog.Logger) *UpdateUseCase {
	return &UpdateUseCase{repo: repo, logger: logger}
}

func (uc *UpdateUseCase) Execute(ctx context.Context, cmd UpdateCommand) (*usr.User, error) {
	uc.logger.Debug("updating user", "user_id", cmd.ID, "tenant_id", cmd.TenantID)

	existing, err := uc.repo.GetByID(ctx, cmd.TenantID, cmd.ID)
	if err != nil {
		uc.logger.Error("failed to fetch user for update", "error", err, "user_id", cmd.ID, "tenant_id", cmd.TenantID)
		return nil, err
	}

	originalEmail := existing.Email

	existing.Name = cmd.Name
	existing.Email = cmd.Email
	existing.Role = usr.Role(cmd.Role)
	existing.Status = usr.Status(cmd.Status)

	if err := existing.Validate(); err != nil {
		uc.logger.Warn("user validation failed", "error", err, "user_id", cmd.ID)
		return nil, err
	}

	if originalEmail != cmd.Email {
		other, err := uc.repo.GetByEmail(ctx, cmd.TenantID, cmd.Email)
		if err == nil && other.ID != cmd.ID {
			uc.logger.Warn("email conflict", "email", cmd.Email, "user_id", cmd.ID, "tenant_id", cmd.TenantID)
			return nil, usr.ErrEmailExists
		}
		if err != nil && !errors.Is(err, usr.ErrNotFound) {
			uc.logger.Error("failed to check email uniqueness", "error", err, "email", cmd.Email)
			return nil, err
		}
	}

	result, err := uc.repo.Update(ctx, existing)
	if err != nil {
		uc.logger.Error("failed to update user", "error", err, "user_id", cmd.ID, "tenant_id", cmd.TenantID)
		return nil, err
	}

	uc.logger.Info("user updated", "user_id", result.ID, "email", result.Email, "tenant_id", cmd.TenantID)
	return result, nil
}
