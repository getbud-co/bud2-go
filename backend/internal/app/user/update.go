package user

import (
	"context"
	"errors"

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
	repo usr.Repository
}

func NewUpdateUseCase(repo usr.Repository) *UpdateUseCase {
	return &UpdateUseCase{repo: repo}
}

func (uc *UpdateUseCase) Execute(ctx context.Context, cmd UpdateCommand) (*usr.User, error) {
	existing, err := uc.repo.GetByID(ctx, cmd.TenantID, cmd.ID)
	if err != nil {
		return nil, err
	}

	originalEmail := existing.Email

	existing.Name = cmd.Name
	existing.Email = cmd.Email
	existing.Role = usr.Role(cmd.Role)
	existing.Status = usr.Status(cmd.Status)

	if err := existing.Validate(); err != nil {
		return nil, err
	}

	if originalEmail != cmd.Email {
		other, err := uc.repo.GetByEmail(ctx, cmd.TenantID, cmd.Email)
		if err == nil && other.ID != cmd.ID {
			return nil, usr.ErrEmailExists
		}
		if err != nil && !errors.Is(err, usr.ErrNotFound) {
			return nil, err
		}
	}

	return uc.repo.Update(ctx, existing)
}
