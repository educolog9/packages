package validations

import (
	"regexp"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
)

// validateCardNumber checks if the card number is valid using Luhn algorithm
func validateCardNumber(fl validator.FieldLevel) bool {
	cardNum := fl.Field().String()
	if matched, _ := regexp.MatchString(`^\d+$`, cardNum); !matched {
		return false
	}

	var sum int
	nDigits := len(cardNum)
	parity := nDigits % 2
	for i := 0; i < nDigits; i++ {
		digit := int(cardNum[i] - '0')
		if i%2 == parity {
			digit = digit * 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}
	return sum%10 == 0
}

// validateCSV checks if the CSV is valid (3 or 4 digits)
func validateCSV(fl validator.FieldLevel) bool {
	csv := fl.Field().String()
	return regexp.MustCompile(`^\d{3,4}$`).MatchString(csv)
}

// validateExpiryMonth checks if the expiry month is valid (1 to 12)
func validateExpiryMonth(fl validator.FieldLevel) bool {
	month, err := strconv.Atoi(fl.Field().String())
	if err != nil {
		return false
	}
	return month >= 1 && month <= 12
}

// validateExpiryYear checks if the expiry year is valid (current year or later in YY format)
func validateExpiryYear(fl validator.FieldLevel) bool {
	year, err := strconv.Atoi(fl.Field().String())
	if err != nil {
		return false
	}
	currentYear := time.Now().Year() % 100 // Get last two digits of the current year
	return year >= currentYear
}
