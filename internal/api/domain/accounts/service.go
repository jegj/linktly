package accounts

import "context"

type AccountService struct {
	ctx        context.Context
	repository accountsRepository
}

func (s *AccountService) GetAccountById(id string) (*Account, error) {
	return s.repository.GetByID(s.ctx, id)
}

func (s *AccountService) CreateAccount(account *Account) (*Account, error) {
	return s.repository.CreateAccount(s.ctx, account)
}
