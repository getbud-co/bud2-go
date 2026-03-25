package user

import (
	"context"
	"errors"

	"github.com/dsbraz/bud2/backend/internal/domain"
	"github.com/dsbraz/bud2/backend/internal/shared/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// userQuerier is the local interface for user queries
// (duck typing - implemented by postgres.Queries)
type userQuerier interface {
	CreateUser(ctx context.Context, arg postgres.CreateUserParams) (postgres.User, error)
	GetUserByID(ctx context.Context, arg postgres.GetUserByIDParams) (postgres.User, error)
	GetUserByEmail(ctx context.Context, arg postgres.GetUserByEmailParams) (postgres.User, error)
	ListUsers(ctx context.Context, arg postgres.ListUsersParams) ([]postgres.User, error)
	ListUsersByStatus(ctx context.Context, arg postgres.ListUsersByStatusParams) ([]postgres.User, error)
	SearchUsers(ctx context.Context, arg postgres.SearchUsersParams) ([]postgres.User, error)
	CountUsers(ctx context.Context, tenantID uuid.UUID) (int64, error)
	CountUsersByStatus(ctx context.Context, arg postgres.CountUsersByStatusParams) (int64, error)
	CountSearchUsers(ctx context.Context, arg postgres.CountSearchUsersParams) (int64, error)
	UpdateUser(ctx context.Context, arg postgres.UpdateUserParams) (postgres.User, error)
}

// RepositoryImpl implements Repository
type RepositoryImpl struct {
	q userQuerier
}

func NewRepository(q userQuerier) *RepositoryImpl {
	return &RepositoryImpl{q: q}
}

func (r *RepositoryImpl) Create(ctx context.Context, u *User) (*User, error) {
	row, err := r.q.CreateUser(ctx, postgres.CreateUserParams{
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

func (r *RepositoryImpl) GetByID(ctx context.Context, tenantID domain.TenantID, id uuid.UUID) (*User, error) {
	row, err := r.q.GetUserByID(ctx, postgres.GetUserByIDParams{
		ID:       id,
		TenantID: tenantID.UUID(),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return userToDomain(row), nil
}

func (r *RepositoryImpl) GetByEmail(ctx context.Context, tenantID domain.TenantID, email string) (*User, error) {
	row, err := r.q.GetUserByEmail(ctx, postgres.GetUserByEmailParams{
		Email:    email,
		TenantID: tenantID.UUID(),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return userToDomain(row), nil
}

func (r *RepositoryImpl) List(ctx context.Context, filter ListFilter) (ListResult, error) {
	limit := int32(filter.Size)
	offset := int32((filter.Page - 1) * filter.Size)
	tid := filter.TenantID.UUID()

	var rows []postgres.User
	var total int64
	var err error

	switch {
	case filter.Search != nil:
		pattern := "%" + *filter.Search + "%"
		rows, err = r.q.SearchUsers(ctx, postgres.SearchUsersParams{
			TenantID: tid,
			Name:     pattern,
			Limit:    limit,
			Offset:   offset,
		})
		if err != nil {
			return ListResult{}, err
		}
		total, err = r.q.CountSearchUsers(ctx, postgres.CountSearchUsersParams{
			TenantID: tid,
			Name:     pattern,
		})
	case filter.Status != nil:
		rows, err = r.q.ListUsersByStatus(ctx, postgres.ListUsersByStatusParams{
			TenantID: tid,
			Status:   string(*filter.Status),
			Limit:    limit,
			Offset:   offset,
		})
		if err != nil {
			return ListResult{}, err
		}
		total, err = r.q.CountUsersByStatus(ctx, postgres.CountUsersByStatusParams{
			TenantID: tid,
			Status:   string(*filter.Status),
		})
	default:
		rows, err = r.q.ListUsers(ctx, postgres.ListUsersParams{
			TenantID: tid,
			Limit:    limit,
			Offset:   offset,
		})
		if err != nil {
			return ListResult{}, err
		}
		total, err = r.q.CountUsers(ctx, tid)
	}

	if err != nil {
		return ListResult{}, err
	}

	users := make([]User, len(rows))
	for i, row := range rows {
		users[i] = *userToDomain(row)
	}
	return ListResult{Users: users, Total: total}, nil
}

func (r *RepositoryImpl) Update(ctx context.Context, u *User) (*User, error) {
	row, err := r.q.UpdateUser(ctx, postgres.UpdateUserParams{
		ID:       u.ID,
		TenantID: u.TenantID.UUID(),
		Name:     u.Name,
		Email:    u.Email,
		Role:     string(u.Role),
		Status:   string(u.Status),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return userToDomain(row), nil
}

func userToDomain(row postgres.User) *User {
	return &User{
		ID:        row.ID,
		TenantID:  domain.TenantID(row.TenantID),
		Name:      row.Name,
		Email:     row.Email,
		Role:      Role(row.Role),
		Status:    Status(row.Status),
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}
}
