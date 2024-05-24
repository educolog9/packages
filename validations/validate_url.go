package validations

import (
	"net/url"
	"reflect"
	"strings"

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

func validateURLArray(fl validator.FieldLevel) bool {
	if reflect.TypeOf(fl.Field().Interface()).Kind() != reflect.Slice {
		return false
	}

	s := reflect.ValueOf(fl.Field().Interface())

	for i := 0; i < s.Len(); i++ {
		urlStr := s.Index(i).String()

		u, err := url.Parse(urlStr)
		if err != nil {
			return false
		}

		if u.Scheme == "" || u.Host == "" {
			return false
		}
	}

	return true
}

func validatePictureURL(fl validator.FieldLevel) bool {
	urlStr := fl.Field().String()

	u, err := url.Parse(urlStr)
	if err != nil {
		return false
	}

	if u.Scheme == "" || u.Host == "" {
		return false
	}

	// Check if the URL is an image URL
	if !strings.HasSuffix(u.Path, ".jpg") && !strings.HasSuffix(u.Path, ".png") && !strings.HasSuffix(u.Path, ".jpeg") && !strings.HasSuffix(u.Path, ".svg") && !strings.HasSuffix(u.Path, ".webp") {
		return false
	}

	// Add any additional URL validation logic here
	return true
}

func validatePicturesURL(fl validator.FieldLevel) bool {
	if reflect.TypeOf(fl.Field().Interface()).Kind() != reflect.Slice {
		return false
	}

	s := reflect.ValueOf(fl.Field().Interface())

	for i := 0; i < s.Len(); i++ {
		urlStr := s.Index(i).String()

		u, err := url.Parse(urlStr)
		if err != nil {
			return false
		}

		if u.Scheme == "" || u.Host == "" {
			return false
		}

		if !strings.HasSuffix(u.Path, ".jpg") && !strings.HasSuffix(u.Path, ".png") && !strings.HasSuffix(u.Path, ".jpeg") && !strings.HasSuffix(u.Path, ".svg") {
			return false
		}
	}

	return true
}
