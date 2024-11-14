package links

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand"
	"net/http"

	"github.com/jegj/linktly/internal/api/types"
)

type LinksService struct {
	Repository linksRepository
}

func (l *LinksService) CreateLink(ctx context.Context, link *Link) (*Link, error) {
	for {
		link.LinktlyCode = createShortCode()
		nlink, err := l.Repository.CreateLink(ctx, link)
		if err == nil || !isConflictApiError(err) {
			return nlink, err
		}
		msg := fmt.Sprintf("Short code conflict %v, generating a new one", link.LinktlyCode)
		slog.Info(msg, "error", err.Error(), "code", link.LinktlyCode)
	}
}

func isConflictApiError(err error) bool {
	if apiErr, ok := err.(types.APIError); ok {
		return apiErr.StatusCode == http.StatusConflict
	} else {
		return false
	}
}

func createShortCode() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
