package membership

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/dsbraz/bud2/backend/internal/domain"
)

type Role string

const (
	RoleAdmin        Role = "admin"
	RoleManager      Role = "manager"
	RoleCollaborator Role = "collaborator"
)

func (r Role) IsValid() bool {
	return r == RoleAdmin || r == RoleManager || r == RoleCollaborator
}

type Status string

const (
	StatusInvited  Status = "invited"
	StatusActive   Status = "active"
	StatusInactive Status = "inactive"
)

func (s Status) IsValid() bool {
	return s == StatusInvited || s == StatusActive || s == StatusInactive
}

type Membership struct {
	ID              uuid.UUID
	OrganizationID  uuid.UUID
	UserID          uuid.UUID
	Role            Role
	Status          Status
	InvitedByUserID *uuid.UUID
	JoinedAt        *time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (m *Membership) Validate() error {
	if m.OrganizationID == uuid.Nil {
		return fmt.Errorf("%w: organization_id is required", domain.ErrValidation)
	}
	if m.UserID == uuid.Nil {
		return fmt.Errorf("%w: user_id is required", domain.ErrValidation)
	}
	if !m.Role.IsValid() {
		return fmt.Errorf("%w: role must be admin, manager or collaborator", domain.ErrValidation)
	}
	if !m.Status.IsValid() {
		return fmt.Errorf("%w: status must be invited, active or inactive", domain.ErrValidation)
	}
	return nil
}

type ListByOrganizationFilter struct {
	OrganizationID uuid.UUID
	Status         *Status
	Page           int
	Size           int
}

type ListByUserFilter struct {
	UserID uuid.UUID
	Status *Status
	Page   int
	Size   int
}

type ListResult struct {
	Memberships []Membership
	Total       int64
}

type Repository interface {
	Create(ctx context.Context, membership *Membership) (*Membership, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Membership, error)
	GetByOrganizationAndUser(ctx context.Context, organizationID, userID uuid.UUID) (*Membership, error)
	ListByOrganization(ctx context.Context, filter ListByOrganizationFilter) (ListResult, error)
	ListByUser(ctx context.Context, filter ListByUserFilter) (ListResult, error)
	Update(ctx context.Context, membership *Membership) (*Membership, error)
}

var (
	ErrNotFound      = errors.New("membership not found")
	ErrAlreadyExists = errors.New("membership already exists")
)
