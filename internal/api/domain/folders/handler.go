package folders

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	linktlyError "github.com/jegj/linktly/internal/api/error"
	"github.com/jegj/linktly/internal/api/response"
)

type FolderHandler struct {
	service FolderService
}

func (f FolderHandler) CreateFolder(w http.ResponseWriter, r *http.Request) error {
	data := &FolderReq{}
	if err := render.Bind(r, data); err != nil {
		return response.InvalidJsonRequest()
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	errs := validate.Struct(data)
	if errs != nil {
		validationErrors := linktlyError.ValidatorFormatting(errs.(validator.ValidationErrors))
		return response.InvalidRequestData(validationErrors)
	}

	folder := &Folder{
		Name:           data.Name,
		ParentFolderId: data.ParentFolderId,
		AccountId:      data.AccountId,
		Description:    data.Description,
	}

	folder, err := f.service.CreateFolder(r.Context(), folder)
	if err != nil {
		return err
	} else {
		resp := &FolderResp{
			Folder: folder,
		}
		return response.WriteJSON(w, r, http.StatusCreated, resp)
	}
}
