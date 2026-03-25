package user

import (
	"context"

	"github.com/dsbraz/bud2/backend/internal/domain"
)

type ListCommand struct {
	TenantID domain.TenantID
	Status   *string
	Search   *string
	Page     int
	Size     int
}

type ListUseCase struct{ repo domain.UserRepository }

func NewListUseCase(repo domain.UserRepository) *ListUseCase {
	return &ListUseCase{repo: repo}
}

func (uc *ListUseCase) Execute(ctx context.Context, cmd ListCommand) (domain.UserListResult, error) {
	if cmd.Size <= 0 {
		cmd.Size = 20
	}
	if cmd.Size > 100 {
		cmd.Size = 100
	}
	if cmd.Page <= 0 {
		cmd.Page = 1
	}

	filter := domain.UserListFilter{
		TenantID: cmd.TenantID,
		Page:     cmd.Page,
		Size:     cmd.Size,
	}
	if cmd.Status != nil {
		s := domain.UserStatus(*cmd.Status)
		filter.Status = &s
	}
	if cmd.Search != nil && *cmd.Search != "" {
		filter.Search = cmd.Search
	}

	return uc.repo.List(ctx, filter)
}
