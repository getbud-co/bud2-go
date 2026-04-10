package organization

import (
	"context"
	"errors"
	"log/slog"

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
	repo   org.Repository
	logger *slog.Logger
}

func NewUpdateUseCase(repo org.Repository, logger *slog.Logger) *UpdateUseCase {
	return &UpdateUseCase{repo: repo, logger: logger}
}

func (uc *UpdateUseCase) Execute(ctx context.Context, cmd UpdateCommand) (*org.Organization, error) {
	uc.logger.Debug("updating organization", "organization_id", cmd.ID)

	existing, err := uc.repo.GetByID(ctx, cmd.ID)
	if err != nil {
		uc.logger.Error("failed to fetch organization for update", "error", err, "organization_id", cmd.ID)
		return nil, err
	}

	originalSlug := existing.Slug

	existing.Name = cmd.Name
	existing.Slug = cmd.Slug
	existing.Status = org.Status(cmd.Status)

	if err := existing.Validate(); err != nil {
		uc.logger.Warn("organization validation failed", "error", err, "organization_id", cmd.ID)
		return nil, err
	}

	if originalSlug != cmd.Slug {
		other, err := uc.repo.GetBySlug(ctx, cmd.Slug)
		if err == nil && other.ID != cmd.ID {
			uc.logger.Warn("slug conflict", "slug", cmd.Slug, "organization_id", cmd.ID)
			return nil, org.ErrSlugExists
		}
		if err != nil && !errors.Is(err, org.ErrNotFound) {
			uc.logger.Error("failed to check slug uniqueness", "error", err, "slug", cmd.Slug)
			return nil, err
		}
	}

	result, err := uc.repo.Update(ctx, existing)
	if err != nil {
		uc.logger.Error("failed to update organization", "error", err, "organization_id", cmd.ID)
		return nil, err
	}

	uc.logger.Info("organization updated", "organization_id", result.ID, "slug", result.Slug)
	return result, nil
}
