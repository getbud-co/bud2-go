package fixtures

import (
	"time"

	"github.com/dsbraz/bud2/backend/internal/domain/organization"
	"github.com/google/uuid"
)

func NewOrganization() *organization.Organization {
	return &organization.Organization{
		ID:        uuid.MustParse("550e8400-e29b-41d4-a716-446655440000"),
		Name:      "Test Organization",
		Slug:      "test-org",
		Status:    organization.StatusActive,
		CreatedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	}
}

func NewOrganizationWithName(name, slug string) *organization.Organization {
	org := NewOrganization()
	org.Name = name
	org.Slug = slug
	return org
}

func NewInactiveOrganization() *organization.Organization {
	org := NewOrganization()
	org.Status = organization.StatusInactive
	return org
}

func NewOrganizationList(count int) []organization.Organization {
	orgs := make([]organization.Organization, count)
	for i := 0; i < count; i++ {
		org := NewOrganization()
		org.ID = uuid.MustParse("550e8400-e29b-41d4-a716-446655440001")
		org.Name = "Org " + string(rune('A'+i))
		org.Slug = "org-" + string(rune('a'+i))
		orgs[i] = *org
	}
	return orgs
}
