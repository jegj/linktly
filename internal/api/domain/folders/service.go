package folders

import "context"

type FolderService struct {
	Repository foldersRepository
}

func (s *FolderService) CreateFolder(ctx context.Context, folder *Folder) (*Folder, error) {
	return s.Repository.CreateFolder(ctx, folder)
}

func (s *FolderService) GetFoldersByUserId(ctx context.Context, userId string) ([]*Folder, error) {
	return s.Repository.GetFolders(ctx, userId)
}
