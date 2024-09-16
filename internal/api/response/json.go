package response

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/jegj/linktly/internal/api/types"
)

func WriteJSON(w http.ResponseWriter, r *http.Request, httpStatus int, responseData render.Renderer) error {
	render.Status(r, httpStatus)
	return render.Render(w, r, responseData)
}

func InvalidRequestData(errors map[string]string) types.APIError {
	return types.APIError{
		Msg:        errors,
		StatusCode: http.StatusBadRequest,
	}
}

func InvalidJsonRequest() types.APIError {
	return types.APIError{
		Msg:        "Invalid JSON request data",
		StatusCode: http.StatusBadRequest,
	}
}
