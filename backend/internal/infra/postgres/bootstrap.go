package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/dsbraz/bud2/backend/internal/domain/membership"
	"github.com/dsbraz/bud2/backend/internal/domain/organization"
	"github.com/dsbraz/bud2/backend/internal/domain/user"
	"github.com/dsbraz/bud2/backend/internal/infra/postgres/sqlc"
)

type TxBootstrapper struct {
	pool *pgxpool.Pool
}

func NewTxBootstrapper(pool *pgxpool.Pool) *TxBootstrapper {
	return &TxBootstrapper{pool: pool}
}

func (tb *TxBootstrapper) WithTx(ctx context.Context, fn func(orgRepo organization.Repository, userRepo user.Repository, membershipRepo membership.Repository) error) error {
	tx, err := tb.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	txQ := sqlc.New(tx)
	orgRepo := NewOrgRepository(txQ)
	userRepo := NewUserRepository(txQ)
	membershipRepo := NewMembershipRepository(txQ)

	if err := fn(orgRepo, userRepo, membershipRepo); err != nil {
		return err
	}

	return tx.Commit(ctx)
}
