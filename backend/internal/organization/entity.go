package organization

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dsbraz/bud2/backend/internal/domain"
	"github.com/google/uuid"
)

type Status string

const (
	StatusActive   Status = "active"
	StatusInactive Status = "inactive"
)

func (s Status) IsValid() bool {
	return s == StatusActive || s == StatusInactive
}

type Organization struct {
	ID        uuid.UUID
	Name      string
	Slug      string
	Status    Status
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (o *Organization) Validate() error {
	if o.Name == "" {
		return fmt.Errorf("%w: name is required", domain.ErrValidation)
	}
	if o.Slug == "" {
		return fmt.Errorf("%w: slug is required", domain.ErrValidation)
	}
	if !o.Status.IsValid() {
		return fmt.Errorf("%w: status must be active or inactive", domain.ErrValidation)
	}
	return nil
}

type ListFilter struct {
	Status *Status
	Page   int
	Size   int
}

type ListResult struct {
	Organizations []Organization
	Total         int64
}

type Repository interface {
	Create(ctx context.Context, org *Organization) (*Organization, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Organization, error)
	GetBySlug(ctx context.Context, slug string) (*Organization, error)
	List(ctx context.Context, filter ListFilter) (ListResult, error)
	Update(ctx context.Context, org *Organization) (*Organization, error)
}

var (
	ErrNotFound   = errors.New("organization not found")
	ErrSlugExists = errors.New("organization slug already exists")
)
