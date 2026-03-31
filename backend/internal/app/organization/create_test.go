package organization

import (
	"context"
	"errors"
	"testing"

	"github.com/dsbraz/bud2/backend/internal/domain"
	"github.com/dsbraz/bud2/backend/internal/domain/organization"
	"github.com/dsbraz/bud2/backend/internal/test/fixtures"
	"github.com/dsbraz/bud2/backend/internal/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUseCase_Execute_Success(t *testing.T) {
	mockRepo := new(mocks.OrganizationRepository)
	uc := NewCreateUseCase(mockRepo)

	cmd := CreateCommand{
		Name:   "Test Organization",
		Slug:   "test-org",
		Status: "active",
	}

	expectedOrg := fixtures.NewOrganization()

	mockRepo.On("GetBySlug", mock.Anything, cmd.Slug).Return(nil, organization.ErrNotFound)
	mockRepo.On("Create", mock.Anything, mock.MatchedBy(func(o *organization.Organization) bool {
		return o.Name == cmd.Name && o.Slug == cmd.Slug && o.Status == organization.StatusActive
	})).Return(expectedOrg, nil)

	result, err := uc.Execute(context.Background(), cmd)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedOrg.ID, result.ID)
	mockRepo.AssertExpectations(t)
}

func TestCreateUseCase_Execute_SlugAlreadyExists(t *testing.T) {
	mockRepo := new(mocks.OrganizationRepository)
	uc := NewCreateUseCase(mockRepo)

	cmd := CreateCommand{
		Name:   "Test Organization",
		Slug:   "existing-slug",
		Status: "active",
	}

	existingOrg := fixtures.NewOrganization()
	existingOrg.Slug = cmd.Slug

	mockRepo.On("GetBySlug", mock.Anything, cmd.Slug).Return(existingOrg, nil)

	result, err := uc.Execute(context.Background(), cmd)

	assert.ErrorIs(t, err, organization.ErrSlugExists)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestCreateUseCase_Execute_InvalidStatus(t *testing.T) {
	mockRepo := new(mocks.OrganizationRepository)
	uc := NewCreateUseCase(mockRepo)

	cmd := CreateCommand{
		Name:   "Test Organization",
		Slug:   "test-org",
		Status: "invalid-status",
	}

	result, err := uc.Execute(context.Background(), cmd)

	assert.ErrorIs(t, err, domain.ErrValidation)
	assert.Nil(t, result)
	mockRepo.AssertNotCalled(t, "GetBySlug")
	mockRepo.AssertNotCalled(t, "Create")
}

func TestCreateUseCase_Execute_MissingName(t *testing.T) {
	mockRepo := new(mocks.OrganizationRepository)
	uc := NewCreateUseCase(mockRepo)

	cmd := CreateCommand{
		Name:   "",
		Slug:   "test-org",
		Status: "active",
	}

	result, err := uc.Execute(context.Background(), cmd)

	assert.ErrorIs(t, err, domain.ErrValidation)
	assert.Nil(t, result)
	mockRepo.AssertNotCalled(t, "GetBySlug")
	mockRepo.AssertNotCalled(t, "Create")
}

func TestCreateUseCase_Execute_MissingSlug(t *testing.T) {
	mockRepo := new(mocks.OrganizationRepository)
	uc := NewCreateUseCase(mockRepo)

	cmd := CreateCommand{
		Name:   "Test Organization",
		Slug:   "",
		Status: "active",
	}

	result, err := uc.Execute(context.Background(), cmd)

	assert.ErrorIs(t, err, domain.ErrValidation)
	assert.Nil(t, result)
	mockRepo.AssertNotCalled(t, "GetBySlug")
	mockRepo.AssertNotCalled(t, "Create")
}

func TestCreateUseCase_Execute_GetBySlugError(t *testing.T) {
	mockRepo := new(mocks.OrganizationRepository)
	uc := NewCreateUseCase(mockRepo)

	cmd := CreateCommand{
		Name:   "Test Organization",
		Slug:   "test-org",
		Status: "active",
	}

	dbError := errors.New("database connection failed")
	mockRepo.On("GetBySlug", mock.Anything, cmd.Slug).Return(nil, dbError)

	result, err := uc.Execute(context.Background(), cmd)

	assert.ErrorIs(t, err, dbError)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
	mockRepo.AssertNotCalled(t, "Create")
}

func TestCreateUseCase_Execute_CreateError(t *testing.T) {
	mockRepo := new(mocks.OrganizationRepository)
	uc := NewCreateUseCase(mockRepo)

	cmd := CreateCommand{
		Name:   "Test Organization",
		Slug:   "test-org",
		Status: "active",
	}

	dbError := errors.New("insert failed")
	mockRepo.On("GetBySlug", mock.Anything, cmd.Slug).Return(nil, organization.ErrNotFound)
	mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*organization.Organization")).Return(nil, dbError)

	result, err := uc.Execute(context.Background(), cmd)

	assert.ErrorIs(t, err, dbError)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestCreateUseCase_Execute_DefaultStatus(t *testing.T) {
	mockRepo := new(mocks.OrganizationRepository)
	uc := NewCreateUseCase(mockRepo)

	cmd := CreateCommand{
		Name: "Test Organization",
		Slug: "test-org",
		// Status not provided - should default to active
	}

	expectedOrg := fixtures.NewOrganization()

	mockRepo.On("GetBySlug", mock.Anything, cmd.Slug).Return(nil, organization.ErrNotFound)
	mockRepo.On("Create", mock.Anything, mock.MatchedBy(func(o *organization.Organization) bool {
		return o.Status == organization.StatusActive
	})).Return(expectedOrg, nil)

	result, err := uc.Execute(context.Background(), cmd)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, organization.StatusActive, result.Status)
	mockRepo.AssertExpectations(t)
}

func TestCreateUseCase_Execute_InactiveStatus(t *testing.T) {
	mockRepo := new(mocks.OrganizationRepository)
	uc := NewCreateUseCase(mockRepo)

	cmd := CreateCommand{
		Name:   "Test Organization",
		Slug:   "test-org",
		Status: "inactive",
	}

	expectedOrg := fixtures.NewInactiveOrganization()

	mockRepo.On("GetBySlug", mock.Anything, cmd.Slug).Return(nil, organization.ErrNotFound)
	mockRepo.On("Create", mock.Anything, mock.MatchedBy(func(o *organization.Organization) bool {
		return o.Status == organization.StatusInactive
	})).Return(expectedOrg, nil)

	result, err := uc.Execute(context.Background(), cmd)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, organization.StatusInactive, result.Status)
	mockRepo.AssertExpectations(t)
}
