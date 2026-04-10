package postgres

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/dsbraz/bud2/backend/internal/domain/user"
	"github.com/dsbraz/bud2/backend/internal/infra/postgres/sqlc"
)

type userQuerier interface {
	CreateUser(ctx context.Context, arg sqlc.CreateUserParams) (sqlc.CreateUserRow, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (sqlc.GetUserByIDRow, error)
	GetUserByEmail(ctx context.Context, lower string) (sqlc.GetUserByEmailRow, error)
	ListUsers(ctx context.Context, arg sqlc.ListUsersParams) ([]sqlc.ListUsersRow, error)
	ListUsersByStatus(ctx context.Context, arg sqlc.ListUsersByStatusParams) ([]sqlc.ListUsersByStatusRow, error)
	SearchUsers(ctx context.Context, arg sqlc.SearchUsersParams) ([]sqlc.SearchUsersRow, error)
	CountUsers(ctx context.Context) (int64, error)
	CountUsersByStatus(ctx context.Context, status string) (int64, error)
	CountSearchUsers(ctx context.Context, name string) (int64, error)
	UpdateUser(ctx context.Context, arg sqlc.UpdateUserParams) (sqlc.UpdateUserRow, error)
}

type UserRepository struct {
	q userQuerier
}

func NewUserRepository(q userQuerier) *UserRepository {
	return &UserRepository{q: q}
}

func (r *UserRepository) Create(ctx context.Context, u *user.User) (*user.User, error) {
	row, err := r.q.CreateUser(ctx, sqlc.CreateUserParams{
		ID:            u.ID,
		Name:          u.Name,
		Email:         u.Email,
		PasswordHash:  u.PasswordHash,
		Status:        string(u.Status),
		IsSystemAdmin: u.IsSystemAdmin,
	})
	if err != nil {
		return nil, err
	}
	return createUserRowToDomain(row), nil
}

func (r *UserRepository) GetByID(ctx context.Context, id uuid.UUID) (*user.User, error) {
	row, err := r.q.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, user.ErrNotFound
		}
		return nil, err
	}
	return getUserByIDRowToDomain(row), nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	row, err := r.q.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, user.ErrNotFound
		}
		return nil, err
	}
	return getUserByEmailRowToDomain(row), nil
}

func (r *UserRepository) List(ctx context.Context, filter user.ListFilter) (user.ListResult, error) {
	limit := int32(filter.Size)
	offset := int32((filter.Page - 1) * filter.Size)
	var usersOut []user.User
	var total int64
	var err error

	switch {
	case filter.Search != nil:
		pattern := "%" + *filter.Search + "%"
		rows, listErr := r.q.SearchUsers(ctx, sqlc.SearchUsersParams{
			Name:   pattern,
			Limit:  limit,
			Offset: offset,
		})
		if listErr != nil {
			return user.ListResult{}, listErr
		}
		usersOut = make([]user.User, len(rows))
		for i, row := range rows {
			usersOut[i] = *searchUsersRowToDomain(row)
		}
		total, err = r.q.CountSearchUsers(ctx, pattern)
	case filter.Status != nil:
		rows, listErr := r.q.ListUsersByStatus(ctx, sqlc.ListUsersByStatusParams{
			Status: string(*filter.Status),
			Limit:  limit,
			Offset: offset,
		})
		if listErr != nil {
			return user.ListResult{}, listErr
		}
		usersOut = make([]user.User, len(rows))
		for i, row := range rows {
			usersOut[i] = *listUsersByStatusRowToDomain(row)
		}
		total, err = r.q.CountUsersByStatus(ctx, string(*filter.Status))
	default:
		rows, listErr := r.q.ListUsers(ctx, sqlc.ListUsersParams{
			Limit:  limit,
			Offset: offset,
		})
		if listErr != nil {
			return user.ListResult{}, listErr
		}
		usersOut = make([]user.User, len(rows))
		for i, row := range rows {
			usersOut[i] = *listUsersRowToDomain(row)
		}
		total, err = r.q.CountUsers(ctx)
	}

	if err != nil {
		return user.ListResult{}, err
	}

	return user.ListResult{Users: usersOut, Total: total}, nil
}

func (r *UserRepository) Update(ctx context.Context, u *user.User) (*user.User, error) {
	row, err := r.q.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:            u.ID,
		Name:          u.Name,
		Email:         u.Email,
		PasswordHash:  u.PasswordHash,
		Status:        string(u.Status),
		IsSystemAdmin: u.IsSystemAdmin,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, user.ErrNotFound
		}
		return nil, err
	}
	return updateUserRowToDomain(row), nil
}

func createUserRowToDomain(row sqlc.CreateUserRow) *user.User {
	return &user.User{
		ID:            row.ID,
		Name:          row.Name,
		Email:         row.Email,
		PasswordHash:  row.PasswordHash,
		Status:        user.Status(row.Status),
		IsSystemAdmin: row.IsSystemAdmin,
		CreatedAt:     row.CreatedAt,
		UpdatedAt:     row.UpdatedAt,
	}
}

func getUserByIDRowToDomain(row sqlc.GetUserByIDRow) *user.User {
	return &user.User{ID: row.ID, Name: row.Name, Email: row.Email, PasswordHash: row.PasswordHash, Status: user.Status(row.Status), IsSystemAdmin: row.IsSystemAdmin, CreatedAt: row.CreatedAt, UpdatedAt: row.UpdatedAt}
}

func getUserByEmailRowToDomain(row sqlc.GetUserByEmailRow) *user.User {
	return &user.User{ID: row.ID, Name: row.Name, Email: row.Email, PasswordHash: row.PasswordHash, Status: user.Status(row.Status), IsSystemAdmin: row.IsSystemAdmin, CreatedAt: row.CreatedAt, UpdatedAt: row.UpdatedAt}
}

func listUsersRowToDomain(row sqlc.ListUsersRow) *user.User {
	return &user.User{ID: row.ID, Name: row.Name, Email: row.Email, PasswordHash: row.PasswordHash, Status: user.Status(row.Status), IsSystemAdmin: row.IsSystemAdmin, CreatedAt: row.CreatedAt, UpdatedAt: row.UpdatedAt}
}

func listUsersByStatusRowToDomain(row sqlc.ListUsersByStatusRow) *user.User {
	return &user.User{ID: row.ID, Name: row.Name, Email: row.Email, PasswordHash: row.PasswordHash, Status: user.Status(row.Status), IsSystemAdmin: row.IsSystemAdmin, CreatedAt: row.CreatedAt, UpdatedAt: row.UpdatedAt}
}

func searchUsersRowToDomain(row sqlc.SearchUsersRow) *user.User {
	return &user.User{ID: row.ID, Name: row.Name, Email: row.Email, PasswordHash: row.PasswordHash, Status: user.Status(row.Status), IsSystemAdmin: row.IsSystemAdmin, CreatedAt: row.CreatedAt, UpdatedAt: row.UpdatedAt}
}

func updateUserRowToDomain(row sqlc.UpdateUserRow) *user.User {
	return &user.User{ID: row.ID, Name: row.Name, Email: row.Email, PasswordHash: row.PasswordHash, Status: user.Status(row.Status), IsSystemAdmin: row.IsSystemAdmin, CreatedAt: row.CreatedAt, UpdatedAt: row.UpdatedAt}
}
