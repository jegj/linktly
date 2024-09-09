package types

import (
	"errors"
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

type LinktlyErrorBuilder struct {
	HTTPStatusCode *int
	StatusText     *string
	Err            *error
	ErrorText      *string
}

func (b *LinktlyErrorBuilder) WithHttpStatusCode(httpStatusCode int) *LinktlyErrorBuilder {
	b.HTTPStatusCode = &httpStatusCode
	return b
}

func (b *LinktlyErrorBuilder) Build() LinktlyError {
	if b.HTTPStatusCode == nil {
		defaultStatusCode := http.StatusInternalServerError
		b.HTTPStatusCode = &defaultStatusCode
	}

	if b.StatusText == nil {
		defaultStatusText := http.StatusText(http.StatusInternalServerError)
		b.StatusText = &defaultStatusText
	}

	if b.Err == nil {
		defaultErr := errors.New("Unknown error")
		defaultErrText := defaultErr.Error()
		b.Err = &defaultErr
		b.ErrorText = &defaultErrText
	}

	return LinktlyError{
		HTTPStatusCode: *b.HTTPStatusCode,
		StatusText:     *b.StatusText,
		Err:            *b.Err,
		ErrorText:      *b.ErrorText,
	}
}
