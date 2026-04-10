package bootstrap

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/dsbraz/bud2/backend/internal/domain"
	"github.com/dsbraz/bud2/backend/internal/domain/organization"
	"github.com/dsbraz/bud2/backend/internal/domain/user"
	"github.com/dsbraz/bud2/backend/internal/test/fixtures"
	"github.com/dsbraz/bud2/backend/internal/test/mocks"
	"github.com/dsbraz/bud2/backend/internal/test/testutil"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockTxManager is a specialized mock that can execute the transaction function
type MockTxManager struct {
	mock.Mock
	orgRepo  *mocks.OrganizationRepository
	userRepo *mocks.UserRepository
}

func (m *MockTxManager) WithTx(ctx context.Context, fn func(orgRepo organization.Repository, userRepo user.Repository) error) error {
	args := m.Called(ctx, mock.Anything)

	// Execute the function with mocked repositories if no error is set
	if args.Error(0) == nil && m.orgRepo != nil && m.userRepo != nil {
		return fn(m.orgRepo, m.userRepo)
	}

	return args.Error(0)
}

func TestUseCase_Execute_Success(t *testing.T) {
	mockOrgRepo := new(mocks.OrganizationRepository)
	mockUserRepo := new(mocks.UserRepository)
	mockTxManager := &MockTxManager{
		orgRepo:  mockOrgRepo,
		userRepo: mockUserRepo,
	}
	mockIssuer := new(mocks.TokenIssuer)

	uc := NewUseCase(mockOrgRepo, mockTxManager, mockIssuer, testutil.NewDiscardLogger())

	cmd := Command{
		OrganizationName: "Test Org",
		OrganizationSlug: "test-org",
		AdminName:        "Admin User",
		AdminEmail:       "admin@example.com",
		AdminPassword:    "admin123",
	}

	createdOrg := fixtures.NewOrganization()
	createdAdmin := fixtures.NewUser()
	createdAdmin.TenantID = domain.TenantID(createdOrg.ID)
	expectedToken := "test-jwt-token"

	// Expectations
	mockOrgRepo.On("CountAll", mock.Anything).Return(int64(0), nil)
	mockTxManager.On("WithTx", mock.Anything, mock.Anything).Return(nil)
	mockOrgRepo.On("Create", mock.Anything, mock.AnythingOfType("*organization.Organization")).Return(createdOrg, nil)
	mockUserRepo.On("Create", mock.Anything, mock.AnythingOfType("*user.User")).Return(createdAdmin, nil)
	mockIssuer.On("IssueToken", mock.AnythingOfType("domain.UserClaims"), 24*time.Hour).Return(expectedToken, nil)

	// Execute
	result, err := uc.Execute(context.Background(), cmd)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, createdOrg.ID, result.Organization.ID)
	assert.Equal(t, createdAdmin.ID, result.Admin.ID)
	assert.Equal(t, expectedToken, result.AccessToken)
}

func TestUseCase_Execute_AlreadyBootstrapped(t *testing.T) {
	mockOrgRepo := new(mocks.OrganizationRepository)
	mockTxManager := new(MockTxManager)
	mockIssuer := new(mocks.TokenIssuer)

	uc := NewUseCase(mockOrgRepo, mockTxManager, mockIssuer, testutil.NewDiscardLogger())

	cmd := Command{
		OrganizationName: "Test Org",
		OrganizationSlug: "test-org",
		AdminName:        "Admin User",
		AdminEmail:       "admin@example.com",
		AdminPassword:    "admin123",
	}

	// Already has organizations
	mockOrgRepo.On("CountAll", mock.Anything).Return(int64(5), nil)

	result, err := uc.Execute(context.Background(), cmd)

	assert.ErrorIs(t, err, ErrAlreadyBootstrapped)
	assert.Nil(t, result)
	mockTxManager.AssertNotCalled(t, "WithTx")
	mockIssuer.AssertNotCalled(t, "IssueToken")
}

func TestUseCase_Execute_CountError(t *testing.T) {
	mockOrgRepo := new(mocks.OrganizationRepository)
	mockTxManager := new(MockTxManager)
	mockIssuer := new(mocks.TokenIssuer)

	uc := NewUseCase(mockOrgRepo, mockTxManager, mockIssuer, testutil.NewDiscardLogger())

	cmd := Command{
		OrganizationName: "Test Org",
		OrganizationSlug: "test-org",
		AdminName:        "Admin User",
		AdminEmail:       "admin@example.com",
		AdminPassword:    "admin123",
	}

	dbError := errors.New("database error")
	mockOrgRepo.On("CountAll", mock.Anything).Return(int64(0), dbError)

	result, err := uc.Execute(context.Background(), cmd)

	assert.ErrorIs(t, err, dbError)
	assert.Nil(t, result)
}

func TestUseCase_Execute_TransactionError(t *testing.T) {
	mockOrgRepo := new(mocks.OrganizationRepository)
	mockTxManager := new(MockTxManager)
	mockIssuer := new(mocks.TokenIssuer)

	uc := NewUseCase(mockOrgRepo, mockTxManager, mockIssuer, testutil.NewDiscardLogger())

	cmd := Command{
		OrganizationName: "Test Org",
		OrganizationSlug: "test-org",
		AdminName:        "Admin User",
		AdminEmail:       "admin@example.com",
		AdminPassword:    "admin123",
	}

	txError := errors.New("transaction failed")
	mockOrgRepo.On("CountAll", mock.Anything).Return(int64(0), nil)
	mockTxManager.On("WithTx", mock.Anything, mock.Anything).Return(txError)

	result, err := uc.Execute(context.Background(), cmd)

	assert.ErrorIs(t, err, txError)
	assert.Nil(t, result)
	mockIssuer.AssertNotCalled(t, "IssueToken")
}

func TestUseCase_Execute_TokenError(t *testing.T) {
	mockOrgRepo := new(mocks.OrganizationRepository)
	mockUserRepo := new(mocks.UserRepository)
	mockTxManager := &MockTxManager{
		orgRepo:  mockOrgRepo,
		userRepo: mockUserRepo,
	}
	mockIssuer := new(mocks.TokenIssuer)

	uc := NewUseCase(mockOrgRepo, mockTxManager, mockIssuer, testutil.NewDiscardLogger())

	cmd := Command{
		OrganizationName: "Test Org",
		OrganizationSlug: "test-org",
		AdminName:        "Admin User",
		AdminEmail:       "admin@example.com",
		AdminPassword:    "admin123",
	}

	createdOrg := fixtures.NewOrganization()
	createdAdmin := fixtures.NewUser()
	createdAdmin.TenantID = domain.TenantID(createdOrg.ID)

	tokenError := errors.New("token issuance failed")
	mockOrgRepo.On("CountAll", mock.Anything).Return(int64(0), nil)
	mockTxManager.On("WithTx", mock.Anything, mock.Anything).Return(nil)
	mockOrgRepo.On("Create", mock.Anything, mock.AnythingOfType("*organization.Organization")).Return(createdOrg, nil)
	mockUserRepo.On("Create", mock.Anything, mock.AnythingOfType("*user.User")).Return(createdAdmin, nil)
	mockIssuer.On("IssueToken", mock.AnythingOfType("domain.UserClaims"), 24*time.Hour).Return("", tokenError)

	result, err := uc.Execute(context.Background(), cmd)

	assert.ErrorIs(t, err, tokenError)
	assert.Nil(t, result)
}

func TestUseCase_Execute_OrganizationCreatedWithCorrectData(t *testing.T) {
	mockOrgRepo := new(mocks.OrganizationRepository)
	mockUserRepo := new(mocks.UserRepository)
	mockTxManager := &MockTxManager{
		orgRepo:  mockOrgRepo,
		userRepo: mockUserRepo,
	}
	mockIssuer := new(mocks.TokenIssuer)

	uc := NewUseCase(mockOrgRepo, mockTxManager, mockIssuer, testutil.NewDiscardLogger())

	cmd := Command{
		OrganizationName: "Acme Corp",
		OrganizationSlug: "acme-corp",
		AdminName:        "John Admin",
		AdminEmail:       "john@acme.com",
		AdminPassword:    "admin123",
	}

	createdOrg := fixtures.NewOrganization()
	createdOrg.Name = cmd.OrganizationName
	createdOrg.Slug = cmd.OrganizationSlug

	createdAdmin := fixtures.NewUser()
	createdAdmin.TenantID = domain.TenantID(createdOrg.ID)
	createdAdmin.Name = cmd.AdminName
	createdAdmin.Email = cmd.AdminEmail

	mockOrgRepo.On("CountAll", mock.Anything).Return(int64(0), nil)
	mockTxManager.On("WithTx", mock.Anything, mock.Anything).Return(nil)

	// Verify organization is created with correct data
	mockOrgRepo.On("Create", mock.Anything, mock.MatchedBy(func(o *organization.Organization) bool {
		return o.Name == cmd.OrganizationName &&
			o.Slug == cmd.OrganizationSlug &&
			o.Status == organization.StatusActive
	})).Return(createdOrg, nil)

	// Verify admin is created with correct data
	mockUserRepo.On("Create", mock.Anything, mock.MatchedBy(func(u *user.User) bool {
		return u.Name == cmd.AdminName &&
			u.Email == cmd.AdminEmail &&
			u.Role == user.RoleAdmin &&
			u.Status == user.StatusActive &&
			u.TenantID == domain.TenantID(createdOrg.ID)
	})).Return(createdAdmin, nil)

	mockIssuer.On("IssueToken", mock.AnythingOfType("domain.UserClaims"), mock.Anything).Return("token", nil)

	result, err := uc.Execute(context.Background(), cmd)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	mockOrgRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestUseCase_Execute_TenantIDMatchesOrganizationID(t *testing.T) {
	mockOrgRepo := new(mocks.OrganizationRepository)
	mockUserRepo := new(mocks.UserRepository)
	mockTxManager := &MockTxManager{
		orgRepo:  mockOrgRepo,
		userRepo: mockUserRepo,
	}
	mockIssuer := new(mocks.TokenIssuer)

	uc := NewUseCase(mockOrgRepo, mockTxManager, mockIssuer, testutil.NewDiscardLogger())

	cmd := Command{
		OrganizationName: "Test Org",
		OrganizationSlug: "test-org",
		AdminName:        "Admin User",
		AdminEmail:       "admin@example.com",
		AdminPassword:    "admin123",
	}

	orgID := uuid.New()
	createdOrg := fixtures.NewOrganization()
	createdOrg.ID = orgID

	var capturedUser *user.User
	mockOrgRepo.On("CountAll", mock.Anything).Return(int64(0), nil)
	mockTxManager.On("WithTx", mock.Anything, mock.Anything).Return(nil)
	mockOrgRepo.On("Create", mock.Anything, mock.Anything).Return(createdOrg, nil)
	mockUserRepo.On("Create", mock.Anything, mock.AnythingOfType("*user.User")).Run(func(args mock.Arguments) {
		capturedUser = args.Get(1).(*user.User)
	}).Return(fixtures.NewUser(), nil)
	mockIssuer.On("IssueToken", mock.Anything, mock.Anything).Return("token", nil)

	uc.Execute(context.Background(), cmd)

	// Verify tenant ID matches organization ID
	assert.Equal(t, domain.TenantID(orgID), capturedUser.TenantID)
}
