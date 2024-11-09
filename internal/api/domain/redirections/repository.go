package redirections

import (
	"context"

	linktlyError "github.com/jegj/linktly/internal/api/error"
	"github.com/jegj/linktly/internal/store"
)

type redirectionsRepository interface {
	GetLinkByCode(ctx context.Context, id string) (*Rlink, error)
}

type PostgresRepository struct {
	store *store.PostgresStore
}

func GetNewRlinkRepository(store *store.PostgresStore) *PostgresRepository {
	return &PostgresRepository{
		store: store,
	}
}

func (repo *PostgresRepository) GetLinkByCode(ctx context.Context, code string) (*Rlink, error) {
	query := `SELECT code FROM linktly.links WHERE code = $1 AND expires_at > NOW()`

	var url string

	err := repo.store.Source.QueryRow(ctx, query, code).Scan(&code)
	if err != nil {
		return nil, linktlyError.PostgresFormatting(err)
	}

	link := Rlink{
		Url: url,
	}

	return &link, nil
}
