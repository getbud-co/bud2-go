package organization

import (
	"context"
	"errors"
	"testing"

	"github.com/dsbraz/bud2/backend/internal/domain"
	"github.com/dsbraz/bud2/backend/internal/domain/organization"
	"github.com/dsbraz/bud2/backend/internal/test/fixtures"
	"github.com/dsbraz/bud2/backend/internal/test/mocks"
	"github.com/dsbraz/bud2/backend/internal/test/testutil"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateUseCase_Execute_Success(t *testing.T) {
	mockRepo := new(mocks.OrganizationRepository)
	uc := NewUpdateUseCase(mockRepo, testutil.NewDiscardLogger())

	existingOrg := fixtures.NewOrganization()
	id := existingOrg.ID

	cmd := UpdateCommand{
		ID:     id,
		Name:   "Updated Name",
		Slug:   existingOrg.Slug, // Same slug, no conflict check needed
		Status: "active",
	}

	updatedOrg := fixtures.NewOrganization()
	updatedOrg.Name = cmd.Name

	mockRepo.On("GetByID", mock.Anything, id).Return(existingOrg, nil)
	mockRepo.On("Update", mock.Anything, mock.MatchedBy(func(o *organization.Organization) bool {
		return o.ID == id && o.Name == cmd.Name
	})).Return(updatedOrg, nil)

	result, err := uc.Execute(context.Background(), cmd)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, cmd.Name, result.Name)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUseCase_Execute_NotFound(t *testing.T) {
	mockRepo := new(mocks.OrganizationRepository)
	uc := NewUpdateUseCase(mockRepo, testutil.NewDiscardLogger())

	id := uuid.New()
	cmd := UpdateCommand{
		ID:     id,
		Name:   "Updated Name",
		Slug:   "updated-slug",
		Status: "active",
	}

	mockRepo.On("GetByID", mock.Anything, id).Return(nil, organization.ErrNotFound)

	result, err := uc.Execute(context.Background(), cmd)

	assert.ErrorIs(t, err, organization.ErrNotFound)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUseCase_Execute_SlugConflict(t *testing.T) {
	mockRepo := new(mocks.OrganizationRepository)
	uc := NewUpdateUseCase(mockRepo, testutil.NewDiscardLogger())

	existingOrg := fixtures.NewOrganization()
	id := existingOrg.ID

	otherOrg := fixtures.NewOrganization()
	otherOrg.ID = uuid.New()
	otherOrg.Slug = "new-slug"

	cmd := UpdateCommand{
		ID:     id,
		Name:   "Updated Name",
		Slug:   "new-slug", // Different from existing
		Status: "active",
	}

	mockRepo.On("GetByID", mock.Anything, id).Return(existingOrg, nil)
	mockRepo.On("GetBySlug", mock.Anything, cmd.Slug).Return(otherOrg, nil)

	result, err := uc.Execute(context.Background(), cmd)

	assert.ErrorIs(t, err, organization.ErrSlugExists)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUseCase_Execute_SameSlugNoConflict(t *testing.T) {
	mockRepo := new(mocks.OrganizationRepository)
	uc := NewUpdateUseCase(mockRepo, testutil.NewDiscardLogger())

	existingOrg := fixtures.NewOrganization()
	id := existingOrg.ID

	cmd := UpdateCommand{
		ID:     id,
		Name:   "Updated Name",
		Slug:   existingOrg.Slug, // Same slug
		Status: "active",
	}

	updatedOrg := fixtures.NewOrganization()
	updatedOrg.Name = cmd.Name

	mockRepo.On("GetByID", mock.Anything, id).Return(existingOrg, nil)
	// GetBySlug should NOT be called when slug hasn't changed
	mockRepo.On("Update", mock.Anything, mock.AnythingOfType("*organization.Organization")).Return(updatedOrg, nil)

	result, err := uc.Execute(context.Background(), cmd)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	mockRepo.AssertNotCalled(t, "GetBySlug")
	mockRepo.AssertExpectations(t)
}

func TestUpdateUseCase_Execute_InvalidStatus(t *testing.T) {
	mockRepo := new(mocks.OrganizationRepository)
	uc := NewUpdateUseCase(mockRepo, testutil.NewDiscardLogger())

	existingOrg := fixtures.NewOrganization()
	id := existingOrg.ID

	cmd := UpdateCommand{
		ID:     id,
		Name:   "Updated Name",
		Slug:   existingOrg.Slug,
		Status: "invalid",
	}

	mockRepo.On("GetByID", mock.Anything, id).Return(existingOrg, nil)

	result, err := uc.Execute(context.Background(), cmd)

	assert.ErrorIs(t, err, domain.ErrValidation)
	assert.Nil(t, result)
	mockRepo.AssertNotCalled(t, "Update")
}

func TestUpdateUseCase_Execute_GetBySlugError(t *testing.T) {
	mockRepo := new(mocks.OrganizationRepository)
	uc := NewUpdateUseCase(mockRepo, testutil.NewDiscardLogger())

	existingOrg := fixtures.NewOrganization()
	id := existingOrg.ID

	cmd := UpdateCommand{
		ID:     id,
		Name:   "Updated Name",
		Slug:   "new-slug",
		Status: "active",
	}

	dbError := errors.New("database error")
	mockRepo.On("GetByID", mock.Anything, id).Return(existingOrg, nil)
	mockRepo.On("GetBySlug", mock.Anything, cmd.Slug).Return(nil, dbError)

	result, err := uc.Execute(context.Background(), cmd)

	assert.ErrorIs(t, err, dbError)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUseCase_Execute_UpdateError(t *testing.T) {
	mockRepo := new(mocks.OrganizationRepository)
	uc := NewUpdateUseCase(mockRepo, testutil.NewDiscardLogger())

	existingOrg := fixtures.NewOrganization()
	id := existingOrg.ID

	cmd := UpdateCommand{
		ID:     id,
		Name:   "Updated Name",
		Slug:   existingOrg.Slug,
		Status: "active",
	}

	dbError := errors.New("update failed")
	mockRepo.On("GetByID", mock.Anything, id).Return(existingOrg, nil)
	mockRepo.On("Update", mock.Anything, mock.AnythingOfType("*organization.Organization")).Return(nil, dbError)

	result, err := uc.Execute(context.Background(), cmd)

	assert.ErrorIs(t, err, dbError)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}
