package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dsbraz/bud2/backend/internal/api/httputil"
	"github.com/dsbraz/bud2/backend/internal/api/validator"
	appauth "github.com/dsbraz/bud2/backend/internal/app/auth"
	"github.com/dsbraz/bud2/backend/internal/domain"
)

type loginUseCase interface {
	Execute(ctx context.Context, cmd appauth.LoginCommand) (*appauth.LoginResult, error)
}

type sessionUseCase interface {
	Execute(ctx context.Context, claims domain.UserClaims) (*appauth.Session, error)
}

type updateSessionUseCase interface {
	Execute(ctx context.Context, claims domain.UserClaims, cmd appauth.SwitchOrganizationCommand) (*appauth.SwitchOrganizationResult, error)
}

type Handler struct {
	login         loginUseCase
	getSession    sessionUseCase
	updateSession updateSessionUseCase
}

func NewHandler(login loginUseCase, getSession sessionUseCase, updateSession updateSessionUseCase) *Handler {
	return &Handler{login: login, getSession: getSession, updateSession: updateSession}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httputil.WriteProblem(w, http.StatusBadRequest, "Bad Request", "invalid JSON body")
		return
	}
	if err := validator.Validate(req); err != nil {
		httputil.WriteProblem(w, http.StatusUnprocessableEntity, "Validation Error", validator.FormatValidationErrors(err))
		return
	}

	result, err := h.login.Execute(r.Context(), req.toCommand())
	if err != nil {
		handleAuthError(w, err)
		return
	}

	resp := responseFromSession(result.Session)
	resp.AccessToken = result.Token
	resp.TokenType = "Bearer"
	httputil.WriteJSON(w, http.StatusOK, resp)
}

func (h *Handler) Session(w http.ResponseWriter, r *http.Request) {
	claims, err := domain.ClaimsFromContext(r.Context())
	if err != nil {
		httputil.WriteProblem(w, http.StatusUnauthorized, "Unauthorized", "authentication required")
		return
	}

	result, err := h.getSession.Execute(r.Context(), claims)
	if err != nil {
		handleAuthError(w, err)
		return
	}

	httputil.WriteJSON(w, http.StatusOK, responseFromSession(*result))
}

func (h *Handler) UpdateSession(w http.ResponseWriter, r *http.Request) {
	claims, err := domain.ClaimsFromContext(r.Context())
	if err != nil {
		httputil.WriteProblem(w, http.StatusUnauthorized, "Unauthorized", "authentication required")
		return
	}

	var req updateSessionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httputil.WriteProblem(w, http.StatusBadRequest, "Bad Request", "invalid JSON body")
		return
	}
	if err := validator.Validate(req); err != nil {
		httputil.WriteProblem(w, http.StatusUnprocessableEntity, "Validation Error", validator.FormatValidationErrors(err))
		return
	}

	cmd, err := req.toCommand()
	if err != nil {
		httputil.WriteProblem(w, http.StatusBadRequest, "Bad Request", "invalid organization_id format")
		return
	}

	result, err := h.updateSession.Execute(r.Context(), claims, cmd)
	if err != nil {
		handleAuthError(w, err)
		return
	}

	resp := responseFromSession(result.Session)
	resp.AccessToken = result.Token
	resp.TokenType = "Bearer"
	httputil.WriteJSON(w, http.StatusOK, resp)
}
