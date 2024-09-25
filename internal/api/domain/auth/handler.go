package auth

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
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

	accessToken, acessTokenExpirationTime, refreshToken, refreshTokenExpirationTime, error := a.service.Login(data.Email, data.Password)
	if error != nil {
		return error
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:     LinktlyAccessTokenCookieName,
			Value:    accessToken,
			Expires:  acessTokenExpirationTime,
			Path:     "/",
			HttpOnly: true,
			// TODO: replace for https envs
			Secure: false, // Set to true if using HTTPS
		})
		http.SetCookie(w, &http.Cookie{
			Name:     LinktlyRefreshTokenCookieName,
			Value:    refreshToken,
			Expires:  refreshTokenExpirationTime,
			Path:     "/",
			HttpOnly: true,
			// TODO: replace for https envs
			Secure: false, // Set to true if using HTTPS
		})
		return nil
	}
}
