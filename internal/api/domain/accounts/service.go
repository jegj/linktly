package accounts

import "context"

type AccountService struct {
	ctx        context.Context
	repository accountsRepository
}

func (s *AccountService) GetAccountById(id string) (*Account, error) {
	return s.repository.GetByID(s.ctx, id)
}
