package user

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"github.com/dsbraz/bud2/backend/internal/api/httputil"
	"github.com/dsbraz/bud2/backend/internal/api/validator"
	appuser "github.com/dsbraz/bud2/backend/internal/app/user"
	"github.com/dsbraz/bud2/backend/internal/domain"
	"github.com/dsbraz/bud2/backend/internal/domain/membership"
	usr "github.com/dsbraz/bud2/backend/internal/domain/user"
)

type createUseCase interface {
	Execute(ctx context.Context, cmd appuser.CreateCommand) (*appuser.Member, error)
}

type getUseCase interface {
	Execute(ctx context.Context, organizationID domain.TenantID, id uuid.UUID) (*appuser.Member, error)
}

type listUseCase interface {
	Execute(ctx context.Context, cmd appuser.ListCommand) (appuser.MemberListResult, error)
}

type updateUseCase interface {
	Execute(ctx context.Context, cmd appuser.UpdateCommand) (*appuser.Member, error)
}

type Handler struct {
	create createUseCase
	get    getUseCase
	list   listUseCase
	update updateUseCase
}

func NewHandler(create createUseCase, get getUseCase, list listUseCase, update updateUseCase) *Handler {
	return &Handler{create: create, get: get, list: list, update: update}
}

type createRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Role     string `json:"role" validate:"required,oneof=admin manager collaborator"`
}

type updateRequest struct {
	Name   string `json:"name" validate:"required,min=2,max=100"`
	Email  string `json:"email" validate:"required,email"`
	Role   string `json:"role" validate:"required,oneof=admin manager collaborator"`
	Status string `json:"status" validate:"required,oneof=invited active inactive"`
}

type Response struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Role          string `json:"role"`
	Status        string `json:"status"`
	IsSystemAdmin bool   `json:"is_system_admin"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

type ListResponse struct {
	Data  []Response `json:"data"`
	Total int64      `json:"total"`
	Page  int        `json:"page"`
	Size  int        `json:"size"`
}

func toResponse(m *appuser.Member) Response {
	return Response{
		ID:            m.User.ID.String(),
		Name:          m.User.Name,
		Email:         m.User.Email,
		Role:          string(m.MembershipRole),
		Status:        string(m.MembershipStatus),
		IsSystemAdmin: m.User.IsSystemAdmin,
		CreatedAt:     m.User.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:     m.User.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	organizationID, err := domain.TenantIDFromContext(r.Context())
	if err != nil {
		httputil.WriteProblem(w, http.StatusUnauthorized, "Unauthorized", err.Error())
		return
	}

	var req createRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httputil.WriteProblem(w, http.StatusBadRequest, "Bad Request", "invalid JSON body")
		return
	}
	if err := validator.Validate(req); err != nil {
		httputil.WriteProblem(w, http.StatusUnprocessableEntity, "Validation Error", validator.FormatValidationErrors(err))
		return
	}

	result, err := h.create.Execute(r.Context(), appuser.CreateCommand{OrganizationID: organizationID, Name: req.Name, Email: req.Email, Password: req.Password, Role: req.Role})
	if err != nil {
		handleError(w, err)
		return
	}

	httputil.WriteJSON(w, http.StatusCreated, toResponse(result))
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	organizationID, err := domain.TenantIDFromContext(r.Context())
	if err != nil {
		httputil.WriteProblem(w, http.StatusUnauthorized, "Unauthorized", err.Error())
		return
	}
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		httputil.WriteProblem(w, http.StatusBadRequest, "Bad Request", "invalid id format")
		return
	}

	result, err := h.get.Execute(r.Context(), organizationID, id)
	if err != nil {
		handleError(w, err)
		return
	}
	httputil.WriteJSON(w, http.StatusOK, toResponse(result))
}

const maxPageSize = 100

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	organizationID, err := domain.TenantIDFromContext(r.Context())
	if err != nil {
		httputil.WriteProblem(w, http.StatusUnauthorized, "Unauthorized", err.Error())
		return
	}
	q := r.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	size, _ := strconv.Atoi(q.Get("size"))
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 20
	}
	if size > maxPageSize {
		size = maxPageSize
	}

	cmd := appuser.ListCommand{OrganizationID: organizationID, Page: page, Size: size}
	if s := q.Get("status"); s != "" {
		cmd.Status = &s
	}

	result, err := h.list.Execute(r.Context(), cmd)
	if err != nil {
		httputil.WriteProblem(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	items := make([]Response, len(result.Members))
	for i := range result.Members {
		items[i] = toResponse(&result.Members[i])
	}
	httputil.WriteJSON(w, http.StatusOK, ListResponse{Data: items, Total: result.Total, Page: page, Size: size})
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	organizationID, err := domain.TenantIDFromContext(r.Context())
	if err != nil {
		httputil.WriteProblem(w, http.StatusUnauthorized, "Unauthorized", err.Error())
		return
	}
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		httputil.WriteProblem(w, http.StatusBadRequest, "Bad Request", "invalid id format")
		return
	}

	var req updateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httputil.WriteProblem(w, http.StatusBadRequest, "Bad Request", "invalid JSON body")
		return
	}
	if err := validator.Validate(req); err != nil {
		httputil.WriteProblem(w, http.StatusUnprocessableEntity, "Validation Error", validator.FormatValidationErrors(err))
		return
	}

	result, err := h.update.Execute(r.Context(), appuser.UpdateCommand{OrganizationID: organizationID, ID: id, Name: req.Name, Email: req.Email, Role: req.Role, Status: req.Status})
	if err != nil {
		handleError(w, err)
		return
	}
	httputil.WriteJSON(w, http.StatusOK, toResponse(result))
}

func handleError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, usr.ErrNotFound), errors.Is(err, membership.ErrNotFound):
		httputil.WriteProblem(w, http.StatusNotFound, "Not Found", err.Error())
	case errors.Is(err, usr.ErrEmailExists), errors.Is(err, membership.ErrAlreadyExists):
		httputil.WriteProblem(w, http.StatusConflict, "Conflict", err.Error())
	case errors.Is(err, domain.ErrValidation):
		httputil.WriteProblem(w, http.StatusUnprocessableEntity, "Unprocessable Entity", err.Error())
	default:
		httputil.WriteProblem(w, http.StatusInternalServerError, "Internal Server Error", "an unexpected error occurred")
	}
}
