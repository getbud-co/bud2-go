package user

import (
	"context"
	"errors"
	"log/slog"

	"github.com/google/uuid"

	apptx "github.com/dsbraz/bud2/backend/internal/app/tx"
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
	txm         apptx.Manager
	logger      *slog.Logger
}

func NewUpdateUseCase(users usr.Repository, memberships membership.Repository, txm apptx.Manager, logger *slog.Logger) *UpdateUseCase {
	return &UpdateUseCase{users: users, memberships: memberships, txm: txm, logger: logger}
}

func (uc *UpdateUseCase) Execute(ctx context.Context, cmd UpdateCommand) (*Member, error) {
	var updatedUser *usr.User
	var updatedMembership *membership.Membership
	err := uc.txm.WithTx(ctx, func(repos apptx.Repositories) error {
		u, txErr := repos.Users().GetByID(ctx, cmd.ID)
		if txErr != nil {
			return txErr
		}
		m, txErr := repos.Memberships().GetByOrganizationAndUser(ctx, cmd.OrganizationID.UUID(), cmd.ID)
		if txErr != nil {
			return txErr
		}

		if u.Email != cmd.Email {
			other, lookupErr := repos.Users().GetByEmail(ctx, cmd.Email)
			if lookupErr == nil && other.ID != cmd.ID {
				return usr.ErrEmailExists
			}
			if lookupErr != nil && !errors.Is(lookupErr, usr.ErrNotFound) {
				return lookupErr
			}
		}

		u.Name = cmd.Name
		u.Email = cmd.Email
		m.Role = membership.Role(cmd.Role)
		m.Status = membership.Status(cmd.Status)

		if txErr = u.Validate(); txErr != nil {
			return txErr
		}
		if txErr = m.Validate(); txErr != nil {
			return txErr
		}

		updatedUser, txErr = repos.Users().Update(ctx, u)
		if txErr != nil {
			return txErr
		}
		updatedMembership, txErr = repos.Memberships().Update(ctx, m)
		return txErr
	})
	if err != nil {
		return nil, err
	}
	return &Member{User: *updatedUser, OrganizationID: updatedMembership.OrganizationID, MembershipRole: updatedMembership.Role, MembershipStatus: updatedMembership.Status}, nil
}
