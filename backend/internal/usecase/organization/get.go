package organization

import (
	"context"

	"github.com/dsbraz/bud2/backend/internal/domain"
	"github.com/google/uuid"
)

type GetUseCase struct {
	repo domain.OrganizationRepository
}

func NewGetUseCase(repo domain.OrganizationRepository) *GetUseCase {
	return &GetUseCase{repo: repo}
}

func (uc *GetUseCase) Execute(ctx context.Context, id uuid.UUID) (*domain.Organization, error) {
	return uc.repo.GetByID(ctx, id)
}
