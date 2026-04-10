package user

import (
	"context"
	"errors"
	"testing"

	"github.com/dsbraz/bud2/backend/internal/domain"
	"github.com/dsbraz/bud2/backend/internal/domain/user"
	"github.com/dsbraz/bud2/backend/internal/test/fixtures"
	"github.com/dsbraz/bud2/backend/internal/test/mocks"
	"github.com/dsbraz/bud2/backend/internal/test/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUseCase_Execute_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	uc := NewCreateUseCase(mockRepo, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()
	cmd := CreateCommand{
		TenantID: tenantID,
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password123",
		Role:     "admin",
	}

	expectedUser := fixtures.NewUser()

	mockRepo.On("GetByEmail", mock.Anything, tenantID, cmd.Email).Return(nil, user.ErrNotFound)
	mockRepo.On("Create", mock.Anything, mock.MatchedBy(func(u *user.User) bool {
		return u.Name == cmd.Name && u.Email == cmd.Email && u.Role == user.RoleAdmin && u.PasswordHash != ""
	})).Return(expectedUser, nil)

	result, err := uc.Execute(context.Background(), cmd)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedUser.ID, result.ID)
	mockRepo.AssertExpectations(t)
}

func TestCreateUseCase_Execute_EmailAlreadyExists(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	uc := NewCreateUseCase(mockRepo, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()
	cmd := CreateCommand{
		TenantID: tenantID,
		Name:     "Test User",
		Email:    "existing@example.com",
		Password: "password123",
		Role:     "admin",
	}

	existingUser := fixtures.NewUser()
	existingUser.Email = cmd.Email

	mockRepo.On("GetByEmail", mock.Anything, tenantID, cmd.Email).Return(existingUser, nil)

	result, err := uc.Execute(context.Background(), cmd)

	assert.ErrorIs(t, err, user.ErrEmailExists)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestCreateUseCase_Execute_InvalidRole(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	uc := NewCreateUseCase(mockRepo, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()
	cmd := CreateCommand{
		TenantID: tenantID,
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password123",
		Role:     "invalid-role",
	}

	result, err := uc.Execute(context.Background(), cmd)

	assert.ErrorIs(t, err, domain.ErrValidation)
	assert.Nil(t, result)
	mockRepo.AssertNotCalled(t, "GetByEmail")
	mockRepo.AssertNotCalled(t, "Create")
}

func TestCreateUseCase_Execute_MissingName(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	uc := NewCreateUseCase(mockRepo, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()
	cmd := CreateCommand{
		TenantID: tenantID,
		Name:     "",
		Email:    "test@example.com",
		Password: "password123",
		Role:     "admin",
	}

	result, err := uc.Execute(context.Background(), cmd)

	assert.ErrorIs(t, err, domain.ErrValidation)
	assert.Nil(t, result)
	mockRepo.AssertNotCalled(t, "GetByEmail")
	mockRepo.AssertNotCalled(t, "Create")
}

func TestCreateUseCase_Execute_MissingEmail(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	uc := NewCreateUseCase(mockRepo, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()
	cmd := CreateCommand{
		TenantID: tenantID,
		Name:     "Test User",
		Email:    "",
		Password: "password123",
		Role:     "admin",
	}

	result, err := uc.Execute(context.Background(), cmd)

	assert.ErrorIs(t, err, domain.ErrValidation)
	assert.Nil(t, result)
	mockRepo.AssertNotCalled(t, "GetByEmail")
	mockRepo.AssertNotCalled(t, "Create")
}

func TestCreateUseCase_Execute_MissingPassword(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	uc := NewCreateUseCase(mockRepo, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()
	cmd := CreateCommand{
		TenantID: tenantID,
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "",
		Role:     "admin",
	}

	result, err := uc.Execute(context.Background(), cmd)

	assert.ErrorIs(t, err, domain.ErrValidation)
	assert.Nil(t, result)
	mockRepo.AssertNotCalled(t, "GetByEmail")
	mockRepo.AssertNotCalled(t, "Create")
}

func TestCreateUseCase_Execute_GetByEmailError(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	uc := NewCreateUseCase(mockRepo, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()
	cmd := CreateCommand{
		TenantID: tenantID,
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password123",
		Role:     "admin",
	}

	dbError := errors.New("database connection failed")
	mockRepo.On("GetByEmail", mock.Anything, tenantID, cmd.Email).Return(nil, dbError)

	result, err := uc.Execute(context.Background(), cmd)

	assert.ErrorIs(t, err, dbError)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
	mockRepo.AssertNotCalled(t, "Create")
}

func TestCreateUseCase_Execute_CreateError(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	uc := NewCreateUseCase(mockRepo, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()
	cmd := CreateCommand{
		TenantID: tenantID,
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password123",
		Role:     "admin",
	}

	dbError := errors.New("insert failed")
	mockRepo.On("GetByEmail", mock.Anything, tenantID, cmd.Email).Return(nil, user.ErrNotFound)
	mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*user.User")).Return(nil, dbError)

	result, err := uc.Execute(context.Background(), cmd)

	assert.ErrorIs(t, err, dbError)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}
