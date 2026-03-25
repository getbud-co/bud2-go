package user

import (
	"context"

	"github.com/dsbraz/bud2/backend/internal/domain"
	"github.com/google/uuid"
)

type GetUseCase struct{ repo domain.UserRepository }

func NewGetUseCase(repo domain.UserRepository) *GetUseCase {
	return &GetUseCase{repo: repo}
}

func (uc *GetUseCase) Execute(ctx context.Context, tenantID domain.TenantID, id uuid.UUID) (*domain.User, error) {
	return uc.repo.GetByID(ctx, tenantID, id)
}
