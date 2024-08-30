package accounts

import (
	"github.com/go-chi/chi/v5"
)

func LoadRoutes(r chi.Router) {
	r.Route("/api/v1/accounts", func(r chi.Router) {
		r.Get("/", GetAccounts)
	})
}
