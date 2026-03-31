package organization

import (
	"context"
	"testing"

	"github.com/dsbraz/bud2/backend/internal/domain/organization"
	"github.com/dsbraz/bud2/backend/internal/test/fixtures"
	"github.com/dsbraz/bud2/backend/internal/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListUseCase_Execute_Success(t *testing.T) {
	mockRepo := new(mocks.OrganizationRepository)
	uc := NewListUseCase(mockRepo)

	orgs := fixtures.NewOrganizationList(3)
	expectedResult := organization.ListResult{
		Organizations: orgs,
		Total:         3,
	}

	mockRepo.On("List", mock.Anything, mock.MatchedBy(func(filter organization.ListFilter) bool {
		return filter.Page == 1 && filter.Size == 20
	})).Return(expectedResult, nil)

	cmd := ListCommand{
		Page: 1,
		Size: 20,
	}

	result, err := uc.Execute(context.Background(), cmd)

	assert.NoError(t, err)
	assert.Equal(t, 3, len(result.Organizations))
	assert.Equal(t, int64(3), result.Total)
	mockRepo.AssertExpectations(t)
}

func TestListUseCase_Execute_WithStatusFilter(t *testing.T) {
	mockRepo := new(mocks.OrganizationRepository)
	uc := NewListUseCase(mockRepo)

	status := "active"
	orgs := fixtures.NewOrganizationList(2)
	expectedResult := organization.ListResult{
		Organizations: orgs,
		Total:         2,
	}

	mockRepo.On("List", mock.Anything, mock.MatchedBy(func(filter organization.ListFilter) bool {
		return filter.Status != nil && *filter.Status == organization.StatusActive
	})).Return(expectedResult, nil)

	cmd := ListCommand{
		Status: &status,
		Page:   1,
		Size:   20,
	}

	result, err := uc.Execute(context.Background(), cmd)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(result.Organizations))
	mockRepo.AssertExpectations(t)
}

func TestListUseCase_Execute_DefaultPagination(t *testing.T) {
	mockRepo := new(mocks.OrganizationRepository)
	uc := NewListUseCase(mockRepo)

	expectedResult := organization.ListResult{
		Organizations: []organization.Organization{},
		Total:         0,
	}

	mockRepo.On("List", mock.Anything, mock.MatchedBy(func(filter organization.ListFilter) bool {
		return filter.Page == 1 && filter.Size == 20
	})).Return(expectedResult, nil)

	cmd := ListCommand{
		Page: 0, // Invalid, should default to 1
		Size: 0, // Invalid, should default to 20
	}

	result, err := uc.Execute(context.Background(), cmd)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestListUseCase_Execute_MaxSizeLimit(t *testing.T) {
	mockRepo := new(mocks.OrganizationRepository)
	uc := NewListUseCase(mockRepo)

	expectedResult := organization.ListResult{
		Organizations: []organization.Organization{},
		Total:         0,
	}

	mockRepo.On("List", mock.Anything, mock.MatchedBy(func(filter organization.ListFilter) bool {
		return filter.Size == 100 // Should be capped at 100
	})).Return(expectedResult, nil)

	cmd := ListCommand{
		Page: 1,
		Size: 200, // Exceeds max, should be capped
	}

	result, err := uc.Execute(context.Background(), cmd)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	mockRepo.AssertExpectations(t)
}
