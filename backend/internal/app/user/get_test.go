package user

import (
	"context"
	"testing"

	"github.com/dsbraz/bud2/backend/internal/domain/user"
	"github.com/dsbraz/bud2/backend/internal/test/fixtures"
	"github.com/dsbraz/bud2/backend/internal/test/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUseCase_Execute_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	uc := NewGetUseCase(mockRepo)

	tenantID := fixtures.NewTestTenantID()
	expectedUser := fixtures.NewUser()
	id := expectedUser.ID

	mockRepo.On("GetByID", mock.Anything, tenantID, id).Return(expectedUser, nil)

	result, err := uc.Execute(context.Background(), tenantID, id)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedUser.ID, result.ID)
	mockRepo.AssertExpectations(t)
}

func TestGetUseCase_Execute_NotFound(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	uc := NewGetUseCase(mockRepo)

	tenantID := fixtures.NewTestTenantID()
	id := uuid.New()

	mockRepo.On("GetByID", mock.Anything, tenantID, id).Return(nil, user.ErrNotFound)

	result, err := uc.Execute(context.Background(), tenantID, id)

	assert.ErrorIs(t, err, user.ErrNotFound)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}
