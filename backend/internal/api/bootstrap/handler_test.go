package bootstrap

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	appbootstrap "github.com/dsbraz/bud2/backend/internal/app/bootstrap"
	"github.com/dsbraz/bud2/backend/internal/domain"
	"github.com/dsbraz/bud2/backend/internal/test/fixtures"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock use case
type mockBootstrapUseCase struct {
	mock.Mock
}

func (m *mockBootstrapUseCase) Execute(ctx context.Context, cmd appbootstrap.Command) (*appbootstrap.Result, error) {
	args := m.Called(ctx, cmd)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*appbootstrap.Result), args.Error(1)
}

func TestHandler_Create_Success(t *testing.T) {
	mockUC := new(mockBootstrapUseCase)
	handler := NewHandler(mockUC)

	reqBody := createRequest{
		OrganizationName: "Test Org",
		OrganizationSlug: "test-org",
		AdminName:        "Admin User",
		AdminEmail:       "admin@example.com",
		AdminPassword:    "admin123",
	}
	body, _ := json.Marshal(reqBody)

	createdOrg := fixtures.NewOrganization()
	createdAdmin := fixtures.NewUser()
	createdAdmin.TenantID = domain.TenantID(createdOrg.ID)

	mockUC.On("Execute", mock.Anything, appbootstrap.Command{
		OrganizationName: reqBody.OrganizationName,
		OrganizationSlug: reqBody.OrganizationSlug,
		AdminName:        reqBody.AdminName,
		AdminEmail:       reqBody.AdminEmail,
		AdminPassword:    reqBody.AdminPassword,
	}).Return(&appbootstrap.Result{
		Organization: *createdOrg,
		Admin:        *createdAdmin,
		AccessToken:  "test-jwt-token",
	}, nil)

	req := httptest.NewRequest(http.MethodPost, "/bootstrap", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.Create(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var response Response
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "test-jwt-token", response.AccessToken)
	assert.Equal(t, "Bearer", response.TokenType)
	assert.Equal(t, createdOrg.ID.String(), response.Organization.ID)
	assert.Equal(t, createdAdmin.ID.String(), response.Admin.ID)
	assert.Equal(t, string(createdAdmin.Role), response.Admin.Role)
}

func TestHandler_Create_InvalidJSON(t *testing.T) {
	mockUC := new(mockBootstrapUseCase)
	handler := NewHandler(mockUC)

	req := httptest.NewRequest(http.MethodPost, "/bootstrap", bytes.NewReader([]byte("invalid")))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.Create(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	mockUC.AssertNotCalled(t, "Execute")
}

func TestHandler_Create_AlreadyBootstrapped(t *testing.T) {
	mockUC := new(mockBootstrapUseCase)
	handler := NewHandler(mockUC)

	reqBody := createRequest{
		OrganizationName: "Test Org",
		OrganizationSlug: "test-org",
		AdminName:        "Admin User",
		AdminEmail:       "admin@example.com",
		AdminPassword:    "admin123",
	}
	body, _ := json.Marshal(reqBody)

	mockUC.On("Execute", mock.Anything, mock.Anything).Return(nil, appbootstrap.ErrAlreadyBootstrapped)

	req := httptest.NewRequest(http.MethodPost, "/bootstrap", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.Create(rr, req)

	assert.Equal(t, http.StatusConflict, rr.Code)
	assert.Contains(t, rr.Body.String(), "already completed")
}

func TestHandler_Create_InternalError(t *testing.T) {
	mockUC := new(mockBootstrapUseCase)
	handler := NewHandler(mockUC)

	reqBody := createRequest{
		OrganizationName: "Test Org",
		OrganizationSlug: "test-org",
		AdminName:        "Admin User",
		AdminEmail:       "admin@example.com",
		AdminPassword:    "admin123",
	}
	body, _ := json.Marshal(reqBody)

	mockUC.On("Execute", mock.Anything, mock.Anything).Return(nil, errors.New("database error"))

	req := httptest.NewRequest(http.MethodPost, "/bootstrap", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.Create(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}
