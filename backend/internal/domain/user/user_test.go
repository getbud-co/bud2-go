package user

import (
	"testing"

	"github.com/dsbraz/bud2/backend/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestRole_IsValid(t *testing.T) {
	tests := []struct {
		name     string
		role     Role
		expected bool
	}{
		{"admin is valid", RoleAdmin, true},
		{"manager is valid", RoleManager, true},
		{"collaborator is valid", RoleCollaborator, true},
		{"invalid role", Role("invalid"), false},
		{"empty role", Role(""), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.role.IsValid())
		})
	}
}

func TestStatus_IsValid(t *testing.T) {
	tests := []struct {
		name     string
		status   Status
		expected bool
	}{
		{"active is valid", StatusActive, true},
		{"inactive is valid", StatusInactive, true},
		{"invalid status", Status("invalid"), false},
		{"empty status", Status(""), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.status.IsValid())
		})
	}
}

func TestUser_Validate(t *testing.T) {
	tests := []struct {
		name        string
		user        User
		expectedErr error
	}{
		{
			name: "valid user",
			user: User{
				Name:   "Test User",
				Email:  "test@example.com",
				Role:   RoleAdmin,
				Status: StatusActive,
			},
			expectedErr: nil,
		},
		{
			name: "missing name",
			user: User{
				Name:   "",
				Email:  "test@example.com",
				Role:   RoleAdmin,
				Status: StatusActive,
			},
			expectedErr: domain.ErrValidation,
		},
		{
			name: "missing email",
			user: User{
				Name:   "Test User",
				Email:  "",
				Role:   RoleAdmin,
				Status: StatusActive,
			},
			expectedErr: domain.ErrValidation,
		},
		{
			name: "invalid role",
			user: User{
				Name:   "Test User",
				Email:  "test@example.com",
				Role:   Role("invalid"),
				Status: StatusActive,
			},
			expectedErr: domain.ErrValidation,
		},
		{
			name: "invalid status",
			user: User{
				Name:   "Test User",
				Email:  "test@example.com",
				Role:   RoleAdmin,
				Status: Status("invalid"),
			},
			expectedErr: domain.ErrValidation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.user.Validate()
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
