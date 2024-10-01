package error

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidatorFormatting(validationsErrors validator.ValidationErrors) map[string]string {
	mappedValidationErrors := make(map[string]string)
	for _, err := range validationsErrors {
		mappedValidationErrors[strings.ToLower(err.Field())] = err.Error()
	}
	return mappedValidationErrors
}
