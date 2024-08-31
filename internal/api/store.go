package api

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

type Store struct {
	source *pgxpool.Pool
}

func NewStore(ctx context.Context, connString string) (*Store, error) {
	pgOnce.Do(func() {
		pgInstance, dbErr = pgxpool.New(ctx, connString)
		if dbErr != nil {
			return
		}
		dbErr = pgInstance.Ping(ctx)
	})

	return &Store{source: pgInstance}, dbErr
}

func (store *Store) Ping(ctx context.Context) error {
	return store.source.Ping(ctx)
}

func (store *Store) Close() {
	store.source.Close()
}
