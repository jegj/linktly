package auth

import (
	"context"

	"github.com/jegj/linktly/internal/api/domain/accounts"
)

type AuthService struct {
	ctx        context.Context
	repository authRepository
}

func (s *AuthService) Login(email string, password string) (*accounts.Account, error) {
	return s.repository.Login(s.ctx, email, password)
}
