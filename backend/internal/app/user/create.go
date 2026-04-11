package user

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/google/uuid"

	apptx "github.com/dsbraz/bud2/backend/internal/app/tx"
	"github.com/dsbraz/bud2/backend/internal/domain"
	domainauth "github.com/dsbraz/bud2/backend/internal/domain/auth"
	"github.com/dsbraz/bud2/backend/internal/domain/membership"
	"github.com/dsbraz/bud2/backend/internal/domain/organization"
	usr "github.com/dsbraz/bud2/backend/internal/domain/user"
)

type CreateCommand struct {
	OrganizationID domain.TenantID
	Name           string
	Email          string
	Password       string
	Role           string
}

type CreateUseCase struct {
	users          usr.Repository
	memberships    membership.Repository
	organizations  organization.Repository
	txm            apptx.Manager
	passwordHasher domainauth.PasswordHasher
	logger         *slog.Logger
}

func NewCreateUseCase(users usr.Repository, memberships membership.Repository, organizations organization.Repository, txm apptx.Manager, passwordHasher domainauth.PasswordHasher, logger *slog.Logger) *CreateUseCase {
	return &CreateUseCase{users: users, memberships: memberships, organizations: organizations, txm: txm, passwordHasher: passwordHasher, logger: logger}
}

func (uc *CreateUseCase) Execute(ctx context.Context, cmd CreateCommand) (*Member, error) {
	role := membership.Role(cmd.Role)
	if !role.IsValid() {
		return nil, domain.ErrValidation
	}

	existingUser, err := uc.users.GetByEmail(ctx, cmd.Email)
	if err != nil && !errors.Is(err, usr.ErrNotFound) {
		return nil, err
	}

	targetUser := existingUser
	passwordHash := ""
	if errors.Is(err, usr.ErrNotFound) {
		if _, orgErr := uc.organizations.GetByDomain(ctx, emailDomain(cmd.Email)); orgErr != nil {
			return nil, fmt.Errorf("native organization not found for email domain: %w", orgErr)
		}
		var hashErr error
		passwordHash, hashErr = uc.passwordHasher.Hash(cmd.Password)
		if hashErr != nil {
			return nil, fmt.Errorf("invalid password: %w", hashErr)
		}
	}

	var createdMembership *membership.Membership
	err = uc.txm.WithTx(ctx, func(repos apptx.Repositories) error {
		var txErr error
		if targetUser == nil {
			targetUser, txErr = repos.Users().Create(ctx, &usr.User{
				ID:            uuid.New(),
				Name:          cmd.Name,
				Email:         cmd.Email,
				PasswordHash:  passwordHash,
				Status:        usr.StatusActive,
				IsSystemAdmin: false,
			})
			if txErr != nil {
				return txErr
			}
		}

		if _, txErr = repos.Memberships().GetByOrganizationAndUser(ctx, cmd.OrganizationID.UUID(), targetUser.ID); txErr == nil {
			return membership.ErrAlreadyExists
		} else if !errors.Is(txErr, membership.ErrNotFound) {
			return txErr
		}

		createdMembership, txErr = repos.Memberships().Create(ctx, &membership.Membership{
			OrganizationID: cmd.OrganizationID.UUID(),
			UserID:         targetUser.ID,
			Role:           role,
			Status:         membership.StatusActive,
		})
		return txErr
	})
	if err != nil {
		return nil, err
	}

	return &Member{User: *targetUser, OrganizationID: createdMembership.OrganizationID, MembershipRole: createdMembership.Role, MembershipStatus: createdMembership.Status}, nil
}
