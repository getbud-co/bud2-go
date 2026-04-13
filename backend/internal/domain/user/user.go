package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/getbud-co/bud2/backend/internal/domain"
)

type Status string

const (
	StatusActive   Status = "active"
	StatusInactive Status = "inactive"
)

func (s Status) IsValid() bool {
	return s == StatusActive || s == StatusInactive
}

type User struct {
	ID            uuid.UUID
	Name          string
	Email         string
	PasswordHash  string
	Status        Status
	IsSystemAdmin bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (u *User) Validate() error {
	if u.Name == "" {
		return fmt.Errorf("%w: name is required", domain.ErrValidation)
	}
	if u.Email == "" {
		return fmt.Errorf("%w: email is required", domain.ErrValidation)
	}
	if !u.Status.IsValid() {
		return fmt.Errorf("%w: status must be active or inactive", domain.ErrValidation)
	}
	return nil
}

type ListFilter struct {
	Status *Status
	Search *string
	Page   int
	Size   int
}

type ListResult struct {
	Users []User
	Total int64
}

type Repository interface {
	Create(ctx context.Context, user *User) (*User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	List(ctx context.Context, filter ListFilter) (ListResult, error)
	Update(ctx context.Context, user *User) (*User, error)
}

var (
	ErrNotFound    = errors.New("user not found")
	ErrEmailExists = errors.New("email already in use")
)
