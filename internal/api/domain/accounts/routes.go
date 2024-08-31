package accounts

import (
	"github.com/go-chi/chi/v5"
	"github.com/jegj/linktly/internal/config"
)

func LoadRoutes(r chi.Router, config config.Config /*store *Store*/) {
	/*
		accountRepository := repository.NewAccountRepository(db, domain.CollectionUser)
		lc := &controller.LoginController{
			LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
			Env:          env,
		}
	*/

	r.Route("/api/v1/accounts", func(r chi.Router) {
		r.Get("/", GetAccountsHandler)
	})
}
