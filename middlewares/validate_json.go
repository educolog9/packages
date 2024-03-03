package middlewares

import (
	"net/http"

	"github.com/educolog7/packages/errors/messages"
	"github.com/educolog7/packages/types"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidateJSON(obj interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		var errors []string

		if err := c.ShouldBind(obj); err != nil {

			for _, err := range err.(validator.ValidationErrors) {
				errors = append(errors, err.Error())
			}

		}

		validate := validator.New()
		if err := validate.Struct(obj); err != nil {
			if errs, ok := err.(validator.ValidationErrors); ok {
				for _, e := range errs {
					errors = append(errors, e.Error())
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
