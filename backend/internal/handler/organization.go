package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/dsbraz/bud2/backend/internal/domain"
	orguc "github.com/dsbraz/bud2/backend/internal/usecase/organization"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// --- Use case interfaces ---

type createOrgUseCase interface {
	Execute(ctx context.Context, cmd orguc.CreateCommand) (*domain.Organization, error)
}

type getOrgUseCase interface {
	Execute(ctx context.Context, id uuid.UUID) (*domain.Organization, error)
}

type listOrgUseCase interface {
	Execute(ctx context.Context, cmd orguc.ListCommand) (domain.OrganizationListResult, error)
}

type updateOrgUseCase interface {
	Execute(ctx context.Context, cmd orguc.UpdateCommand) (*domain.Organization, error)
}

// --- Handler ---

type OrganizationHandler struct {
	create createOrgUseCase
	get    getOrgUseCase
	list   listOrgUseCase
	update updateOrgUseCase
}

func NewOrganizationHandler(
	create createOrgUseCase,
	get getOrgUseCase,
	list listOrgUseCase,
	update updateOrgUseCase,
) *OrganizationHandler {
	return &OrganizationHandler{create: create, get: get, list: list, update: update}
}

// --- DTOs ---

type createOrgRequest struct {
	Name   string `json:"name"`
	Slug   string `json:"slug"`
	Status string `json:"status"`
}

type updateOrgRequest struct {
	Name   string `json:"name"`
	Slug   string `json:"slug"`
	Status string `json:"status"`
}

type orgResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type listOrgResponse struct {
	Data  []orgResponse `json:"data"`
	Total int64         `json:"total"`
	Page  int           `json:"page"`
	Size  int           `json:"size"`
}

func toOrgResponse(org *domain.Organization) orgResponse {
	return orgResponse{
		ID:        org.ID,
		Name:      org.Name,
		Slug:      org.Slug,
		Status:    string(org.Status),
		CreatedAt: org.CreatedAt,
		UpdatedAt: org.UpdatedAt,
	}
}

// --- Handlers ---

func (h *OrganizationHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req createOrgRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeProblem(w, http.StatusBadRequest, "Bad Request", "invalid JSON body")
		return
	}

	org, err := h.create.Execute(r.Context(), orguc.CreateCommand{
		Name:   req.Name,
		Slug:   req.Slug,
		Status: req.Status,
	})
	if err != nil {
		handleOrgError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, toOrgResponse(org))
}

func (h *OrganizationHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		writeProblem(w, http.StatusBadRequest, "Bad Request", "invalid id format")
		return
	}

	org, err := h.get.Execute(r.Context(), id)
	if err != nil {
		handleOrgError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, toOrgResponse(org))
}

func (h *OrganizationHandler) List(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	page, _ := strconv.Atoi(q.Get("page"))
	size, _ := strconv.Atoi(q.Get("size"))

	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 20
	}

	cmd := orguc.ListCommand{Page: page, Size: size}
	if s := q.Get("status"); s != "" {
		cmd.Status = &s
	}

	result, err := h.list.Execute(r.Context(), cmd)
	if err != nil {
		writeProblem(w, http.StatusInternalServerError, "Internal Server Error", err.Error())
		return
	}

	items := make([]orgResponse, len(result.Organizations))
	for i := range result.Organizations {
		items[i] = toOrgResponse(&result.Organizations[i])
	}

	writeJSON(w, http.StatusOK, listOrgResponse{
		Data:  items,
		Total: result.Total,
		Page:  page,
		Size:  size,
	})
}

func (h *OrganizationHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		writeProblem(w, http.StatusBadRequest, "Bad Request", "invalid id format")
		return
	}

	var req updateOrgRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeProblem(w, http.StatusBadRequest, "Bad Request", "invalid JSON body")
		return
	}

	org, err := h.update.Execute(r.Context(), orguc.UpdateCommand{
		ID:     id,
		Name:   req.Name,
		Slug:   req.Slug,
		Status: req.Status,
	})
	if err != nil {
		handleOrgError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, toOrgResponse(org))
}

func handleOrgError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, domain.ErrOrganizationNotFound):
		writeProblem(w, http.StatusNotFound, "Not Found", err.Error())
	case errors.Is(err, domain.ErrOrganizationSlugExists):
		writeProblem(w, http.StatusConflict, "Conflict", err.Error())
	case errors.Is(err, domain.ErrValidation):
		writeProblem(w, http.StatusUnprocessableEntity, "Unprocessable Entity", err.Error())
	default:
		writeProblem(w, http.StatusInternalServerError, "Internal Server Error", "an unexpected error occurred")
	}
}
