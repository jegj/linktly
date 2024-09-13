package auth

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/jegj/linktly/internal/config"
	"github.com/jegj/linktly/internal/store"
)

func LoadRoutes(ctx context.Context, r chi.Router, config config.Config, store *store.PostgresStore) {
	accountRepository := GetNewAuthRepository(store)
	accountService := AuthService{
		ctx:        ctx,
		repository: accountRepository,
	}
	authHandler := AuthHandler{
		service: accountService,
	}

	r.Route("/api/v1/auth", func(r chi.Router) {
		r.Post("/login", authHandler.Login)
	})
}
