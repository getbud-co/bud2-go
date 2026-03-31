package organization

import (
	"testing"

	"github.com/dsbraz/bud2/backend/internal/domain"
	"github.com/stretchr/testify/assert"
)

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

func TestOrganization_Validate(t *testing.T) {
	tests := []struct {
		name        string
		org         Organization
		expectedErr error
	}{
		{
			name: "valid organization",
			org: Organization{
				Name:   "Test Org",
				Slug:   "test-org",
				Status: StatusActive,
			},
			expectedErr: nil,
		},
		{
			name: "missing name",
			org: Organization{
				Name:   "",
				Slug:   "test-org",
				Status: StatusActive,
			},
			expectedErr: domain.ErrValidation,
		},
		{
			name: "missing slug",
			org: Organization{
				Name:   "Test Org",
				Slug:   "",
				Status: StatusActive,
			},
			expectedErr: domain.ErrValidation,
		},
		{
			name: "invalid status",
			org: Organization{
				Name:   "Test Org",
				Slug:   "test-org",
				Status: Status("invalid"),
			},
			expectedErr: domain.ErrValidation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.org.Validate()
			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
