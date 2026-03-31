package bootstrap

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dsbraz/bud2/backend/internal/domain"
	"github.com/dsbraz/bud2/backend/internal/domain/organization"
	"github.com/dsbraz/bud2/backend/internal/domain/user"
	"github.com/dsbraz/bud2/backend/internal/infra/auth"
	"github.com/google/uuid"
)

var ErrAlreadyBootstrapped = errors.New("bootstrap already completed")

type Command struct {
	OrganizationName string
	OrganizationSlug string
	AdminName        string
	AdminEmail       string
	AdminPassword    string
}

type Result struct {
	Organization organization.Organization
	Admin        user.User
	AccessToken  string
}

type TxManager interface {
	WithTx(ctx context.Context, fn func(orgRepo organization.Repository, userRepo user.Repository) error) error
}

type TokenIssuer interface {
	IssueToken(claims domain.UserClaims, ttl time.Duration) (string, error)
}

type UseCase struct {
	orgRepo  organization.Repository
	txm      TxManager
	issuer   TokenIssuer
	tokenTTL time.Duration
}

func NewUseCase(orgRepo organization.Repository, txm TxManager, issuer TokenIssuer) *UseCase {
	return &UseCase{
		orgRepo:  orgRepo,
		txm:      txm,
		issuer:   issuer,
		tokenTTL: 24 * time.Hour,
	}
}

func (uc *UseCase) Execute(ctx context.Context, cmd Command) (*Result, error) {
	count, err := uc.orgRepo.CountAll(ctx)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, ErrAlreadyBootstrapped
	}

	var createdOrg *organization.Organization
	var createdAdmin *user.User

	// Hash admin password
	passwordHash, err := auth.HashPassword(cmd.AdminPassword)
	if err != nil {
		return nil, fmt.Errorf("invalid admin password: %w", err)
	}

	err = uc.txm.WithTx(ctx, func(orgRepo organization.Repository, userRepo user.Repository) error {
		var txErr error
		createdOrg, txErr = orgRepo.Create(ctx, &organization.Organization{
			Name:   cmd.OrganizationName,
			Slug:   cmd.OrganizationSlug,
			Status: organization.StatusActive,
		})
		if txErr != nil {
			return txErr
		}

		admin := &user.User{
			ID:           uuid.New(),
			TenantID:     domain.TenantID(createdOrg.ID),
			Name:         cmd.AdminName,
			Email:        cmd.AdminEmail,
			PasswordHash: passwordHash,
			Role:         user.RoleAdmin,
			Status:       user.StatusActive,
		}
		if txErr = admin.Validate(); txErr != nil {
			return txErr
		}

		createdAdmin, txErr = userRepo.Create(ctx, admin)
		return txErr
	})
	if err != nil {
		return nil, err
	}

	token, err := uc.issuer.IssueToken(domain.UserClaims{
		UserID:   domain.UserID(createdAdmin.ID),
		TenantID: createdAdmin.TenantID,
		Role:     string(createdAdmin.Role),
	}, uc.tokenTTL)
	if err != nil {
		return nil, err
	}

	return &Result{
		Organization: *createdOrg,
		Admin:        *createdAdmin,
		AccessToken:  token,
	}, nil
}
