package validations

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func validateCountryISOCode(fl validator.FieldLevel) bool {
	countryCode := fl.Field().String()

	// Check if the country code is valid (only "US" and "DO" are valid)
	validCodes := []string{"US", "DO"}
	for _, code := range validCodes {
		if strings.EqualFold(code, countryCode) {
			return true
		}
	}
	return false
}
