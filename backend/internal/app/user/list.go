package user

import (
	"context"
	"log/slog"

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
	repo   usr.Repository
	logger *slog.Logger
}

func NewListUseCase(repo usr.Repository, logger *slog.Logger) *ListUseCase {
	return &ListUseCase{repo: repo, logger: logger}
}

func (uc *ListUseCase) Execute(ctx context.Context, cmd ListCommand) (usr.ListResult, error) {
	uc.logger.Debug("listing users", "tenant_id", cmd.TenantID, "page", cmd.Page, "size", cmd.Size)

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

	result, err := uc.repo.List(ctx, filter)
	if err != nil {
		uc.logger.Error("failed to list users", "error", err, "tenant_id", cmd.TenantID)
		return usr.ListResult{}, err
	}

	uc.logger.Debug("users listed", "count", len(result.Users), "total", result.Total, "tenant_id", cmd.TenantID)
	return result, nil
}
