package auth

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/jegj/linktly/internal/api/jwt"
	"github.com/jegj/linktly/internal/api/response"
	"github.com/jegj/linktly/internal/api/types"
	"github.com/jegj/linktly/internal/api/validations"
	"github.com/jegj/linktly/internal/config"
)

type AuthHandler struct {
	service AuthService
	config  config.Config
}

func (a AuthHandler) Login(w http.ResponseWriter, r *http.Request) error {
	data := &LoginReq{}
	if err := render.Bind(r, data); err != nil {
		return response.InvalidJsonRequest()
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.RegisterValidation("password", validations.PasswordValidation); err != nil {
		return types.APIError{
			Msg:        err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	errs := validate.Struct(data)
	if errs != nil {
		validationErrors := make(map[string]string)
		// Cast the error to a ValidationErrors type
		for _, err := range errs.(validator.ValidationErrors) {
			// Extract the field name and error message
			validationErrors[err.Field()] = err.Error()
		}
		return response.InvalidRequestData(validationErrors)
	}

	accessToken, acessTokenExpirationTime, refreshToken, refreshTokenExpirationTime, error := a.service.Login(r.Context(), data.Email, data.Password)
	if error != nil {
		return error
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:     LinktlyAccessTokenCookieName,
			Value:    accessToken,
			Expires:  acessTokenExpirationTime,
			Path:     "/",
			HttpOnly: true,
			Secure:   a.config.HTTPCookieSecure,
			SameSite: http.SameSiteStrictMode,
		})
		http.SetCookie(w, &http.Cookie{
			Name:     LinktlyRefreshTokenCookieName,
			Value:    refreshToken,
			Expires:  refreshTokenExpirationTime,
			Path:     "/api/v1/auth/refresh",
			HttpOnly: true,
			Secure:   a.config.HTTPCookieSecure,
			SameSite: http.SameSiteStrictMode,
		})
		return nil
	}
}

func (a AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie(LinktlyRefreshTokenCookieName)
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			return types.APIError{
				Msg:        err.Error(),
				StatusCode: http.StatusUnauthorized,
			}
		default:
			return types.APIError{
				Msg:        err.Error(),
				StatusCode: http.StatusInternalServerError,
			}
		}
	}

	accessToken, acessTokenExpirationTime, refreshToken, refreshTokenExpirationTime, error := a.service.Refresh(r.Context(), cookie.Value)

	if error != nil {
		return error
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:     LinktlyAccessTokenCookieName,
			Value:    accessToken,
			Expires:  acessTokenExpirationTime,
			Path:     "/",
			HttpOnly: true,
			Secure:   a.config.HTTPCookieSecure,
			SameSite: http.SameSiteStrictMode,
		})
		http.SetCookie(w, &http.Cookie{
			Name:     LinktlyRefreshTokenCookieName,
			Value:    refreshToken,
			Expires:  refreshTokenExpirationTime,
			Path:     "/api/v1/auth/refresh",
			HttpOnly: true,
			Secure:   a.config.HTTPCookieSecure,
			SameSite: http.SameSiteStrictMode,
		})
		return nil
	}
}

func (a AuthHandler) Logout(w http.ResponseWriter, r *http.Request) error {
	context := r.Context()
	userId := context.Value(jwt.UserIdContextKey).(string)
	err := a.service.Logout(context, userId)
	if err != nil {
		return err
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:     LinktlyAccessTokenCookieName,
			Value:    "",
			Expires:  time.Now().Add(-time.Hour),
			Path:     "/",
			HttpOnly: true,
			Secure:   a.config.HTTPCookieSecure,
			SameSite: http.SameSiteStrictMode,
		})
		http.SetCookie(w, &http.Cookie{
			Name:     LinktlyRefreshTokenCookieName,
			Value:    "",
			Expires:  time.Now().Add(-time.Hour),
			Path:     "/api/v1/auth/refresh",
			HttpOnly: true,
			Secure:   a.config.HTTPCookieSecure,
			SameSite: http.SameSiteStrictMode,
		})
		return nil
	}
}
