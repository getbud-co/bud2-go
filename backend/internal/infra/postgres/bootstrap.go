package postgres

import (
	"context"

	"github.com/dsbraz/bud2/backend/internal/domain/organization"
	"github.com/dsbraz/bud2/backend/internal/domain/user"
	"github.com/dsbraz/bud2/backend/internal/infra/postgres/sqlc"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TxBootstrapper struct {
	pool *pgxpool.Pool
}

func NewTxBootstrapper(pool *pgxpool.Pool) *TxBootstrapper {
	return &TxBootstrapper{pool: pool}
}

func (tb *TxBootstrapper) WithTx(ctx context.Context, fn func(orgRepo organization.Repository, userRepo user.Repository) error) error {
	tx, err := tb.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	txQ := sqlc.New(tx)
	orgRepo := NewOrgRepository(txQ)
	userRepo := NewUserRepository(txQ)

	if err := fn(orgRepo, userRepo); err != nil {
		return err
	}

	return tx.Commit(ctx)
}
