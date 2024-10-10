package jwt

import (
	"github.com/golang-jwt/jwt/v5"
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
