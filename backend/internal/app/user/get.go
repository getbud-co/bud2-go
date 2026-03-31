package user

import (
	"context"

	"github.com/dsbraz/bud2/backend/internal/domain"
	usr "github.com/dsbraz/bud2/backend/internal/domain/user"
	"github.com/google/uuid"
)

type GetUseCase struct {
	repo usr.Repository
}

func NewGetUseCase(repo usr.Repository) *GetUseCase {
	return &GetUseCase{repo: repo}
}

func (uc *GetUseCase) Execute(ctx context.Context, tenantID domain.TenantID, id uuid.UUID) (*usr.User, error) {
	return uc.repo.GetByID(ctx, tenantID, id)
}
