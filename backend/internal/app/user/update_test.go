package user

import (
	"context"
	"errors"
	"testing"

	"github.com/dsbraz/bud2/backend/internal/domain"
	"github.com/dsbraz/bud2/backend/internal/domain/user"
	"github.com/dsbraz/bud2/backend/internal/test/fixtures"
	"github.com/dsbraz/bud2/backend/internal/test/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateUseCase_Execute_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	uc := NewUpdateUseCase(mockRepo)

	tenantID := fixtures.NewTestTenantID()
	existingUser := fixtures.NewUser()
	existingUser.TenantID = tenantID
	id := existingUser.ID

	cmd := UpdateCommand{
		TenantID: tenantID,
		ID:       id,
		Name:     "Updated Name",
		Email:    existingUser.Email, // Same email, no conflict check needed
		Role:     "manager",
		Status:   "active",
	}

	updatedUser := fixtures.NewUser()
	updatedUser.TenantID = tenantID
	updatedUser.Name = cmd.Name
	updatedUser.Role = user.RoleManager

	mockRepo.On("GetByID", mock.Anything, tenantID, id).Return(existingUser, nil)
	mockRepo.On("Update", mock.Anything, mock.MatchedBy(func(u *user.User) bool {
		return u.ID == id && u.Name == cmd.Name && u.Role == user.RoleManager
	})).Return(updatedUser, nil)

	result, err := uc.Execute(context.Background(), cmd)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, cmd.Name, result.Name)
	assert.Equal(t, user.RoleManager, result.Role)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUseCase_Execute_NotFound(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	uc := NewUpdateUseCase(mockRepo)

	tenantID := fixtures.NewTestTenantID()
	id := uuid.New()

	cmd := UpdateCommand{
		TenantID: tenantID,
		ID:       id,
		Name:     "Updated Name",
		Email:    "updated@example.com",
		Role:     "admin",
		Status:   "active",
	}

	mockRepo.On("GetByID", mock.Anything, tenantID, id).Return(nil, user.ErrNotFound)

	result, err := uc.Execute(context.Background(), cmd)

	assert.ErrorIs(t, err, user.ErrNotFound)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUseCase_Execute_EmailConflict(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	uc := NewUpdateUseCase(mockRepo)

	tenantID := fixtures.NewTestTenantID()
	existingUser := fixtures.NewUser()
	existingUser.TenantID = tenantID
	id := existingUser.ID

	otherUser := fixtures.NewUser()
	otherUser.TenantID = tenantID
	otherUser.ID = uuid.New()
	otherUser.Email = "new@example.com"

	cmd := UpdateCommand{
		TenantID: tenantID,
		ID:       id,
		Name:     "Updated Name",
		Email:    "new@example.com", // Different from existing
		Role:     "admin",
		Status:   "active",
	}

	mockRepo.On("GetByID", mock.Anything, tenantID, id).Return(existingUser, nil)
	mockRepo.On("GetByEmail", mock.Anything, tenantID, cmd.Email).Return(otherUser, nil)

	result, err := uc.Execute(context.Background(), cmd)

	assert.ErrorIs(t, err, user.ErrEmailExists)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUseCase_Execute_SameEmailNoConflict(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	uc := NewUpdateUseCase(mockRepo)

	tenantID := fixtures.NewTestTenantID()
	existingUser := fixtures.NewUser()
	existingUser.TenantID = tenantID
	id := existingUser.ID

	cmd := UpdateCommand{
		TenantID: tenantID,
		ID:       id,
		Name:     "Updated Name",
		Email:    existingUser.Email, // Same email
		Role:     "manager",
		Status:   "active",
	}

	updatedUser := fixtures.NewUser()
	updatedUser.TenantID = tenantID
	updatedUser.Name = cmd.Name
	updatedUser.Role = user.RoleManager

	mockRepo.On("GetByID", mock.Anything, tenantID, id).Return(existingUser, nil)
	// GetByEmail should NOT be called when email hasn't changed
	mockRepo.On("Update", mock.Anything, mock.AnythingOfType("*user.User")).Return(updatedUser, nil)

	result, err := uc.Execute(context.Background(), cmd)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	mockRepo.AssertNotCalled(t, "GetByEmail")
	mockRepo.AssertExpectations(t)
}

func TestUpdateUseCase_Execute_InvalidRole(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	uc := NewUpdateUseCase(mockRepo)

	tenantID := fixtures.NewTestTenantID()
	existingUser := fixtures.NewUser()
	existingUser.TenantID = tenantID
	id := existingUser.ID

	cmd := UpdateCommand{
		TenantID: tenantID,
		ID:       id,
		Name:     "Updated Name",
		Email:    existingUser.Email,
		Role:     "invalid",
		Status:   "active",
	}

	mockRepo.On("GetByID", mock.Anything, tenantID, id).Return(existingUser, nil)

	result, err := uc.Execute(context.Background(), cmd)

	assert.ErrorIs(t, err, domain.ErrValidation)
	assert.Nil(t, result)
	mockRepo.AssertNotCalled(t, "Update")
}

func TestUpdateUseCase_Execute_InvalidStatus(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	uc := NewUpdateUseCase(mockRepo)

	tenantID := fixtures.NewTestTenantID()
	existingUser := fixtures.NewUser()
	existingUser.TenantID = tenantID
	id := existingUser.ID

	cmd := UpdateCommand{
		TenantID: tenantID,
		ID:       id,
		Name:     "Updated Name",
		Email:    existingUser.Email,
		Role:     "admin",
		Status:   "invalid",
	}

	mockRepo.On("GetByID", mock.Anything, tenantID, id).Return(existingUser, nil)

	result, err := uc.Execute(context.Background(), cmd)

	assert.ErrorIs(t, err, domain.ErrValidation)
	assert.Nil(t, result)
	mockRepo.AssertNotCalled(t, "Update")
}

func TestUpdateUseCase_Execute_GetByEmailError(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	uc := NewUpdateUseCase(mockRepo)

	tenantID := fixtures.NewTestTenantID()
	existingUser := fixtures.NewUser()
	existingUser.TenantID = tenantID
	id := existingUser.ID

	cmd := UpdateCommand{
		TenantID: tenantID,
		ID:       id,
		Name:     "Updated Name",
		Email:    "new@example.com",
		Role:     "admin",
		Status:   "active",
	}

	dbError := errors.New("database error")
	mockRepo.On("GetByID", mock.Anything, tenantID, id).Return(existingUser, nil)
	mockRepo.On("GetByEmail", mock.Anything, tenantID, cmd.Email).Return(nil, dbError)

	result, err := uc.Execute(context.Background(), cmd)

	assert.ErrorIs(t, err, dbError)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUseCase_Execute_UpdateError(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	uc := NewUpdateUseCase(mockRepo)

	tenantID := fixtures.NewTestTenantID()
	existingUser := fixtures.NewUser()
	existingUser.TenantID = tenantID
	id := existingUser.ID

	cmd := UpdateCommand{
		TenantID: tenantID,
		ID:       id,
		Name:     "Updated Name",
		Email:    existingUser.Email,
		Role:     "admin",
		Status:   "active",
	}

	dbError := errors.New("update failed")
	mockRepo.On("GetByID", mock.Anything, tenantID, id).Return(existingUser, nil)
	mockRepo.On("Update", mock.Anything, mock.AnythingOfType("*user.User")).Return(nil, dbError)

	result, err := uc.Execute(context.Background(), cmd)

	assert.ErrorIs(t, err, dbError)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}
