package folders

import (
	"context"

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
	return nil, nil
}
