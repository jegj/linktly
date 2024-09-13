package auth

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/jegj/linktly/internal/api/types"
	"github.com/jegj/linktly/internal/api/validations"
)

type AuthHandler struct {
	service AuthService
}

func (a AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// slog.Debug("Debug message accounts")
	// slog.Info("Info message accounts")
	data := &LoginReq{}
	if err := render.Bind(r, data); err != nil {
		err := render.Render(w, r, types.NewLinktlyError(err, http.StatusBadRequest, http.StatusText(http.StatusBadRequest)))
		if err != nil {
			slog.Error(err.Error())
		}
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.RegisterValidation("password", validations.PasswordValidation); err != nil {
		err := render.Render(w, r, types.NewLinktlyError(err, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)))
		if err != nil {
			slog.Error(err.Error())
		}
		return
	}

	errs := validate.Struct(data)
	if errs != nil {
		builder := types.LinktlyErrorBuilder{}
		renderError := builder.WithError(errs).WithHttpStatusCode(http.StatusBadRequest).Build()
		err := render.Render(w, r, &renderError)
		if err != nil {
			slog.Error(err.Error())
		}
		return
	}

	_, error := a.service.Login(data.Email, data.Password)
	if error != nil {
		builder := types.LinktlyErrorBuilder{}
		renderError := builder.WithError(error).Build()
		error := render.Render(w, r, &renderError)
		if error != nil {
			slog.Error(error.Error())
		}
		return
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:     "example_cookie",
			Value:    "cookie_value",
			Expires:  time.Now().Add(24 * time.Hour), // Set expiration time
			Path:     "/",
			HttpOnly: true,  // For security, HttpOnly cookies are not accessible via JavaScript
			Secure:   false, // Set to true if using HTTPS
		})
	}
}
