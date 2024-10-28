package folders

import (
	"context"
	"net/http"
	"time"

	linktlyError "github.com/jegj/linktly/internal/api/error"
	"github.com/jegj/linktly/internal/api/types"
	"github.com/jegj/linktly/internal/store"
)

type foldersRepository interface {
	CreateFolder(ctx context.Context, folder *Folder) (*Folder, error)
	GetFolders(ctx context.Context, userId string) ([]*Folder, error)
}

type PostgresRepository struct {
	store *store.PostgresStore
}

func GetNewFolderRepository(store *store.PostgresStore) *PostgresRepository {
	return &PostgresRepository{
		store: store,
	}
}

func (repo *PostgresRepository) CreateFolder(ctx context.Context, folder *Folder) (*Folder, error) {
	var id string
	var createdAt time.Time
	var parentFolder *string = nil

	if folder.ParentFolderId != nil {
		parentFolder = folder.ParentFolderId
	}

	query := `INSERT INTO linktly.folders(name, description, account_id, parent_folder_id ) VALUES($1,$2,$3,$4) RETURNING id, created_at`
	err := repo.store.Source.QueryRow(ctx, query, folder.Name, folder.Description, folder.AccountId, parentFolder).Scan(&id, &createdAt)
	if err != nil {
		return nil, linktlyError.PostgresFormatting(err)
	}

	folder.Id = id
	folder.CreatedAt = createdAt
	return folder, nil
}

func (repo *PostgresRepository) GetFolders(ctx context.Context, userId string) ([]*Folder, error) {
	query := `SELECT id, name, description, parent_folder_id, created_at FROM linktly.folders WHERE account_id = $1`
	rows, error := repo.store.Source.Query(ctx, query, userId)

	if error != nil {
		return nil, linktlyError.PostgresFormatting(error)
	}

	var result []*Folder

	for rows.Next() {
		var folder Folder
		err := rows.Scan(&folder.Id, &folder.Name, &folder.Description, &folder.ParentFolderId, &folder.CreatedAt)
		if err != nil {
			return nil, types.APIError{
				Msg:        err.Error(),
				StatusCode: http.StatusInternalServerError,
			}
		}
		result = append(result, &folder)
	}

	if rows.Err() != nil {
		return nil, types.APIError{
			Msg:        rows.Err().Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return result, nil
}
