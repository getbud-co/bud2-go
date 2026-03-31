package auth

import (
	"context"
	"errors"
	"fmt"
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
}

func NewLoginUseCase(repo user.Repository, issuer *auth.TokenIssuer) *LoginUseCase {
	return &LoginUseCase{
		repo:     repo,
		issuer:   issuer,
		tokenTTL: 7 * 24 * time.Hour, // 7 days
	}
}

func (uc *LoginUseCase) Execute(ctx context.Context, cmd LoginCommand) (*LoginResult, error) {
	// Find user by email (global search, email is unique)
	u, err := uc.repo.GetByEmailForLogin(ctx, cmd.Email)
	if err != nil {
		if errors.Is(err, user.ErrNotFound) {
			// Return generic error to not reveal if email exists
			return nil, ErrInvalidCredentials
		}
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	// Verify password
	if !auth.VerifyPassword(cmd.Password, u.PasswordHash) {
		return nil, ErrInvalidCredentials
	}

	// Check if user is active
	if u.Status != user.StatusActive {
		return nil, ErrUserInactive
	}

	// Generate JWT token
	claims := domain.UserClaims{
		UserID:   domain.UserID(u.ID),
		TenantID: u.TenantID,
		Role:     string(u.Role),
	}

	token, err := uc.issuer.IssueToken(claims, uc.tokenTTL)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	// Don't return password hash
	u.PasswordHash = ""

	return &LoginResult{
		Token: token,
		User:  *u,
	}, nil
}
