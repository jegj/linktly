package folders

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/jegj/linktly/internal/api/domain/links"
	linktlyError "github.com/jegj/linktly/internal/api/error"
	"github.com/jegj/linktly/internal/api/jwt"
	"github.com/jegj/linktly/internal/api/response"
	"github.com/jegj/linktly/internal/api/types"
	"github.com/jegj/linktly/internal/api/validations"
)

type FolderHandler struct {
	service     FolderService
	linkService links.LinksService
}

func (f FolderHandler) CreateFolder(w http.ResponseWriter, r *http.Request) error {
	userId := r.Context().Value(jwt.UserIdContextKey).(string)

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
		AccountId:      userId,
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

func (f FolderHandler) GetFoldersByUserId(w http.ResponseWriter, r *http.Request) error {
	userId := r.Context().Value(jwt.UserIdContextKey).(string)

	folders, err := f.service.GetFoldersByUserId(r.Context(), userId)
	if err != nil {
		return err
	} else {
		folderResponses := make([]render.Renderer, len(folders))
		for i, folder := range folders {
			folderResponses[i] = &FolderResp{
				Folder: folder,
			}
		}
		w.Header().Set("Cache-Control", "public, max-age=20")
		return response.WriteJSONCollection(w, r, http.StatusOK, folderResponses)
	}
}

func (f FolderHandler) DeleteFoldersByIdAndUserId(w http.ResponseWriter, r *http.Request) error {
	userId := r.Context().Value(jwt.UserIdContextKey).(string)
	folderId := chi.URLParam(r, "id")

	err := f.service.DeleteFoldersByIdAndUserId(r.Context(), folderId, userId)
	if err != nil {
		return err
	} else {
		resp := &FolderDeleteResp{
			Id: folderId,
		}
		return response.WriteJSON(w, r, http.StatusOK, resp)
	}
}

func (f FolderHandler) PatchFoldersByIdAndUserId(w http.ResponseWriter, r *http.Request) error {
	userId := r.Context().Value(jwt.UserIdContextKey).(string)
	folderId := chi.URLParam(r, "id")

	data := &FolderPatchReq{}
	if err := render.Bind(r, data); err != nil {
		return response.InvalidJsonRequest()
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	errs := validate.Struct(data)
	if errs != nil {
		validationErrors := linktlyError.ValidatorFormatting(errs.(validator.ValidationErrors))
		return response.InvalidRequestData(validationErrors)
	}

	folderReq := &Folder{
		Name:        data.Name,
		Description: data.Description,
	}

	folder, err := f.service.PatchFolderByIdAndUserId(r.Context(), folderId, userId, folderReq)
	if err != nil {
		return err
	} else {
		resp := &FolderResp{
			Folder: folder,
		}
		return response.WriteJSON(w, r, http.StatusOK, resp)
	}
}

func (f FolderHandler) GetFolderByIdAndUserId(w http.ResponseWriter, r *http.Request) error {
	userId := r.Context().Value(jwt.UserIdContextKey).(string)
	folderId := chi.URLParam(r, "id")

	folder, err := f.service.GetFolderByIdAndUserId(r.Context(), folderId, userId)
	if err != nil {
		return err
	} else {
		resp := &FolderResp{
			Folder: folder,
		}
		w.Header().Set("Cache-Control", "public, max-age=10")
		return response.WriteJSON(w, r, http.StatusOK, resp)
	}
}

func (l FolderHandler) CreateLink(w http.ResponseWriter, r *http.Request) error {
	userId := r.Context().Value(jwt.UserIdContextKey).(string)
	folderId := chi.URLParam(r, "id")

	// TODO: validate folderId as uuid

	data := &links.LinkReq{}
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

	link := &links.Link{
		Name:        data.Name,
		Description: data.Description,
		Url:         data.Url,
		FolderId:    &folderId,
		AccountId:   userId,
		ExpiresAt:   data.ExpiresAt,
	}

	link, err := l.linkService.CreateLink(r.Context(), link)
	if err != nil {
		return err
	} else {
		resp := &links.LinkResp{
			Link: link,
		}
		return response.WriteJSON(w, r, http.StatusCreated, resp)
	}
}
