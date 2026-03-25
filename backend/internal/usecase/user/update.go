package user

import (
	"context"
	"errors"

	"github.com/dsbraz/bud2/backend/internal/domain"
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

type UpdateUseCase struct{ repo domain.UserRepository }

func NewUpdateUseCase(repo domain.UserRepository) *UpdateUseCase {
	return &UpdateUseCase{repo: repo}
}

func (uc *UpdateUseCase) Execute(ctx context.Context, cmd UpdateCommand) (*domain.User, error) {
	existing, err := uc.repo.GetByID(ctx, cmd.TenantID, cmd.ID)
	if err != nil {
		return nil, err
	}

	originalEmail := existing.Email

	existing.Name = cmd.Name
	existing.Email = cmd.Email
	existing.Role = domain.UserRole(cmd.Role)
	existing.Status = domain.UserStatus(cmd.Status)

	if err := existing.Validate(); err != nil {
		return nil, err
	}

	if originalEmail != cmd.Email {
		other, err := uc.repo.GetByEmail(ctx, cmd.TenantID, cmd.Email)
		if err == nil && other.ID != cmd.ID {
			return nil, domain.ErrUserEmailExists
		}
		if err != nil && !errors.Is(err, domain.ErrUserNotFound) {
			return nil, err
		}
	}

	return uc.repo.Update(ctx, existing)
}
