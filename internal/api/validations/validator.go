package validations

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// Password validation function
// At least 8 characters, one uppercase letter, one lowercase letter, one digit, and one special character
func PasswordValidation(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	return ValidatePassword(password)
}

func ExpiresAtValidation(fl validator.FieldLevel) bool {
	// Ensure the field is a Time type
	expiresAt, ok := fl.Field().Interface().(time.Time)
	now := time.Now()
	if !ok {
		return false
	}

	return ValidateExpiresAt(expiresAt, now)
}
