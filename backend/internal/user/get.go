package user

import (
	"context"

	"github.com/dsbraz/bud2/backend/internal/domain"
	"github.com/google/uuid"
)

type GetUseCase struct {
	repo Repository
}

func NewGetUseCase(repo Repository) *GetUseCase {
	return &GetUseCase{repo: repo}
}

func (uc *GetUseCase) Execute(ctx context.Context, tenantID domain.TenantID, id uuid.UUID) (*User, error) {
	return uc.repo.GetByID(ctx, tenantID, id)
}
