package jwt

import (
	"context"
	"crypto/rsa"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/jegj/linktly/internal/api/types"
)

type contextKey string

const (
	UserIdContextKey    contextKey = "userId"
	UserEmailContextKey contextKey = "userEmail"
	UserRolesContextKey contextKey = "userRole"
)

// TODO: Do something with ctx
func AuthMiddleware(publicKey rsa.PublicKey) func(http.Handler) http.Handler {
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
				err = json.NewEncoder(w).Encode(jsonResp)
				if err != nil {
					slog.Error(err.Error())
				}
				return
			}

			refreshTokenClaims, err := VerifyJwt(accessTokenCookie.Value, &publicKey)
			if err != nil {
				jsonResp := types.APIError{
					Msg:        "Invalid access token",
					StatusCode: http.StatusUnauthorized,
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				err = json.NewEncoder(w).Encode(jsonResp)
				if err != nil {
					slog.Error(err.Error())
				}
				return
			}
			ctx := r.Context()
			ctx = context.WithValue(ctx, UserIdContextKey, refreshTokenClaims.Subject)
			ctx = context.WithValue(ctx, UserEmailContextKey, refreshTokenClaims.Email)
			ctx = context.WithValue(ctx, UserRolesContextKey, refreshTokenClaims.Role)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
