package auth

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/dsbraz/bud2/backend/internal/domain"
	"github.com/dsbraz/bud2/backend/internal/domain/user"
	"github.com/dsbraz/bud2/backend/internal/infra/auth"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserInactive       = errors.New("user account is inactive")
)

type LoginCommand struct {
	Email    string
	Password string
}

type LoginResult struct {
	Token string
	User  user.User
}

type LoginUseCase struct {
	repo     user.Repository
	issuer   *auth.TokenIssuer
	tokenTTL time.Duration
	logger   *slog.Logger
}

func NewLoginUseCase(repo user.Repository, issuer *auth.TokenIssuer, logger *slog.Logger) *LoginUseCase {
	return &LoginUseCase{
		repo:     repo,
		issuer:   issuer,
		tokenTTL: 7 * 24 * time.Hour,
		logger:   logger,
	}
}

func (uc *LoginUseCase) Execute(ctx context.Context, cmd LoginCommand) (*LoginResult, error) {
	uc.logger.Debug("login attempt", "email", cmd.Email)

	u, err := uc.repo.GetByEmailForLogin(ctx, cmd.Email)
	if err != nil {
		if errors.Is(err, user.ErrNotFound) {
			uc.logger.Warn("login failed - invalid credentials", "email", cmd.Email)
			return nil, ErrInvalidCredentials
		}
		uc.logger.Error("failed to find user", "error", err, "email", cmd.Email)
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	if !auth.VerifyPassword(cmd.Password, u.PasswordHash) {
		uc.logger.Warn("login failed - invalid credentials", "email", cmd.Email)
		return nil, ErrInvalidCredentials
	}

	if u.Status != user.StatusActive {
		uc.logger.Warn("login failed - user inactive", "email", cmd.Email, "user_id", u.ID, "status", u.Status)
		return nil, ErrUserInactive
	}

	claims := domain.UserClaims{
		UserID:   domain.UserID(u.ID),
		TenantID: u.TenantID,
		Role:     string(u.Role),
	}

	token, err := uc.issuer.IssueToken(claims, uc.tokenTTL)
	if err != nil {
		uc.logger.Error("failed to generate token", "error", err, "user_id", u.ID)
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	uc.logger.Info("login successful", "user_id", u.ID, "tenant_id", u.TenantID, "email", cmd.Email)

	u.PasswordHash = ""

	return &LoginResult{
		Token: token,
		User:  *u,
	}, nil
}
