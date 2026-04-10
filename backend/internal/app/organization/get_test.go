package organization

import (
	"context"
	"testing"

	"github.com/dsbraz/bud2/backend/internal/domain/organization"
	"github.com/dsbraz/bud2/backend/internal/test/fixtures"
	"github.com/dsbraz/bud2/backend/internal/test/mocks"
	"github.com/dsbraz/bud2/backend/internal/test/testutil"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUseCase_Execute_Success(t *testing.T) {
	mockRepo := new(mocks.OrganizationRepository)
	uc := NewGetUseCase(mockRepo, testutil.NewDiscardLogger())

	expectedOrg := fixtures.NewOrganization()
	id := expectedOrg.ID

	mockRepo.On("GetByID", mock.Anything, id).Return(expectedOrg, nil)

	result, err := uc.Execute(context.Background(), id)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedOrg.ID, result.ID)
	mockRepo.AssertExpectations(t)
}

func TestGetUseCase_Execute_NotFound(t *testing.T) {
	mockRepo := new(mocks.OrganizationRepository)
	uc := NewGetUseCase(mockRepo, testutil.NewDiscardLogger())

	id := uuid.New()

	mockRepo.On("GetByID", mock.Anything, id).Return(nil, organization.ErrNotFound)

	result, err := uc.Execute(context.Background(), id)

	assert.ErrorIs(t, err, organization.ErrNotFound)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}
