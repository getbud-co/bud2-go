package organization

import (
	"context"

	"github.com/dsbraz/bud2/backend/internal/domain"
)

type ListCommand struct {
	Status *string
	Page   int
	Size   int
}

type ListUseCase struct {
	repo domain.OrganizationRepository
}

func NewListUseCase(repo domain.OrganizationRepository) *ListUseCase {
	return &ListUseCase{repo: repo}
}

func (uc *ListUseCase) Execute(ctx context.Context, cmd ListCommand) (domain.OrganizationListResult, error) {
	if cmd.Size <= 0 {
		cmd.Size = 20
	}
	if cmd.Size > 100 {
		cmd.Size = 100
	}
	if cmd.Page <= 0 {
		cmd.Page = 1
	}

	filter := domain.OrganizationListFilter{
		Page: cmd.Page,
		Size: cmd.Size,
	}
	if cmd.Status != nil {
		s := domain.OrganizationStatus(*cmd.Status)
		filter.Status = &s
	}

	return uc.repo.List(ctx, filter)
}
