package auth

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/dsbraz/bud2/backend/internal/api/httputil"
	appauth "github.com/dsbraz/bud2/backend/internal/app/auth"
)

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	AccessToken string       `json:"access_token"`
	TokenType   string       `json:"token_type"`
	User        userResponse `json:"user"`
}

type userResponse struct {
	ID       string `json:"id"`
	TenantID string `json:"tenant_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Status   string `json:"status"`
}

type loginUseCase interface {
	Execute(ctx context.Context, cmd appauth.LoginCommand) (*appauth.LoginResult, error)
}

type Handler struct {
	login loginUseCase
}

func NewHandler(login loginUseCase) *Handler {
	return &Handler{login: login}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httputil.WriteProblem(w, http.StatusBadRequest, "Bad Request", "invalid JSON body")
		return
	}

	// Validate required fields
	if req.Email == "" || req.Password == "" {
		httputil.WriteProblem(w, http.StatusUnprocessableEntity, "Unprocessable Entity", "email and password are required")
		return
	}

	result, err := h.login.Execute(r.Context(), appauth.LoginCommand{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		switch {
		case errors.Is(err, appauth.ErrInvalidCredentials):
			httputil.WriteProblem(w, http.StatusUnauthorized, "Unauthorized", "invalid credentials")
		case errors.Is(err, appauth.ErrUserInactive):
			httputil.WriteProblem(w, http.StatusForbidden, "Forbidden", "user account is inactive")
		default:
			httputil.WriteProblem(w, http.StatusInternalServerError, "Internal Server Error", "an unexpected error occurred")
		}
		return
	}

	resp := loginResponse{
		AccessToken: result.Token,
		TokenType:   "Bearer",
		User: userResponse{
			ID:       result.User.ID.String(),
			TenantID: result.User.TenantID.String(),
			Name:     result.User.Name,
			Email:    result.User.Email,
			Role:     string(result.User.Role),
			Status:   string(result.User.Status),
		},
	}

	httputil.WriteJSON(w, http.StatusOK, resp)
}
