package user

import (
	"context"
	"log/slog"

	"github.com/google/uuid"

	"github.com/getbud-co/bud2/backend/internal/domain"
	"github.com/getbud-co/bud2/backend/internal/domain/membership"
	usr "github.com/getbud-co/bud2/backend/internal/domain/user"
)

type GetUseCase struct {
	users       usr.Repository
	memberships membership.Repository
	logger      *slog.Logger
}

func NewGetUseCase(users usr.Repository, memberships membership.Repository, logger *slog.Logger) *GetUseCase {
	return &GetUseCase{users: users, memberships: memberships, logger: logger}
}

func (uc *GetUseCase) Execute(ctx context.Context, organizationID domain.TenantID, id uuid.UUID) (*Member, error) {
	m, err := uc.memberships.GetByOrganizationAndUser(ctx, organizationID.UUID(), id)
	if err != nil {
		return nil, err
	}
	u, err := uc.users.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &Member{User: *u, OrganizationID: m.OrganizationID, MembershipRole: m.Role, MembershipStatus: m.Status}, nil
}
