package organization

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/dsbraz/bud2/backend/internal/domain"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// Use case interfaces

type createUseCase interface {
	Execute(ctx context.Context, cmd CreateCommand) (*Organization, error)
}

type getUseCase interface {
	Execute(ctx context.Context, id uuid.UUID) (*Organization, error)
}

type listUseCase interface {
	Execute(ctx context.Context, cmd ListCommand) (ListResult, error)
}

type updateUseCase interface {
	Execute(ctx context.Context, cmd UpdateCommand) (*Organization, error)
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
	Name   string `json:"name"`
	Slug   string `json:"slug"`
	Status string `json:"status"`
}

type updateRequest struct {
	Name   string `json:"name"`
	Slug   string `json:"slug"`
	Status string `json:"status"`
}

type Response struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ListResponse struct {
	Data  []Response `json:"data"`
	Total int64      `json:"total"`
	Page  int        `json:"page"`
	Size  int        `json:"size"`
}

func toResponse(org *Organization) Response {
	return Response{
		ID:        org.ID,
		Name:      org.Name,
		Slug:      org.Slug,
		Status:    string(org.Status),
		CreatedAt: org.CreatedAt,
		UpdatedAt: org.UpdatedAt,
	}
}

// Handlers

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req createRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeProblem(w, http.StatusBadRequest, "Bad Request", "invalid JSON body")
		return
	}

	org, err := h.create.Execute(r.Context(), CreateCommand{
		Name:   req.Name,
		Slug:   req.Slug,
		Status: req.Status,
	})
	if err != nil {
		handleError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, toResponse(org))
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		writeProblem(w, http.StatusBadRequest, "Bad Request", "invalid id format")
		return
	}

	org, err := h.get.Execute(r.Context(), id)
	if err != nil {
		handleError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, toResponse(org))
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	page, _ := strconv.Atoi(q.Get("page"))
	size, _ := strconv.Atoi(q.Get("size"))

	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 20
	}

	cmd := ListCommand{Page: page, Size: size}
	if s := q.Get("status"); s != "" {
		cmd.Status = &s
	}

	result, err := h.list.Execute(r.Context(), cmd)
	if err != nil {
		writeProblem(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	items := make([]Response, len(result.Organizations))
	for i := range result.Organizations {
		items[i] = toResponse(&result.Organizations[i])
	}

	writeJSON(w, http.StatusOK, ListResponse{
		Data:  items,
		Total: result.Total,
		Page:  page,
		Size:  size,
	})
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
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

	org, err := h.update.Execute(r.Context(), UpdateCommand{
		ID:     id,
		Name:   req.Name,
		Slug:   req.Slug,
		Status: req.Status,
	})
	if err != nil {
		handleError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, toResponse(org))
}

func handleError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, ErrNotFound):
		writeProblem(w, http.StatusNotFound, "Not Found", err.Error())
	case errors.Is(err, ErrSlugExists):
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
