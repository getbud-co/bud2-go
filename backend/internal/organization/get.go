package organization

import (
	"context"

	"github.com/google/uuid"
)

type GetUseCase struct {
	repo Repository
}

func NewGetUseCase(repo Repository) *GetUseCase {
	return &GetUseCase{repo: repo}
}

func (uc *GetUseCase) Execute(ctx context.Context, id uuid.UUID) (*Organization, error) {
	return uc.repo.GetByID(ctx, id)
}
