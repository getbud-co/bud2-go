package organization

import (
	"context"
)

type ListCommand struct {
	Status *string
	Page   int
	Size   int
}

type ListUseCase struct {
	repo Repository
}

func NewListUseCase(repo Repository) *ListUseCase {
	return &ListUseCase{repo: repo}
}

func (uc *ListUseCase) Execute(ctx context.Context, cmd ListCommand) (ListResult, error) {
	if cmd.Size <= 0 {
		cmd.Size = 20
	}
	if cmd.Size > 100 {
		cmd.Size = 100
	}
	if cmd.Page <= 0 {
		cmd.Page = 1
	}

	filter := ListFilter{
		Page: cmd.Page,
		Size: cmd.Size,
	}
	if cmd.Status != nil {
		s := Status(*cmd.Status)
		filter.Status = &s
	}

	return uc.repo.List(ctx, filter)
}
