package accounts

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/jegj/linktly/internal/api/handlers"
	"github.com/jegj/linktly/internal/api/jwt"
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

	// TODO: Do something better here
	publicKey, error := config.GetPublicKey()
	if error != nil {
		panic("oh no!")
	}

	r.Route("/api/v1/accounts", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(jwt.AuthMiddleware(*publicKey))
			r.Method("GET", "/{id}", handlers.CentralizedErrorHandler(accountHandler.GetAccountByIdHandler))
			r.Method("POST", "/", handlers.CentralizedErrorHandler(accountHandler.CreateAccount))
		})
	})
}
