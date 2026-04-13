package user

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	appuser "github.com/getbud-co/bud2/backend/internal/app/user"
	"github.com/getbud-co/bud2/backend/internal/domain"
	"github.com/getbud-co/bud2/backend/internal/domain/membership"
	usr "github.com/getbud-co/bud2/backend/internal/domain/user"
	"github.com/getbud-co/bud2/backend/internal/test/fixtures"
)

type mockCreateUseCase struct{ mock.Mock }

func (m *mockCreateUseCase) Execute(ctx context.Context, cmd appuser.CreateCommand) (*appuser.Member, error) {
	args := m.Called(ctx, cmd)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*appuser.Member), args.Error(1)
}

type mockGetUseCase struct{ mock.Mock }

func (m *mockGetUseCase) Execute(ctx context.Context, organizationID domain.TenantID, id uuid.UUID) (*appuser.Member, error) {
	args := m.Called(ctx, organizationID, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*appuser.Member), args.Error(1)
}

type mockListUseCase struct{ mock.Mock }

func (m *mockListUseCase) Execute(ctx context.Context, cmd appuser.ListCommand) (appuser.MemberListResult, error) {
	args := m.Called(ctx, cmd)
	return args.Get(0).(appuser.MemberListResult), args.Error(1)
}

type mockUpdateUseCase struct{ mock.Mock }

func (m *mockUpdateUseCase) Execute(ctx context.Context, cmd appuser.UpdateCommand) (*appuser.Member, error) {
	args := m.Called(ctx, cmd)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*appuser.Member), args.Error(1)
}

func newTestMember() *appuser.Member {
	u := fixtures.NewUser()
	m := fixtures.NewMembership()
	return &appuser.Member{
		User:             *u,
		OrganizationID:   m.OrganizationID,
		MembershipRole:   m.Role,
		MembershipStatus: m.Status,
	}
}

func TestHandler_Create_Success(t *testing.T) {
	createUC := new(mockCreateUseCase)
	handler := NewHandler(createUC, nil, nil, nil)

	tenantID := fixtures.NewTestTenantID()
	member := newTestMember()
	createUC.On("Execute", mock.Anything, appuser.CreateCommand{
		OrganizationID: tenantID, Name: "Test User", Email: "test@example.com", Password: "password123", Role: "admin",
	}).Return(member, nil)

	body, _ := json.Marshal(createRequest{Name: "Test User", Email: "test@example.com", Password: "password123", Role: "admin"})
	req := httptest.NewRequest(http.MethodPost, "/organizations/"+tenantID.UUID().String()+"/users", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req = req.WithContext(fixtures.NewContextWithTenant(tenantID))
	rr := httptest.NewRecorder()

	handler.Create(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	var resp Response
	assert.NoError(t, json.Unmarshal(rr.Body.Bytes(), &resp))
	assert.Equal(t, member.User.ID.String(), resp.ID)
	assert.Equal(t, "admin", resp.Role)
}

func TestHandler_Create_MissingTenant(t *testing.T) {
	handler := NewHandler(new(mockCreateUseCase), nil, nil, nil)

	body, _ := json.Marshal(createRequest{Name: "Test User", Email: "test@example.com", Password: "password123", Role: "admin"})
	req := httptest.NewRequest(http.MethodPost, "/organizations/x/users", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.Create(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
}

func TestHandler_Create_InvalidJSON(t *testing.T) {
	handler := NewHandler(new(mockCreateUseCase), nil, nil, nil)

	tenantID := fixtures.NewTestTenantID()
	req := httptest.NewRequest(http.MethodPost, "/organizations/x/users", bytes.NewReader([]byte("invalid")))
	req.Header.Set("Content-Type", "application/json")
	req = req.WithContext(fixtures.NewContextWithTenant(tenantID))
	rr := httptest.NewRecorder()

	handler.Create(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestHandler_Create_RequestValidationError(t *testing.T) {
	handler := NewHandler(new(mockCreateUseCase), nil, nil, nil)

	tenantID := fixtures.NewTestTenantID()
	body, _ := json.Marshal(createRequest{Name: "T", Email: "not-email", Password: "short", Role: "invalid"})
	req := httptest.NewRequest(http.MethodPost, "/organizations/x/users", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req = req.WithContext(fixtures.NewContextWithTenant(tenantID))
	rr := httptest.NewRecorder()

	handler.Create(rr, req)

	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
}

func TestHandler_Create_MembershipConflict(t *testing.T) {
	createUC := new(mockCreateUseCase)
	handler := NewHandler(createUC, nil, nil, nil)

	tenantID := fixtures.NewTestTenantID()
	createUC.On("Execute", mock.Anything, mock.Anything).Return(nil, membership.ErrAlreadyExists)

	body, _ := json.Marshal(createRequest{Name: "Test User", Email: "test@example.com", Password: "password123", Role: "admin"})
	req := httptest.NewRequest(http.MethodPost, "/organizations/x/users", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req = req.WithContext(fixtures.NewContextWithTenant(tenantID))
	rr := httptest.NewRecorder()

	handler.Create(rr, req)

	assert.Equal(t, http.StatusConflict, rr.Code)
}

func TestHandler_Get_Success(t *testing.T) {
	getUC := new(mockGetUseCase)
	handler := NewHandler(nil, getUC, nil, nil)

	tenantID := fixtures.NewTestTenantID()
	member := newTestMember()
	getUC.On("Execute", mock.Anything, tenantID, member.User.ID).Return(member, nil)

	req := httptest.NewRequest(http.MethodGet, "/organizations/x/users/"+member.User.ID.String(), nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", member.User.ID.String())
	req = req.WithContext(context.WithValue(fixtures.NewContextWithTenant(tenantID), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	handler.Get(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var resp Response
	assert.NoError(t, json.Unmarshal(rr.Body.Bytes(), &resp))
	assert.Equal(t, member.User.ID.String(), resp.ID)
}

func TestHandler_Get_MissingTenant(t *testing.T) {
	handler := NewHandler(nil, new(mockGetUseCase), nil, nil)

	req := httptest.NewRequest(http.MethodGet, "/organizations/x/users/"+uuid.New().String(), nil)
	rr := httptest.NewRecorder()

	handler.Get(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
}

func TestHandler_Get_InvalidID(t *testing.T) {
	handler := NewHandler(nil, new(mockGetUseCase), nil, nil)

	tenantID := fixtures.NewTestTenantID()
	req := httptest.NewRequest(http.MethodGet, "/organizations/x/users/invalid", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "invalid")
	req = req.WithContext(context.WithValue(fixtures.NewContextWithTenant(tenantID), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	handler.Get(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestHandler_Get_NotFound(t *testing.T) {
	getUC := new(mockGetUseCase)
	handler := NewHandler(nil, getUC, nil, nil)

	tenantID := fixtures.NewTestTenantID()
	id := uuid.New()
	getUC.On("Execute", mock.Anything, tenantID, id).Return(nil, membership.ErrNotFound)

	req := httptest.NewRequest(http.MethodGet, "/organizations/x/users/"+id.String(), nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", id.String())
	req = req.WithContext(context.WithValue(fixtures.NewContextWithTenant(tenantID), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	handler.Get(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestHandler_List_Success(t *testing.T) {
	listUC := new(mockListUseCase)
	handler := NewHandler(nil, nil, listUC, nil)

	tenantID := fixtures.NewTestTenantID()
	member := *newTestMember()
	listUC.On("Execute", mock.Anything, appuser.ListCommand{
		OrganizationID: tenantID, Page: 1, Size: 20,
	}).Return(appuser.MemberListResult{Members: []appuser.Member{member}, Total: 1}, nil)

	req := httptest.NewRequest(http.MethodGet, "/organizations/x/users", nil)
	req = req.WithContext(fixtures.NewContextWithTenant(tenantID))
	rr := httptest.NewRecorder()

	handler.List(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var resp ListResponse
	assert.NoError(t, json.Unmarshal(rr.Body.Bytes(), &resp))
	assert.Len(t, resp.Data, 1)
}

func TestHandler_Update_Success(t *testing.T) {
	updateUC := new(mockUpdateUseCase)
	handler := NewHandler(nil, nil, nil, updateUC)

	tenantID := fixtures.NewTestTenantID()
	member := newTestMember()
	updateUC.On("Execute", mock.Anything, appuser.UpdateCommand{
		OrganizationID: tenantID, ID: member.User.ID, Name: "Updated", Email: "test@example.com", Role: "manager", Status: "active",
	}).Return(member, nil)

	body, _ := json.Marshal(updateRequest{Name: "Updated", Email: "test@example.com", Role: "manager", Status: "active"})
	req := httptest.NewRequest(http.MethodPut, "/organizations/x/users/"+member.User.ID.String(), bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", member.User.ID.String())
	req = req.WithContext(context.WithValue(fixtures.NewContextWithTenant(tenantID), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	handler.Update(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestHandler_Update_NotFound(t *testing.T) {
	updateUC := new(mockUpdateUseCase)
	handler := NewHandler(nil, nil, nil, updateUC)

	tenantID := fixtures.NewTestTenantID()
	id := uuid.New()
	updateUC.On("Execute", mock.Anything, mock.Anything).Return(nil, usr.ErrNotFound)

	body, _ := json.Marshal(updateRequest{Name: "Updated", Email: "test@example.com", Role: "manager", Status: "active"})
	req := httptest.NewRequest(http.MethodPut, "/organizations/x/users/"+id.String(), bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", id.String())
	req = req.WithContext(context.WithValue(fixtures.NewContextWithTenant(tenantID), chi.RouteCtxKey, rctx))
	rr := httptest.NewRecorder()

	handler.Update(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestHandler_Create_InternalError(t *testing.T) {
	createUC := new(mockCreateUseCase)
	handler := NewHandler(createUC, nil, nil, nil)

	tenantID := fixtures.NewTestTenantID()
	createUC.On("Execute", mock.Anything, mock.Anything).Return(nil, errors.New("internal error"))

	body, _ := json.Marshal(createRequest{Name: "Test User", Email: "test@example.com", Password: "password123", Role: "admin"})
	req := httptest.NewRequest(http.MethodPost, "/organizations/x/users", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req = req.WithContext(fixtures.NewContextWithTenant(tenantID))
	rr := httptest.NewRecorder()

	handler.Create(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}
