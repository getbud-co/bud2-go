package postgres

import (
	"context"
	"errors"

	"github.com/dsbraz/bud2/backend/internal/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	q *Queries
}

func NewUserRepository(q *Queries) *UserRepository {
	return &UserRepository{q: q}
}

func (r *UserRepository) Create(ctx context.Context, u *domain.User) (*domain.User, error) {
	row, err := r.q.CreateUser(ctx, CreateUserParams{
		ID:       u.ID,
		TenantID: u.TenantID.UUID(),
		Name:     u.Name,
		Email:    u.Email,
		Role:     string(u.Role),
		Status:   string(u.Status),
	})
	if err != nil {
		return nil, err
	}
	return userToDomain(row), nil
}

func (r *UserRepository) GetByID(ctx context.Context, tenantID domain.TenantID, id uuid.UUID) (*domain.User, error) {
	row, err := r.q.GetUserByID(ctx, GetUserByIDParams{
		ID:       id,
		TenantID: tenantID.UUID(),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}
	return userToDomain(row), nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, tenantID domain.TenantID, email string) (*domain.User, error) {
	row, err := r.q.GetUserByEmail(ctx, GetUserByEmailParams{
		Email:    email,
		TenantID: tenantID.UUID(),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}
	return userToDomain(row), nil
}

func (r *UserRepository) List(ctx context.Context, filter domain.UserListFilter) (domain.UserListResult, error) {
	limit := int32(filter.Size)
	offset := int32((filter.Page - 1) * filter.Size)
	tid := filter.TenantID.UUID()

	var rows []User
	var total int64
	var err error

	switch {
	case filter.Search != nil:
		pattern := "%" + *filter.Search + "%"
		rows, err = r.q.SearchUsers(ctx, SearchUsersParams{
			TenantID: tid,
			Name:     pattern,
			Limit:    limit,
			Offset:   offset,
		})
		if err != nil {
			return domain.UserListResult{}, err
		}
		total, err = r.q.CountSearchUsers(ctx, CountSearchUsersParams{
			TenantID: tid,
			Name:     pattern,
		})
	case filter.Status != nil:
		rows, err = r.q.ListUsersByStatus(ctx, ListUsersByStatusParams{
			TenantID: tid,
			Status:   string(*filter.Status),
			Limit:    limit,
			Offset:   offset,
		})
		if err != nil {
			return domain.UserListResult{}, err
		}
		total, err = r.q.CountUsersByStatus(ctx, CountUsersByStatusParams{
			TenantID: tid,
			Status:   string(*filter.Status),
		})
	default:
		rows, err = r.q.ListUsers(ctx, ListUsersParams{
			TenantID: tid,
			Limit:    limit,
			Offset:   offset,
		})
		if err != nil {
			return domain.UserListResult{}, err
		}
		total, err = r.q.CountUsers(ctx, tid)
	}

	if err != nil {
		return domain.UserListResult{}, err
	}

	users := make([]domain.User, len(rows))
	for i, row := range rows {
		users[i] = *userToDomain(row)
	}
	return domain.UserListResult{Users: users, Total: total}, nil
}

func (r *UserRepository) Update(ctx context.Context, u *domain.User) (*domain.User, error) {
	row, err := r.q.UpdateUser(ctx, UpdateUserParams{
		ID:       u.ID,
		TenantID: u.TenantID.UUID(),
		Name:     u.Name,
		Email:    u.Email,
		Role:     string(u.Role),
		Status:   string(u.Status),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}
	return userToDomain(row), nil
}

func userToDomain(row User) *domain.User {
	return &domain.User{
		ID:        row.ID,
		TenantID:  domain.TenantID(row.TenantID),
		Name:      row.Name,
		Email:     row.Email,
		Role:      domain.UserRole(row.Role),
		Status:    domain.UserStatus(row.Status),
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}
}
