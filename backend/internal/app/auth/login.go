package auth

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	domainauth "github.com/dsbraz/bud2/backend/internal/domain/auth"
	"github.com/dsbraz/bud2/backend/internal/domain/membership"
	"github.com/dsbraz/bud2/backend/internal/domain/organization"
	"github.com/dsbraz/bud2/backend/internal/domain/user"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserInactive       = errors.New("user account is inactive")
	ErrNoOrganizations    = errors.New("user has no accessible organizations")
)

type LoginCommand struct {
	Email    string
	Password string
}

type LoginResult struct {
	Token   string
	Session Session
}

type LoginUseCase struct {
	support authSupport
}

func NewLoginUseCase(users user.Repository, memberships membership.Repository, organizations organization.Repository, issuer tokenIssuer, passwordHasher domainauth.PasswordHasher, logger *slog.Logger) *LoginUseCase {
	return &LoginUseCase{support: newAuthSupport(users, memberships, organizations, issuer, passwordHasher, logger, 7*24*time.Hour)}
}

func (uc *LoginUseCase) Execute(ctx context.Context, cmd LoginCommand) (*LoginResult, error) {
	uc.support.logger.Debug("login attempt", "email", cmd.Email)

	u, err := uc.support.users.GetByEmail(ctx, cmd.Email)
	if err != nil {
		if errors.Is(err, user.ErrNotFound) {
			uc.support.logger.Warn("login failed - invalid credentials", "email", cmd.Email)
			return nil, ErrInvalidCredentials
		}
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	if !uc.support.passwordHasher.Verify(cmd.Password, u.PasswordHash) {
		uc.support.logger.Warn("login failed - invalid credentials", "email", cmd.Email)
		return nil, ErrInvalidCredentials
	}
	if u.Status != user.StatusActive {
		uc.support.logger.Warn("login failed - user inactive", "email", cmd.Email, "user_id", u.ID)
		return nil, ErrUserInactive
	}

	session, err := uc.support.loadSession(ctx, u)
	if err != nil {
		if errors.Is(err, ErrNoOrganizations) {
			return nil, err
		}
		return nil, fmt.Errorf("failed to load session: %w", err)
	}

	token, err := uc.support.issueTokenForSession(session)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	session.User = setUserPasswordHash(&session.User, "")
	uc.support.logger.Info("login successful", "user_id", session.User.ID, "active_organization", session.ActiveOrganization != nil)
	return &LoginResult{Token: token, Session: *session}, nil
}
