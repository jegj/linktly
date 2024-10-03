package accounts

import (
	"context"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jegj/linktly/internal/api/handlers"
	"github.com/jegj/linktly/internal/api/jwt"
	"github.com/jegj/linktly/internal/api/types"
	"github.com/jegj/linktly/internal/config"
	"github.com/jegj/linktly/internal/store"
)

// TODO: MOVE TO JWT PACKAGE
func jwtMiddleware(publicKey rsa.PublicKey) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			accessTokenCookie, err := r.Cookie("access_token")
			if err != nil {
				jsonResp := types.APIError{
					Msg:        "Can't access to the cookie",
					StatusCode: http.StatusUnauthorized,
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(jsonResp)
				return
			}

			refreshTokenClaims, err := jwt.VerifyJwt(accessTokenCookie.Value, &publicKey)
			fmt.Printf("===>%v", refreshTokenClaims)
			if err != nil {
				jsonResp := types.APIError{
					Msg:        "Invalid access token",
					StatusCode: http.StatusUnauthorized,
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(jsonResp)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

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
			r.Use(jwtMiddleware(*publicKey))
			r.Method("GET", "/{id}", handlers.CentralizedErrorHandler(accountHandler.GetAccountByIdHandler))
			r.Method("POST", "/", handlers.CentralizedErrorHandler(accountHandler.CreateAccount))
		})
	})
}
