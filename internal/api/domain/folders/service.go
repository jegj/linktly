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

func (s *FolderService) DeleteFoldersByIdAndUserId(ctx context.Context, folderId string, userId string) error {
	return s.Repository.DeleteFoldersByIdAndUserId(ctx, folderId, userId)
}

func (s *FolderService) PatchFolderByIdAndUserId(ctx context.Context, folderId string, userId string, folder *Folder) (*Folder, error) {
	return s.Repository.PatchFolderByIdAndUserId(ctx, folderId, userId, folder)
}
