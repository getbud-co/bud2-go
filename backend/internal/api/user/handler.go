package user

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/dsbraz/bud2/backend/internal/api/httputil"
	"github.com/dsbraz/bud2/backend/internal/api/validator"
	appuser "github.com/dsbraz/bud2/backend/internal/app/user"
	"github.com/dsbraz/bud2/backend/internal/domain"
	usr "github.com/dsbraz/bud2/backend/internal/domain/user"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type createUseCase interface {
	Execute(ctx context.Context, cmd appuser.CreateCommand) (*usr.User, error)
}

type getUseCase interface {
	Execute(ctx context.Context, tenantID domain.TenantID, id uuid.UUID) (*usr.User, error)
}

type listUseCase interface {
	Execute(ctx context.Context, cmd appuser.ListCommand) (usr.ListResult, error)
}

type updateUseCase interface {
	Execute(ctx context.Context, cmd appuser.UpdateCommand) (*usr.User, error)
}

type Handler struct {
	create createUseCase
	get    getUseCase
	list   listUseCase
	update updateUseCase
}

func NewHandler(
	create createUseCase,
	get getUseCase,
	list listUseCase,
	update updateUseCase,
) *Handler {
	return &Handler{create: create, get: get, list: list, update: update}
}

// DTOs

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
	Status string `json:"status" validate:"required,oneof=active inactive"`
}

type Response struct {
	ID        string `json:"id"`
	TenantID  string `json:"tenant_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ListResponse struct {
	Data  []Response `json:"data"`
	Total int64      `json:"total"`
	Page  int        `json:"page"`
	Size  int        `json:"size"`
}

func toResponse(u *usr.User) Response {
	return Response{
		ID:        u.ID.String(),
		TenantID:  u.TenantID.String(),
		Name:      u.Name,
		Email:     u.Email,
		Role:      string(u.Role),
		Status:    string(u.Status),
		CreatedAt: u.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt: u.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}

// Handlers

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	tenantID, err := domain.TenantIDFromContext(r.Context())
	if err != nil {
		httputil.WriteProblem(w, http.StatusUnauthorized, "Unauthorized", err.Error())
		return
	}

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

	u, err := h.create.Execute(r.Context(), appuser.CreateCommand{
		TenantID: tenantID,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	})
	if err != nil {
		handleError(w, err)
		return
	}

	httputil.WriteJSON(w, http.StatusCreated, toResponse(u))
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	tenantID, err := domain.TenantIDFromContext(r.Context())
	if err != nil {
		httputil.WriteProblem(w, http.StatusUnauthorized, "Unauthorized", err.Error())
		return
	}

	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		httputil.WriteProblem(w, http.StatusBadRequest, "Bad Request", "invalid id format")
		return
	}

	u, err := h.get.Execute(r.Context(), tenantID, id)
	if err != nil {
		handleError(w, err)
		return
	}

	httputil.WriteJSON(w, http.StatusOK, toResponse(u))
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	tenantID, err := domain.TenantIDFromContext(r.Context())
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

	cmd := appuser.ListCommand{TenantID: tenantID, Page: page, Size: size}
	if s := q.Get("status"); s != "" {
		cmd.Status = &s
	}
	if s := q.Get("search"); s != "" {
		cmd.Search = &s
	}

	result, err := h.list.Execute(r.Context(), cmd)
	if err != nil {
		httputil.WriteProblem(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	items := make([]Response, len(result.Users))
	for i := range result.Users {
		items[i] = toResponse(&result.Users[i])
	}

	httputil.WriteJSON(w, http.StatusOK, ListResponse{
		Data:  items,
		Total: result.Total,
		Page:  page,
		Size:  size,
	})
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	tenantID, err := domain.TenantIDFromContext(r.Context())
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

	// Validate request format
	if err := validator.Validate(req); err != nil {
		httputil.WriteProblem(w, http.StatusUnprocessableEntity, "Validation Error", validator.FormatValidationErrors(err))
		return
	}

	u, err := h.update.Execute(r.Context(), appuser.UpdateCommand{
		TenantID: tenantID,
		ID:       id,
		Name:     req.Name,
		Email:    req.Email,
		Role:     req.Role,
		Status:   req.Status,
	})
	if err != nil {
		handleError(w, err)
		return
	}

	httputil.WriteJSON(w, http.StatusOK, toResponse(u))
}

func handleError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, usr.ErrNotFound):
		httputil.WriteProblem(w, http.StatusNotFound, "Not Found", err.Error())
	case errors.Is(err, usr.ErrEmailExists):
		httputil.WriteProblem(w, http.StatusConflict, "Conflict", err.Error())
	case errors.Is(err, domain.ErrValidation):
		httputil.WriteProblem(w, http.StatusUnprocessableEntity, "Unprocessable Entity", err.Error())
	default:
		httputil.WriteProblem(w, http.StatusInternalServerError, "Internal Server Error", "an unexpected error occurred")
	}
}
