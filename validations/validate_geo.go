package validations

import (
	"github.com/go-playground/validator/v10"
)

func validateLatitude(fl validator.FieldLevel) bool {
	lat := fl.Field().Float()

	// Check if the latitude is a valid (between -90 and 90)
	return lat >= -90 && lat <= 90
}

func validateLongitude(fl validator.FieldLevel) bool {
	lon := fl.Field().Float()

	// Check if the longitude is a valid (between -180 and 180)
	return lon >= -180 && lon <= 180
}
