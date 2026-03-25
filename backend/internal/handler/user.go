package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/dsbraz/bud2/backend/internal/domain"
	useruc "github.com/dsbraz/bud2/backend/internal/usecase/user"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// Use case interfaces

type createUserUseCase interface {
	Execute(ctx context.Context, cmd useruc.CreateCommand) (*domain.User, error)
}

type getUserUseCase interface {
	Execute(ctx context.Context, tenantID domain.TenantID, id uuid.UUID) (*domain.User, error)
}

type listUserUseCase interface {
	Execute(ctx context.Context, cmd useruc.ListCommand) (domain.UserListResult, error)
}

type updateUserUseCase interface {
	Execute(ctx context.Context, cmd useruc.UpdateCommand) (*domain.User, error)
}

// DTOs

type createUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type updateUserRequest struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	Status string `json:"status"`
}

type userResponse struct {
	ID        string `json:"id"`
	TenantID  string `json:"tenant_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type listUserResponse struct {
	Data  []userResponse `json:"data"`
	Total int64          `json:"total"`
	Page  int            `json:"page"`
	Size  int            `json:"size"`
}

func toUserResponse(u *domain.User) userResponse {
	return userResponse{
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

// Handler

type UserHandler struct {
	create createUserUseCase
	get    getUserUseCase
	list   listUserUseCase
	update updateUserUseCase
}

func NewUserHandler(
	create createUserUseCase,
	get getUserUseCase,
	list listUserUseCase,
	update updateUserUseCase,
) *UserHandler {
	return &UserHandler{create: create, get: get, list: list, update: update}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	tenantID, err := domain.TenantIDFromContext(r.Context())
	if err != nil {
		writeProblem(w, http.StatusUnauthorized, "Unauthorized", err.Error())
		return
	}

	var req createUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeProblem(w, http.StatusBadRequest, "Bad Request", "invalid JSON body")
		return
	}

	user, err := h.create.Execute(r.Context(), useruc.CreateCommand{
		TenantID: tenantID,
		Name:     req.Name,
		Email:    req.Email,
		Role:     req.Role,
	})
	if err != nil {
		handleUserError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, toUserResponse(user))
}

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
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
		handleUserError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, toUserResponse(user))
}

func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
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

	cmd := useruc.ListCommand{TenantID: tenantID, Page: page, Size: size}
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

	items := make([]userResponse, len(result.Users))
	for i := range result.Users {
		items[i] = toUserResponse(&result.Users[i])
	}

	writeJSON(w, http.StatusOK, listUserResponse{
		Data:  items,
		Total: result.Total,
		Page:  page,
		Size:  size,
	})
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
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

	var req updateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeProblem(w, http.StatusBadRequest, "Bad Request", "invalid JSON body")
		return
	}

	user, err := h.update.Execute(r.Context(), useruc.UpdateCommand{
		TenantID: tenantID,
		ID:       id,
		Name:     req.Name,
		Email:    req.Email,
		Role:     req.Role,
		Status:   req.Status,
	})
	if err != nil {
		handleUserError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, toUserResponse(user))
}

func handleUserError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, domain.ErrUserNotFound):
		writeProblem(w, http.StatusNotFound, "Not Found", err.Error())
	case errors.Is(err, domain.ErrUserEmailExists):
		writeProblem(w, http.StatusConflict, "Conflict", err.Error())
	case errors.Is(err, domain.ErrValidation):
		writeProblem(w, http.StatusUnprocessableEntity, "Unprocessable Entity", err.Error())
	default:
		writeProblem(w, http.StatusInternalServerError, "Internal Server Error", "an unexpected error occurred")
	}
}
