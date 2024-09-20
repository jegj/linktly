package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jegj/linktly/internal/config"
)

func CreateJwt(config config.Config, expirationTime time.Time) (string, error) {
	privateKey, err := config.GetPrivateKey()
	if err != nil {
		return "", err
	}

	t := jwt.NewWithClaims(jwt.SigningMethodRS256,
		jwt.MapClaims{
			"iss": "my-auth-server",
			"sub": "john",
			"foo": 2,
			"exp": expirationTime,
		})

	tokenString, err := t.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
