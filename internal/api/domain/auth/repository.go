package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jegj/linktly/internal/api/domain/accounts"
	"github.com/jegj/linktly/internal/api/types"
	"github.com/jegj/linktly/internal/store"
	"golang.org/x/crypto/bcrypt"
)

type authRepository interface {
	Login(ctx context.Context, email string, password string) (*accounts.Account, error)
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
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, types.APIError{
				Msg:        fmt.Sprintf("Acount not found for email %s", email),
				StatusCode: http.StatusNotFound,
			}
		} else {
			return nil, types.APIError{
				Msg:        err.Error(),
				StatusCode: http.StatusInternalServerError,
			}
		}
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
