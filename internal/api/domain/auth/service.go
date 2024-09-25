package auth

import (
	"context"
	"net/http"
	"time"

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

	accessToken, error := CreateJwt(privateKey, accessTokenExpirationTime, claims)
	if error != nil {
		return "", time.Time{}, "", time.Time{}, types.APIError{
			Msg:        error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	refreshToken, error := CreateJwt(privateKey, refreshTokenExpirationTime, claims)
	if error != nil {
		return "", time.Time{}, "", time.Time{}, types.APIError{
			Msg:        error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return accessToken, accessTokenExpirationTime, refreshToken, refreshTokenExpirationTime, nil
}
