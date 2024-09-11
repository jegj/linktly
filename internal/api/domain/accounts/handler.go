package accounts

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/jegj/linktly/internal/api/types"
	"github.com/jegj/linktly/internal/api/validations"
)

type AccountHandler struct {
	service AccountService
}

func (s AccountHandler) GetAccountByIdHandler(w http.ResponseWriter, r *http.Request) {
	// slog.Debug("Debug message accounts")
	// slog.Info("Info message accounts")

	id := chi.URLParam(r, "id")
	req := GetAccountByIdHandlerReq{
		Id: id,
	}
	validate := validator.New(validator.WithRequiredStructEnabled())
	errs := validate.Struct(req)

	if errs != nil {
		builder := types.LinktlyErrorBuilder{}
		renderError := builder.WithError(errs).WithHttpStatusCode(http.StatusBadRequest).Build()
		err := render.Render(w, r, &renderError)
		if err != nil {
			slog.Error(err.Error())
		}
		return
	}

	account, error := s.service.GetAccountById(id)
	builder := types.LinktlyErrorBuilder{}
	renderError := builder.WithError(error).Build()
	if error != nil {
		err := render.Render(w, r, &renderError)
		if err != nil {
			slog.Error(err.Error())
		}
		return
	} else {
		resp := &AccountResp{
			Account: account,
		}
		render.Status(r, http.StatusOK)
		err := render.Render(w, r, resp)
		if err != nil {
			slog.Error(err.Error())
		}
		return
	}
}

func (s AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	data := &CreateAccountReq{}
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

	account := &Account{
		Name:     data.Name,
		LastName: data.LastName,
		Password: data.Password,
		Email:    data.Email,
	}

	newId, err := s.service.CreateAccount(account)
	if err != nil {

		builder := types.LinktlyErrorBuilder{}
		renderError := builder.WithError(err).Build()
		error := render.Render(w, r, &renderError)
		if error != nil {
			slog.Error(error.Error())
		}
		return
	} else {
		account.Id = newId
		resp := &AccountResp{
			Account: account,
		}
		render.Status(r, http.StatusCreated)
		err = render.Render(w, r, resp)
		if err != nil {
			slog.Error(err.Error())
		}
		return
	}
}
