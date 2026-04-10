package auth

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	domainauth "github.com/dsbraz/bud2/backend/internal/domain/auth"
	"github.com/dsbraz/bud2/backend/internal/domain/membership"
	"github.com/dsbraz/bud2/backend/internal/domain/organization"
	"github.com/dsbraz/bud2/backend/internal/domain/user"

	"github.com/dsbraz/bud2/backend/internal/domain"
)

type GetSessionUseCase struct {
	support authSupport
}

func NewGetSessionUseCase(users user.Repository, memberships membership.Repository, organizations organization.Repository, issuer tokenIssuer, passwordHasher domainauth.PasswordHasher, logger *slog.Logger) *GetSessionUseCase {
	return &GetSessionUseCase{support: newAuthSupport(users, memberships, organizations, issuer, passwordHasher, logger, 7*24*time.Hour)}
}

func (uc *GetSessionUseCase) Execute(ctx context.Context, claims domain.UserClaims) (*Session, error) {
	u, err := uc.support.users.GetByID(ctx, claims.UserID.UUID())
	if err != nil {
		return nil, fmt.Errorf("failed to load user: %w", err)
	}
	session, err := uc.support.loadSession(ctx, u)
	if err != nil {
		return nil, err
	}
	if claims.HasActiveOrganization {
		if active := findAccessibleOrganization(session.Organizations, claims.ActiveOrganizationID.UUID()); active != nil {
			session.ActiveOrganization = active
		}
	}
	session.User = setUserPasswordHash(&session.User, "")
	return session, nil
}
