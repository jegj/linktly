package accounts

import "context"

type AccountService struct {
	Repository accountsRepository
}

func (s *AccountService) GetAccountById(ctx context.Context, id string) (*Account, error) {
	return s.Repository.GetByID(ctx, id)
}

func (s *AccountService) CreateAccount(ctx context.Context, account *Account) (*Account, error) {
	return s.Repository.CreateAccount(ctx, account)
}
