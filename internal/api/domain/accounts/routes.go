package accounts

import (
	"log/slog"

	"github.com/go-chi/chi/v5"
	"github.com/jegj/linktly/internal/api/handlers"
	"github.com/jegj/linktly/internal/api/jwt"
	"github.com/jegj/linktly/internal/config"
	"github.com/jegj/linktly/internal/store"
)

func LoadRoutes(r chi.Router, config config.Config, store *store.PostgresStore) {
	accountRepository := GetNewAccountRepository(store)
	accountService := AccountService{
		Repository: accountRepository,
	}
	accountHandler := AccountHandler{
		service: accountService,
	}

	publicKey, error := config.GetPublicKey()
	if error != nil {
		slog.Error(error.Error())
	} else {
		r.Route("/api/v1/accounts", func(r chi.Router) {
			r.Group(func(r chi.Router) {
				r.Use(jwt.AuthMiddleware(*publicKey))
				r.Method("GET", "/{id}", handlers.CentralizedErrorHandler(accountHandler.GetAccountByIdHandler))
				r.Method("POST", "/", handlers.CentralizedErrorHandler(accountHandler.CreateAccount))
			})
		})
	}
}
