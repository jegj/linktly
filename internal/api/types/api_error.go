package types

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)

// TODO: Set up friendly message to the user
type APIError struct {
	Msg        any `json:"msg"`
	StatusCode int `json:"statusCode"`
}

func (e APIError) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.StatusCode)
	return nil
}

func (e APIError) Error() string {
	return fmt.Sprintf("HTTP Status Code: %v - %v", e.StatusCode, e.Msg)
}
