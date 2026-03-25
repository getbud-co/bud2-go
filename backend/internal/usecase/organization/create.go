package organization

import (
	"context"
	"errors"

	"github.com/dsbraz/bud2/backend/internal/domain"
)

type CreateCommand struct {
	Name   string
	Slug   string
	Status string
}

type CreateUseCase struct {
	repo domain.OrganizationRepository
}

func NewCreateUseCase(repo domain.OrganizationRepository) *CreateUseCase {
	return &CreateUseCase{repo: repo}
}

func (uc *CreateUseCase) Execute(ctx context.Context, cmd CreateCommand) (*domain.Organization, error) {
	status := domain.OrganizationStatus(cmd.Status)
	if cmd.Status == "" {
		status = domain.OrganizationStatusActive
	}

	org := &domain.Organization{
		Name:   cmd.Name,
		Slug:   cmd.Slug,
		Status: status,
	}

	if err := org.Validate(); err != nil {
		return nil, err
	}

	_, err := uc.repo.GetBySlug(ctx, cmd.Slug)
	if err == nil {
		return nil, domain.ErrOrganizationSlugExists
	}
	if !errors.Is(err, domain.ErrOrganizationNotFound) {
		return nil, err
	}

	return uc.repo.Create(ctx, org)
}
