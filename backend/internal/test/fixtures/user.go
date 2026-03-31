package fixtures

import (
	"time"

	"github.com/dsbraz/bud2/backend/internal/domain"
	"github.com/dsbraz/bud2/backend/internal/domain/user"
	"github.com/google/uuid"
)

func NewUser() *user.User {
	return &user.User{
		ID:        uuid.MustParse("660e8400-e29b-41d4-a716-446655440000"),
		TenantID:  domain.TenantID(uuid.MustParse("550e8400-e29b-41d4-a716-446655440000")),
		Name:      "Test User",
		Email:     "test@example.com",
		Role:      user.RoleAdmin,
		Status:    user.StatusActive,
		CreatedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	}
}

func NewUserWithEmail(email string) *user.User {
	u := NewUser()
	u.Email = email
	return u
}

func NewManager() *user.User {
	u := NewUser()
	u.Role = user.RoleManager
	return u
}

func NewCollaborator() *user.User {
	u := NewUser()
	u.Role = user.RoleCollaborator
	return u
}

func NewInactiveUser() *user.User {
	u := NewUser()
	u.Status = user.StatusInactive
	return u
}

func NewUserList(count int) []user.User {
	users := make([]user.User, count)
	for i := 0; i < count; i++ {
		u := NewUser()
		u.ID = uuid.MustParse("660e8400-e29b-41d4-a716-446655440001")
		u.Email = "user" + string(rune('0'+i)) + "@example.com"
		users[i] = *u
	}
	return users
}
