package postgres

import (
	"context"
	"errors"
	"testing"

	"github.com/dsbraz/bud2/backend/internal/domain"
	"github.com/dsbraz/bud2/backend/internal/domain/organization"
	"github.com/dsbraz/bud2/backend/internal/domain/user"
	"github.com/dsbraz/bud2/backend/internal/infra/postgres/sqlc"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupBootstrapper(t *testing.T) *TxBootstrapper {
	if testDB == nil {
		t.Skip("Skipping integration test: TEST_DATABASE_URL not set")
	}
	return NewTxBootstrapper(testDB)
}

func TestTxBootstrapper_WithTx_Success(t *testing.T) {
	bootstrapper := setupBootstrapper(t)
	ctx := context.Background()

	org := &organization.Organization{
		ID:     uuid.New(),
		Name:   "Test Org",
		Slug:   "test-org-" + uuid.New().String()[:8],
		Status: organization.StatusActive,
	}

	u := &user.User{
		ID:           uuid.New(),
		TenantID:     domain.TenantID(org.ID),
		Name:         "Admin User",
		Email:        "admin-" + uuid.New().String()[:8] + "@example.com",
		PasswordHash: "hashedpassword",
		Role:         user.RoleAdmin,
		Status:       user.StatusActive,
	}

	err := bootstrapper.WithTx(ctx, func(orgRepo organization.Repository, userRepo user.Repository) error {
		// Create organization
		_, err := orgRepo.Create(ctx, org)
		if err != nil {
			return err
		}

		// Create user
		_, err = userRepo.Create(ctx, u)
		if err != nil {
			return err
		}

		return nil
	})

	require.NoError(t, err)

	// Verify both were created
	queries := sqlc.New(testDB)
	orgRepo := NewOrgRepository(queries)
	userRepo := NewUserRepository(queries)

	foundOrg, err := orgRepo.GetByID(ctx, org.ID)
	require.NoError(t, err)
	assert.Equal(t, org.Name, foundOrg.Name)

	foundUser, err := userRepo.GetByID(ctx, u.TenantID, u.ID)
	require.NoError(t, err)
	assert.Equal(t, u.Name, foundUser.Name)
}

func TestTxBootstrapper_WithTx_RollbackOnError(t *testing.T) {
	bootstrapper := setupBootstrapper(t)
	ctx := context.Background()

	org := &organization.Organization{
		ID:     uuid.New(),
		Name:   "Test Org Rollback",
		Slug:   "test-org-rollback-" + uuid.New().String()[:8],
		Status: organization.StatusActive,
	}

	err := bootstrapper.WithTx(ctx, func(orgRepo organization.Repository, userRepo user.Repository) error {
		// Create organization
		_, err := orgRepo.Create(ctx, org)
		if err != nil {
			return err
		}

		// Return error to trigger rollback
		return errors.New("intentional error")
	})

	assert.Error(t, err)

	// Verify organization was NOT created (rolled back)
	queries := sqlc.New(testDB)
	orgRepo := NewOrgRepository(queries)

	_, err = orgRepo.GetByID(ctx, org.ID)
	assert.ErrorIs(t, err, organization.ErrNotFound)
}

func TestTxBootstrapper_WithTx_OrgRepoError(t *testing.T) {
	bootstrapper := setupBootstrapper(t)
	ctx := context.Background()

	// Try to create with invalid data (duplicate slug will fail)
	org1 := &organization.Organization{
		ID:     uuid.New(),
		Name:   "First Org",
		Slug:   "duplicate-slug-test",
		Status: organization.StatusActive,
	}

	// First create successfully
	err := bootstrapper.WithTx(ctx, func(orgRepo organization.Repository, userRepo user.Repository) error {
		_, err := orgRepo.Create(ctx, org1)
		return err
	})
	require.NoError(t, err)

	// Try to create another with same slug (should fail)
	org2 := &organization.Organization{
		ID:     uuid.New(),
		Name:   "Second Org",
		Slug:   "duplicate-slug-test",
		Status: organization.StatusActive,
	}

	err = bootstrapper.WithTx(ctx, func(orgRepo organization.Repository, userRepo user.Repository) error {
		_, err := orgRepo.Create(ctx, org2)
		return err
	})

	assert.Error(t, err)
}

func TestTxBootstrapper_WithTx_UserRepoError(t *testing.T) {
	bootstrapper := setupBootstrapper(t)
	ctx := context.Background()

	org := &organization.Organization{
		ID:     uuid.New(),
		Name:   "Test Org",
		Slug:   "test-org-user-error-" + uuid.New().String()[:8],
		Status: organization.StatusActive,
	}

	err := bootstrapper.WithTx(ctx, func(orgRepo organization.Repository, userRepo user.Repository) error {
		// Create organization successfully
		_, err := orgRepo.Create(ctx, org)
		if err != nil {
			return err
		}

		// Try to create user with invalid data (empty email should fail validation)
		u := &user.User{
			ID:           uuid.New(),
			TenantID:     domain.TenantID(org.ID),
			Name:         "",
			Email:        "", // Invalid - empty email
			PasswordHash: "hashedpassword",
			Role:         user.RoleAdmin,
			Status:       user.StatusActive,
		}
		_, err = userRepo.Create(ctx, u)
		if err != nil {
			return err
		}

		return nil
	})

	// Should error due to empty name/email
	assert.Error(t, err)

	// Organization should be rolled back
	queries := sqlc.New(testDB)
	orgRepo := NewOrgRepository(queries)

	_, err = orgRepo.GetByID(ctx, org.ID)
	assert.ErrorIs(t, err, organization.ErrNotFound)
}
