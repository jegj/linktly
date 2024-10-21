package accounts

import (
	"context"
	"net/http"
	"path"
	"testing"

	"github.com/jegj/linktly/internal/api/types"
	"github.com/jegj/linktly/internal/store"
	"github.com/jegj/linktly/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccountRepository(t *testing.T) {
	ctx := context.Background()

	// TODO: Probably creating snapshot gonna depend on the test
	pgContainer, err := testutils.CreatePostgresContainer(
		ctx,
		path.Join("../../../database/testdb/up.sql"),
	)
	require.NoError(t, err)

	t.Cleanup(func() {
		err := pgContainer.Terminate(ctx)
		require.NoError(t, err)
	})

	connectionString := pgContainer.ConnectionString
	t.Run("GetByID must return NotFound API error when there is not a match by id", func(t *testing.T) {
		t.Cleanup(func() {
			err = pgContainer.Restore(ctx)
			require.NoError(t, err)
		})

		store, err := store.NewPostgresStoreForTesting(ctx, connectionString)
		require.NoError(t, err)
		defer store.Close()

		accountRepository := GetNewAccountRepository(store)
		id := "0192aeb2-ac7e-7a52-ac5a-3b3354c60187"
		response, err := accountRepository.GetByID(ctx, id)

		assert.Nil(t, response, "GetByID returned a valid record from the database instead of an error")
		if assert.NotNil(t, err) {
			assert.IsType(t, types.APIError{}, err, "Error should be of type ApiError")
			apiErr, _ := err.(types.APIError)
			assert.Equal(t, http.StatusNotFound, apiErr.StatusCode, "Error status code should be NotFound 404")
		}
	})
}
