package validations

import (
	"net/http"

	"github.com/educolog9/packages/errors/messages"
	languagues "github.com/educolog9/packages/languages"
	"github.com/educolog9/packages/types"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidateJSON(obj interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		var errors []string

		language, ok := c.Get("language")
		if !ok {
			language = languagues.Spanish.String()
		}

		trans, _ := Uni.GetTranslator(language.(string))

		if err := c.ShouldBind(obj); err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				errors = append(errors, err.Translate(trans))
			}
		}

		if err := Validate.Struct(obj); err != nil {
			if errs, ok := err.(validator.ValidationErrors); ok {
				for _, e := range errs {
					errors = append(errors, e.Translate(trans))
				}
			} else {
				errors = append(errors, err.Error())
			}
		}

		if len(errors) > 0 {
			response := types.ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: messages.ValidationFailed,
				Errors:  errors,
			}

			c.JSON(http.StatusBadRequest, response)
			c.Abort()
			return
		}

		c.Set("validatedJSON", obj)
		c.Next()
	}
}
