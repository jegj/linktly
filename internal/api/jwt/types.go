package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

const (
	LinktlyAccessTokenCookieName  = "access_token"
	LinktlyRefreshTokenCookieName = "refresh_token"
)

type JwtCustomClaims struct {
	Email string
	Sub   string
	Role  int
}

type JwtClaims struct {
	*JwtCustomClaims
	jwt.RegisteredClaims
}
