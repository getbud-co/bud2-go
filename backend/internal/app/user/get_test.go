package user

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/getbud-co/bud2/backend/internal/domain/membership"
	usr "github.com/getbud-co/bud2/backend/internal/domain/user"
	"github.com/getbud-co/bud2/backend/internal/test/fixtures"
	"github.com/getbud-co/bud2/backend/internal/test/mocks"
	"github.com/getbud-co/bud2/backend/internal/test/testutil"
)

func TestGetUseCase_Execute_Success(t *testing.T) {
	users := new(mocks.UserRepository)
	memberships := new(mocks.MembershipRepository)
	uc := NewGetUseCase(users, memberships, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()
	testUser := fixtures.NewUser()
	testMembership := fixtures.NewMembership()

	memberships.On("GetByOrganizationAndUser", mock.Anything, tenantID.UUID(), testUser.ID).Return(testMembership, nil)
	users.On("GetByID", mock.Anything, testUser.ID).Return(testUser, nil)

	result, err := uc.Execute(context.Background(), tenantID, testUser.ID)

	assert.NoError(t, err)
	assert.Equal(t, testUser.Email, result.User.Email)
	assert.Equal(t, testMembership.Role, result.MembershipRole)
}

func TestGetUseCase_Execute_MembershipNotFound(t *testing.T) {
	users := new(mocks.UserRepository)
	memberships := new(mocks.MembershipRepository)
	uc := NewGetUseCase(users, memberships, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()
	userID := uuid.New()

	memberships.On("GetByOrganizationAndUser", mock.Anything, tenantID.UUID(), userID).Return(nil, membership.ErrNotFound)

	result, err := uc.Execute(context.Background(), tenantID, userID)

	assert.ErrorIs(t, err, membership.ErrNotFound)
	assert.Nil(t, result)
	users.AssertNotCalled(t, "GetByID")
}

func TestGetUseCase_Execute_UserNotFound(t *testing.T) {
	users := new(mocks.UserRepository)
	memberships := new(mocks.MembershipRepository)
	uc := NewGetUseCase(users, memberships, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()
	userID := uuid.New()
	testMembership := fixtures.NewMembership()
	testMembership.UserID = userID

	memberships.On("GetByOrganizationAndUser", mock.Anything, tenantID.UUID(), userID).Return(testMembership, nil)
	users.On("GetByID", mock.Anything, userID).Return(nil, usr.ErrNotFound)

	result, err := uc.Execute(context.Background(), tenantID, userID)

	assert.ErrorIs(t, err, usr.ErrNotFound)
	assert.Nil(t, result)
}
