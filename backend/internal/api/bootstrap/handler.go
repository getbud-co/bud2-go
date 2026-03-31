package bootstrap

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/dsbraz/bud2/backend/internal/api/httputil"
	"github.com/dsbraz/bud2/backend/internal/api/validator"
	appbootstrap "github.com/dsbraz/bud2/backend/internal/app/bootstrap"
)

type createRequest struct {
	OrganizationName string `json:"organization_name" validate:"required,min=2,max=100"`
	OrganizationSlug string `json:"organization_slug" validate:"required,min=2,max=100,slug"`
	AdminName        string `json:"admin_name" validate:"required,min=2,max=100"`
	AdminEmail       string `json:"admin_email" validate:"required,email"`
	AdminPassword    string `json:"admin_password" validate:"required,min=8"`
}

type Response struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	Organization struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"organization"`
	Admin struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
		Role  string `json:"role"`
	} `json:"admin"`
}

type bootstrapUseCase interface {
	Execute(ctx context.Context, cmd appbootstrap.Command) (*appbootstrap.Result, error)
}

type Handler struct {
	uc bootstrapUseCase
}

func NewHandler(uc bootstrapUseCase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req createRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httputil.WriteProblem(w, http.StatusBadRequest, "Bad Request", "invalid JSON body")
		return
	}

	// Validate request format
	if err := validator.Validate(req); err != nil {
		httputil.WriteProblem(w, http.StatusUnprocessableEntity, "Validation Error", validator.FormatValidationErrors(err))
		return
	}

	result, err := h.uc.Execute(r.Context(), appbootstrap.Command{
		OrganizationName: req.OrganizationName,
		OrganizationSlug: req.OrganizationSlug,
		AdminName:        req.AdminName,
		AdminEmail:       req.AdminEmail,
		AdminPassword:    req.AdminPassword,
	})
	if err != nil {
		switch {
		case errors.Is(err, appbootstrap.ErrAlreadyBootstrapped):
			httputil.WriteProblem(w, http.StatusConflict, "Conflict", err.Error())
		default:
			httputil.WriteProblem(w, http.StatusInternalServerError, "Internal Server Error", "an unexpected error occurred")
		}
		return
	}

	resp := Response{AccessToken: result.AccessToken, TokenType: "Bearer"}
	resp.Organization.ID = result.Organization.ID.String()
	resp.Organization.Name = result.Organization.Name
	resp.Organization.Slug = result.Organization.Slug
	resp.Admin.ID = result.Admin.ID.String()
	resp.Admin.Name = result.Admin.Name
	resp.Admin.Email = result.Admin.Email
	resp.Admin.Role = string(result.Admin.Role)

	httputil.WriteJSON(w, http.StatusCreated, resp)
}
