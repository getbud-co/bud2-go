package auth

import (
	"context"
	"errors"
	"testing"

	"github.com/dsbraz/bud2/backend/internal/domain/user"
	"github.com/dsbraz/bud2/backend/internal/infra/auth"
	"github.com/dsbraz/bud2/backend/internal/test/fixtures"
	"github.com/dsbraz/bud2/backend/internal/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLoginUseCase_Execute_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	tokenIssuer := auth.NewTokenIssuer("test-secret")
	uc := NewLoginUseCase(mockRepo, tokenIssuer)

	cmd := LoginCommand{
		Email:    "admin@example.com",
		Password: "password123",
	}

	// Create user with hashed password
	testUser := fixtures.NewUser()
	testUser.Email = cmd.Email
	testUser.Status = user.StatusActive
	hash, _ := auth.HashPassword(cmd.Password)
	testUser.PasswordHash = hash

	mockRepo.On("GetByEmailForLogin", mock.Anything, cmd.Email).Return(testUser, nil)

	result, err := uc.Execute(context.Background(), cmd)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotEmpty(t, result.Token)
	assert.Equal(t, testUser.ID, result.User.ID)
	assert.Empty(t, result.User.PasswordHash) // Password hash should be cleared
}

func TestLoginUseCase_Execute_InvalidCredentials(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	tokenIssuer := auth.NewTokenIssuer("test-secret")
	uc := NewLoginUseCase(mockRepo, tokenIssuer)

	cmd := LoginCommand{
		Email:    "admin@example.com",
		Password: "wrongpassword",
	}

	testUser := fixtures.NewUser()
	testUser.Email = cmd.Email
	testUser.Status = user.StatusActive
	hash, _ := auth.HashPassword("correctpassword")
	testUser.PasswordHash = hash

	mockRepo.On("GetByEmailForLogin", mock.Anything, cmd.Email).Return(testUser, nil)

	result, err := uc.Execute(context.Background(), cmd)

	assert.ErrorIs(t, err, ErrInvalidCredentials)
	assert.Nil(t, result)
}

func TestLoginUseCase_Execute_UserNotFound(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	tokenIssuer := auth.NewTokenIssuer("test-secret")
	uc := NewLoginUseCase(mockRepo, tokenIssuer)

	cmd := LoginCommand{
		Email:    "nonexistent@example.com",
		Password: "password123",
	}

	mockRepo.On("GetByEmailForLogin", mock.Anything, cmd.Email).Return(nil, user.ErrNotFound)

	result, err := uc.Execute(context.Background(), cmd)

	assert.ErrorIs(t, err, ErrInvalidCredentials)
	assert.Nil(t, result)
}

func TestLoginUseCase_Execute_UserInactive(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	tokenIssuer := auth.NewTokenIssuer("test-secret")
	uc := NewLoginUseCase(mockRepo, tokenIssuer)

	cmd := LoginCommand{
		Email:    "admin@example.com",
		Password: "password123",
	}

	testUser := fixtures.NewUser()
	testUser.Email = cmd.Email
	testUser.Status = user.StatusInactive // User is inactive
	hash, _ := auth.HashPassword(cmd.Password)
	testUser.PasswordHash = hash

	mockRepo.On("GetByEmailForLogin", mock.Anything, cmd.Email).Return(testUser, nil)

	result, err := uc.Execute(context.Background(), cmd)

	assert.ErrorIs(t, err, ErrUserInactive)
	assert.Nil(t, result)
}

func TestLoginUseCase_Execute_RepositoryError(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	tokenIssuer := auth.NewTokenIssuer("test-secret")
	uc := NewLoginUseCase(mockRepo, tokenIssuer)

	cmd := LoginCommand{
		Email:    "admin@example.com",
		Password: "password123",
	}

	mockRepo.On("GetByEmailForLogin", mock.Anything, cmd.Email).Return(nil, errors.New("db error"))

	result, err := uc.Execute(context.Background(), cmd)

	assert.Error(t, err)
	assert.Nil(t, result)
}
