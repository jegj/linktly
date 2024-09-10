package types

import (
	"net/http"
	"testing"

	"github.com/jackc/pgx/v5"
)

func TestNewLinktlyErrorBuilderDefaultValues(t *testing.T) {
	builder := LinktlyErrorBuilder{}
	linktlyError := builder.Build()

	if linktlyError.HTTPStatusCode != http.StatusInternalServerError {
		t.Errorf("expected http.StatusInternalServerError but got '%d'", linktlyError.HTTPStatusCode)
	}

	if linktlyError.StatusText != http.StatusText(http.StatusInternalServerError) {
		t.Errorf("expected Status text for http.StatusInternalServerError but got '%s'", linktlyError.StatusText)
	}

	if linktlyError.ErrorText != "Unknown error" {
		t.Errorf("expected 'Unknown error' by default but got '%s'", linktlyError.ErrorText)
	}
}

func TestNewLinktlyErrorBuilderNoRowsFromDatabaseError(t *testing.T) {
	builder := LinktlyErrorBuilder{}
	err := pgx.ErrNoRows
	builder.WithError(err)
	linktlyError := builder.Build()

	if linktlyError.HTTPStatusCode != http.StatusNotFound {
		t.Errorf("expected http.NotFound code but got '%d'", linktlyError.HTTPStatusCode)
	}

	if linktlyError.StatusText != http.StatusText(http.StatusNotFound) {
		t.Errorf("expected Status text for http.StatusNotFound but got '%s'", linktlyError.StatusText)
	}

	if linktlyError.ErrorText != "Not found entity" {
		t.Errorf("expected 'Not found entity' by default but got '%s'", linktlyError.ErrorText)
	}
}
