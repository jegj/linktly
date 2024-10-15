package auth

import (
	"context"
	"testing"

	"github.com/jegj/linktly/internal/test"
	"github.com/stretchr/testify/require"
)

func TestAuthRepository(t *testing.T) {
	ctx := context.Background()

	pgContainer, err := test.CreatePostgresContainer(ctx)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate pgContainer: %s", err)
		}
	})

	conn := pgContainer.Conn
	defer conn.Close(context.Background())

	var test bool
	err = conn.QueryRow(
		context.Background(),
		"SELECT 1=1",
	).Scan(&test)
	require.NoError(t, err)
	if !test {
		t.Errorf("not true")
	}
}
