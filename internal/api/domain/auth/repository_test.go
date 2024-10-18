package auth

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

	connectionString := pgContainer.ConnectionString

	t.Run("Login must return NotFound API error when the email does not match", func(t *testing.T) {
		t.Cleanup(func() {
			err = pgContainer.Restore(ctx)
			require.NoError(t, err)
		})

		store, err := store.NewPostgresStore(ctx, connectionString)
		require.NoError(t, err)

		defer store.Source.Close()

		accountRepository := GetNewAuthRepository(store)
		email := "fake@email.com"
		password := "test_fake_password"
		response, err := accountRepository.Login(ctx, email, password)

		assert.Nil(t, response, "Login returned a valid record from the database")
		if assert.NotNil(t, err) {
			assert.IsType(t, types.APIError{}, err, "Error should be of type ApiError")
			apiErr, _ := err.(types.APIError)
			assert.Equal(t, http.StatusNotFound, apiErr.StatusCode, "Error status code should be 404")
		}
	})
}
