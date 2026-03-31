package organization

import (
	"context"

	org "github.com/dsbraz/bud2/backend/internal/domain/organization"
	"github.com/google/uuid"
)

type GetUseCase struct {
	repo org.Repository
}

func NewGetUseCase(repo org.Repository) *GetUseCase {
	return &GetUseCase{repo: repo}
}

func (uc *GetUseCase) Execute(ctx context.Context, id uuid.UUID) (*org.Organization, error) {
	return uc.repo.GetByID(ctx, id)
}
