package auth

import (
	"net/http"
	"time"

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

	account, error := a.service.Login(data.Email, data.Password)
	if error != nil {
		return error
	} else {
		claims := GetClaimsFromAccount(*account)
		expirationTime := time.Now().Add(a.config.AccessTokenExpTime)

		privateKey, err := a.config.GetPrivateKey()
		if err != nil {
			return types.APIError{
				Msg:        err.Error(),
				StatusCode: http.StatusInternalServerError,
			}
		}

		jwt, error := CreateJwt(privateKey, expirationTime, claims)
		if error != nil {
			return types.APIError{
				Msg:        error.Error(),
				StatusCode: http.StatusInternalServerError,
			}
		} else {
			http.SetCookie(w, &http.Cookie{
				Name:     LinktlyAuthCookieName,
				Value:    jwt,
				Expires:  expirationTime,
				Path:     "/",
				HttpOnly: true,
				// TODO: replace for https envs
				Secure: false, // Set to true if using HTTPS
			})
			return nil
		}
	}
}
