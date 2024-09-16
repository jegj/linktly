package accounts

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jegj/linktly/internal/api/types"
	"github.com/jegj/linktly/internal/store"
	"golang.org/x/crypto/bcrypt"
)

type accountsRepository interface {
	GetByID(ctx context.Context, id string) (*Account, error)
	CreateAccount(ctx context.Context, account *Account) (*Account, error)
}

type PostgresRepository struct {
	store *store.PostgresStore
}

func GetNewAccountRepository(store *store.PostgresStore) *PostgresRepository {
	return &PostgresRepository{
		store: store,
	}
}

func (repo *PostgresRepository) GetByID(ctx context.Context, id string) (*Account, error) {
	query := `SELECT name, lastname, email, role, created_at FROM linktly.accounts WHERE id = $1`

	var name string
	var lastname string
	var email string
	var role int
	var createdAt time.Time

	err := repo.store.Source.QueryRow(ctx, query, id).Scan(&name, &lastname, &email, &role, &createdAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, types.APIError{
				Msg:        fmt.Sprintf("Acount not found for id %v", id),
				StatusCode: http.StatusNotFound,
			}
		} else {
			return nil, types.APIError{
				Msg:        err.Error(),
				StatusCode: http.StatusInternalServerError,
			}
		}
	}

	account := Account{
		Id:        id,
		Name:      name,
		LastName:  lastname,
		Email:     email,
		Role:      role,
		CreatedAt: createdAt,
	}

	return &account, nil
}

func (repo *PostgresRepository) CreateAccount(ctx context.Context, account *Account) (*Account, error) {
	var id string
	var createdAt time.Time
	var role int

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(account.Password), 15)
	if err != nil {
		return nil, types.APIError{
			Msg:        err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	query := `INSERT INTO linktly.accounts(name, lastname, email, password ) VALUES($1,$2,$3,$4) RETURNING id, created_at, role`
	err = repo.store.Source.QueryRow(ctx, query, account.Name, account.LastName, account.Email, string(hashedPassword)).Scan(&id, &createdAt, &role)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" {
				return nil, types.APIError{
					Msg:        err.Error(),
					StatusCode: http.StatusConflict,
				}
			} else {
				return nil, types.APIError{
					Msg:        err.Error(),
					StatusCode: http.StatusInternalServerError,
				}
			}
		} else {
			return nil, types.APIError{
				Msg:        err.Error(),
				StatusCode: http.StatusInternalServerError,
			}
		}
	}

	account.Id = id
	account.CreatedAt = createdAt
	account.Role = role
	return account, nil
}
