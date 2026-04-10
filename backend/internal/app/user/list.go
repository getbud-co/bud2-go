package user

import (
	"context"
	"log/slog"

	"github.com/dsbraz/bud2/backend/internal/domain"
	"github.com/dsbraz/bud2/backend/internal/domain/membership"
	usr "github.com/dsbraz/bud2/backend/internal/domain/user"
)

type ListCommand struct {
	OrganizationID domain.TenantID
	Status         *string
	Page           int
	Size           int
}

type ListUseCase struct {
	users       usr.Repository
	memberships membership.Repository
	logger      *slog.Logger
}

func NewListUseCase(users usr.Repository, memberships membership.Repository, logger *slog.Logger) *ListUseCase {
	return &ListUseCase{users: users, memberships: memberships, logger: logger}
}

func (uc *ListUseCase) Execute(ctx context.Context, cmd ListCommand) (MemberListResult, error) {
	if cmd.Size <= 0 {
		cmd.Size = 20
	}
	if cmd.Size > 100 {
		cmd.Size = 100
	}
	if cmd.Page <= 0 {
		cmd.Page = 1
	}

	filter := membership.ListByOrganizationFilter{OrganizationID: cmd.OrganizationID.UUID(), Page: cmd.Page, Size: cmd.Size}
	if cmd.Status != nil {
		status := membership.Status(*cmd.Status)
		filter.Status = &status
	}

	result, err := uc.memberships.ListByOrganization(ctx, filter)
	if err != nil {
		return MemberListResult{}, err
	}

	members := make([]Member, 0, len(result.Memberships))
	for _, membershipItem := range result.Memberships {
		u, err := uc.users.GetByID(ctx, membershipItem.UserID)
		if err != nil {
			return MemberListResult{}, err
		}
		members = append(members, Member{User: *u, OrganizationID: membershipItem.OrganizationID, MembershipRole: membershipItem.Role, MembershipStatus: membershipItem.Status})
	}

	return MemberListResult{Members: members, Total: result.Total}, nil
}
