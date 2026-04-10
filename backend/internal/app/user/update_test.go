package user

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/dsbraz/bud2/backend/internal/domain/membership"
	usr "github.com/dsbraz/bud2/backend/internal/domain/user"
	"github.com/dsbraz/bud2/backend/internal/test/fixtures"
	"github.com/dsbraz/bud2/backend/internal/test/mocks"
	"github.com/dsbraz/bud2/backend/internal/test/testutil"
)

func TestUpdateUseCase_Execute_Success(t *testing.T) {
	users := new(mocks.UserRepository)
	memberships := new(mocks.MembershipRepository)
	uc := NewUpdateUseCase(users, memberships, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()
	testUser := fixtures.NewUser()
	testMembership := fixtures.NewMembership()
	updatedUser := fixtures.NewUser()
	updatedUser.Name = "Updated Name"
	updatedMembership := fixtures.NewMembership()
	updatedMembership.Role = membership.RoleManager

	users.On("GetByID", mock.Anything, testUser.ID).Return(testUser, nil)
	memberships.On("GetByOrganizationAndUser", mock.Anything, tenantID.UUID(), testUser.ID).Return(testMembership, nil)
	users.On("Update", mock.Anything, mock.Anything).Return(updatedUser, nil)
	memberships.On("Update", mock.Anything, mock.Anything).Return(updatedMembership, nil)

	result, err := uc.Execute(context.Background(), UpdateCommand{
		OrganizationID: tenantID, ID: testUser.ID, Name: "Updated Name", Email: testUser.Email, Role: "manager", Status: "active",
	})

	assert.NoError(t, err)
	assert.Equal(t, "Updated Name", result.User.Name)
	assert.Equal(t, membership.RoleManager, result.MembershipRole)
}

func TestUpdateUseCase_Execute_UserNotFound(t *testing.T) {
	users := new(mocks.UserRepository)
	memberships := new(mocks.MembershipRepository)
	uc := NewUpdateUseCase(users, memberships, testutil.NewDiscardLogger())

	userID := uuid.New()
	users.On("GetByID", mock.Anything, userID).Return(nil, usr.ErrNotFound)

	result, err := uc.Execute(context.Background(), UpdateCommand{
		OrganizationID: fixtures.NewTestTenantID(), ID: userID, Name: "Test", Email: "test@example.com", Role: "admin", Status: "active",
	})

	assert.ErrorIs(t, err, usr.ErrNotFound)
	assert.Nil(t, result)
}

func TestUpdateUseCase_Execute_MembershipNotFound(t *testing.T) {
	users := new(mocks.UserRepository)
	memberships := new(mocks.MembershipRepository)
	uc := NewUpdateUseCase(users, memberships, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()
	testUser := fixtures.NewUser()

	users.On("GetByID", mock.Anything, testUser.ID).Return(testUser, nil)
	memberships.On("GetByOrganizationAndUser", mock.Anything, tenantID.UUID(), testUser.ID).Return(nil, membership.ErrNotFound)

	result, err := uc.Execute(context.Background(), UpdateCommand{
		OrganizationID: tenantID, ID: testUser.ID, Name: "Test", Email: testUser.Email, Role: "admin", Status: "active",
	})

	assert.ErrorIs(t, err, membership.ErrNotFound)
	assert.Nil(t, result)
}

func TestUpdateUseCase_Execute_EmailConflict(t *testing.T) {
	users := new(mocks.UserRepository)
	memberships := new(mocks.MembershipRepository)
	uc := NewUpdateUseCase(users, memberships, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()
	testUser := fixtures.NewUser()
	testMembership := fixtures.NewMembership()
	otherUser := fixtures.NewUser()
	otherUser.ID = uuid.New()
	otherUser.Email = "other@example.com"

	users.On("GetByID", mock.Anything, testUser.ID).Return(testUser, nil)
	memberships.On("GetByOrganizationAndUser", mock.Anything, tenantID.UUID(), testUser.ID).Return(testMembership, nil)
	users.On("GetByEmail", mock.Anything, "other@example.com").Return(otherUser, nil)

	result, err := uc.Execute(context.Background(), UpdateCommand{
		OrganizationID: tenantID, ID: testUser.ID, Name: testUser.Name, Email: "other@example.com", Role: "admin", Status: "active",
	})

	assert.ErrorIs(t, err, usr.ErrEmailExists)
	assert.Nil(t, result)
}

func TestUpdateUseCase_Execute_SameEmailNoConflict(t *testing.T) {
	users := new(mocks.UserRepository)
	memberships := new(mocks.MembershipRepository)
	uc := NewUpdateUseCase(users, memberships, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()
	testUser := fixtures.NewUser()
	testMembership := fixtures.NewMembership()

	users.On("GetByID", mock.Anything, testUser.ID).Return(testUser, nil)
	memberships.On("GetByOrganizationAndUser", mock.Anything, tenantID.UUID(), testUser.ID).Return(testMembership, nil)
	users.On("Update", mock.Anything, mock.Anything).Return(testUser, nil)
	memberships.On("Update", mock.Anything, mock.Anything).Return(testMembership, nil)

	result, err := uc.Execute(context.Background(), UpdateCommand{
		OrganizationID: tenantID, ID: testUser.ID, Name: testUser.Name, Email: testUser.Email, Role: "admin", Status: "active",
	})

	assert.NoError(t, err)
	assert.NotNil(t, result)
	users.AssertNotCalled(t, "GetByEmail")
}

func TestUpdateUseCase_Execute_InvalidRole(t *testing.T) {
	users := new(mocks.UserRepository)
	memberships := new(mocks.MembershipRepository)
	uc := NewUpdateUseCase(users, memberships, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()
	testUser := fixtures.NewUser()
	testMembership := fixtures.NewMembership()

	users.On("GetByID", mock.Anything, testUser.ID).Return(testUser, nil)
	memberships.On("GetByOrganizationAndUser", mock.Anything, tenantID.UUID(), testUser.ID).Return(testMembership, nil)

	result, err := uc.Execute(context.Background(), UpdateCommand{
		OrganizationID: tenantID, ID: testUser.ID, Name: testUser.Name, Email: testUser.Email, Role: "invalid", Status: "active",
	})

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestUpdateUseCase_Execute_InvalidMembershipStatus(t *testing.T) {
	users := new(mocks.UserRepository)
	memberships := new(mocks.MembershipRepository)
	uc := NewUpdateUseCase(users, memberships, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()
	testUser := fixtures.NewUser()
	testMembership := fixtures.NewMembership()

	users.On("GetByID", mock.Anything, testUser.ID).Return(testUser, nil)
	memberships.On("GetByOrganizationAndUser", mock.Anything, tenantID.UUID(), testUser.ID).Return(testMembership, nil)

	result, err := uc.Execute(context.Background(), UpdateCommand{
		OrganizationID: tenantID, ID: testUser.ID, Name: testUser.Name, Email: testUser.Email, Role: "admin", Status: "invalid",
	})

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestUpdateUseCase_Execute_GetByEmailError(t *testing.T) {
	users := new(mocks.UserRepository)
	memberships := new(mocks.MembershipRepository)
	uc := NewUpdateUseCase(users, memberships, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()
	testUser := fixtures.NewUser()
	testMembership := fixtures.NewMembership()

	users.On("GetByID", mock.Anything, testUser.ID).Return(testUser, nil)
	memberships.On("GetByOrganizationAndUser", mock.Anything, tenantID.UUID(), testUser.ID).Return(testMembership, nil)
	users.On("GetByEmail", mock.Anything, "new@example.com").Return(nil, errors.New("db error"))

	result, err := uc.Execute(context.Background(), UpdateCommand{
		OrganizationID: tenantID, ID: testUser.ID, Name: testUser.Name, Email: "new@example.com", Role: "admin", Status: "active",
	})

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestUpdateUseCase_Execute_UserUpdateError(t *testing.T) {
	users := new(mocks.UserRepository)
	memberships := new(mocks.MembershipRepository)
	uc := NewUpdateUseCase(users, memberships, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()
	testUser := fixtures.NewUser()
	testMembership := fixtures.NewMembership()

	users.On("GetByID", mock.Anything, testUser.ID).Return(testUser, nil)
	memberships.On("GetByOrganizationAndUser", mock.Anything, tenantID.UUID(), testUser.ID).Return(testMembership, nil)
	users.On("Update", mock.Anything, mock.Anything).Return(nil, errors.New("db error"))

	result, err := uc.Execute(context.Background(), UpdateCommand{
		OrganizationID: tenantID, ID: testUser.ID, Name: testUser.Name, Email: testUser.Email, Role: "admin", Status: "active",
	})

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestUpdateUseCase_Execute_MembershipUpdateError(t *testing.T) {
	users := new(mocks.UserRepository)
	memberships := new(mocks.MembershipRepository)
	uc := NewUpdateUseCase(users, memberships, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()
	testUser := fixtures.NewUser()
	testMembership := fixtures.NewMembership()

	users.On("GetByID", mock.Anything, testUser.ID).Return(testUser, nil)
	memberships.On("GetByOrganizationAndUser", mock.Anything, tenantID.UUID(), testUser.ID).Return(testMembership, nil)
	users.On("Update", mock.Anything, mock.Anything).Return(testUser, nil)
	memberships.On("Update", mock.Anything, mock.Anything).Return(nil, errors.New("db error"))

	result, err := uc.Execute(context.Background(), UpdateCommand{
		OrganizationID: tenantID, ID: testUser.ID, Name: testUser.Name, Email: testUser.Email, Role: "admin", Status: "active",
	})

	assert.Error(t, err)
	assert.Nil(t, result)
}
