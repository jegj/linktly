package links

import (
	"context"
	"time"

	linktlyError "github.com/jegj/linktly/internal/api/error"
	"github.com/jegj/linktly/internal/store"
)

type linksRepository interface {
	CreateLink(ctx context.Context, link *Link) (*Link, error)
	GetLink(ctx context.Context, id string, userId string) (*Link, error)
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
	// TODO: Add a hash column so this function can reuse code for similar codes ?? what about other user ownerships?

	query := `INSERT INTO linktly.links(name, description, account_id, folder_id, linktly_code, url, expires_at ) VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING id, created_at`
	err := repo.store.Source.QueryRow(ctx, query, link.Name, link.Description, link.AccountId, link.FolderId, link.LinktlyCode, link.Url, link.ExpiresAt).Scan(&id, &createdAt)
	if err != nil {
		return nil, linktlyError.PostgresFormatting(err)
	}

	link.Id = id
	link.CreatedAt = createdAt
	return link, nil
}

func (repo *PostgresRepository) GetLink(ctx context.Context, id string, userId string) (*Link, error) {
	var link Link
	query := `SELECT id, name, description, account_id, folder_id, linktly_code, url, expires_at, created_at FROM linktly.links WHERE id = $1 AND account_id = $2`
	err := repo.store.Source.QueryRow(ctx, query, id, userId).Scan(&link.Id, &link.Name, &link.Description, &link.AccountId, &link.FolderId, &link.LinktlyCode, &link.Url, &link.ExpiresAt, &link.CreatedAt)
	if err != nil {
		return nil, linktlyError.PostgresFormatting(err)
	}
	return &link, nil
}
