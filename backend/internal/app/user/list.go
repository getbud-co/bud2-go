package user

import (
	"context"

	"github.com/dsbraz/bud2/backend/internal/domain"
	usr "github.com/dsbraz/bud2/backend/internal/domain/user"
)

type ListCommand struct {
	TenantID domain.TenantID
	Status   *string
	Search   *string
	Page     int
	Size     int
}

type ListUseCase struct {
	repo usr.Repository
}

func NewListUseCase(repo usr.Repository) *ListUseCase {
	return &ListUseCase{repo: repo}
}

func (uc *ListUseCase) Execute(ctx context.Context, cmd ListCommand) (usr.ListResult, error) {
	if cmd.Size <= 0 {
		cmd.Size = 20
	}
	if cmd.Size > 100 {
		cmd.Size = 100
	}
	if cmd.Page <= 0 {
		cmd.Page = 1
	}

	filter := usr.ListFilter{
		TenantID: cmd.TenantID,
		Page:     cmd.Page,
		Size:     cmd.Size,
	}
	if cmd.Status != nil {
		s := usr.Status(*cmd.Status)
		filter.Status = &s
	}
	if cmd.Search != nil && *cmd.Search != "" {
		filter.Search = cmd.Search
	}

	return uc.repo.List(ctx, filter)
}
