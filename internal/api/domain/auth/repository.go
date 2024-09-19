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
	// UpdateRefreshToken(ctx context.Context, email string) (*accounts.Account, error)
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
