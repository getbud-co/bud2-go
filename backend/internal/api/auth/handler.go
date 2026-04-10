package auth

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"

	"github.com/dsbraz/bud2/backend/internal/api/httputil"
	"github.com/dsbraz/bud2/backend/internal/api/validator"
	appauth "github.com/dsbraz/bud2/backend/internal/app/auth"
	"github.com/dsbraz/bud2/backend/internal/domain"
)

type loginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type switchOrganizationRequest struct {
	OrganizationID string `json:"organization_id" validate:"required,uuid4"`
}

type authResponse struct {
	AccessToken        string                 `json:"access_token,omitempty"`
	TokenType          string                 `json:"token_type,omitempty"`
	User               userResponse           `json:"user"`
	ActiveOrganization *organizationResponse  `json:"active_organization,omitempty"`
	Organizations      []organizationResponse `json:"organizations"`
}

type userResponse struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Status        string `json:"status"`
	IsSystemAdmin bool   `json:"is_system_admin"`
}

type organizationResponse struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Domain           string `json:"domain"`
	Workspace        string `json:"workspace"`
	Status           string `json:"status"`
	MembershipRole   string `json:"membership_role,omitempty"`
	MembershipStatus string `json:"membership_status,omitempty"`
}

type loginUseCase interface {
	Execute(ctx context.Context, cmd appauth.LoginCommand) (*appauth.LoginResult, error)
}

type sessionUseCase interface {
	Execute(ctx context.Context, claims domain.UserClaims) (*appauth.Session, error)
}

type switchOrganizationUseCase interface {
	Execute(ctx context.Context, claims domain.UserClaims, cmd appauth.SwitchOrganizationCommand) (*appauth.SwitchOrganizationResult, error)
}

type Handler struct {
	login              loginUseCase
	getSession         sessionUseCase
	switchOrganization switchOrganizationUseCase
}

func NewHandler(login loginUseCase, getSession sessionUseCase, switchOrganization switchOrganizationUseCase) *Handler {
	return &Handler{login: login, getSession: getSession, switchOrganization: switchOrganization}
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

	result, err := h.login.Execute(r.Context(), appauth.LoginCommand{Email: req.Email, Password: req.Password})
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

func (h *Handler) SwitchOrganization(w http.ResponseWriter, r *http.Request) {
	claims, err := domain.ClaimsFromContext(r.Context())
	if err != nil {
		httputil.WriteProblem(w, http.StatusUnauthorized, "Unauthorized", "authentication required")
		return
	}

	var req switchOrganizationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httputil.WriteProblem(w, http.StatusBadRequest, "Bad Request", "invalid JSON body")
		return
	}
	if err := validator.Validate(req); err != nil {
		httputil.WriteProblem(w, http.StatusUnprocessableEntity, "Validation Error", validator.FormatValidationErrors(err))
		return
	}

	organizationID, err := uuid.Parse(req.OrganizationID)
	if err != nil {
		httputil.WriteProblem(w, http.StatusBadRequest, "Bad Request", "invalid organization_id format")
		return
	}

	result, err := h.switchOrganization.Execute(r.Context(), claims, appauth.SwitchOrganizationCommand{OrganizationID: organizationID})
	if err != nil {
		handleAuthError(w, err)
		return
	}

	resp := responseFromSession(result.Session)
	resp.AccessToken = result.Token
	resp.TokenType = "Bearer"
	httputil.WriteJSON(w, http.StatusOK, resp)
}

func responseFromSession(session appauth.Session) authResponse {
	resp := authResponse{
		User: userResponse{
			ID:            session.User.ID.String(),
			Name:          session.User.Name,
			Email:         session.User.Email,
			Status:        string(session.User.Status),
			IsSystemAdmin: session.User.IsSystemAdmin,
		},
		Organizations: make([]organizationResponse, len(session.Organizations)),
	}
	for i := range session.Organizations {
		resp.Organizations[i] = organizationResponse{
			ID:               session.Organizations[i].ID.String(),
			Name:             session.Organizations[i].Name,
			Domain:           session.Organizations[i].Domain,
			Workspace:        session.Organizations[i].Workspace,
			Status:           string(session.Organizations[i].Status),
			MembershipRole:   session.Organizations[i].MembershipRole,
			MembershipStatus: session.Organizations[i].MembershipStatus,
		}
	}
	if session.ActiveOrganization != nil {
		resp.ActiveOrganization = &organizationResponse{
			ID:               session.ActiveOrganization.ID.String(),
			Name:             session.ActiveOrganization.Name,
			Domain:           session.ActiveOrganization.Domain,
			Workspace:        session.ActiveOrganization.Workspace,
			Status:           string(session.ActiveOrganization.Status),
			MembershipRole:   session.ActiveOrganization.MembershipRole,
			MembershipStatus: session.ActiveOrganization.MembershipStatus,
		}
	}
	return resp
}

func handleAuthError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, appauth.ErrInvalidCredentials):
		httputil.WriteProblem(w, http.StatusUnauthorized, "Unauthorized", "invalid credentials")
	case errors.Is(err, appauth.ErrUserInactive):
		httputil.WriteProblem(w, http.StatusForbidden, "Forbidden", "user account is inactive")
	case errors.Is(err, appauth.ErrNoOrganizations):
		httputil.WriteProblem(w, http.StatusForbidden, "Forbidden", "user has no accessible organizations")
	default:
		httputil.WriteProblem(w, http.StatusInternalServerError, "Internal Server Error", "an unexpected error occurred")
	}
}
