package organization

import (
	"context"
	"errors"
	"log/slog"

	org "github.com/dsbraz/bud2/backend/internal/domain/organization"
)

type CreateCommand struct {
	Name   string
	Slug   string
	Status string
}

type CreateUseCase struct {
	repo   org.Repository
	logger *slog.Logger
}

func NewCreateUseCase(repo org.Repository, logger *slog.Logger) *CreateUseCase {
	return &CreateUseCase{repo: repo, logger: logger}
}

func (uc *CreateUseCase) Execute(ctx context.Context, cmd CreateCommand) (*org.Organization, error) {
	uc.logger.Debug("creating organization", "name", cmd.Name, "slug", cmd.Slug)

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
		uc.logger.Warn("organization validation failed", "error", err, "slug", cmd.Slug)
		return nil, err
	}

	_, err := uc.repo.GetBySlug(ctx, cmd.Slug)
	if err == nil {
		uc.logger.Warn("slug conflict", "slug", cmd.Slug)
		return nil, org.ErrSlugExists
	}
	if !errors.Is(err, org.ErrNotFound) {
		uc.logger.Error("failed to check slug uniqueness", "error", err, "slug", cmd.Slug)
		return nil, err
	}

	result, err := uc.repo.Create(ctx, o)
	if err != nil {
		uc.logger.Error("failed to create organization", "error", err, "slug", cmd.Slug)
		return nil, err
	}

	uc.logger.Info("organization created", "organization_id", result.ID, "slug", result.Slug)
	return result, nil
}
