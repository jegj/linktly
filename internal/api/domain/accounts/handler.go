package accounts

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/jegj/linktly/internal/api/types"
)

type AccountHandler struct {
	service AccountService
}

func (s AccountHandler) GetAccountByIdHandler(w http.ResponseWriter, r *http.Request) {
	slog.Debug("Debug message accounts")
	slog.Info("Info message accounts")

	id := chi.URLParam(r, "id")
	validate := validator.New(validator.WithRequiredStructEnabled())
	// FIXME: errs comes with an empty key
	// fmt.Println(errs) // output: Key: "" Error:Field validation for "" failed on the "email" tag
	errs := validate.Var(id, "required,uuid")

	if errs != nil {
		builder := types.LinktlyErrorBuilder{}
		renderError := builder.WithError(errs).Build()
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
	// TODO: Add validations
	data := &AccountReq{}
	if err := render.Bind(r, data); err != nil {
		err := render.Render(w, r, types.NewLinktlyError(err, http.StatusBadRequest, http.StatusText(http.StatusBadRequest)))
		if err != nil {
			slog.Error(err.Error())
		}
		return
	}
	account := data.Account
	newId, err := s.service.CreateAccount(account)
	if err != nil {
		error := render.Render(w, r, types.NewLinktlyError(err, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)))
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
