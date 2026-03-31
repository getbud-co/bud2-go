package organization

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/dsbraz/bud2/backend/internal/api/httputil"
	"github.com/dsbraz/bud2/backend/internal/api/validator"
	apporg "github.com/dsbraz/bud2/backend/internal/app/organization"
	"github.com/dsbraz/bud2/backend/internal/domain"
	org "github.com/dsbraz/bud2/backend/internal/domain/organization"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type createUseCase interface {
	Execute(ctx context.Context, cmd apporg.CreateCommand) (*org.Organization, error)
}

type getUseCase interface {
	Execute(ctx context.Context, id uuid.UUID) (*org.Organization, error)
}

type listUseCase interface {
	Execute(ctx context.Context, cmd apporg.ListCommand) (org.ListResult, error)
}

type updateUseCase interface {
	Execute(ctx context.Context, cmd apporg.UpdateCommand) (*org.Organization, error)
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
	Name   string `json:"name" validate:"required,min=2,max=100"`
	Slug   string `json:"slug" validate:"required,min=2,max=100,slug"`
	Status string `json:"status" validate:"omitempty,oneof=active inactive"`
}

type updateRequest struct {
	Name   string `json:"name" validate:"required,min=2,max=100"`
	Slug   string `json:"slug" validate:"required,min=2,max=100,slug"`
	Status string `json:"status" validate:"required,oneof=active inactive"`
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

func toResponse(o *org.Organization) Response {
	return Response{
		ID:        o.ID,
		Name:      o.Name,
		Slug:      o.Slug,
		Status:    string(o.Status),
		CreatedAt: o.CreatedAt,
		UpdatedAt: o.UpdatedAt,
	}
}

// Handlers

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

	result, err := h.create.Execute(r.Context(), apporg.CreateCommand{
		Name:   req.Name,
		Slug:   req.Slug,
		Status: req.Status,
	})
	if err != nil {
		handleError(w, err)
		return
	}

	httputil.WriteJSON(w, http.StatusCreated, toResponse(result))
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		httputil.WriteProblem(w, http.StatusBadRequest, "Bad Request", "invalid id format")
		return
	}

	result, err := h.get.Execute(r.Context(), id)
	if err != nil {
		handleError(w, err)
		return
	}

	httputil.WriteJSON(w, http.StatusOK, toResponse(result))
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

	cmd := apporg.ListCommand{Page: page, Size: size}
	if s := q.Get("status"); s != "" {
		cmd.Status = &s
	}

	result, err := h.list.Execute(r.Context(), cmd)
	if err != nil {
		httputil.WriteProblem(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	items := make([]Response, len(result.Organizations))
	for i := range result.Organizations {
		items[i] = toResponse(&result.Organizations[i])
	}

	httputil.WriteJSON(w, http.StatusOK, ListResponse{
		Data:  items,
		Total: result.Total,
		Page:  page,
		Size:  size,
	})
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
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

	result, err := h.update.Execute(r.Context(), apporg.UpdateCommand{
		ID:     id,
		Name:   req.Name,
		Slug:   req.Slug,
		Status: req.Status,
	})
	if err != nil {
		handleError(w, err)
		return
	}

	httputil.WriteJSON(w, http.StatusOK, toResponse(result))
}

func handleError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, org.ErrNotFound):
		httputil.WriteProblem(w, http.StatusNotFound, "Not Found", err.Error())
	case errors.Is(err, org.ErrSlugExists):
		httputil.WriteProblem(w, http.StatusConflict, "Conflict", err.Error())
	case errors.Is(err, domain.ErrValidation):
		httputil.WriteProblem(w, http.StatusUnprocessableEntity, "Unprocessable Entity", err.Error())
	default:
		httputil.WriteProblem(w, http.StatusInternalServerError, "Internal Server Error", "an unexpected error occurred")
	}
}
