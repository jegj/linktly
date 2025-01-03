package links

import (
	"context"
	"net/http"
	"time"

	linktlyError "github.com/jegj/linktly/internal/api/error"
	"github.com/jegj/linktly/internal/api/types"
	"github.com/jegj/linktly/internal/store"
)

type linksRepository interface {
	CreateLink(ctx context.Context, link *Link) (*Link, error)
	GetLink(ctx context.Context, id string, userId string) (*Link, error)
	GetLinkByFolderId(ctx context.Context, id string, userId string, folderId string) (*Link, error)
	GetLinksByFolderId(ctx context.Context, folderId string, userId string) ([]*Link, error)
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

func (repo *PostgresRepository) GetLinkByFolderId(ctx context.Context, id string, userId string, folderId string) (*Link, error) {
	var link Link
	query := `SELECT id, name, description, account_id, folder_id, linktly_code, url, expires_at, created_at FROM linktly.links WHERE id = $1 AND account_id = $2 AND folder_id = $3`
	err := repo.store.Source.QueryRow(ctx, query, id, userId, folderId).Scan(&link.Id, &link.Name, &link.Description, &link.AccountId, &link.FolderId, &link.LinktlyCode, &link.Url, &link.ExpiresAt, &link.CreatedAt)
	if err != nil {
		return nil, linktlyError.PostgresFormatting(err)
	}
	return &link, nil
}

func (repo *PostgresRepository) GetLinksByFolderId(ctx context.Context, folderId string, userId string) ([]*Link, error) {
	query := `SELECT id, name, description, account_id, folder_id, linktly_code, url, expires_at, created_at FROM linktly.links WHERE folder_id = $1 AND account_id = $2`
	rows, err := repo.store.Source.Query(ctx, query, folderId, userId)
	if err != nil {
		return nil, linktlyError.PostgresFormatting(err)
	}

	var result []*Link

	for rows.Next() {
		var link Link

		err := rows.Scan(&link.Id, &link.Name, &link.Description, &link.AccountId, &link.FolderId, &link.LinktlyCode, &link.Url, &link.ExpiresAt, &link.CreatedAt)
		if err != nil {
			return nil, types.APIError{
				Msg:        err.Error(),
				StatusCode: http.StatusInternalServerError,
			}
		}
		result = append(result, &link)
	}

	if rows.Err() != nil {
		return nil, types.APIError{
			Msg:        rows.Err().Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return result, nil
}
