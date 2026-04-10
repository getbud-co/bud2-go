package organization

import (
	"context"
	"log/slog"

	org "github.com/dsbraz/bud2/backend/internal/domain/organization"
	"github.com/google/uuid"
)

type GetUseCase struct {
	repo   org.Repository
	logger *slog.Logger
}

func NewGetUseCase(repo org.Repository, logger *slog.Logger) *GetUseCase {
	return &GetUseCase{repo: repo, logger: logger}
}

func (uc *GetUseCase) Execute(ctx context.Context, id uuid.UUID) (*org.Organization, error) {
	uc.logger.Debug("fetching organization", "organization_id", id)
	result, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		uc.logger.Error("failed to fetch organization", "error", err, "organization_id", id)
		return nil, err
	}
	return result, nil
}
