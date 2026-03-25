package domain

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type OrganizationStatus string

const (
	OrganizationStatusActive   OrganizationStatus = "active"
	OrganizationStatusInactive OrganizationStatus = "inactive"
)

func (s OrganizationStatus) IsValid() bool {
	return s == OrganizationStatusActive || s == OrganizationStatusInactive
}

type Organization struct {
	ID        uuid.UUID
	Name      string
	Slug      string
	Status    OrganizationStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (o *Organization) Validate() error {
	if o.Name == "" {
		return fmt.Errorf("%w: name is required", ErrValidation)
	}
	if o.Slug == "" {
		return fmt.Errorf("%w: slug is required", ErrValidation)
	}
	if !o.Status.IsValid() {
		return fmt.Errorf("%w: status must be active or inactive", ErrValidation)
	}
	return nil
}

type OrganizationListFilter struct {
	Status *OrganizationStatus
	Page   int
	Size   int
}

type OrganizationListResult struct {
	Organizations []Organization
	Total         int64
}

type OrganizationRepository interface {
	Create(ctx context.Context, org *Organization) (*Organization, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Organization, error)
	GetBySlug(ctx context.Context, slug string) (*Organization, error)
	List(ctx context.Context, filter OrganizationListFilter) (OrganizationListResult, error)
	Update(ctx context.Context, org *Organization) (*Organization, error)
}

var (
	ErrValidation             = errors.New("validation error")
	ErrOrganizationNotFound   = errors.New("organization not found")
	ErrOrganizationSlugExists = errors.New("organization slug already exists")
)
