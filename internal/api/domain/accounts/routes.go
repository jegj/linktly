package accounts

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/jegj/linktly/internal/config"
	"github.com/jegj/linktly/internal/store"
)

func LoadRoutes(ctx context.Context, r chi.Router, config config.Config, store *store.PostgresStore) {
	accountRepository := GetNewAccountRepository(store)
	accountService := AccountService{
		ctx:        ctx,
		repository: accountRepository,
	}
	accountHandler := AccountHandler{
		service: accountService,
	}

	r.Route("/api/v1/accounts", func(r chi.Router) {
		r.Get("/", accountHandler.GetAccountsHandler)
	})
}
