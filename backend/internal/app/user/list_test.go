package user

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/getbud-co/bud2/backend/internal/domain/membership"
	"github.com/getbud-co/bud2/backend/internal/test/fixtures"
	"github.com/getbud-co/bud2/backend/internal/test/mocks"
	"github.com/getbud-co/bud2/backend/internal/test/testutil"
)

func TestListUseCase_Execute_Success(t *testing.T) {
	users := new(mocks.UserRepository)
	memberships := new(mocks.MembershipRepository)
	uc := NewListUseCase(users, memberships, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()
	testUser := fixtures.NewUser()
	testMembership := fixtures.NewMembership()

	memberships.On("ListByOrganization", mock.Anything, membership.ListByOrganizationFilter{
		OrganizationID: tenantID.UUID(), Page: 1, Size: 20,
	}).Return(membership.ListResult{Memberships: []membership.Membership{*testMembership}, Total: 1}, nil)
	users.On("GetByID", mock.Anything, testMembership.UserID).Return(testUser, nil)

	result, err := uc.Execute(context.Background(), ListCommand{
		OrganizationID: tenantID, Page: 1, Size: 20,
	})

	assert.NoError(t, err)
	assert.Len(t, result.Members, 1)
	assert.Equal(t, int64(1), result.Total)
}

func TestListUseCase_Execute_WithStatusFilter(t *testing.T) {
	users := new(mocks.UserRepository)
	memberships := new(mocks.MembershipRepository)
	uc := NewListUseCase(users, memberships, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()
	status := "active"
	membershipStatus := membership.StatusActive

	memberships.On("ListByOrganization", mock.Anything, membership.ListByOrganizationFilter{
		OrganizationID: tenantID.UUID(), Status: &membershipStatus, Page: 1, Size: 20,
	}).Return(membership.ListResult{}, nil)

	_, err := uc.Execute(context.Background(), ListCommand{
		OrganizationID: tenantID, Status: &status, Page: 1, Size: 20,
	})

	assert.NoError(t, err)
	memberships.AssertExpectations(t)
}

func TestListUseCase_Execute_DefaultPagination(t *testing.T) {
	users := new(mocks.UserRepository)
	memberships := new(mocks.MembershipRepository)
	uc := NewListUseCase(users, memberships, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()

	memberships.On("ListByOrganization", mock.Anything, membership.ListByOrganizationFilter{
		OrganizationID: tenantID.UUID(), Page: 1, Size: 20,
	}).Return(membership.ListResult{}, nil)

	_, err := uc.Execute(context.Background(), ListCommand{OrganizationID: tenantID})

	assert.NoError(t, err)
	memberships.AssertExpectations(t)
}

func TestListUseCase_Execute_MaxSizeLimit(t *testing.T) {
	users := new(mocks.UserRepository)
	memberships := new(mocks.MembershipRepository)
	uc := NewListUseCase(users, memberships, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()

	memberships.On("ListByOrganization", mock.Anything, membership.ListByOrganizationFilter{
		OrganizationID: tenantID.UUID(), Page: 1, Size: 100,
	}).Return(membership.ListResult{}, nil)

	_, err := uc.Execute(context.Background(), ListCommand{
		OrganizationID: tenantID, Page: 1, Size: 500,
	})

	assert.NoError(t, err)
	memberships.AssertExpectations(t)
}

func TestListUseCase_Execute_MembershipListError(t *testing.T) {
	users := new(mocks.UserRepository)
	memberships := new(mocks.MembershipRepository)
	uc := NewListUseCase(users, memberships, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()

	memberships.On("ListByOrganization", mock.Anything, mock.Anything).Return(membership.ListResult{}, errors.New("db error"))

	result, err := uc.Execute(context.Background(), ListCommand{
		OrganizationID: tenantID, Page: 1, Size: 20,
	})

	assert.Error(t, err)
	assert.Empty(t, result.Members)
}

func TestListUseCase_Execute_UserFetchError(t *testing.T) {
	users := new(mocks.UserRepository)
	memberships := new(mocks.MembershipRepository)
	uc := NewListUseCase(users, memberships, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()
	testMembership := fixtures.NewMembership()

	memberships.On("ListByOrganization", mock.Anything, mock.Anything).Return(membership.ListResult{
		Memberships: []membership.Membership{*testMembership}, Total: 1,
	}, nil)
	users.On("GetByID", mock.Anything, testMembership.UserID).Return(nil, errors.New("db error"))

	result, err := uc.Execute(context.Background(), ListCommand{
		OrganizationID: tenantID, Page: 1, Size: 20,
	})

	assert.Error(t, err)
	assert.Empty(t, result.Members)
}
