package redirections

import (
	"context"
)

type LinksService struct {
	Repository redirectionsRepository
}

func (l *LinksService) GetLinkByCode(ctx context.Context, code string) (*Rlink, error) {
	// TODO: handle collisions
	return l.Repository.GetLinkByCode(ctx, code)
}
