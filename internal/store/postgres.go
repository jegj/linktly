package store

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	pgInstance *pgxpool.Pool
	pgOnce     sync.Once
	dbErr      error
)

type PostgresStore struct {
	Source *pgxpool.Pool
}

func NewPostgresStore(ctx context.Context, connString string) (*PostgresStore, error) {
	pgOnce.Do(func() {
		pgInstance, dbErr = pgxpool.New(ctx, connString)
		// TODO: DEFINE POOL CONFIGURATION
		if dbErr != nil {
			return
		}
		dbErr = pgInstance.Ping(ctx)
	})

	return &PostgresStore{Source: pgInstance}, dbErr
}

func NewPostgresStoreForTesting(ctx context.Context, connString string) (*PostgresStore, error) {
	pgInstance, dbErr = pgxpool.New(ctx, connString)
	return &PostgresStore{Source: pgInstance}, dbErr
}

func (store *PostgresStore) Ping(ctx context.Context) error {
	return store.Source.Ping(ctx)
}

func (store *PostgresStore) Close() {
	store.Source.Close()
}
