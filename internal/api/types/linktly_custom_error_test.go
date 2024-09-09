package types

import (
	"net/http"
	"testing"
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
