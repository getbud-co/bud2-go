package postgres

import (
	"context"
	"errors"

	"github.com/dsbraz/bud2/backend/internal/domain"
	"github.com/dsbraz/bud2/backend/internal/domain/user"
	"github.com/dsbraz/bud2/backend/internal/infra/postgres/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type userQuerier interface {
	CreateUser(ctx context.Context, arg sqlc.CreateUserParams) (sqlc.User, error)
	GetUserByID(ctx context.Context, arg sqlc.GetUserByIDParams) (sqlc.User, error)
	GetUserByEmail(ctx context.Context, arg sqlc.GetUserByEmailParams) (sqlc.User, error)
	GetUserByEmailForLogin(ctx context.Context, email string) (sqlc.User, error)
	ListUsers(ctx context.Context, arg sqlc.ListUsersParams) ([]sqlc.User, error)
	ListUsersByStatus(ctx context.Context, arg sqlc.ListUsersByStatusParams) ([]sqlc.User, error)
	SearchUsers(ctx context.Context, arg sqlc.SearchUsersParams) ([]sqlc.User, error)
	CountUsers(ctx context.Context, tenantID uuid.UUID) (int64, error)
	CountUsersByStatus(ctx context.Context, arg sqlc.CountUsersByStatusParams) (int64, error)
	CountSearchUsers(ctx context.Context, arg sqlc.CountSearchUsersParams) (int64, error)
	UpdateUser(ctx context.Context, arg sqlc.UpdateUserParams) (sqlc.User, error)
}

type UserRepository struct {
	q userQuerier
}

func NewUserRepository(q userQuerier) *UserRepository {
	return &UserRepository{q: q}
}

func (r *UserRepository) Create(ctx context.Context, u *user.User) (*user.User, error) {
	row, err := r.q.CreateUser(ctx, sqlc.CreateUserParams{
		ID:           u.ID,
		TenantID:     u.TenantID.UUID(),
		Name:         u.Name,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
		Role:         string(u.Role),
		Status:       string(u.Status),
	})
	if err != nil {
		return nil, err
	}
	return userToDomain(row), nil
}

func (r *UserRepository) GetByID(ctx context.Context, tenantID domain.TenantID, id uuid.UUID) (*user.User, error) {
	row, err := r.q.GetUserByID(ctx, sqlc.GetUserByIDParams{
		ID:       id,
		TenantID: tenantID.UUID(),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, user.ErrNotFound
		}
		return nil, err
	}
	return userToDomain(row), nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, tenantID domain.TenantID, email string) (*user.User, error) {
	row, err := r.q.GetUserByEmail(ctx, sqlc.GetUserByEmailParams{
		Email:    email,
		TenantID: tenantID.UUID(),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, user.ErrNotFound
		}
		return nil, err
	}
	return userToDomain(row), nil
}

func (r *UserRepository) GetByEmailForLogin(ctx context.Context, email string) (*user.User, error) {
	row, err := r.q.GetUserByEmailForLogin(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, user.ErrNotFound
		}
		return nil, err
	}
	return userToDomain(row), nil
}

func (r *UserRepository) List(ctx context.Context, filter user.ListFilter) (user.ListResult, error) {
	limit := int32(filter.Size)
	offset := int32((filter.Page - 1) * filter.Size)
	tid := filter.TenantID.UUID()

	var rows []sqlc.User
	var total int64
	var err error

	switch {
	case filter.Search != nil:
		pattern := "%" + *filter.Search + "%"
		rows, err = r.q.SearchUsers(ctx, sqlc.SearchUsersParams{
			TenantID: tid,
			Name:     pattern,
			Limit:    limit,
			Offset:   offset,
		})
		if err != nil {
			return user.ListResult{}, err
		}
		total, err = r.q.CountSearchUsers(ctx, sqlc.CountSearchUsersParams{
			TenantID: tid,
			Name:     pattern,
		})
	case filter.Status != nil:
		rows, err = r.q.ListUsersByStatus(ctx, sqlc.ListUsersByStatusParams{
			TenantID: tid,
			Status:   string(*filter.Status),
			Limit:    limit,
			Offset:   offset,
		})
		if err != nil {
			return user.ListResult{}, err
		}
		total, err = r.q.CountUsersByStatus(ctx, sqlc.CountUsersByStatusParams{
			TenantID: tid,
			Status:   string(*filter.Status),
		})
	default:
		rows, err = r.q.ListUsers(ctx, sqlc.ListUsersParams{
			TenantID: tid,
			Limit:    limit,
			Offset:   offset,
		})
		if err != nil {
			return user.ListResult{}, err
		}
		total, err = r.q.CountUsers(ctx, tid)
	}

	if err != nil {
		return user.ListResult{}, err
	}

	users := make([]user.User, len(rows))
	for i, row := range rows {
		users[i] = *userToDomain(row)
	}
	return user.ListResult{Users: users, Total: total}, nil
}

func (r *UserRepository) Update(ctx context.Context, u *user.User) (*user.User, error) {
	row, err := r.q.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:       u.ID,
		TenantID: u.TenantID.UUID(),
		Name:     u.Name,
		Email:    u.Email,
		Role:     string(u.Role),
		Status:   string(u.Status),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, user.ErrNotFound
		}
		return nil, err
	}
	return userToDomain(row), nil
}

func userToDomain(row sqlc.User) *user.User {
	return &user.User{
		ID:           row.ID,
		TenantID:     domain.TenantID(row.TenantID),
		Name:         row.Name,
		Email:        row.Email,
		PasswordHash: row.PasswordHash,
		Role:         user.Role(row.Role),
		Status:       user.Status(row.Status),
		CreatedAt:    row.CreatedAt,
		UpdatedAt:    row.UpdatedAt,
	}
}
