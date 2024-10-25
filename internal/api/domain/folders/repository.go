package folders

import (
	"context"
	"time"

	linktlyError "github.com/jegj/linktly/internal/api/error"
	"github.com/jegj/linktly/internal/store"
)

type foldersRepository interface {
	CreateFolder(ctx context.Context, folder *Folder) (*Folder, error)
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

	if folder.ParentFolderId != "" {
		parentFolder = &folder.ParentFolderId
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
