package store

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jegj/linktly/internal/config"
)

var (
	pgInstance *pgxpool.Pool
	pgOnce     sync.Once
	dbErr      error
)

type PostgresStore struct {
	Source *pgxpool.Pool
}

func NewPostgresStore(ctx context.Context, appConfig config.Config) (*PostgresStore, error) {
	pgOnce.Do(func() {
		config, err := pgxpool.ParseConfig(appConfig.GetDBConnectionString())
		config.MaxConns = appConfig.PgPoolMaxConn
		config.MinConns = appConfig.PgPoolMinConn
		config.MaxConnLifetime = appConfig.PgPoolConnLifeTime
		config.MaxConnIdleTime = appConfig.PgPoolMaxConnIdleTime
		config.HealthCheckPeriod = appConfig.PgPoolHealthCheckPeriod

		if err != nil {
			return
		}
		pgInstance, dbErr = pgxpool.NewWithConfig(ctx, config)
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
