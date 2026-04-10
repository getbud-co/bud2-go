package auth

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/dsbraz/bud2/backend/internal/domain/membership"
	"github.com/dsbraz/bud2/backend/internal/domain/organization"
	"github.com/dsbraz/bud2/backend/internal/domain/user"
	infraauth "github.com/dsbraz/bud2/backend/internal/infra/auth"
	"github.com/dsbraz/bud2/backend/internal/test/mocks"
	"github.com/dsbraz/bud2/backend/internal/test/testutil"
)

type mockMembershipRepository struct{ mock.Mock }

func (m *mockMembershipRepository) Create(ctx context.Context, member *membership.Membership) (*membership.Membership, error) {
	args := m.Called(ctx, member)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*membership.Membership), args.Error(1)
}
func (m *mockMembershipRepository) GetByID(ctx context.Context, id uuid.UUID) (*membership.Membership, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*membership.Membership), args.Error(1)
}
func (m *mockMembershipRepository) GetByOrganizationAndUser(ctx context.Context, organizationID, userID uuid.UUID) (*membership.Membership, error) {
	args := m.Called(ctx, organizationID, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*membership.Membership), args.Error(1)
}
func (m *mockMembershipRepository) ListByOrganization(ctx context.Context, filter membership.ListByOrganizationFilter) (membership.ListResult, error) {
	args := m.Called(ctx, filter)
	return args.Get(0).(membership.ListResult), args.Error(1)
}
func (m *mockMembershipRepository) ListByUser(ctx context.Context, filter membership.ListByUserFilter) (membership.ListResult, error) {
	args := m.Called(ctx, filter)
	return args.Get(0).(membership.ListResult), args.Error(1)
}
func (m *mockMembershipRepository) Update(ctx context.Context, member *membership.Membership) (*membership.Membership, error) {
	args := m.Called(ctx, member)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*membership.Membership), args.Error(1)
}

func TestLoginUseCase_Execute_Success(t *testing.T) {
	users := new(mocks.UserRepository)
	memberships := new(mockMembershipRepository)
	organizations := new(mocks.OrganizationRepository)
	tokenIssuer := infraauth.NewTokenIssuer("test-secret")
	passwordHasher := infraauth.NewDefaultBcryptPasswordHasher()
	uc := NewLoginUseCase(users, memberships, organizations, tokenIssuer, passwordHasher, testutil.NewDiscardLogger())

	hash, _ := passwordHasher.Hash("password123")
	testUser := &user.User{ID: uuid.New(), Name: "Admin", Email: "admin@example.com", PasswordHash: hash, Status: user.StatusActive}
	testOrg := &organization.Organization{ID: uuid.New(), Name: "Example", Domain: "example.com", Workspace: "example", Status: organization.StatusActive}
	testMembership := membership.Membership{ID: uuid.New(), OrganizationID: testOrg.ID, UserID: testUser.ID, Role: membership.RoleAdmin, Status: membership.StatusActive, CreatedAt: time.Now(), UpdatedAt: time.Now()}

	users.On("GetByEmail", mock.Anything, testUser.Email).Return(testUser, nil)
	memberships.On("ListByUser", mock.Anything, mock.Anything).Return(membership.ListResult{Memberships: []membership.Membership{testMembership}, Total: 1}, nil)
	organizations.On("GetByID", mock.Anything, testOrg.ID).Return(testOrg, nil)

	result, err := uc.Execute(context.Background(), LoginCommand{Email: testUser.Email, Password: "password123"})

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotEmpty(t, result.Token)
	assert.Equal(t, testOrg.ID, result.Session.ActiveOrganization.ID)
	assert.Empty(t, result.Session.User.PasswordHash)
}

func TestLoginUseCase_Execute_UserNotFound(t *testing.T) {
	users := new(mocks.UserRepository)
	uc := NewLoginUseCase(users, new(mockMembershipRepository), new(mocks.OrganizationRepository), infraauth.NewTokenIssuer("test-secret"), infraauth.NewDefaultBcryptPasswordHasher(), testutil.NewDiscardLogger())
	users.On("GetByEmail", mock.Anything, "missing@example.com").Return(nil, user.ErrNotFound)

	result, err := uc.Execute(context.Background(), LoginCommand{Email: "missing@example.com", Password: "password123"})
	assert.ErrorIs(t, err, ErrInvalidCredentials)
	assert.Nil(t, result)
}

func TestLoginUseCase_Execute_NoOrganizations(t *testing.T) {
	users := new(mocks.UserRepository)
	memberships := new(mockMembershipRepository)
	passwordHasher := infraauth.NewDefaultBcryptPasswordHasher()
	hash, _ := passwordHasher.Hash("password123")
	testUser := &user.User{ID: uuid.New(), Name: "Admin", Email: "admin@example.com", PasswordHash: hash, Status: user.StatusActive}

	uc := NewLoginUseCase(users, memberships, new(mocks.OrganizationRepository), infraauth.NewTokenIssuer("test-secret"), passwordHasher, testutil.NewDiscardLogger())
	users.On("GetByEmail", mock.Anything, testUser.Email).Return(testUser, nil)
	memberships.On("ListByUser", mock.Anything, mock.Anything).Return(membership.ListResult{}, nil)

	result, err := uc.Execute(context.Background(), LoginCommand{Email: testUser.Email, Password: "password123"})
	assert.ErrorIs(t, err, ErrNoOrganizations)
	assert.Nil(t, result)
}

func TestLoginUseCase_Execute_RepositoryError(t *testing.T) {
	users := new(mocks.UserRepository)
	uc := NewLoginUseCase(users, new(mockMembershipRepository), new(mocks.OrganizationRepository), infraauth.NewTokenIssuer("test-secret"), infraauth.NewDefaultBcryptPasswordHasher(), testutil.NewDiscardLogger())
	users.On("GetByEmail", mock.Anything, "admin@example.com").Return(nil, errors.New("db error"))

	result, err := uc.Execute(context.Background(), LoginCommand{Email: "admin@example.com", Password: "password123"})
	assert.Error(t, err)
	assert.Nil(t, result)
}
