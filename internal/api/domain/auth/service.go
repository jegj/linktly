package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jegj/linktly/internal/api/jwt"
	"github.com/jegj/linktly/internal/api/types"
	"github.com/jegj/linktly/internal/config"
)

// TODO: context should be the one defined by the router
type AuthService struct {
	ctx        context.Context
	repository authRepository
	config     config.Config
}

func (s *AuthService) Login(email string, password string) (string, time.Time, string, time.Time, error) {
	// TODO: Pass request context instead
	account, err := s.repository.Login(s.ctx, email, password)
	if err != nil {
		return "", time.Time{}, "", time.Time{}, err
	}

	claims := jwt.GetClaimsFromAccountData(account.Id, account.Email, account.Role)
	accessTokenExpirationTime := time.Now().Add(s.config.AccessTokenExpTime)
	refreshTokenExpirationTime := time.Now().Add(s.config.RefreshTokenExpTime)

	privateKey, err := s.config.GetPrivateKey()
	if err != nil {
		return "", time.Time{}, "", time.Time{}, types.APIError{
			Msg:        err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	accessToken, error := jwt.CreateJwt(privateKey, accessTokenExpirationTime, claims, nil)
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

	refreshToken, error := jwt.CreateJwt(privateKey, refreshTokenExpirationTime, claims, &jtiRef)
	if error != nil {
		return "", time.Time{}, "", time.Time{}, types.APIError{
			Msg:        error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	// TODO: Pass request context instead
	error = s.repository.UpdateRefreshTokenJtiByUserId(s.ctx, jtiRef, account.Id)
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

	refreshTokenClaims, err := jwt.VerifyJwt(refreshToken, publicKey)
	if err != nil {
		return "", time.Time{}, "", time.Time{}, types.APIError{
			Msg:        err.Error(),
			StatusCode: http.StatusUnauthorized,
		}
	}

	cookieJti := refreshTokenClaims.ID
	cookieUserId := refreshTokenClaims.Subject

	customClaims := jwt.GetClaimsFromJwtClaims(*refreshTokenClaims)

	accessTokenExpirationTime := time.Now().Add(s.config.AccessTokenExpTime)
	refreshTokenExpirationTime := time.Now().Add(s.config.RefreshTokenExpTime)

	accessToken, err := jwt.CreateJwt(privateKey, accessTokenExpirationTime, customClaims, nil)
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

	refreshToken, err = jwt.CreateJwt(privateKey, refreshTokenExpirationTime, customClaims, &newJtiRef)
	if err != nil {
		return "", time.Time{}, "", time.Time{}, types.APIError{
			Msg:        err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	// TODO: Pass request context instead
	err = s.repository.UpdateRefreshTokenJtiByUserIdAndJti(s.ctx, cookieUserId, cookieJti, newJtiRef)
	if err != nil {
		return "", time.Time{}, "", time.Time{}, err
	}

	return accessToken, accessTokenExpirationTime, refreshToken, refreshTokenExpirationTime, nil
}

func (s *AuthService) Logout(ctx context.Context, userId string) error {
	err := s.repository.UpdateRefreshTokenJtiByUserId(ctx, "", userId)
	if err != nil {
		return err
	} else {
		return nil
	}
}
