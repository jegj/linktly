package handlers

import (
	"log/slog"
	"net/http"

	"github.com/jegj/linktly/internal/api/response"
	"github.com/jegj/linktly/internal/api/types"
)

type CentralizedErrorHandler func(w http.ResponseWriter, r *http.Request) error

func (h CentralizedErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		if apiErr, ok := err.(types.APIError); ok {
			if err := response.WriteJSON(w, r, apiErr.StatusCode, apiErr); err != nil {
				slog.Error("Error writing JSON response", "err", err.Error())
			}
		} else {
			defaultErr := types.APIError{
				StatusCode: http.StatusInternalServerError,
				Msg:        "Internal Server error",
			}
			if err := response.WriteJSON(w, r, http.StatusInternalServerError, defaultErr); err != nil {
				slog.Error("Error writing JSON response", "err", err.Error())
			}
		}
		slog.Error("HTTP API error", "err", err.Error(), "path", r.URL.Path)
	}
}
