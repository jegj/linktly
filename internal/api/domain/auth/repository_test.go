package auth

import (
	"context"
	"fmt"
	"path"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jegj/linktly/internal/testutils"
	"github.com/stretchr/testify/require"
)

func TestAuthRepository(t *testing.T) {
	ctx := context.Background()

	pgContainer, err := testutils.CreatePostgresContainer(
		ctx,
		path.Join("../../../database/testdb/up.sql"),
	)
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

		var id string
		err = conn.QueryRow(
			ctx,
			"SELECT id from linktly.accounts WHERE email = 'jegj57@gmail.com'",
		).Scan(&id)
		require.NoError(t, err)
		fmt.Printf("===========>%v\n", id)
		if len(id) < 1 {
			t.Errorf("not true")
		}
	})
}
