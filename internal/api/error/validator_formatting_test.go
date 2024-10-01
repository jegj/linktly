package error

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

type inputReq struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6,max=30"`
	Name     string `validate:"required,min=6,max=255"`
}

func compareMaps(map1, map2 map[string]string) bool {
	// Step 1: Check if the lengths are the same
	if len(map1) != len(map2) {
		return false
	}

	for key, value1 := range map1 {
		if value2, exists := map2[key]; !exists || value1 != value2 {
			return false
		}
	}

	return true
}

func TestValidatorFormatting(t *testing.T) {
	tests := []struct {
		want     map[string]string
		input    inputReq
		testName string
	}{
		{
			testName: "Must return a map with the error messages per tag",
			input: inputReq{
				Email:    "no_at_email",
				Password: "123",
				Name:     "n",
			},
			want: map[string]string{
				"email":    "Key: 'inputReq.Email' Error:Field validation for 'Email' failed on the 'email' tag",
				"name":     "Key: 'inputReq.Name' Error:Field validation for 'Name' failed on the 'min' tag",
				"password": "Key: 'inputReq.Password' Error:Field validation for 'Password' failed on the 'min' tag",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			validate := validator.New(validator.WithRequiredStructEnabled())
			errs := validate.Struct(tt.input)
			got := ValidatorFormatting(errs.(validator.ValidationErrors))
			if !compareMaps(got, tt.want) {
				t.Errorf("ValidatorFormatting(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
