package user

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/dsbraz/bud2/backend/internal/domain"
	"github.com/dsbraz/bud2/backend/internal/domain/membership"
	"github.com/dsbraz/bud2/backend/internal/domain/organization"
	usr "github.com/dsbraz/bud2/backend/internal/domain/user"
	"github.com/dsbraz/bud2/backend/internal/test/fixtures"
	"github.com/dsbraz/bud2/backend/internal/test/mocks"
	"github.com/dsbraz/bud2/backend/internal/test/testutil"
)

func TestCreateUseCase_Execute_Success_NewUser(t *testing.T) {
	users := new(mocks.UserRepository)
	memberships := new(mocks.MembershipRepository)
	organizations := new(mocks.OrganizationRepository)
	hasher := new(mocks.PasswordHasher)
	uc := NewCreateUseCase(users, memberships, organizations, hasher, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()
	testOrg := fixtures.NewOrganization()

	users.On("GetByEmail", mock.Anything, "new@example.com").Return(nil, usr.ErrNotFound)
	organizations.On("GetByDomain", mock.Anything, "example.com").Return(testOrg, nil)
	hasher.On("Hash", "password123").Return("hashed", nil)
	users.On("Create", mock.Anything, mock.Anything).Return(fixtures.NewUserWithEmail("new@example.com"), nil)
	memberships.On("GetByOrganizationAndUser", mock.Anything, tenantID.UUID(), mock.Anything).Return(nil, membership.ErrNotFound)
	memberships.On("Create", mock.Anything, mock.Anything).Return(fixtures.NewMembership(), nil)

	result, err := uc.Execute(context.Background(), CreateCommand{
		OrganizationID: tenantID, Name: "New User", Email: "new@example.com", Password: "password123", Role: "admin",
	})

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "new@example.com", result.User.Email)
}

func TestCreateUseCase_Execute_Success_ExistingUser(t *testing.T) {
	users := new(mocks.UserRepository)
	memberships := new(mocks.MembershipRepository)
	organizations := new(mocks.OrganizationRepository)
	hasher := new(mocks.PasswordHasher)
	uc := NewCreateUseCase(users, memberships, organizations, hasher, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()
	existingUser := fixtures.NewUser()

	users.On("GetByEmail", mock.Anything, existingUser.Email).Return(existingUser, nil)
	memberships.On("GetByOrganizationAndUser", mock.Anything, tenantID.UUID(), existingUser.ID).Return(nil, membership.ErrNotFound)
	memberships.On("Create", mock.Anything, mock.Anything).Return(fixtures.NewMembership(), nil)

	result, err := uc.Execute(context.Background(), CreateCommand{
		OrganizationID: tenantID, Name: existingUser.Name, Email: existingUser.Email, Password: "password123", Role: "admin",
	})

	assert.NoError(t, err)
	assert.NotNil(t, result)
	organizations.AssertNotCalled(t, "GetByDomain")
	hasher.AssertNotCalled(t, "Hash")
	users.AssertNotCalled(t, "Create")
}

func TestCreateUseCase_Execute_MembershipAlreadyExists(t *testing.T) {
	users := new(mocks.UserRepository)
	memberships := new(mocks.MembershipRepository)
	organizations := new(mocks.OrganizationRepository)
	hasher := new(mocks.PasswordHasher)
	uc := NewCreateUseCase(users, memberships, organizations, hasher, testutil.NewDiscardLogger())

	tenantID := fixtures.NewTestTenantID()
	existingUser := fixtures.NewUser()

	users.On("GetByEmail", mock.Anything, existingUser.Email).Return(existingUser, nil)
	memberships.On("GetByOrganizationAndUser", mock.Anything, tenantID.UUID(), existingUser.ID).Return(fixtures.NewMembership(), nil)

	result, err := uc.Execute(context.Background(), CreateCommand{
		OrganizationID: tenantID, Name: existingUser.Name, Email: existingUser.Email, Password: "password123", Role: "admin",
	})

	assert.ErrorIs(t, err, membership.ErrAlreadyExists)
	assert.Nil(t, result)
}

func TestCreateUseCase_Execute_InvalidRole(t *testing.T) {
	users := new(mocks.UserRepository)
	memberships := new(mocks.MembershipRepository)
	organizations := new(mocks.OrganizationRepository)
	hasher := new(mocks.PasswordHasher)
	uc := NewCreateUseCase(users, memberships, organizations, hasher, testutil.NewDiscardLogger())

	result, err := uc.Execute(context.Background(), CreateCommand{
		OrganizationID: fixtures.NewTestTenantID(), Name: "Test", Email: "test@example.com", Password: "password123", Role: "invalid",
	})

	assert.ErrorIs(t, err, domain.ErrValidation)
	assert.Nil(t, result)
}

func TestCreateUseCase_Execute_EmailDomainOrgNotFound(t *testing.T) {
	users := new(mocks.UserRepository)
	memberships := new(mocks.MembershipRepository)
	organizations := new(mocks.OrganizationRepository)
	hasher := new(mocks.PasswordHasher)
	uc := NewCreateUseCase(users, memberships, organizations, hasher, testutil.NewDiscardLogger())

	users.On("GetByEmail", mock.Anything, "test@unknown.com").Return(nil, usr.ErrNotFound)
	organizations.On("GetByDomain", mock.Anything, "unknown.com").Return(nil, organization.ErrNotFound)

	result, err := uc.Execute(context.Background(), CreateCommand{
		OrganizationID: fixtures.NewTestTenantID(), Name: "Test", Email: "test@unknown.com", Password: "password123", Role: "admin",
	})

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestCreateUseCase_Execute_PasswordHashError(t *testing.T) {
	users := new(mocks.UserRepository)
	memberships := new(mocks.MembershipRepository)
	organizations := new(mocks.OrganizationRepository)
	hasher := new(mocks.PasswordHasher)
	uc := NewCreateUseCase(users, memberships, organizations, hasher, testutil.NewDiscardLogger())

	users.On("GetByEmail", mock.Anything, "test@example.com").Return(nil, usr.ErrNotFound)
	organizations.On("GetByDomain", mock.Anything, "example.com").Return(fixtures.NewOrganization(), nil)
	hasher.On("Hash", "password123").Return("", errors.New("hash error"))

	result, err := uc.Execute(context.Background(), CreateCommand{
		OrganizationID: fixtures.NewTestTenantID(), Name: "Test", Email: "test@example.com", Password: "password123", Role: "admin",
	})

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestCreateUseCase_Execute_UserCreateError(t *testing.T) {
	users := new(mocks.UserRepository)
	memberships := new(mocks.MembershipRepository)
	organizations := new(mocks.OrganizationRepository)
	hasher := new(mocks.PasswordHasher)
	uc := NewCreateUseCase(users, memberships, organizations, hasher, testutil.NewDiscardLogger())

	users.On("GetByEmail", mock.Anything, "test@example.com").Return(nil, usr.ErrNotFound)
	organizations.On("GetByDomain", mock.Anything, "example.com").Return(fixtures.NewOrganization(), nil)
	hasher.On("Hash", "password123").Return("hashed", nil)
	users.On("Create", mock.Anything, mock.Anything).Return(nil, errors.New("db error"))

	result, err := uc.Execute(context.Background(), CreateCommand{
		OrganizationID: fixtures.NewTestTenantID(), Name: "Test", Email: "test@example.com", Password: "password123", Role: "admin",
	})

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestCreateUseCase_Execute_GetByEmailError(t *testing.T) {
	users := new(mocks.UserRepository)
	memberships := new(mocks.MembershipRepository)
	organizations := new(mocks.OrganizationRepository)
	hasher := new(mocks.PasswordHasher)
	uc := NewCreateUseCase(users, memberships, organizations, hasher, testutil.NewDiscardLogger())

	users.On("GetByEmail", mock.Anything, "test@example.com").Return(nil, errors.New("db error"))

	result, err := uc.Execute(context.Background(), CreateCommand{
		OrganizationID: fixtures.NewTestTenantID(), Name: "Test", Email: "test@example.com", Password: "password123", Role: "admin",
	})

	assert.Error(t, err)
	assert.Nil(t, result)
}
