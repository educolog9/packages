package validations

import "github.com/go-playground/validator/v10"

var Validate *validator.Validate

func init() {
	// Initialize the validator and register the custom URL validation function
	Validate = validator.New()
	_ = Validate.RegisterValidation("customurl", validateURL)
}
