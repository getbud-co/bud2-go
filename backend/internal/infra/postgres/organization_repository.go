package postgres

import (
	"context"
	"errors"

	"github.com/dsbraz/bud2/backend/internal/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type OrganizationRepository struct {
	q *Queries
}

func NewOrganizationRepository(q *Queries) *OrganizationRepository {
	return &OrganizationRepository{q: q}
}

func (r *OrganizationRepository) Create(ctx context.Context, org *domain.Organization) (*domain.Organization, error) {
	row, err := r.q.CreateOrganization(ctx, CreateOrganizationParams{
		Name:   org.Name,
		Slug:   org.Slug,
		Status: string(org.Status),
	})
	if err != nil {
		return nil, err
	}
	return toDomain(row), nil
}

func (r *OrganizationRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Organization, error) {
	row, err := r.q.GetOrganizationByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrOrganizationNotFound
		}
		return nil, err
	}
	return toDomain(row), nil
}

func (r *OrganizationRepository) GetBySlug(ctx context.Context, slug string) (*domain.Organization, error) {
	row, err := r.q.GetOrganizationBySlug(ctx, slug)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrOrganizationNotFound
		}
		return nil, err
	}
	return toDomain(row), nil
}

func (r *OrganizationRepository) List(ctx context.Context, filter domain.OrganizationListFilter) (domain.OrganizationListResult, error) {
	limit := int32(filter.Size)
	offset := int32((filter.Page - 1) * filter.Size)

	var rows []Organization
	var total int64
	var err error

	if filter.Status != nil {
		status := string(*filter.Status)
		rows, err = r.q.ListOrganizationsByStatus(ctx, ListOrganizationsByStatusParams{
			Status: status,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return domain.OrganizationListResult{}, err
		}
		total, err = r.q.CountOrganizationsByStatus(ctx, status)
	} else {
		rows, err = r.q.ListOrganizations(ctx, ListOrganizationsParams{
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return domain.OrganizationListResult{}, err
		}
		total, err = r.q.CountOrganizations(ctx)
	}

	if err != nil {
		return domain.OrganizationListResult{}, err
	}

	orgs := make([]domain.Organization, len(rows))
	for i, row := range rows {
		orgs[i] = *toDomain(row)
	}

	return domain.OrganizationListResult{
		Organizations: orgs,
		Total:         total,
	}, nil
}

func (r *OrganizationRepository) Update(ctx context.Context, org *domain.Organization) (*domain.Organization, error) {
	row, err := r.q.UpdateOrganization(ctx, UpdateOrganizationParams{
		ID:     org.ID,
		Name:   org.Name,
		Slug:   org.Slug,
		Status: string(org.Status),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrOrganizationNotFound
		}
		return nil, err
	}
	return toDomain(row), nil
}

func toDomain(row Organization) *domain.Organization {
	return &domain.Organization{
		ID:        row.ID,
		Name:      row.Name,
		Slug:      row.Slug,
		Status:    domain.OrganizationStatus(row.Status),
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}
}
