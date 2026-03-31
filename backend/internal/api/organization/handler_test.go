package organization

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	apporg "github.com/dsbraz/bud2/backend/internal/app/organization"
	"github.com/dsbraz/bud2/backend/internal/domain"
	"github.com/dsbraz/bud2/backend/internal/domain/organization"
	"github.com/dsbraz/bud2/backend/internal/test/fixtures"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockCreateUseCase struct {
	mock.Mock
}

func (m *mockCreateUseCase) Execute(ctx context.Context, cmd apporg.CreateCommand) (*organization.Organization, error) {
	args := m.Called(ctx, cmd)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*organization.Organization), args.Error(1)
}

type mockGetUseCase struct {
	mock.Mock
}

func (m *mockGetUseCase) Execute(ctx context.Context, id uuid.UUID) (*organization.Organization, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*organization.Organization), args.Error(1)
}

type mockListUseCase struct {
	mock.Mock
}

func (m *mockListUseCase) Execute(ctx context.Context, cmd apporg.ListCommand) (organization.ListResult, error) {
	args := m.Called(ctx, cmd)
	return args.Get(0).(organization.ListResult), args.Error(1)
}

type mockUpdateUseCase struct {
	mock.Mock
}

func (m *mockUpdateUseCase) Execute(ctx context.Context, cmd apporg.UpdateCommand) (*organization.Organization, error) {
	args := m.Called(ctx, cmd)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*organization.Organization), args.Error(1)
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

	reqBody := createRequest{Name: "Test Org", Slug: "test-org", Status: "active"}
	body, _ := json.Marshal(reqBody)
	expectedOrg := fixtures.NewOrganization()

	mockCreate.On("Execute", mock.Anything, apporg.CreateCommand{
		Name: reqBody.Name, Slug: reqBody.Slug, Status: reqBody.Status,
	}).Return(expectedOrg, nil)

	req := httptest.NewRequest(http.MethodPost, "/organizations", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.Create(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	var response Response
	json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Equal(t, expectedOrg.ID, response.ID)
}

func TestHandler_Create_InvalidJSON(t *testing.T) {
	handler, mockCreate, _, _, _ := setupHandler()
	req := httptest.NewRequest(http.MethodPost, "/organizations", bytes.NewReader([]byte("invalid")))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler.Create(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	mockCreate.AssertNotCalled(t, "Execute")
}

func TestHandler_Create_ValidationError(t *testing.T) {
	handler, mockCreate, _, _, _ := setupHandler()
	reqBody := createRequest{Name: "Test", Slug: "test", Status: "active"}
	body, _ := json.Marshal(reqBody)
	mockCreate.On("Execute", mock.Anything, mock.Anything).Return(nil, domain.ErrValidation)
	req := httptest.NewRequest(http.MethodPost, "/organizations", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler.Create(rr, req)
	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
}

func TestHandler_Create_Conflict(t *testing.T) {
	handler, mockCreate, _, _, _ := setupHandler()
	reqBody := createRequest{Name: "Test", Slug: "test", Status: "active"}
	body, _ := json.Marshal(reqBody)
	mockCreate.On("Execute", mock.Anything, mock.Anything).Return(nil, organization.ErrSlugExists)
	req := httptest.NewRequest(http.MethodPost, "/organizations", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler.Create(rr, req)
	assert.Equal(t, http.StatusConflict, rr.Code)
}

func TestHandler_Create_InternalError(t *testing.T) {
	handler, mockCreate, _, _, _ := setupHandler()
	reqBody := createRequest{Name: "Test", Slug: "test", Status: "active"}
	body, _ := json.Marshal(reqBody)
	mockCreate.On("Execute", mock.Anything, mock.Anything).Return(nil, errors.New("db error"))
	req := httptest.NewRequest(http.MethodPost, "/organizations", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler.Create(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestHandler_Get_Success(t *testing.T) {
	handler, _, mockGet, _, _ := setupHandler()
	expectedOrg := fixtures.NewOrganization()
	id := expectedOrg.ID.String()
	mockGet.On("Execute", mock.Anything, expectedOrg.ID).Return(expectedOrg, nil)
	req := httptest.NewRequest(http.MethodGet, "/organizations/"+id, nil)
	rr := httptest.NewRecorder()
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", id)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	handler.Get(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestHandler_Get_InvalidID(t *testing.T) {
	handler, _, _, _, _ := setupHandler()
	req := httptest.NewRequest(http.MethodGet, "/organizations/invalid", nil)
	rr := httptest.NewRecorder()
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "invalid")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	handler.Get(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestHandler_Get_NotFound(t *testing.T) {
	handler, _, mockGet, _, _ := setupHandler()
	id := uuid.New()
	mockGet.On("Execute", mock.Anything, id).Return(nil, organization.ErrNotFound)
	req := httptest.NewRequest(http.MethodGet, "/organizations/"+id.String(), nil)
	rr := httptest.NewRecorder()
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", id.String())
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	handler.Get(rr, req)
	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestHandler_List_Success(t *testing.T) {
	handler, _, _, mockList, _ := setupHandler()
	orgs := fixtures.NewOrganizationList(2)
	result := organization.ListResult{Organizations: orgs, Total: 2}
	mockList.On("Execute", mock.Anything, apporg.ListCommand{Page: 1, Size: 20}).Return(result, nil)
	req := httptest.NewRequest(http.MethodGet, "/organizations", nil)
	rr := httptest.NewRecorder()
	handler.List(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	var response ListResponse
	json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Equal(t, 2, len(response.Data))
	assert.Equal(t, int64(2), response.Total)
}

func TestHandler_List_WithPagination(t *testing.T) {
	handler, _, _, mockList, _ := setupHandler()
	result := organization.ListResult{Organizations: []organization.Organization{}, Total: 0}
	mockList.On("Execute", mock.Anything, apporg.ListCommand{Page: 2, Size: 10}).Return(result, nil)
	req := httptest.NewRequest(http.MethodGet, "/organizations?page=2&size=10", nil)
	rr := httptest.NewRecorder()
	handler.List(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestHandler_List_WithStatusFilter(t *testing.T) {
	handler, _, _, mockList, _ := setupHandler()
	status := "active"
	result := organization.ListResult{Organizations: []organization.Organization{}, Total: 0}
	mockList.On("Execute", mock.Anything, apporg.ListCommand{Page: 1, Size: 20, Status: &status}).Return(result, nil)
	req := httptest.NewRequest(http.MethodGet, "/organizations?status=active", nil)
	rr := httptest.NewRecorder()
	handler.List(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestHandler_Update_Success(t *testing.T) {
	handler, _, _, _, mockUpdate := setupHandler()
	id := uuid.New()
	reqBody := updateRequest{Name: "Updated", Slug: "updated", Status: "active"}
	body, _ := json.Marshal(reqBody)
	updatedOrg := fixtures.NewOrganization()
	updatedOrg.Name = reqBody.Name
	mockUpdate.On("Execute", mock.Anything, apporg.UpdateCommand{
		ID: id, Name: reqBody.Name, Slug: reqBody.Slug, Status: reqBody.Status,
	}).Return(updatedOrg, nil)
	req := httptest.NewRequest(http.MethodPut, "/organizations/"+id.String(), bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", id.String())
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	handler.Update(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestHandler_Update_InvalidID(t *testing.T) {
	handler, _, _, _, _ := setupHandler()
	reqBody := updateRequest{Name: "Updated", Slug: "updated", Status: "active"}
	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPut, "/organizations/invalid", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "invalid")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	handler.Update(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestHandler_Update_NotFound(t *testing.T) {
	handler, _, _, _, mockUpdate := setupHandler()
	id := uuid.New()
	reqBody := updateRequest{Name: "Updated", Slug: "updated", Status: "active"}
	body, _ := json.Marshal(reqBody)
	mockUpdate.On("Execute", mock.Anything, mock.Anything).Return(nil, organization.ErrNotFound)
	req := httptest.NewRequest(http.MethodPut, "/organizations/"+id.String(), bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", id.String())
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	handler.Update(rr, req)
	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestHandler_Update_InvalidJSON(t *testing.T) {
	handler, _, _, _, _ := setupHandler()
	id := uuid.New()
	req := httptest.NewRequest(http.MethodPut, "/organizations/"+id.String(), bytes.NewReader([]byte("invalid")))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", id.String())
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	handler.Update(rr, req)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}
