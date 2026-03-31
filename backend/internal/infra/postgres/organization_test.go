package postgres

import (
	"context"
	"testing"

	"github.com/dsbraz/bud2/backend/internal/domain/organization"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOrgRepository_Create(t *testing.T) {
	repo := setupOrgRepo(t)
	ctx := context.Background()

	org := &organization.Organization{
		ID:     uuid.New(),
		Name:   "Test Org",
		Slug:   "test-org-" + uuid.New().String()[:8],
		Status: organization.StatusActive,
	}

	created, err := repo.Create(ctx, org)
	require.NoError(t, err)
	assert.NotNil(t, created)
	assert.Equal(t, org.Name, created.Name)
	assert.Equal(t, org.Slug, created.Slug)
	assert.Equal(t, org.Status, created.Status)
}

func TestOrgRepository_GetByID(t *testing.T) {
	repo := setupOrgRepo(t)
	ctx := context.Background()

	// Create first
	org := &organization.Organization{
		ID:     uuid.New(),
		Name:   "Test Org",
		Slug:   "test-org-" + uuid.New().String()[:8],
		Status: organization.StatusActive,
	}
	created, err := repo.Create(ctx, org)
	require.NoError(t, err)

	// Get by ID
	found, err := repo.GetByID(ctx, created.ID)
	require.NoError(t, err)
	assert.Equal(t, created.ID, found.ID)
	assert.Equal(t, created.Name, found.Name)
}

func TestOrgRepository_GetByID_NotFound(t *testing.T) {
	repo := setupOrgRepo(t)
	ctx := context.Background()

	_, err := repo.GetByID(ctx, uuid.New())
	assert.ErrorIs(t, err, organization.ErrNotFound)
}

func TestOrgRepository_GetBySlug(t *testing.T) {
	repo := setupOrgRepo(t)
	ctx := context.Background()

	org := &organization.Organization{
		ID:     uuid.New(),
		Name:   "Test Org",
		Slug:   "test-org-" + uuid.New().String()[:8],
		Status: organization.StatusActive,
	}
	created, err := repo.Create(ctx, org)
	require.NoError(t, err)

	found, err := repo.GetBySlug(ctx, created.Slug)
	require.NoError(t, err)
	assert.Equal(t, created.ID, found.ID)
}

func TestOrgRepository_GetBySlug_NotFound(t *testing.T) {
	repo := setupOrgRepo(t)
	ctx := context.Background()

	_, err := repo.GetBySlug(ctx, "non-existent-slug")
	assert.ErrorIs(t, err, organization.ErrNotFound)
}

func TestOrgRepository_Update(t *testing.T) {
	repo := setupOrgRepo(t)
	ctx := context.Background()

	org := &organization.Organization{
		ID:     uuid.New(),
		Name:   "Test Org",
		Slug:   "test-org-" + uuid.New().String()[:8],
		Status: organization.StatusActive,
	}
	created, err := repo.Create(ctx, org)
	require.NoError(t, err)

	created.Name = "Updated Name"
	updated, err := repo.Update(ctx, created)
	require.NoError(t, err)
	assert.Equal(t, "Updated Name", updated.Name)
	assert.Equal(t, created.ID, updated.ID)
}

func TestOrgRepository_Update_NotFound(t *testing.T) {
	repo := setupOrgRepo(t)
	ctx := context.Background()

	org := &organization.Organization{
		ID:     uuid.New(),
		Name:   "Test Org",
		Slug:   "test-org-" + uuid.New().String()[:8],
		Status: organization.StatusActive,
	}

	_, err := repo.Update(ctx, org)
	assert.ErrorIs(t, err, organization.ErrNotFound)
}

func TestOrgRepository_List(t *testing.T) {
	repo := setupOrgRepo(t)
	ctx := context.Background()

	// Create multiple orgs
	for i := 0; i < 3; i++ {
		org := &organization.Organization{
			ID:     uuid.New(),
			Name:   "Test Org " + string(rune('A'+i)),
			Slug:   "test-org-" + uuid.New().String()[:8],
			Status: organization.StatusActive,
		}
		_, err := repo.Create(ctx, org)
		require.NoError(t, err)
	}

	result, err := repo.List(ctx, organization.ListFilter{Page: 1, Size: 10})
	require.NoError(t, err)
	assert.True(t, result.Total >= 3)
	assert.True(t, len(result.Organizations) >= 3)
}

func TestOrgRepository_List_WithStatusFilter(t *testing.T) {
	repo := setupOrgRepo(t)
	ctx := context.Background()

	// Create active org
	activeOrg := &organization.Organization{
		ID:     uuid.New(),
		Name:   "Active Org",
		Slug:   "active-org-" + uuid.New().String()[:8],
		Status: organization.StatusActive,
	}
	_, err := repo.Create(ctx, activeOrg)
	require.NoError(t, err)

	// Create inactive org
	inactiveOrg := &organization.Organization{
		ID:     uuid.New(),
		Name:   "Inactive Org",
		Slug:   "inactive-org-" + uuid.New().String()[:8],
		Status: organization.StatusInactive,
	}
	_, err = repo.Create(ctx, inactiveOrg)
	require.NoError(t, err)

	// List only active
	status := organization.StatusActive
	result, err := repo.List(ctx, organization.ListFilter{Page: 1, Size: 10, Status: &status})
	require.NoError(t, err)

	for _, org := range result.Organizations {
		assert.Equal(t, organization.StatusActive, org.Status)
	}
}

func TestOrgRepository_CountAll(t *testing.T) {
	repo := setupOrgRepo(t)
	ctx := context.Background()

	// Get initial count
	initialCount, err := repo.CountAll(ctx)
	require.NoError(t, err)

	// Create org
	org := &organization.Organization{
		ID:     uuid.New(),
		Name:   "Test Org",
		Slug:   "test-org-" + uuid.New().String()[:8],
		Status: organization.StatusActive,
	}
	_, err = repo.Create(ctx, org)
	require.NoError(t, err)

	// Verify count increased
	newCount, err := repo.CountAll(ctx)
	require.NoError(t, err)
	assert.Equal(t, initialCount+1, newCount)
}
