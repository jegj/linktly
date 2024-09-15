package accounts

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/jegj/linktly/internal/api/response"
)

type AccountHandler struct {
	service AccountService
}

func (s AccountHandler) GetAccountByIdHandler(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")
	req := GetAccountByIdHandlerReq{
		Id: id,
	}
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(req)
	if err != nil {
		// TODO: MAKE BETTER WAY TO COLLECT ERRORS HERE
		validationErrors := make(map[string]string)
		validationErrors["Id"] = err.Error()
		return response.InvalidRequestData(validationErrors)
	}

	account, err := s.service.GetAccountById(id)
	if err != nil {
		return err
	} else {
		resp := &AccountResp{
			Account: account,
		}
		return response.WriteJSON(w, r, http.StatusOK, resp)
	}
}

/*
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

	account, err := s.service.CreateAccount(account)
	if err != nil {

		builder := types.LinktlyErrorBuilder{}
		renderError := builder.WithError(err).Build()
		error := render.Render(w, r, &renderError)
		if error != nil {
			slog.Error(error.Error())
		}
		return
	} else {
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
*/
