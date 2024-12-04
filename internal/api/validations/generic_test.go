package validations

import (
	"testing"
	"time"
)

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		testName string
		input    string
		want     bool
	}{
		{
			testName: "Must contain at leats a lowercase letter",
			input:    "TEST",
			want:     false,
		},
		{
			testName: "Must contain at leats a uppercase letter",
			input:    "test",
			want:     false,
		},
		{
			testName: "Must contain at leats a digit",
			input:    "Test",
			want:     false,
		},
		{
			testName: "Must contain at leats a special character",
			input:    "Test1",
			want:     false,
		},
		{
			testName: "Must contain at least 8 characters",
			input:    "Test_1",
			want:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			got := ValidatePassword(tt.input)
			if got != tt.want {
				t.Errorf("PasswordValidation( %v ) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestValidateExpiresAt(t *testing.T) {
	tests := []struct {
		inputExpireAt time.Time
		inputNow      time.Time
		testName      string
		want          bool
	}{
		{
			testName:      "ExpireAt after Now",
			inputExpireAt: time.Date(2024, 12, 5, 12, 0, 0, 0, time.UTC),
			inputNow:      time.Date(2024, 12, 4, 12, 0, 0, 0, time.UTC),
			want:          true,
		},
		{
			testName:      "ExpireAt equals Now",
			inputExpireAt: time.Date(2024, 12, 4, 12, 0, 0, 0, time.UTC),
			inputNow:      time.Date(2024, 12, 4, 12, 0, 0, 0, time.UTC),
			want:          false,
		},
		{
			testName:      "ExpireAt before Now",
			inputExpireAt: time.Date(2024, 12, 3, 12, 0, 0, 0, time.UTC),
			inputNow:      time.Date(2024, 12, 4, 12, 0, 0, 0, time.UTC),
			want:          false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			got := ValidateExpiresAt(tt.inputExpireAt, tt.inputNow)
			if got != tt.want {
				t.Errorf("Test %q failed: got %v, want %v", tt.testName, got, tt.want)
			}
		})
	}
}
