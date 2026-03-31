package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	appauth "github.com/dsbraz/bud2/backend/internal/app/auth"
	"github.com/dsbraz/bud2/backend/internal/test/fixtures"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock login use case
type mockLoginUseCase struct {
	mock.Mock
}

func (m *mockLoginUseCase) Execute(ctx context.Context, cmd appauth.LoginCommand) (*appauth.LoginResult, error) {
	args := m.Called(ctx, cmd)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*appauth.LoginResult), args.Error(1)
}

func TestHandler_Login_Success(t *testing.T) {
	mockLogin := new(mockLoginUseCase)
	handler := NewHandler(mockLogin)

	reqBody := loginRequest{
		Email:    "admin@example.com",
		Password: "password123",
	}
	body, _ := json.Marshal(reqBody)

	testUser := fixtures.NewUser()
	mockLogin.On("Execute", mock.Anything, appauth.LoginCommand{
		Email:    reqBody.Email,
		Password: reqBody.Password,
	}).Return(&appauth.LoginResult{
		Token: "test-jwt-token",
		User:  *testUser,
	}, nil)

	req := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.Login(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response loginResponse
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "test-jwt-token", response.AccessToken)
	assert.Equal(t, "Bearer", response.TokenType)
	assert.Equal(t, testUser.ID.String(), response.User.ID)
}

func TestHandler_Login_InvalidJSON(t *testing.T) {
	mockLogin := new(mockLoginUseCase)
	handler := NewHandler(mockLogin)

	req := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewReader([]byte("invalid")))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.Login(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	mockLogin.AssertNotCalled(t, "Execute")
}

func TestHandler_Login_MissingFields(t *testing.T) {
	mockLogin := new(mockLoginUseCase)
	handler := NewHandler(mockLogin)

	// Missing password
	reqBody := loginRequest{Email: "admin@example.com"}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.Login(rr, req)

	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	mockLogin.AssertNotCalled(t, "Execute")
}

func TestHandler_Login_InvalidCredentials(t *testing.T) {
	mockLogin := new(mockLoginUseCase)
	handler := NewHandler(mockLogin)

	reqBody := loginRequest{
		Email:    "admin@example.com",
		Password: "wrongpassword",
	}
	body, _ := json.Marshal(reqBody)

	mockLogin.On("Execute", mock.Anything, mock.Anything).Return(nil, appauth.ErrInvalidCredentials)

	req := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.Login(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
}

func TestHandler_Login_UserInactive(t *testing.T) {
	mockLogin := new(mockLoginUseCase)
	handler := NewHandler(mockLogin)

	reqBody := loginRequest{
		Email:    "admin@example.com",
		Password: "password123",
	}
	body, _ := json.Marshal(reqBody)

	mockLogin.On("Execute", mock.Anything, mock.Anything).Return(nil, appauth.ErrUserInactive)

	req := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.Login(rr, req)

	assert.Equal(t, http.StatusForbidden, rr.Code)
}

func TestHandler_Login_InternalError(t *testing.T) {
	mockLogin := new(mockLoginUseCase)
	handler := NewHandler(mockLogin)

	reqBody := loginRequest{
		Email:    "admin@example.com",
		Password: "password123",
	}
	body, _ := json.Marshal(reqBody)

	mockLogin.On("Execute", mock.Anything, mock.Anything).Return(nil, errors.New("internal error"))

	req := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.Login(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}
