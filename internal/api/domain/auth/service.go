package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jegj/linktly/internal/api/types"
	"github.com/jegj/linktly/internal/config"
)

type AuthService struct {
	ctx        context.Context
	repository authRepository
	config     config.Config
}

func (s *AuthService) Login(email string, password string) (string, time.Time, string, time.Time, error) {
	account, err := s.repository.Login(s.ctx, email, password)
	if err != nil {
		return "", time.Time{}, "", time.Time{}, err
	}

	claims := GetClaimsFromAccount(*account)
	accessTokenExpirationTime := time.Now().Add(s.config.AccessTokenExpTime)
	refreshTokenExpirationTime := time.Now().Add(s.config.RefreshTokenExpTime)

	privateKey, err := s.config.GetPrivateKey()
	if err != nil {
		return "", time.Time{}, "", time.Time{}, types.APIError{
			Msg:        err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	accessToken, error := CreateJwt(privateKey, accessTokenExpirationTime, claims, nil)
	if error != nil {
		return "", time.Time{}, "", time.Time{}, types.APIError{
			Msg:        error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	jti, error := uuid.NewV7()
	if error != nil {
		return "", time.Time{}, "", time.Time{}, error
	}

	jtiRef := jti.String()

	refreshToken, error := CreateJwt(privateKey, refreshTokenExpirationTime, claims, &jtiRef)
	if error != nil {
		return "", time.Time{}, "", time.Time{}, types.APIError{
			Msg:        error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	error = s.repository.UpdateRefreshToken(s.ctx, jtiRef, email)
	if error != nil {
		return "", time.Time{}, "", time.Time{}, error
	}

	return accessToken, accessTokenExpirationTime, refreshToken, refreshTokenExpirationTime, nil
}

func (s *AuthService) Refresh(refreshToken string) (string, time.Time, string, time.Time, error) {
	publicKey, err := s.config.GetPublicKey()
	if err != nil {
		return "", time.Time{}, "", time.Time{}, types.APIError{
			Msg:        err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	privateKey, err := s.config.GetPrivateKey()
	if err != nil {
		return "", time.Time{}, "", time.Time{}, types.APIError{
			Msg:        err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	claims, err := VerifyJwt(refreshToken, publicKey)
	if err != nil {
		return "", time.Time{}, "", time.Time{}, types.APIError{
			Msg:        err.Error(),
			StatusCode: http.StatusUnauthorized,
		}
	}

	// FIXME: Improve this part of code realted to sub claim
	cookieJti := claims.ID
	cookieUserId := claims.Subject

	customClaims := GetClaimsFromJwtClaims(claims)

	accessTokenExpirationTime := time.Now().Add(s.config.AccessTokenExpTime)
	refreshTokenExpirationTime := time.Now().Add(s.config.RefreshTokenExpTime)

	accessToken, err := CreateJwt(privateKey, accessTokenExpirationTime, customClaims, nil)
	if err != nil {
		return "", time.Time{}, "", time.Time{}, types.APIError{
			Msg:        err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	jti, err := uuid.NewV7()
	if err != nil {
		return "", time.Time{}, "", time.Time{}, err
	}

	newJtiRef := jti.String()

	refreshToken, err = CreateJwt(privateKey, refreshTokenExpirationTime, customClaims, &newJtiRef)
	if err != nil {
		return "", time.Time{}, "", time.Time{}, types.APIError{
			Msg:        err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	err = s.repository.UpdateRefreshTokenJtiBySubAndJti(s.ctx, cookieUserId, cookieJti, newJtiRef)
	if err != nil {
		return "", time.Time{}, "", time.Time{}, err
	}

	return accessToken, accessTokenExpirationTime, refreshToken, refreshTokenExpirationTime, nil
}
