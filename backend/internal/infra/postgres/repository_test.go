package postgres

import (
	"context"
	"os"
	"testing"

	"github.com/dsbraz/bud2/backend/internal/domain"
	"github.com/dsbraz/bud2/backend/internal/domain/organization"
	"github.com/dsbraz/bud2/backend/internal/domain/user"
	"github.com/dsbraz/bud2/backend/internal/infra/postgres/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	// Check if test database URL is set
	dbURL := os.Getenv("TEST_DATABASE_URL")
	if dbURL == "" {
		// Skip integration tests if no test database
		os.Exit(0)
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		panic(err)
	}
	testDB = pool

	// Run migrations or clean up
	cleanup(ctx, pool)

	code := m.Run()
	
	cleanup(ctx, pool)
	pool.Close()
	
	os.Exit(code)
}

func cleanup(ctx context.Context, pool *pgxpool.Pool) {
	// Clean up test data
	pool.Exec(ctx, "DELETE FROM users")
	pool.Exec(ctx, "DELETE FROM organizations")
}

func setupOrgRepo(t *testing.T) *OrgRepository {
	if testDB == nil {
		t.Skip("Skipping integration test: TEST_DATABASE_URL not set")
	}
	queries := sqlc.New(testDB)
	return NewOrgRepository(queries)
}

func setupUserRepo(t *testing.T) *UserRepository {
	if testDB == nil {
		t.Skip("Skipping integration test: TEST_DATABASE_URL not set")
	}
	queries := sqlc.New(testDB)
	return NewUserRepository(queries)
}

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

func TestUserRepository_Create(t *testing.T) {
	repo := setupUserRepo(t)
	ctx := context.Background()

	u := &user.User{
		ID:       uuid.New(),
		TenantID: domain.TenantID(uuid.New()),
		Name:     "Test User",
		Email:    "test-" + uuid.New().String()[:8] + "@example.com",
		Role:     user.RoleAdmin,
		Status:   user.StatusActive,
	}

	created, err := repo.Create(ctx, u)
	require.NoError(t, err)
	assert.NotNil(t, created)
	assert.Equal(t, u.Name, created.Name)
	assert.Equal(t, u.Email, created.Email)
}

func TestUserRepository_GetByID(t *testing.T) {
	repo := setupUserRepo(t)
	ctx := context.Background()

	tenantID := domain.TenantID(uuid.New())
	u := &user.User{
		ID:       uuid.New(),
		TenantID: tenantID,
		Name:     "Test User",
		Email:    "test-" + uuid.New().String()[:8] + "@example.com",
		Role:     user.RoleAdmin,
		Status:   user.StatusActive,
	}
	created, err := repo.Create(ctx, u)
	require.NoError(t, err)

	found, err := repo.GetByID(ctx, tenantID, created.ID)
	require.NoError(t, err)
	assert.Equal(t, created.ID, found.ID)
}

func TestUserRepository_GetByEmail(t *testing.T) {
	repo := setupUserRepo(t)
	ctx := context.Background()

	tenantID := domain.TenantID(uuid.New())
	email := "test-" + uuid.New().String()[:8] + "@example.com"
	u := &user.User{
		ID:       uuid.New(),
		TenantID: tenantID,
		Name:     "Test User",
		Email:    email,
		Role:     user.RoleAdmin,
		Status:   user.StatusActive,
	}
	created, err := repo.Create(ctx, u)
	require.NoError(t, err)

	found, err := repo.GetByEmail(ctx, tenantID, email)
	require.NoError(t, err)
	assert.Equal(t, created.ID, found.ID)
}
