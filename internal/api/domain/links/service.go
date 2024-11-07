package links

import (
	"context"
	"math/rand"
)

type LinksService struct {
	Repository linksRepository
}

func (l *LinksService) CreateLink(ctx context.Context, link *Link) (*Link, error) {
	link.LinktlyUrl = CreateShortCode()
	return l.Repository.CreateLink(ctx, link)
}

func CreateShortCode() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
