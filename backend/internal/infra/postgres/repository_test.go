package postgres

import (
	"context"
	"os"
	"testing"

	"github.com/dsbraz/bud2/backend/internal/infra/postgres/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
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
