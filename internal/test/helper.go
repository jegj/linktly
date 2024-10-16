package test

import (
	"context"
	"path/filepath"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type PostgresContainer struct {
	*postgres.PostgresContainer
	ConnectionString string
}

func CreatePostgresContainer(ctx context.Context) (*PostgresContainer, error) {
	pgContainer, err := postgres.Run(ctx, "jegj/postgres_16_uuidv7",
		postgres.WithInitScripts(filepath.Join("..", "..", "..", "database", "testdata", "up.sql")),
		postgres.WithDatabase("linktly_test"),
		postgres.WithUsername("linktly_admin"),
		postgres.WithPassword("linktly_pw"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
		postgres.WithSQLDriver("pgx"),
	)
	if err != nil {
		return nil, err
	}

	err = pgContainer.Snapshot(ctx, postgres.WithSnapshotName("linktly_test_snapshot"))
	if err != nil {
		return nil, err
	}

	connectionString, err := pgContainer.ConnectionString(ctx)
	if err != nil {
		return nil, err
	}

	return &PostgresContainer{
		PostgresContainer: pgContainer,
		ConnectionString:  connectionString,
	}, nil
}
