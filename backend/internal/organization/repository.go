package organization

import (
	"context"
	"errors"

	"github.com/dsbraz/bud2/backend/internal/shared/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// orgQuerier is the local interface for organization queries
// (duck typing - implemented by postgres.Queries)
type orgQuerier interface {
	CreateOrganization(ctx context.Context, arg postgres.CreateOrganizationParams) (postgres.Organization, error)
	GetOrganizationByID(ctx context.Context, id uuid.UUID) (postgres.Organization, error)
	GetOrganizationBySlug(ctx context.Context, slug string) (postgres.Organization, error)
	ListOrganizations(ctx context.Context, arg postgres.ListOrganizationsParams) ([]postgres.Organization, error)
	ListOrganizationsByStatus(ctx context.Context, arg postgres.ListOrganizationsByStatusParams) ([]postgres.Organization, error)
	CountOrganizations(ctx context.Context) (int64, error)
	CountOrganizationsByStatus(ctx context.Context, status string) (int64, error)
	UpdateOrganization(ctx context.Context, arg postgres.UpdateOrganizationParams) (postgres.Organization, error)
}

// RepositoryImpl implements Repository
type RepositoryImpl struct {
	q orgQuerier
}

func NewRepository(q orgQuerier) *RepositoryImpl {
	return &RepositoryImpl{q: q}
}

func (r *RepositoryImpl) Create(ctx context.Context, org *Organization) (*Organization, error) {
	row, err := r.q.CreateOrganization(ctx, postgres.CreateOrganizationParams{
		Name:   org.Name,
		Slug:   org.Slug,
		Status: string(org.Status),
	})
	if err != nil {
		return nil, err
	}
	return toDomain(row), nil
}

func (r *RepositoryImpl) GetByID(ctx context.Context, id uuid.UUID) (*Organization, error) {
	row, err := r.q.GetOrganizationByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return toDomain(row), nil
}

func (r *RepositoryImpl) GetBySlug(ctx context.Context, slug string) (*Organization, error) {
	row, err := r.q.GetOrganizationBySlug(ctx, slug)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return toDomain(row), nil
}

func (r *RepositoryImpl) List(ctx context.Context, filter ListFilter) (ListResult, error) {
	limit := int32(filter.Size)
	offset := int32((filter.Page - 1) * filter.Size)

	var rows []postgres.Organization
	var total int64
	var err error

	if filter.Status != nil {
		status := string(*filter.Status)
		rows, err = r.q.ListOrganizationsByStatus(ctx, postgres.ListOrganizationsByStatusParams{
			Status: status,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return ListResult{}, err
		}
		total, err = r.q.CountOrganizationsByStatus(ctx, status)
	} else {
		rows, err = r.q.ListOrganizations(ctx, postgres.ListOrganizationsParams{
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return ListResult{}, err
		}
		total, err = r.q.CountOrganizations(ctx)
	}

	if err != nil {
		return ListResult{}, err
	}

	orgs := make([]Organization, len(rows))
	for i, row := range rows {
		orgs[i] = *toDomain(row)
	}

	return ListResult{
		Organizations: orgs,
		Total:         total,
	}, nil
}

func (r *RepositoryImpl) Update(ctx context.Context, org *Organization) (*Organization, error) {
	row, err := r.q.UpdateOrganization(ctx, postgres.UpdateOrganizationParams{
		ID:     org.ID,
		Name:   org.Name,
		Slug:   org.Slug,
		Status: string(org.Status),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return toDomain(row), nil
}

func toDomain(row postgres.Organization) *Organization {
	return &Organization{
		ID:        row.ID,
		Name:      row.Name,
		Slug:      row.Slug,
		Status:    Status(row.Status),
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}
}
