package links

import (
	"context"
	"time"

	linktlyError "github.com/jegj/linktly/internal/api/error"
	"github.com/jegj/linktly/internal/store"
)

type linksRepository interface {
	CreateLink(ctx context.Context, link *Link) (*Link, error)
}

type PostgresRepository struct {
	store *store.PostgresStore
}

func GetNewLinkRepository(store *store.PostgresStore) *PostgresRepository {
	return &PostgresRepository{
		store: store,
	}
}

func (repo *PostgresRepository) CreateLink(ctx context.Context, link *Link) (*Link, error) {
	var id string
	var createdAt time.Time

	query := `INSERT INTO linktly.links(name, description, account_id, folder_id, linktly_url, url, expires_at ) VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING id, created_at`
	err := repo.store.Source.QueryRow(ctx, query, link.Name, link.Description, link.AccountId, link.FolderId, link.LinktlyUrl, link.Url, link.ExpiresAt).Scan(&id, &createdAt)
	if err != nil {
		return nil, linktlyError.PostgresFormatting(err)
	}

	link.Id = id
	link.CreatedAt = createdAt
	return link, nil
}
