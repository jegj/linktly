package redirections

import (
	"context"
)

type RedirectionsService struct {
	Repository redirectionsRepository
}

func (l *RedirectionsService) GetLinkByCode(ctx context.Context, code string) (*Rlink, error) {
	// TODO: handle collisions
	return l.Repository.GetLinkByCode(ctx, code)
}
