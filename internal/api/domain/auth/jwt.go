package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jegj/linktly/internal/api/domain/accounts"
	"github.com/jegj/linktly/internal/config"
)

func GetClaimsFromAccount(account accounts.Account) JwtClaims {
	return JwtClaims{
		Sub:   account.Id,
		Email: account.Email,
		Role:  account.Role,
	}
}

func CreateJwt(config config.Config, expirationTime time.Time, claims JwtClaims) (string, error) {
	fmt.Println("-------------------")
	privateKey, err := config.GetPrivateKey()
	if err != nil {
		fmt.Printf("-->%v", err)
		return "", err
	}

	t := jwt.NewWithClaims(jwt.SigningMethodRS256,
		jwt.MapClaims{
			"iat":   time.Now(),
			"sub":   claims.Sub,
			"email": claims.Email,
			"role":  claims.Role,
			"exp":   expirationTime,
		})

	tokenString, err := t.SignedString(privateKey)
	if err != nil {
		fmt.Printf("token-->%v", err)

		return "", err
	}

	return tokenString, nil
}
