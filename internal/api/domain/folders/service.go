package folders

import "context"

type FolderService struct {
	Repository foldersRepository
}

func (s *FolderService) CreateFolder(ctx context.Context, folder *Folder) (*Folder, error) {
	return s.Repository.CreateFolder(ctx, folder)
}
