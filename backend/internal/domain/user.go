package domain

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type UserRole string

const (
	UserRoleAdmin        UserRole = "admin"
	UserRoleManager      UserRole = "manager"
	UserRoleCollaborator UserRole = "collaborator"
)

func (r UserRole) IsValid() bool {
	return r == UserRoleAdmin || r == UserRoleManager || r == UserRoleCollaborator
}

type UserStatus string

const (
	UserStatusActive   UserStatus = "active"
	UserStatusInactive UserStatus = "inactive"
)

func (s UserStatus) IsValid() bool {
	return s == UserStatusActive || s == UserStatusInactive
}

type User struct {
	ID        uuid.UUID
	TenantID  TenantID
	Name      string
	Email     string
	Role      UserRole
	Status    UserStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) Validate() error {
	if u.Name == "" {
		return fmt.Errorf("%w: name is required", ErrValidation)
	}
	if u.Email == "" {
		return fmt.Errorf("%w: email is required", ErrValidation)
	}
	if !u.Role.IsValid() {
		return fmt.Errorf("%w: role must be admin, manager or collaborator", ErrValidation)
	}
	if !u.Status.IsValid() {
		return fmt.Errorf("%w: status must be active or inactive", ErrValidation)
	}
	return nil
}

type UserListFilter struct {
	TenantID TenantID
	Status   *UserStatus
	Search   *string
	Page     int
	Size     int
}

type UserListResult struct {
	Users []User
	Total int64
}

type UserRepository interface {
	Create(ctx context.Context, user *User) (*User, error)
	GetByID(ctx context.Context, tenantID TenantID, id uuid.UUID) (*User, error)
	GetByEmail(ctx context.Context, tenantID TenantID, email string) (*User, error)
	List(ctx context.Context, filter UserListFilter) (UserListResult, error)
	Update(ctx context.Context, user *User) (*User, error)
}

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrUserEmailExists = errors.New("email already in use")
)
