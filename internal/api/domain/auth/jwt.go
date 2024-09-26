package auth

import (
	"crypto/rsa"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jegj/linktly/internal/api/domain/accounts"
)

func GetClaimsFromAccount(account accounts.Account) JwtClaims {
	return JwtClaims{
		Sub:   account.Id,
		Email: account.Email,
		Role:  account.Role,
	}
}

func CreateJwt(privateKey *rsa.PrivateKey, expirationTime time.Time, claims JwtClaims, jti *string) (string, error) {
	var jwtClaims jwt.MapClaims
	if jti != nil {
		jwtClaims = jwt.MapClaims{
			"iat":   time.Now(),
			"sub":   claims.Sub,
			"email": claims.Email,
			"role":  claims.Role,
			"exp":   expirationTime,
			"jti":   *jti,
		}
	} else {
		jwtClaims = jwt.MapClaims{
			"iat":   time.Now(),
			"sub":   claims.Sub,
			"email": claims.Email,
			"role":  claims.Role,
			"exp":   expirationTime,
		}
	}
	t := jwt.NewWithClaims(jwt.SigningMethodRS256, jwtClaims)

	tokenString, err := t.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
