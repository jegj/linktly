package validations

import (
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

// TODO: Tests
// Password validation function
// At least 8 characters, one uppercase letter, one lowercase letter, one digit, and one special character
func PasswordValidation(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	return containsLowercase(password) && containsUppercase(password)
}

// Custom validation for lowercase letters
func containsLowercase(password string) bool {
	re := regexp.MustCompile(`[a-z]`)
	return re.MatchString(password)
}

// Custom validation for uppercase letters
func containsUppercase(password string) bool {
	re := regexp.MustCompile(`[A-Z]`)
	return re.MatchString(password)
}

// Custom validation for digits
func containsDigit(fl validator.FieldLevel) bool {
	// Regular expression to check for at least one digit
	re := regexp.MustCompile(`\d`)
	return re.MatchString(fl.Field().String())
}

// Custom validation for special characters
func containsSpecial(fl validator.FieldLevel) bool {
	// Regular expression to check for at least one special character
	re := regexp.MustCompile(`[^a-zA-Z0-9]`)
	return re.MatchString(fl.Field().String())
}

func ExpiresAtValidation(fl validator.FieldLevel) bool {
	// Ensure the field is a Time type
	expiresAt, ok := fl.Field().Interface().(time.Time)
	if !ok {
		return false
	}

	// Check if the time is after the current time
	return expiresAt.After(time.Now())
}
