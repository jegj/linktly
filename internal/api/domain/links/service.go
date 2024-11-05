package links

import "context"

type LinksService struct {
	Repository linksRepository
}

func (l *LinksService) CreateLink(ctx context.Context, link *Link) (*Link, error) {
	return l.Repository.CreateLink(ctx, link)
}
