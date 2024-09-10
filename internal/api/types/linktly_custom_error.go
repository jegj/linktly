package types

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
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
	statusText := http.StatusText(httpStatusCode)
	b.StatusText = &statusText
	return b
}

func (b *LinktlyErrorBuilder) WithError(err error) *LinktlyErrorBuilder {
	errMessage := err.Error()
	errorText := ""
	httpStatusCode := 0
	httpStatusText := ""
	var validationError *validator.InvalidValidationError

	fmt.Printf("====>%v", err)
	fmt.Printf("====>%v", errors.As(err, &validationError))

	if invalidErr, ok := err.(*validator.InvalidValidationError); ok {
		fmt.Println("is invalid error vlaidation\n")
	}

	if errors.As(err, &validationError) {
		httpStatusCode = http.StatusBadRequest
		httpStatusText = http.StatusText(http.StatusBadRequest)
		errorText = "Not found entity"
	} else if errors.Is(err, pgx.ErrNoRows) {
		httpStatusCode = http.StatusNotFound
		httpStatusText = http.StatusText(http.StatusNotFound)
		errorText = "Not found entity"
	} else {
		httpStatusCode = http.StatusInternalServerError
		httpStatusText = http.StatusText(http.StatusInternalServerError)
		errorText = "Something worng happeneded"
	}

	b.Err = &err
	b.ErrorText = &errMessage
	b.ErrorText = &errorText
	b.HTTPStatusCode = &httpStatusCode
	b.StatusText = &httpStatusText

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
