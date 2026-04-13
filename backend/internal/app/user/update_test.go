package user

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	apptx "github.com/getbud-co/bud2/backend/internal/app/tx"
	"github.com/getbud-co/bud2/backend/internal/domain/membership"
	"github.com/getbud-co/bud2/backend/internal/domain/organization"
	usr "github.com/getbud-co/bud2/backend/internal/domain/user"
	"github.com/getbud-co/bud2/backend/internal/test/fixtures"
	"github.com/getbud-co/bud2/backend/internal/test/mocks"
	"github.com/getbud-co/bud2/backend/internal/test/testutil"
)

type updateTestTxRepos struct {
	userRepo       usr.Repository
	membershipRepo membership.Repository
}

func (r updateTestTxRepos) Organizations() organization.Repository { return nil }
func (r updateTestTxRepos) Users() usr.Repository                  { return r.userRepo }
func (r updateTestTxRepos) Memberships() membership.Repository     { return r.membershipRepo }

type updateTestTxManager struct {
	repos     apptx.Repositories
	called    bool
	returnErr error
}

func (m *updateTestTxManager) WithTx(ctx context.Context, fn func(repos apptx.Repositories) error) error {
	m.called = true
	if m.returnErr != nil {
		return m.returnErr
	}
	return fn(m.repos)
}

func TestUpdateUseCase_Execute_Success(t *testing.T) {
	users := new(mocks.UserRepository)
	memberships := new(mocks.MembershipRepository)
	txUsers := new(mocks.UserRepository)
	txMemberships := new(mocks.MembershipRepository)
	txm := &updateTestTxManager{repos: updateTestTxRepos{userRepo: txUsers, membershipRepo: txMemberships}}
	uc := NewUpdateUseCase(users, memberships, txm, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()
	testUser := fixtures.NewUser()
	testMembership := fixtures.NewMembership()
	updatedUser := fixtures.NewUser()
	updatedUser.Name = "Updated Name"
	updatedMembership := fixtures.NewMembership()
	updatedMembership.Role = membership.RoleManager

	txUsers.On("GetByID", mock.Anything, testUser.ID).Return(testUser, nil)
	txMemberships.On("GetByOrganizationAndUser", mock.Anything, tenantID.UUID(), testUser.ID).Return(testMembership, nil)
	txUsers.On("Update", mock.Anything, mock.Anything).Return(updatedUser, nil)
	txMemberships.On("Update", mock.Anything, mock.Anything).Return(updatedMembership, nil)

	result, err := uc.Execute(context.Background(), UpdateCommand{
		OrganizationID: tenantID, ID: testUser.ID, Name: "Updated Name", Email: testUser.Email, Role: "manager", Status: "active",
	})

	assert.NoError(t, err)
	assert.Equal(t, "Updated Name", result.User.Name)
	assert.Equal(t, membership.RoleManager, result.MembershipRole)
	assert.True(t, txm.called)
}

func TestUpdateUseCase_Execute_UserNotFound(t *testing.T) {
	users := new(mocks.UserRepository)
	memberships := new(mocks.MembershipRepository)
	txm := &updateTestTxManager{returnErr: usr.ErrNotFound}
	uc := NewUpdateUseCase(users, memberships, txm, testutil.NewDiscardLogger())

	userID := uuid.New()

	result, err := uc.Execute(context.Background(), UpdateCommand{
		OrganizationID: fixtures.NewTestTenantID(), ID: userID, Name: "Test", Email: "test@example.com", Role: "admin", Status: "active",
	})

	assert.ErrorIs(t, err, usr.ErrNotFound)
	assert.Nil(t, result)
}

func TestUpdateUseCase_Execute_MembershipNotFound(t *testing.T) {
	users := new(mocks.UserRepository)
	memberships := new(mocks.MembershipRepository)
	txUsers := new(mocks.UserRepository)
	txMemberships := new(mocks.MembershipRepository)
	txm := &updateTestTxManager{repos: updateTestTxRepos{userRepo: txUsers, membershipRepo: txMemberships}}
	uc := NewUpdateUseCase(users, memberships, txm, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()
	testUser := fixtures.NewUser()

	txUsers.On("GetByID", mock.Anything, testUser.ID).Return(testUser, nil)
	txMemberships.On("GetByOrganizationAndUser", mock.Anything, tenantID.UUID(), testUser.ID).Return(nil, membership.ErrNotFound)

	result, err := uc.Execute(context.Background(), UpdateCommand{
		OrganizationID: tenantID, ID: testUser.ID, Name: "Test", Email: testUser.Email, Role: "admin", Status: "active",
	})

	assert.ErrorIs(t, err, membership.ErrNotFound)
	assert.Nil(t, result)
}

func TestUpdateUseCase_Execute_EmailConflict(t *testing.T) {
	users := new(mocks.UserRepository)
	memberships := new(mocks.MembershipRepository)
	txUsers := new(mocks.UserRepository)
	txMemberships := new(mocks.MembershipRepository)
	txm := &updateTestTxManager{repos: updateTestTxRepos{userRepo: txUsers, membershipRepo: txMemberships}}
	uc := NewUpdateUseCase(users, memberships, txm, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()
	testUser := fixtures.NewUser()
	testMembership := fixtures.NewMembership()
	otherUser := fixtures.NewUser()
	otherUser.ID = uuid.New()
	otherUser.Email = "other@example.com"

	txUsers.On("GetByID", mock.Anything, testUser.ID).Return(testUser, nil)
	txMemberships.On("GetByOrganizationAndUser", mock.Anything, tenantID.UUID(), testUser.ID).Return(testMembership, nil)
	txUsers.On("GetByEmail", mock.Anything, "other@example.com").Return(otherUser, nil)

	result, err := uc.Execute(context.Background(), UpdateCommand{
		OrganizationID: tenantID, ID: testUser.ID, Name: testUser.Name, Email: "other@example.com", Role: "admin", Status: "active",
	})

	assert.ErrorIs(t, err, usr.ErrEmailExists)
	assert.Nil(t, result)
}

func TestUpdateUseCase_Execute_SameEmailNoConflict(t *testing.T) {
	users := new(mocks.UserRepository)
	memberships := new(mocks.MembershipRepository)
	txUsers := new(mocks.UserRepository)
	txMemberships := new(mocks.MembershipRepository)
	txm := &updateTestTxManager{repos: updateTestTxRepos{userRepo: txUsers, membershipRepo: txMemberships}}
	uc := NewUpdateUseCase(users, memberships, txm, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()
	testUser := fixtures.NewUser()
	testMembership := fixtures.NewMembership()

	txUsers.On("GetByID", mock.Anything, testUser.ID).Return(testUser, nil)
	txMemberships.On("GetByOrganizationAndUser", mock.Anything, tenantID.UUID(), testUser.ID).Return(testMembership, nil)
	txUsers.On("Update", mock.Anything, mock.Anything).Return(testUser, nil)
	txMemberships.On("Update", mock.Anything, mock.Anything).Return(testMembership, nil)

	result, err := uc.Execute(context.Background(), UpdateCommand{
		OrganizationID: tenantID, ID: testUser.ID, Name: testUser.Name, Email: testUser.Email, Role: "admin", Status: "active",
	})

	assert.NoError(t, err)
	assert.NotNil(t, result)
	txUsers.AssertNotCalled(t, "GetByEmail")
}

func TestUpdateUseCase_Execute_InvalidRole(t *testing.T) {
	users := new(mocks.UserRepository)
	memberships := new(mocks.MembershipRepository)
	txUsers := new(mocks.UserRepository)
	txMemberships := new(mocks.MembershipRepository)
	txm := &updateTestTxManager{repos: updateTestTxRepos{userRepo: txUsers, membershipRepo: txMemberships}}
	uc := NewUpdateUseCase(users, memberships, txm, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()
	testUser := fixtures.NewUser()
	testMembership := fixtures.NewMembership()

	txUsers.On("GetByID", mock.Anything, testUser.ID).Return(testUser, nil)
	txMemberships.On("GetByOrganizationAndUser", mock.Anything, tenantID.UUID(), testUser.ID).Return(testMembership, nil)

	result, err := uc.Execute(context.Background(), UpdateCommand{
		OrganizationID: tenantID, ID: testUser.ID, Name: testUser.Name, Email: testUser.Email, Role: "invalid", Status: "active",
	})

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestUpdateUseCase_Execute_InvalidMembershipStatus(t *testing.T) {
	users := new(mocks.UserRepository)
	memberships := new(mocks.MembershipRepository)
	txUsers := new(mocks.UserRepository)
	txMemberships := new(mocks.MembershipRepository)
	txm := &updateTestTxManager{repos: updateTestTxRepos{userRepo: txUsers, membershipRepo: txMemberships}}
	uc := NewUpdateUseCase(users, memberships, txm, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()
	testUser := fixtures.NewUser()
	testMembership := fixtures.NewMembership()

	txUsers.On("GetByID", mock.Anything, testUser.ID).Return(testUser, nil)
	txMemberships.On("GetByOrganizationAndUser", mock.Anything, tenantID.UUID(), testUser.ID).Return(testMembership, nil)

	result, err := uc.Execute(context.Background(), UpdateCommand{
		OrganizationID: tenantID, ID: testUser.ID, Name: testUser.Name, Email: testUser.Email, Role: "admin", Status: "invalid",
	})

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestUpdateUseCase_Execute_GetByEmailError(t *testing.T) {
	users := new(mocks.UserRepository)
	memberships := new(mocks.MembershipRepository)
	txUsers := new(mocks.UserRepository)
	txMemberships := new(mocks.MembershipRepository)
	txm := &updateTestTxManager{repos: updateTestTxRepos{userRepo: txUsers, membershipRepo: txMemberships}}
	uc := NewUpdateUseCase(users, memberships, txm, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()
	testUser := fixtures.NewUser()
	testMembership := fixtures.NewMembership()

	txUsers.On("GetByID", mock.Anything, testUser.ID).Return(testUser, nil)
	txMemberships.On("GetByOrganizationAndUser", mock.Anything, tenantID.UUID(), testUser.ID).Return(testMembership, nil)
	txUsers.On("GetByEmail", mock.Anything, "new@example.com").Return(nil, errors.New("db error"))

	result, err := uc.Execute(context.Background(), UpdateCommand{
		OrganizationID: tenantID, ID: testUser.ID, Name: testUser.Name, Email: "new@example.com", Role: "admin", Status: "active",
	})

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestUpdateUseCase_Execute_UserUpdateError(t *testing.T) {
	users := new(mocks.UserRepository)
	memberships := new(mocks.MembershipRepository)
	txUsers := new(mocks.UserRepository)
	txMemberships := new(mocks.MembershipRepository)
	txm := &updateTestTxManager{repos: updateTestTxRepos{userRepo: txUsers, membershipRepo: txMemberships}}
	uc := NewUpdateUseCase(users, memberships, txm, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()
	testUser := fixtures.NewUser()
	testMembership := fixtures.NewMembership()

	txUsers.On("GetByID", mock.Anything, testUser.ID).Return(testUser, nil)
	txMemberships.On("GetByOrganizationAndUser", mock.Anything, tenantID.UUID(), testUser.ID).Return(testMembership, nil)
	txUsers.On("Update", mock.Anything, mock.Anything).Return(nil, errors.New("db error"))

	result, err := uc.Execute(context.Background(), UpdateCommand{
		OrganizationID: tenantID, ID: testUser.ID, Name: testUser.Name, Email: testUser.Email, Role: "admin", Status: "active",
	})

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestUpdateUseCase_Execute_MembershipUpdateError(t *testing.T) {
	users := new(mocks.UserRepository)
	memberships := new(mocks.MembershipRepository)
	txUsers := new(mocks.UserRepository)
	txMemberships := new(mocks.MembershipRepository)
	txm := &updateTestTxManager{repos: updateTestTxRepos{userRepo: txUsers, membershipRepo: txMemberships}}
	uc := NewUpdateUseCase(users, memberships, txm, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()
	testUser := fixtures.NewUser()
	testMembership := fixtures.NewMembership()

	txUsers.On("GetByID", mock.Anything, testUser.ID).Return(testUser, nil)
	txMemberships.On("GetByOrganizationAndUser", mock.Anything, tenantID.UUID(), testUser.ID).Return(testMembership, nil)
	txUsers.On("Update", mock.Anything, mock.Anything).Return(testUser, nil)
	txMemberships.On("Update", mock.Anything, mock.Anything).Return(nil, errors.New("db error"))

	result, err := uc.Execute(context.Background(), UpdateCommand{
		OrganizationID: tenantID, ID: testUser.ID, Name: testUser.Name, Email: testUser.Email, Role: "admin", Status: "active",
	})

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestUpdateUseCase_Execute_TransactionError(t *testing.T) {
	users := new(mocks.UserRepository)
	memberships := new(mocks.MembershipRepository)
	txm := &updateTestTxManager{returnErr: errors.New("tx error")}
	uc := NewUpdateUseCase(users, memberships, txm, testutil.NewDiscardLogger())

	result, err := uc.Execute(context.Background(), UpdateCommand{
		OrganizationID: fixtures.NewTestTenantID(), ID: uuid.New(), Name: "Test", Email: "test@example.com", Role: "admin", Status: "active",
	})

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.True(t, txm.called)
}
