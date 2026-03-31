package organization

import (
	"context"
	"errors"

	org "github.com/dsbraz/bud2/backend/internal/domain/organization"
	"github.com/google/uuid"
)

type UpdateCommand struct {
	ID     uuid.UUID
	Name   string
	Slug   string
	Status string
}

type UpdateUseCase struct {
	repo org.Repository
}

func NewUpdateUseCase(repo org.Repository) *UpdateUseCase {
	return &UpdateUseCase{repo: repo}
}

func (uc *UpdateUseCase) Execute(ctx context.Context, cmd UpdateCommand) (*org.Organization, error) {
	existing, err := uc.repo.GetByID(ctx, cmd.ID)
	if err != nil {
		return nil, err
	}

	originalSlug := existing.Slug

	existing.Name = cmd.Name
	existing.Slug = cmd.Slug
	existing.Status = org.Status(cmd.Status)

	if err := existing.Validate(); err != nil {
		return nil, err
	}

	if originalSlug != cmd.Slug {
		other, err := uc.repo.GetBySlug(ctx, cmd.Slug)
		if err == nil && other.ID != cmd.ID {
			return nil, org.ErrSlugExists
		}
		if err != nil && !errors.Is(err, org.ErrNotFound) {
			return nil, err
		}
	}

	return uc.repo.Update(ctx, existing)
}
