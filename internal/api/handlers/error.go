package handlers

import (
	"log/slog"
	"net/http"

	"github.com/jegj/linktly/internal/api/response"
	"github.com/jegj/linktly/internal/api/types"
)

type CustomHandler func(w http.ResponseWriter, r *http.Request) error

func (h CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		if apiErr, ok := err.(types.APIError); ok {
			response.WriteJSON(w, r, apiErr.StatusCode, apiErr)
		} else {
			defaultErr := types.APIError{
				StatusCode: http.StatusInternalServerError,
				Msg:        "Internal Server error",
			}
			response.WriteJSON(w, r, http.StatusInternalServerError, defaultErr)
		}
		slog.Error("HTTP API error", "err", err.Error(), "path", r.URL.Path)
	}
}
