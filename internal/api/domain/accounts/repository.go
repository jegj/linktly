package accounts

import (
	"context"
	"time"

	"github.com/jegj/linktly/internal/store"
)

type accountsRepository interface {
	GetByID(ctx context.Context, id string) (*Account, error)
	CreateAccount(ctx context.Context, account *Account) (string, error)
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
		return nil, err
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

func (repo *PostgresRepository) CreateAccount(ctx context.Context, account *Account) (string, error) {
	var id string
	var createdAt time.Time
	query := `INSERT INTO linktly.accounts(name, lastname, email, password ) VALUES($1,$2,$3,$4) RETURNING id, created_at`
	err := repo.store.Source.QueryRow(ctx, query, account.Name, account.LastName, account.Email, account.Password).Scan(&id, &createdAt)
	if err != nil {
		return "", err
	}

	return id, nil
}
