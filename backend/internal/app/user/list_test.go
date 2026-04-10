package user

import (
	"context"
	"testing"

	"github.com/dsbraz/bud2/backend/internal/domain/user"
	"github.com/dsbraz/bud2/backend/internal/test/fixtures"
	"github.com/dsbraz/bud2/backend/internal/test/mocks"
	"github.com/dsbraz/bud2/backend/internal/test/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListUseCase_Execute_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	uc := NewListUseCase(mockRepo, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()
	users := fixtures.NewUserList(3)
	expectedResult := user.ListResult{
		Users: users,
		Total: 3,
	}

	mockRepo.On("List", mock.Anything, mock.MatchedBy(func(filter user.ListFilter) bool {
		return filter.TenantID == tenantID && filter.Page == 1 && filter.Size == 20
	})).Return(expectedResult, nil)

	cmd := ListCommand{
		TenantID: tenantID,
		Page:     1,
		Size:     20,
	}

	result, err := uc.Execute(context.Background(), cmd)

	assert.NoError(t, err)
	assert.Equal(t, 3, len(result.Users))
	assert.Equal(t, int64(3), result.Total)
	mockRepo.AssertExpectations(t)
}

func TestListUseCase_Execute_WithStatusFilter(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	uc := NewListUseCase(mockRepo, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()
	status := "active"
	users := fixtures.NewUserList(2)
	expectedResult := user.ListResult{
		Users: users,
		Total: 2,
	}

	mockRepo.On("List", mock.Anything, mock.MatchedBy(func(filter user.ListFilter) bool {
		return filter.TenantID == tenantID && filter.Status != nil && *filter.Status == user.StatusActive
	})).Return(expectedResult, nil)

	cmd := ListCommand{
		TenantID: tenantID,
		Status:   &status,
		Page:     1,
		Size:     20,
	}

	result, err := uc.Execute(context.Background(), cmd)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(result.Users))
	mockRepo.AssertExpectations(t)
}

func TestListUseCase_Execute_WithSearchFilter(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	uc := NewListUseCase(mockRepo, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()
	search := "john"
	users := fixtures.NewUserList(1)
	expectedResult := user.ListResult{
		Users: users,
		Total: 1,
	}

	mockRepo.On("List", mock.Anything, mock.MatchedBy(func(filter user.ListFilter) bool {
		return filter.TenantID == tenantID && filter.Search != nil && *filter.Search == "john"
	})).Return(expectedResult, nil)

	cmd := ListCommand{
		TenantID: tenantID,
		Search:   &search,
		Page:     1,
		Size:     20,
	}

	result, err := uc.Execute(context.Background(), cmd)

	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.Users))
	mockRepo.AssertExpectations(t)
}

func TestListUseCase_Execute_DefaultPagination(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	uc := NewListUseCase(mockRepo, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()
	expectedResult := user.ListResult{
		Users: []user.User{},
		Total: 0,
	}

	mockRepo.On("List", mock.Anything, mock.MatchedBy(func(filter user.ListFilter) bool {
		return filter.Page == 1 && filter.Size == 20
	})).Return(expectedResult, nil)

	cmd := ListCommand{
		TenantID: tenantID,
		Page:     0, // Invalid, should default to 1
		Size:     0, // Invalid, should default to 20
	}

	result, err := uc.Execute(context.Background(), cmd)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestListUseCase_Execute_MaxSizeLimit(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	uc := NewListUseCase(mockRepo, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()
	expectedResult := user.ListResult{
		Users: []user.User{},
		Total: 0,
	}

	mockRepo.On("List", mock.Anything, mock.MatchedBy(func(filter user.ListFilter) bool {
		return filter.Size == 100 // Should be capped at 100
	})).Return(expectedResult, nil)

	cmd := ListCommand{
		TenantID: tenantID,
		Page:     1,
		Size:     200, // Exceeds max, should be capped
	}

	result, err := uc.Execute(context.Background(), cmd)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	mockRepo.AssertExpectations(t)
}
