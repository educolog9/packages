package validations

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// validatePhoneNumber validates if a given phone number is in a valid format.
// It uses a regular expression to match phone numbers in the format +1234567890, 123-456-7890, (123) 456-7890, and 123 456 7890.
// The function takes a validator.FieldLevel parameter, which provides access to the field being validated.
// It returns a boolean value indicating whether the phone number is valid or not.
func validatePhoneNumber(fl validator.FieldLevel) bool {
	phoneNumber := fl.Field().String()

	// Check if the phone number is valid
	// This regex matches phone numbers in the format +1234567890, 123-456-7890, (123) 456-7890, and 123 456 7890
	phoneRegex := `^\+?(\d{1,3})?[-. (]*(\d{1,3})[-. )]*(\d{1,4})[-. ]*(\d{1,9})$`
	re := regexp.MustCompile(phoneRegex)
	return re.MatchString(phoneNumber)
}
