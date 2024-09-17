package auth

import (
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/jegj/linktly/internal/api/response"
	"github.com/jegj/linktly/internal/api/types"
	"github.com/jegj/linktly/internal/api/validations"
)

type AuthHandler struct {
	service AuthService
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

	_, error := a.service.Login(data.Email, data.Password)
	if error != nil {
		return error
	} else {
		// TODO: IMPROVE THIS
		http.SetCookie(w, &http.Cookie{
			Name:     "example_cookie",
			Value:    "cookie_value",
			Expires:  time.Now().Add(24 * time.Hour), // Set expiration time
			Path:     "/",
			HttpOnly: true,  // For security, HttpOnly cookies are not accessible via JavaScript
			Secure:   false, // Set to true if using HTTPS
		})
		return nil
	}
}
