package auth

import (
	"log/slog"

	"github.com/go-chi/chi/v5"
	"github.com/jegj/linktly/internal/api/domain/accounts"
	"github.com/jegj/linktly/internal/api/handlers"
	"github.com/jegj/linktly/internal/api/jwt"
	"github.com/jegj/linktly/internal/config"
	"github.com/jegj/linktly/internal/store"
)

func LoadRoutes(r chi.Router, config config.Config, store *store.PostgresStore) {
	authRepository := GetNewAuthRepository(store)
	authService := AuthService{
		repository: authRepository,
		config:     config,
	}

	accountRepository := accounts.GetNewAccountRepository(store)
	accountService := accounts.AccountService{
		Repository: accountRepository,
	}

	authHandler := AuthHandler{
		service:        authService,
		accountService: accountService,
		config:         config,
	}

	publicKey, error := config.GetPublicKey()
	if error != nil {
		slog.Error(error.Error())
	} else {
		r.Route("/api/v1/auth", func(r chi.Router) {
			r.Method("POST", "/signup", handlers.CentralizedErrorHandler(authHandler.Signup))
			r.Method("POST", "/login", handlers.CentralizedErrorHandler(authHandler.Login))
			r.Method("POST", "/refresh", handlers.CentralizedErrorHandler(authHandler.Refresh))
			r.Group(func(r chi.Router) {
				r.Use(jwt.AuthMiddleware(*publicKey))
				r.Method("POST", "/logout", handlers.CentralizedErrorHandler(authHandler.Logout))
			})
		})
	}
}
