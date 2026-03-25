package user

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/dsbraz/bud2/backend/internal/domain"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// Use case interfaces

type createUseCase interface {
	Execute(ctx context.Context, cmd CreateCommand) (*User, error)
}

type getUseCase interface {
	Execute(ctx context.Context, tenantID domain.TenantID, id uuid.UUID) (*User, error)
}

type listUseCase interface {
	Execute(ctx context.Context, cmd ListCommand) (ListResult, error)
}

type updateUseCase interface {
	Execute(ctx context.Context, cmd UpdateCommand) (*User, error)
}

// Handler

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
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type updateRequest struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	Status string `json:"status"`
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

func toResponse(u *User) Response {
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
		writeProblem(w, http.StatusUnauthorized, "Unauthorized", err.Error())
		return
	}

	var req createRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeProblem(w, http.StatusBadRequest, "Bad Request", "invalid JSON body")
		return
	}

	user, err := h.create.Execute(r.Context(), CreateCommand{
		TenantID: tenantID,
		Name:     req.Name,
		Email:    req.Email,
		Role:     req.Role,
	})
	if err != nil {
		handleError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, toResponse(user))
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	tenantID, err := domain.TenantIDFromContext(r.Context())
	if err != nil {
		writeProblem(w, http.StatusUnauthorized, "Unauthorized", err.Error())
		return
	}

	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		writeProblem(w, http.StatusBadRequest, "Bad Request", "invalid id format")
		return
	}

	user, err := h.get.Execute(r.Context(), tenantID, id)
	if err != nil {
		handleError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, toResponse(user))
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	tenantID, err := domain.TenantIDFromContext(r.Context())
	if err != nil {
		writeProblem(w, http.StatusUnauthorized, "Unauthorized", err.Error())
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

	cmd := ListCommand{TenantID: tenantID, Page: page, Size: size}
	if s := q.Get("status"); s != "" {
		cmd.Status = &s
	}
	if s := q.Get("search"); s != "" {
		cmd.Search = &s
	}

	result, err := h.list.Execute(r.Context(), cmd)
	if err != nil {
		writeProblem(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	items := make([]Response, len(result.Users))
	for i := range result.Users {
		items[i] = toResponse(&result.Users[i])
	}

	writeJSON(w, http.StatusOK, ListResponse{
		Data:  items,
		Total: result.Total,
		Page:  page,
		Size:  size,
	})
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	tenantID, err := domain.TenantIDFromContext(r.Context())
	if err != nil {
		writeProblem(w, http.StatusUnauthorized, "Unauthorized", err.Error())
		return
	}

	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		writeProblem(w, http.StatusBadRequest, "Bad Request", "invalid id format")
		return
	}

	var req updateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeProblem(w, http.StatusBadRequest, "Bad Request", "invalid JSON body")
		return
	}

	user, err := h.update.Execute(r.Context(), UpdateCommand{
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

	writeJSON(w, http.StatusOK, toResponse(user))
}

func handleError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, ErrNotFound):
		writeProblem(w, http.StatusNotFound, "Not Found", err.Error())
	case errors.Is(err, ErrEmailExists):
		writeProblem(w, http.StatusConflict, "Conflict", err.Error())
	case errors.Is(err, domain.ErrValidation):
		writeProblem(w, http.StatusUnprocessableEntity, "Unprocessable Entity", err.Error())
	default:
		writeProblem(w, http.StatusInternalServerError, "Internal Server Error", "an unexpected error occurred")
	}
}

// Helper functions (RFC 7807)

type problemDetail struct {
	Type   string `json:"type"`
	Title  string `json:"title"`
	Status int    `json:"status"`
	Detail string `json:"detail,omitempty"`
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeProblem(w http.ResponseWriter, status int, title, detail string) {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(problemDetail{
		Type:   "about:blank",
		Title:  title,
		Status: status,
		Detail: detail,
	})
}
