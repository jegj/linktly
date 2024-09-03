package types

import (
	"net/http"

	"github.com/go-chi/render"
)

type LinktlyError struct {
	Err            error  `json:"-"`
	StatusText     string `json:"status"`
	ErrorText      string `json:"error"`
	HTTPStatusCode int    `json:"-"`
}

func (e *LinktlyError) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func NewLinktlyError(err error, code int, statusText string) *LinktlyError {
	return &LinktlyError{
		Err:            err,
		HTTPStatusCode: code,
		StatusText:     statusText,
		ErrorText:      err.Error(),
	}
}

// TODO: MAKE Factory for error from database o interal code errors
