package user

import (
	"context"
	"errors"
	"log/slog"

	"github.com/google/uuid"

	"github.com/dsbraz/bud2/backend/internal/domain"
	"github.com/dsbraz/bud2/backend/internal/domain/membership"
	usr "github.com/dsbraz/bud2/backend/internal/domain/user"
)

type UpdateCommand struct {
	OrganizationID domain.TenantID
	ID             uuid.UUID
	Name           string
	Email          string
	Role           string
	Status         string
}

type UpdateUseCase struct {
	users       usr.Repository
	memberships membership.Repository
	logger      *slog.Logger
}

func NewUpdateUseCase(users usr.Repository, memberships membership.Repository, logger *slog.Logger) *UpdateUseCase {
	return &UpdateUseCase{users: users, memberships: memberships, logger: logger}
}

func (uc *UpdateUseCase) Execute(ctx context.Context, cmd UpdateCommand) (*Member, error) {
	u, err := uc.users.GetByID(ctx, cmd.ID)
	if err != nil {
		return nil, err
	}
	m, err := uc.memberships.GetByOrganizationAndUser(ctx, cmd.OrganizationID.UUID(), cmd.ID)
	if err != nil {
		return nil, err
	}

	if u.Email != cmd.Email {
		other, err := uc.users.GetByEmail(ctx, cmd.Email)
		if err == nil && other.ID != cmd.ID {
			return nil, usr.ErrEmailExists
		}
		if err != nil && !errors.Is(err, usr.ErrNotFound) {
			return nil, err
		}
	}

	u.Name = cmd.Name
	u.Email = cmd.Email
	m.Role = membership.Role(cmd.Role)
	m.Status = membership.Status(cmd.Status)

	if err := u.Validate(); err != nil {
		return nil, err
	}
	if err := m.Validate(); err != nil {
		return nil, err
	}

	updatedUser, err := uc.users.Update(ctx, u)
	if err != nil {
		return nil, err
	}
	updatedMembership, err := uc.memberships.Update(ctx, m)
	if err != nil {
		return nil, err
	}
	return &Member{User: *updatedUser, OrganizationID: updatedMembership.OrganizationID, MembershipRole: updatedMembership.Role, MembershipStatus: updatedMembership.Status}, nil
}
