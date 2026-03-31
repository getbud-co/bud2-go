package organization

import (
	"context"

	org "github.com/dsbraz/bud2/backend/internal/domain/organization"
)

type ListCommand struct {
	Status *string
	Page   int
	Size   int
}

type ListUseCase struct {
	repo org.Repository
}

func NewListUseCase(repo org.Repository) *ListUseCase {
	return &ListUseCase{repo: repo}
}

func (uc *ListUseCase) Execute(ctx context.Context, cmd ListCommand) (org.ListResult, error) {
	if cmd.Size <= 0 {
		cmd.Size = 20
	}
	if cmd.Size > 100 {
		cmd.Size = 100
	}
	if cmd.Page <= 0 {
		cmd.Page = 1
	}

	filter := org.ListFilter{
		Page: cmd.Page,
		Size: cmd.Size,
	}
	if cmd.Status != nil {
		s := org.Status(*cmd.Status)
		filter.Status = &s
	}

	return uc.repo.List(ctx, filter)
}
