package accounts

import "context"

type AccountService struct {
	repository accountsRepository
}

func (s *AccountService) GetAccountById(ctx context.Context, id string) (*Account, error) {
	return s.repository.GetByID(ctx, id)
}

func (s *AccountService) CreateAccount(ctx context.Context, account *Account) (*Account, error) {
	return s.repository.CreateAccount(ctx, account)
}
