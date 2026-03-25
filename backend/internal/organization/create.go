package organization

import (
	"context"
	"errors"
)

type CreateCommand struct {
	Name   string
	Slug   string
	Status string
}

type CreateUseCase struct {
	repo Repository
}

func NewCreateUseCase(repo Repository) *CreateUseCase {
	return &CreateUseCase{repo: repo}
}

func (uc *CreateUseCase) Execute(ctx context.Context, cmd CreateCommand) (*Organization, error) {
	status := StatusActive
	if cmd.Status != "" {
		status = Status(cmd.Status)
	}

	org := &Organization{
		Name:   cmd.Name,
		Slug:   cmd.Slug,
		Status: status,
	}

	if err := org.Validate(); err != nil {
		return nil, err
	}

	_, err := uc.repo.GetBySlug(ctx, cmd.Slug)
	if err == nil {
		return nil, ErrSlugExists
	}
	if !errors.Is(err, ErrNotFound) {
		return nil, err
	}

	return uc.repo.Create(ctx, org)
}
