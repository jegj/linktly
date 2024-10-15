package test

import (
	"context"
	"path/filepath"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type PostgresContainer struct {
	*postgres.PostgresContainer
	*pgx.Conn
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
	dbURL, err := pgContainer.ConnectionString(ctx)
	if err != nil {
		return nil, err
	}

	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		return nil, err
	}

	return &PostgresContainer{
		PostgresContainer: pgContainer,
		Conn:              conn,
	}, nil
}
