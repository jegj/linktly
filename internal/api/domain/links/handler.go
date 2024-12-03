package links

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	linktlyError "github.com/jegj/linktly/internal/api/error"
	"github.com/jegj/linktly/internal/api/jwt"
	"github.com/jegj/linktly/internal/api/response"
	"github.com/jegj/linktly/internal/api/types"
	"github.com/jegj/linktly/internal/api/validations"
)

type LinksHandler struct {
	service LinksService
}

func (l LinksHandler) CreateLink(w http.ResponseWriter, r *http.Request) error {
	userId := r.Context().Value(jwt.UserIdContextKey).(string)

	data := &LinkReq{}
	if err := render.Bind(r, data); err != nil {
		return response.InvalidJsonRequest()
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := validate.RegisterValidation("expires_at", validations.ExpiresAtValidation); err != nil {
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

	link := &Link{
		Name:        data.Name,
		Description: data.Description,
		Url:         data.Url,
		FolderId:    data.FolderId,
		AccountId:   userId,
		ExpiresAt:   data.ExpiresAt,
	}

	link, err := l.service.CreateLink(r.Context(), link)
	if err != nil {
		return err
	} else {
		resp := &LinkResp{
			Link: link,
		}
		return response.WriteJSON(w, r, http.StatusCreated, resp)
	}
}

func (l LinksHandler) GetLink(w http.ResponseWriter, r *http.Request) error {
	userId := r.Context().Value(jwt.UserIdContextKey).(string)
	id := chi.URLParam(r, "id")

	req := GetLinktByIdHandlerReq{
		Id: id,
	}
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(req)
	if err != nil {
		validationErrors := linktlyError.ValidatorFormatting(err.(validator.ValidationErrors))
		return response.InvalidRequestData(validationErrors)
	}

	link, err := l.service.GetLink(r.Context(), id, userId)
	if err != nil {
		return err
	} else {
		resp := &LinkResp{
			Link: link,
		}
		w.Header().Set("Cache-Control", "public, max-age=10")
		return response.WriteJSON(w, r, http.StatusCreated, resp)
	}
}
