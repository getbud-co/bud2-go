package postgres

import (
	"context"
	"testing"

	"github.com/dsbraz/bud2/backend/internal/domain"
	"github.com/dsbraz/bud2/backend/internal/domain/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserRepository_Create(t *testing.T) {
	repo := setupUserRepo(t)
	ctx := context.Background()

	u := &user.User{
		ID:           uuid.New(),
		TenantID:     domain.TenantID(uuid.New()),
		Name:         "Test User",
		Email:        "test-" + uuid.New().String()[:8] + "@example.com",
		PasswordHash: "hashedpassword",
		Role:         user.RoleAdmin,
		Status:       user.StatusActive,
	}

	created, err := repo.Create(ctx, u)
	require.NoError(t, err)
	assert.NotNil(t, created)
	assert.Equal(t, u.Name, created.Name)
	assert.Equal(t, u.Email, created.Email)
	assert.Equal(t, u.PasswordHash, created.PasswordHash)
}

func TestUserRepository_GetByID(t *testing.T) {
	repo := setupUserRepo(t)
	ctx := context.Background()

	tenantID := domain.TenantID(uuid.New())
	u := &user.User{
		ID:           uuid.New(),
		TenantID:     tenantID,
		Name:         "Test User",
		Email:        "test-" + uuid.New().String()[:8] + "@example.com",
		PasswordHash: "hashedpassword",
		Role:         user.RoleAdmin,
		Status:       user.StatusActive,
	}
	created, err := repo.Create(ctx, u)
	require.NoError(t, err)

	found, err := repo.GetByID(ctx, tenantID, created.ID)
	require.NoError(t, err)
	assert.Equal(t, created.ID, found.ID)
	assert.Equal(t, created.Name, found.Name)
}

func TestUserRepository_GetByID_NotFound(t *testing.T) {
	repo := setupUserRepo(t)
	ctx := context.Background()

	tenantID := domain.TenantID(uuid.New())
	_, err := repo.GetByID(ctx, tenantID, uuid.New())
	assert.ErrorIs(t, err, user.ErrNotFound)
}

func TestUserRepository_GetByID_WrongTenant(t *testing.T) {
	repo := setupUserRepo(t)
	ctx := context.Background()

	tenantID := domain.TenantID(uuid.New())
	u := &user.User{
		ID:           uuid.New(),
		TenantID:     tenantID,
		Name:         "Test User",
		Email:        "test-" + uuid.New().String()[:8] + "@example.com",
		PasswordHash: "hashedpassword",
		Role:         user.RoleAdmin,
		Status:       user.StatusActive,
	}
	created, err := repo.Create(ctx, u)
	require.NoError(t, err)

	// Try to find with wrong tenant
	wrongTenantID := domain.TenantID(uuid.New())
	_, err = repo.GetByID(ctx, wrongTenantID, created.ID)
	assert.ErrorIs(t, err, user.ErrNotFound)
}

func TestUserRepository_GetByEmail(t *testing.T) {
	repo := setupUserRepo(t)
	ctx := context.Background()

	tenantID := domain.TenantID(uuid.New())
	email := "test-" + uuid.New().String()[:8] + "@example.com"
	u := &user.User{
		ID:           uuid.New(),
		TenantID:     tenantID,
		Name:         "Test User",
		Email:        email,
		PasswordHash: "hashedpassword",
		Role:         user.RoleAdmin,
		Status:       user.StatusActive,
	}
	created, err := repo.Create(ctx, u)
	require.NoError(t, err)

	found, err := repo.GetByEmail(ctx, tenantID, email)
	require.NoError(t, err)
	assert.Equal(t, created.ID, found.ID)
	assert.Equal(t, email, found.Email)
}

func TestUserRepository_GetByEmail_NotFound(t *testing.T) {
	repo := setupUserRepo(t)
	ctx := context.Background()

	tenantID := domain.TenantID(uuid.New())
	_, err := repo.GetByEmail(ctx, tenantID, "nonexistent@example.com")
	assert.ErrorIs(t, err, user.ErrNotFound)
}

func TestUserRepository_GetByEmailForLogin(t *testing.T) {
	repo := setupUserRepo(t)
	ctx := context.Background()

	tenantID := domain.TenantID(uuid.New())
	email := "test-" + uuid.New().String()[:8] + "@example.com"
	u := &user.User{
		ID:           uuid.New(),
		TenantID:     tenantID,
		Name:         "Test User",
		Email:        email,
		PasswordHash: "hashedpassword",
		Role:         user.RoleAdmin,
		Status:       user.StatusActive,
	}
	created, err := repo.Create(ctx, u)
	require.NoError(t, err)

	// GetByEmailForLogin should find user regardless of tenant (for login)
	found, err := repo.GetByEmailForLogin(ctx, email)
	require.NoError(t, err)
	assert.Equal(t, created.ID, found.ID)
}

func TestUserRepository_Update(t *testing.T) {
	repo := setupUserRepo(t)
	ctx := context.Background()

	tenantID := domain.TenantID(uuid.New())
	u := &user.User{
		ID:           uuid.New(),
		TenantID:     tenantID,
		Name:         "Test User",
		Email:        "test-" + uuid.New().String()[:8] + "@example.com",
		PasswordHash: "hashedpassword",
		Role:         user.RoleAdmin,
		Status:       user.StatusActive,
	}
	created, err := repo.Create(ctx, u)
	require.NoError(t, err)

	created.Name = "Updated Name"
	updated, err := repo.Update(ctx, created)
	require.NoError(t, err)
	assert.Equal(t, "Updated Name", updated.Name)
	assert.Equal(t, created.ID, updated.ID)
}

func TestUserRepository_Update_NotFound(t *testing.T) {
	repo := setupUserRepo(t)
	ctx := context.Background()

	tenantID := domain.TenantID(uuid.New())
	u := &user.User{
		ID:           uuid.New(),
		TenantID:     tenantID,
		Name:         "Test User",
		Email:        "test-" + uuid.New().String()[:8] + "@example.com",
		PasswordHash: "hashedpassword",
		Role:         user.RoleAdmin,
		Status:       user.StatusActive,
	}

	_, err := repo.Update(ctx, u)
	assert.ErrorIs(t, err, user.ErrNotFound)
}

func TestUserRepository_List(t *testing.T) {
	repo := setupUserRepo(t)
	ctx := context.Background()

	tenantID := domain.TenantID(uuid.New())

	// Create multiple users
	for i := 0; i < 3; i++ {
		u := &user.User{
			ID:           uuid.New(),
			TenantID:     tenantID,
			Name:         "Test User " + string(rune('A'+i)),
			Email:        "test-" + uuid.New().String()[:8] + "@example.com",
			PasswordHash: "hashedpassword",
			Role:         user.RoleAdmin,
			Status:       user.StatusActive,
		}
		_, err := repo.Create(ctx, u)
		require.NoError(t, err)
	}

	result, err := repo.List(ctx, user.ListFilter{TenantID: tenantID, Page: 1, Size: 10})
	require.NoError(t, err)
	assert.True(t, result.Total >= 3)
	assert.True(t, len(result.Users) >= 3)
}

func TestUserRepository_List_WithStatusFilter(t *testing.T) {
	repo := setupUserRepo(t)
	ctx := context.Background()

	tenantID := domain.TenantID(uuid.New())

	// Create active user
	activeUser := &user.User{
		ID:           uuid.New(),
		TenantID:     tenantID,
		Name:         "Active User",
		Email:        "active-" + uuid.New().String()[:8] + "@example.com",
		PasswordHash: "hashedpassword",
		Role:         user.RoleAdmin,
		Status:       user.StatusActive,
	}
	_, err := repo.Create(ctx, activeUser)
	require.NoError(t, err)

	// Create inactive user
	inactiveUser := &user.User{
		ID:           uuid.New(),
		TenantID:     tenantID,
		Name:         "Inactive User",
		Email:        "inactive-" + uuid.New().String()[:8] + "@example.com",
		PasswordHash: "hashedpassword",
		Role:         user.RoleAdmin,
		Status:       user.StatusInactive,
	}
	_, err = repo.Create(ctx, inactiveUser)
	require.NoError(t, err)

	// List only active
	status := user.StatusActive
	result, err := repo.List(ctx, user.ListFilter{TenantID: tenantID, Page: 1, Size: 10, Status: &status})
	require.NoError(t, err)

	for _, u := range result.Users {
		assert.Equal(t, user.StatusActive, u.Status)
	}
}

func TestUserRepository_List_WithSearch(t *testing.T) {
	repo := setupUserRepo(t)
	ctx := context.Background()

	tenantID := domain.TenantID(uuid.New())

	// Create user with specific name
	u := &user.User{
		ID:           uuid.New(),
		TenantID:     tenantID,
		Name:         "John Doe",
		Email:        "john-" + uuid.New().String()[:8] + "@example.com",
		PasswordHash: "hashedpassword",
		Role:         user.RoleAdmin,
		Status:       user.StatusActive,
	}
	_, err := repo.Create(ctx, u)
	require.NoError(t, err)

	// Search for user
	search := "John"
	result, err := repo.List(ctx, user.ListFilter{TenantID: tenantID, Page: 1, Size: 10, Search: &search})
	require.NoError(t, err)
	assert.True(t, result.Total >= 1)
}
