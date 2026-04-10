package bootstrap

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
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
	logger   *slog.Logger
}

func NewUseCase(orgRepo organization.Repository, txm TxManager, issuer TokenIssuer, logger *slog.Logger) *UseCase {
	return &UseCase{
		orgRepo:  orgRepo,
		txm:      txm,
		issuer:   issuer,
		tokenTTL: 24 * time.Hour,
		logger:   logger,
	}
}

func (uc *UseCase) Execute(ctx context.Context, cmd Command) (*Result, error) {
	uc.logger.Debug("starting bootstrap", "organization_name", cmd.OrganizationName, "admin_email", cmd.AdminEmail)

	count, err := uc.orgRepo.CountAll(ctx)
	if err != nil {
		uc.logger.Error("failed to count organizations", "error", err)
		return nil, err
	}
	if count > 0 {
		uc.logger.Warn("bootstrap already completed", "organization_count", count)
		return nil, ErrAlreadyBootstrapped
	}

	var createdOrg *organization.Organization
	var createdAdmin *user.User

	passwordHash, err := auth.HashPassword(cmd.AdminPassword)
	if err != nil {
		uc.logger.Error("failed to hash admin password", "error", err)
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
			uc.logger.Error("failed to create organization in transaction", "error", txErr)
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
			uc.logger.Warn("admin validation failed", "error", txErr, "email", cmd.AdminEmail)
			return txErr
		}

		createdAdmin, txErr = userRepo.Create(ctx, admin)
		return txErr
	})
	if err != nil {
		uc.logger.Error("bootstrap transaction failed", "error", err)
		return nil, err
	}

	token, err := uc.issuer.IssueToken(domain.UserClaims{
		UserID:   domain.UserID(createdAdmin.ID),
		TenantID: createdAdmin.TenantID,
		Role:     string(createdAdmin.Role),
	}, uc.tokenTTL)
	if err != nil {
		uc.logger.Error("failed to generate bootstrap token", "error", err, "user_id", createdAdmin.ID)
		return nil, err
	}

	uc.logger.Info("bootstrap completed", "organization_id", createdOrg.ID, "admin_id", createdAdmin.ID, "organization_name", cmd.OrganizationName)

	return &Result{
		Organization: *createdOrg,
		Admin:        *createdAdmin,
		AccessToken:  token,
	}, nil
}
