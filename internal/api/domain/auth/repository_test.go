package auth

import (
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jegj/linktly/internal/test"
	"github.com/stretchr/testify/require"
)

func TestAuthRepository(t *testing.T) {
	ctx := context.Background()

	pgContainer, err := test.CreatePostgresContainer(ctx)
	require.NoError(t, err)

	t.Cleanup(func() {
		err := pgContainer.Terminate(ctx)
		require.NoError(t, err)
	})

	connextionString := pgContainer.ConnectionString

	t.Run("test", func(t *testing.T) {
		t.Cleanup(func() {
			err = pgContainer.Restore(ctx)
			require.NoError(t, err)
		})

		conn, err := pgx.Connect(ctx, connextionString)
		require.NoError(t, err)

		defer conn.Close(ctx)

		var test bool
		err = conn.QueryRow(
			context.Background(),
			"SELECT 1=1",
		).Scan(&test)
		require.NoError(t, err)
		fmt.Printf("===========>%v\n", test)
		if !test {
			t.Errorf("not true")
		}
	})
}
