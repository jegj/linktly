package testutils

import (
	"context"
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

func CreatePostgresContainer(ctx context.Context, initialScripts string) (*PostgresContainer, error) {
	dbname := "linktly_test"
	dbuser := "linktly_admin"
	options := []testcontainers.ContainerCustomizer{
		postgres.WithDatabase(dbname),
		postgres.WithUsername(dbuser),
		postgres.WithPassword("linktly_pw"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5 * time.Second)),
		postgres.WithSQLDriver("pgx"),
		postgres.WithInitScripts(initialScripts),
	}

	pgContainer, err := postgres.Run(ctx, "jegj/postgres_16_uuidv7", options...)
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
