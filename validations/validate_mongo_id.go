package validations

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func validateMongoID(fl validator.FieldLevel) bool {
	idStr := fl.Field().String()

	// Check if the string is a valid MongoDB ObjectID
	_, err := primitive.ObjectIDFromHex(idStr)
	return err == nil
}
