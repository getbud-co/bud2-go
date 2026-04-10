package user

import (
	"context"
	"log/slog"

	"github.com/dsbraz/bud2/backend/internal/domain"
	usr "github.com/dsbraz/bud2/backend/internal/domain/user"
	"github.com/google/uuid"
)

type GetUseCase struct {
	repo   usr.Repository
	logger *slog.Logger
}

func NewGetUseCase(repo usr.Repository, logger *slog.Logger) *GetUseCase {
	return &GetUseCase{repo: repo, logger: logger}
}

func (uc *GetUseCase) Execute(ctx context.Context, tenantID domain.TenantID, id uuid.UUID) (*usr.User, error) {
	uc.logger.Debug("fetching user", "user_id", id, "tenant_id", tenantID)
	result, err := uc.repo.GetByID(ctx, tenantID, id)
	if err != nil {
		uc.logger.Error("failed to fetch user", "error", err, "user_id", id, "tenant_id", tenantID)
		return nil, err
	}
	return result, nil
}
