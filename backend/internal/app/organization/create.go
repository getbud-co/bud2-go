package organization

import (
	"context"
	"errors"

	org "github.com/dsbraz/bud2/backend/internal/domain/organization"
)

type CreateCommand struct {
	Name   string
	Slug   string
	Status string
}

type CreateUseCase struct {
	repo org.Repository
}

func NewCreateUseCase(repo org.Repository) *CreateUseCase {
	return &CreateUseCase{repo: repo}
}

func (uc *CreateUseCase) Execute(ctx context.Context, cmd CreateCommand) (*org.Organization, error) {
	status := org.StatusActive
	if cmd.Status != "" {
		status = org.Status(cmd.Status)
	}

	o := &org.Organization{
		Name:   cmd.Name,
		Slug:   cmd.Slug,
		Status: status,
	}

	if err := o.Validate(); err != nil {
		return nil, err
	}

	_, err := uc.repo.GetBySlug(ctx, cmd.Slug)
	if err == nil {
		return nil, org.ErrSlugExists
	}
	if !errors.Is(err, org.ErrNotFound) {
		return nil, err
	}

	return uc.repo.Create(ctx, o)
}
