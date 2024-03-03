package validations

import (
	"net/url"

	"github.com/go-playground/validator/v10"
)

func validateURL(fl validator.FieldLevel) bool {
	urlStr := fl.Field().String()

	// Check if the URL can be parsed
	u, err := url.Parse(urlStr)
	if err != nil {
		return false
	}

	// Check if the URL has a scheme and host
	if u.Scheme == "" || u.Host == "" {
		return false
	}

	// Add any additional URL validation logic here
	return true
}
