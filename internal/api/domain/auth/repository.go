package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/jegj/linktly/internal/api/domain/accounts"
	linktlyError "github.com/jegj/linktly/internal/api/error"
	"github.com/jegj/linktly/internal/api/types"
	"github.com/jegj/linktly/internal/store"
	"golang.org/x/crypto/bcrypt"
)

type authRepository interface {
	Login(ctx context.Context, email string, password string) (*accounts.Account, error)
	UpdateRefreshTokenJtiByUserId(ctx context.Context, jti string, accountId string) error
	UpdateRefreshTokenJtiByUserIdAndJti(ctx context.Context, accountId string, jti string, newJti string) error
}

type PostgresRepository struct {
	store *store.PostgresStore
}

func GetNewAuthRepository(store *store.PostgresStore) *PostgresRepository {
	return &PostgresRepository{
		store: store,
	}
}

func (repo *PostgresRepository) Login(ctx context.Context, email string, password string) (*accounts.Account, error) {
	query := `SELECT id, password, name, lastname, role, created_at FROM linktly.accounts WHERE email = $1`

	var id string
	var dbPassword string
	var name string
	var lastname string
	var role int
	var createdAt time.Time

	err := repo.store.Source.QueryRow(ctx, query, email).Scan(&id, &dbPassword, &name, &lastname, &role, &createdAt)
	if err != nil {
		return nil, linktlyError.PostgresFormatting(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))
	if err != nil {
		return nil, types.APIError{
			Msg:        err.Error(),
			StatusCode: http.StatusUnauthorized,
		}
	}

	account := accounts.Account{
		Id:        id,
		Name:      name,
		LastName:  lastname,
		Email:     email,
		Role:      role,
		CreatedAt: createdAt,
	}

	return &account, nil
}

func (repo *PostgresRepository) UpdateRefreshTokenJtiByUserId(ctx context.Context, jti string, accountId string) error {
	query := "UPDATE linktly.accounts SET refresh_token_jti = $1 WHERE id = $2"

	_, err := repo.store.Source.Exec(ctx, query, jti, accountId)
	if err != nil {
		return linktlyError.PostgresFormatting(err)
	}

	return nil
}

func (repo *PostgresRepository) UpdateRefreshTokenJtiByUserIdAndJti(ctx context.Context, accountId string, jti string, newJti string) error {
	query := "UPDATE linktly.accounts SET refresh_token_jti = $1 WHERE id = $2 AND refresh_token_jti = $3"

	cmdTag, err := repo.store.Source.Exec(ctx, query, newJti, accountId, jti)
	if err != nil {
		return linktlyError.PostgresFormatting(err)
	}

	if cmdTag.RowsAffected() == 0 {
		return types.APIError{
			Msg:        "Refresh token does not match",
			StatusCode: http.StatusUnauthorized,
		}
	}

	return nil
}
