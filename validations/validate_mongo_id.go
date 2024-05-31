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

func validateMongoIDs(fl validator.FieldLevel) bool {
	ids := fl.Field()

	for i := 0; i < ids.Len(); i++ {
		idStr := ids.Index(i).String()

		// Check if the string is a valid MongoDB ObjectID
		_, err := primitive.ObjectIDFromHex(idStr)
		if err != nil {
			return false
		}
	}

	return true
}
