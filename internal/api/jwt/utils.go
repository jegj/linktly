package jwt

import (
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GetClaimsFromAccountData(accounId string, accountEmail string, accountRole int) JwtCustomClaims {
	return JwtCustomClaims{
		Sub:   accounId,
		Email: accountEmail,
		Role:  accountRole,
	}
}

func GetClaimsFromJwtClaims(claims JwtClaims) JwtCustomClaims {
	return JwtCustomClaims{
		Sub:   claims.Subject,
		Email: claims.Email,
		Role:  claims.Role,
	}
}

func CreateJwt(privateKey *rsa.PrivateKey, expirationTime time.Time, claims JwtCustomClaims, jti *string) (string, error) {
	var jwtClaims jwt.MapClaims
	if jti != nil {
		jwtClaims = jwt.MapClaims{
			"iat":   time.Now().Unix(),
			"sub":   claims.Sub,
			"email": claims.Email,
			"role":  claims.Role,
			"exp":   expirationTime.Unix(),
			"jti":   *jti,
		}
	} else {
		jwtClaims = jwt.MapClaims{
			"iat":   time.Now().Unix(),
			"sub":   claims.Sub,
			"email": claims.Email,
			"role":  claims.Role,
			"exp":   expirationTime.Unix(),
		}
	}
	t := jwt.NewWithClaims(jwt.SigningMethodRS256, jwtClaims)

	tokenString, err := t.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJwt(tokenString string, publicKey *rsa.PublicKey) (*JwtClaims, error) {
	claims := &JwtClaims{}
	// Parse the token with the secret key
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	// Check for verification errors
	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Return the claims
	return claims, nil
}
