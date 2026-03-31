package user

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	appuser "github.com/dsbraz/bud2/backend/internal/app/user"
	"github.com/dsbraz/bud2/backend/internal/domain"
	"github.com/dsbraz/bud2/backend/internal/domain/user"
	"github.com/dsbraz/bud2/backend/internal/test/fixtures"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockCreateUseCase struct {
	mock.Mock
}

func (m *mockCreateUseCase) Execute(ctx context.Context, cmd appuser.CreateCommand) (*user.User, error) {
	args := m.Called(ctx, cmd)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*user.User), args.Error(1)
}

type mockGetUseCase struct {
	mock.Mock
}

func (m *mockGetUseCase) Execute(ctx context.Context, tenantID domain.TenantID, id uuid.UUID) (*user.User, error) {
	args := m.Called(ctx, tenantID, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*user.User), args.Error(1)
}

type mockListUseCase struct {
	mock.Mock
}

func (m *mockListUseCase) Execute(ctx context.Context, cmd appuser.ListCommand) (user.ListResult, error) {
	args := m.Called(ctx, cmd)
	return args.Get(0).(user.ListResult), args.Error(1)
}

type mockUpdateUseCase struct {
	mock.Mock
}

func (m *mockUpdateUseCase) Execute(ctx context.Context, cmd appuser.UpdateCommand) (*user.User, error) {
	args := m.Called(ctx, cmd)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*user.User), args.Error(1)
}

func setupHandler() (*Handler, *mockCreateUseCase, *mockGetUseCase, *mockListUseCase, *mockUpdateUseCase) {
	mockCreate := new(mockCreateUseCase)
	mockGet := new(mockGetUseCase)
	mockList := new(mockListUseCase)
	mockUpdate := new(mockUpdateUseCase)
	handler := NewHandler(mockCreate, mockGet, mockList, mockUpdate)
	return handler, mockCreate, mockGet, mockList, mockUpdate
}

func TestHandler_Create_Success(t *testing.T) {
	handler, mockCreate, _, _, _ := setupHandler()
	tenantID := fixtures.NewTestTenantID()

	reqBody := createRequest{Name: "Test User", Email: "test@example.com", Role: "admin"}
	body, _ := json.Marshal(reqBody)
	expectedUser := fixtures.NewUser()

	mockCreate.On("Execute", mock.Anything, appuser.CreateCommand{
		TenantID: tenantID, Name: reqBody.Name, Email: reqBody.Email, Role: reqBody.Role,
	}).Return(expectedUser, nil)

	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req = req.WithContext(fixtures.NewContextWithTenant(tenantID))
	rr := httptest.NewRecorder()

	handler.Create(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestHandler_Create_MissingTenant(t *testing.T) {
	handler, mockCreate, _, _, _ := setupHandler()
	reqBody := createRequest{Name: "Test", Email: "test@example.com", Role: "admin"}
	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler.Create(rr, req)
	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	mockCreate.AssertNotCalled(t, "Execute")
}

func TestHandler_Create_InvalidJSON(t *testing.T) {
	handler, _, _, _, _ := setupHandler()
	tenantID := fixtures.NewTestTenantID()
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader([]byte("invalid")))
	req.Header.Set("Content-Type", "application/json")
	req = req.WithContext(fixtures.NewContextWithTenant(tenantID))
	rr := httptest.NewRecorder()
	handler.Create(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestHandler_Create_ValidationError(t *testing.T) {
	handler, mockCreate, _, _, _ := setupHandler()
	tenantID := fixtures.NewTestTenantID()
	reqBody := createRequest{Name: "Test", Email: "test@example.com", Role: "admin"}
	body, _ := json.Marshal(reqBody)
	mockCreate.On("Execute", mock.Anything, mock.Anything).Return(nil, domain.ErrValidation)
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req = req.WithContext(fixtures.NewContextWithTenant(tenantID))
	rr := httptest.NewRecorder()
	handler.Create(rr, req)
	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
}

func TestHandler_Create_Conflict(t *testing.T) {
	handler, mockCreate, _, _, _ := setupHandler()
	tenantID := fixtures.NewTestTenantID()
	reqBody := createRequest{Name: "Test", Email: "test@example.com", Role: "admin"}
	body, _ := json.Marshal(reqBody)
	mockCreate.On("Execute", mock.Anything, mock.Anything).Return(nil, user.ErrEmailExists)
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req = req.WithContext(fixtures.NewContextWithTenant(tenantID))
	rr := httptest.NewRecorder()
	handler.Create(rr, req)
	assert.Equal(t, http.StatusConflict, rr.Code)
}

func TestHandler_Get_Success(t *testing.T) {
	handler, _, mockGet, _, _ := setupHandler()
	tenantID := fixtures.NewTestTenantID()
	expectedUser := fixtures.NewUser()
	id := expectedUser.ID

	mockGet.On("Execute", mock.Anything, tenantID, id).Return(expectedUser, nil)

	req := httptest.NewRequest(http.MethodGet, "/users/"+id.String(), nil)
	req = req.WithContext(fixtures.NewContextWithTenant(tenantID))
	rr := httptest.NewRecorder()
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", id.String())
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler.Get(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestHandler_Get_MissingTenant(t *testing.T) {
	handler, _, _, _, _ := setupHandler()
	id := uuid.New()
	req := httptest.NewRequest(http.MethodGet, "/users/"+id.String(), nil)
	rr := httptest.NewRecorder()
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", id.String())
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	handler.Get(rr, req)
	assert.Equal(t, http.StatusUnauthorized, rr.Code)
}

func TestHandler_Get_InvalidID(t *testing.T) {
	handler, _, _, _, _ := setupHandler()
	tenantID := fixtures.NewTestTenantID()
	req := httptest.NewRequest(http.MethodGet, "/users/invalid", nil)
	req = req.WithContext(fixtures.NewContextWithTenant(tenantID))
	rr := httptest.NewRecorder()
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "invalid")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	handler.Get(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestHandler_Get_NotFound(t *testing.T) {
	handler, _, mockGet, _, _ := setupHandler()
	tenantID := fixtures.NewTestTenantID()
	id := uuid.New()
	mockGet.On("Execute", mock.Anything, tenantID, id).Return(nil, user.ErrNotFound)
	req := httptest.NewRequest(http.MethodGet, "/users/"+id.String(), nil)
	req = req.WithContext(fixtures.NewContextWithTenant(tenantID))
	rr := httptest.NewRecorder()
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", id.String())
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	handler.Get(rr, req)
	assert.Equal(t, http.StatusNotFound, rr.Code)
}
