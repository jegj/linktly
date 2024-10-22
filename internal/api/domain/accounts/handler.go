package accounts

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	linktlyError "github.com/jegj/linktly/internal/api/error"
	"github.com/jegj/linktly/internal/api/response"
	"github.com/jegj/linktly/internal/api/types"
	"github.com/jegj/linktly/internal/api/validations"
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
		validationErrors := linktlyError.ValidatorFormatting(err.(validator.ValidationErrors))
		return response.InvalidRequestData(validationErrors)
	}

	account, err := s.service.GetAccountById(r.Context(), id)
	if err != nil {
		return err
	} else {
		resp := &AccountResp{
			Account: account,
		}
		return response.WriteJSON(w, r, http.StatusOK, resp)
	}
}

func (s AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) error {
	data := &CreateAccountReq{}
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
		validationErrors := linktlyError.ValidatorFormatting(errs.(validator.ValidationErrors))
		return response.InvalidRequestData(validationErrors)
	}

	account := &Account{
		Name:     data.Name,
		LastName: data.LastName,
		Password: data.Password,
		Email:    data.Email,
	}

	account, err := s.service.CreateAccount(r.Context(), account)
	if err != nil {
		return err
	} else {
		resp := &AccountResp{
			Account: account,
		}
		return response.WriteJSON(w, r, http.StatusCreated, resp)
	}
}
