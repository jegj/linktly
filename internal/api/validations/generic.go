package validations

import (
	"regexp"
	"time"
)

func ValidatePassword(password string) bool {
	return containsLowercase(password) && containsUppercase(password) && containsDigit(password) && containsSpecial(password)
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
func containsDigit(password string) bool {
	// Regular expression to check for at least one digit
	re := regexp.MustCompile(`\d`)
	return re.MatchString(password)
}

// Custom validation for special characters
func containsSpecial(password string) bool {
	// Regular expression to check for at least one special character
	re := regexp.MustCompile(`[^a-zA-Z0-9]`)
	return re.MatchString(password)
}

func ValidateExpiresAt(expiresAt time.Time, now time.Time) bool {
	return expiresAt.After(now)
}
