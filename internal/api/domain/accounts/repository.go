package accounts

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jegj/linktly/internal/store"
)

type accountsRepository interface {
	GetByID(ctx context.Context, id string) (*Account, error)
}

type PostgresRepository struct {
	store *store.Store
}

func GetNewAccountRepository(store *store.Store) *PostgresRepository {
	return &PostgresRepository{
		store: store,
	}
}

func (repo *PostgresRepository) GetByID(ctx context.Context, id string) (*Account, error) {
	query := `SELECT id, name, lastname, email, role, created_at FROM linktly.accounts where id = $1`

	rows, err := repo.store.Source.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}

	account, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[Account])
	if err != nil {
		return nil, err
	}

	return &account, nil
}
