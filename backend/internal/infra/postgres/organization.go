package postgres

import (
	"context"
	"errors"

	"github.com/dsbraz/bud2/backend/internal/domain/organization"
	"github.com/dsbraz/bud2/backend/internal/infra/postgres/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type orgQuerier interface {
	CreateOrganization(ctx context.Context, arg sqlc.CreateOrganizationParams) (sqlc.Organization, error)
	GetOrganizationByID(ctx context.Context, id uuid.UUID) (sqlc.Organization, error)
	GetOrganizationBySlug(ctx context.Context, slug string) (sqlc.Organization, error)
	ListOrganizations(ctx context.Context, arg sqlc.ListOrganizationsParams) ([]sqlc.Organization, error)
	ListOrganizationsByStatus(ctx context.Context, arg sqlc.ListOrganizationsByStatusParams) ([]sqlc.Organization, error)
	CountOrganizations(ctx context.Context) (int64, error)
	CountOrganizationsByStatus(ctx context.Context, status string) (int64, error)
	UpdateOrganization(ctx context.Context, arg sqlc.UpdateOrganizationParams) (sqlc.Organization, error)
}

type OrgRepository struct {
	q orgQuerier
}

func NewOrgRepository(q orgQuerier) *OrgRepository {
	return &OrgRepository{q: q}
}

func (r *OrgRepository) Create(ctx context.Context, org *organization.Organization) (*organization.Organization, error) {
	row, err := r.q.CreateOrganization(ctx, sqlc.CreateOrganizationParams{
		Name:   org.Name,
		Slug:   org.Slug,
		Status: string(org.Status),
	})
	if err != nil {
		return nil, err
	}
	return orgToDomain(row), nil
}

func (r *OrgRepository) GetByID(ctx context.Context, id uuid.UUID) (*organization.Organization, error) {
	row, err := r.q.GetOrganizationByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, organization.ErrNotFound
		}
		return nil, err
	}
	return orgToDomain(row), nil
}

func (r *OrgRepository) GetBySlug(ctx context.Context, slug string) (*organization.Organization, error) {
	row, err := r.q.GetOrganizationBySlug(ctx, slug)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, organization.ErrNotFound
		}
		return nil, err
	}
	return orgToDomain(row), nil
}

func (r *OrgRepository) List(ctx context.Context, filter organization.ListFilter) (organization.ListResult, error) {
	limit := int32(filter.Size)
	offset := int32((filter.Page - 1) * filter.Size)

	var rows []sqlc.Organization
	var total int64
	var err error

	if filter.Status != nil {
		status := string(*filter.Status)
		rows, err = r.q.ListOrganizationsByStatus(ctx, sqlc.ListOrganizationsByStatusParams{
			Status: status,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return organization.ListResult{}, err
		}
		total, err = r.q.CountOrganizationsByStatus(ctx, status)
	} else {
		rows, err = r.q.ListOrganizations(ctx, sqlc.ListOrganizationsParams{
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return organization.ListResult{}, err
		}
		total, err = r.q.CountOrganizations(ctx)
	}

	if err != nil {
		return organization.ListResult{}, err
	}

	orgs := make([]organization.Organization, len(rows))
	for i, row := range rows {
		orgs[i] = *orgToDomain(row)
	}

	return organization.ListResult{
		Organizations: orgs,
		Total:         total,
	}, nil
}

func (r *OrgRepository) Update(ctx context.Context, org *organization.Organization) (*organization.Organization, error) {
	row, err := r.q.UpdateOrganization(ctx, sqlc.UpdateOrganizationParams{
		ID:     org.ID,
		Name:   org.Name,
		Slug:   org.Slug,
		Status: string(org.Status),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, organization.ErrNotFound
		}
		return nil, err
	}
	return orgToDomain(row), nil
}

func (r *OrgRepository) CountAll(ctx context.Context) (int64, error) {
	return r.q.CountOrganizations(ctx)
}

func orgToDomain(row sqlc.Organization) *organization.Organization {
	return &organization.Organization{
		ID:        row.ID,
		Name:      row.Name,
		Slug:      row.Slug,
		Status:    organization.Status(row.Status),
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}
}
