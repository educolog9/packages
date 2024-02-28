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
		if err := c.ShouldBind(obj); err != nil {
			var errors []string

			for _, err := range err.(validator.ValidationErrors) {
				errors = append(errors, err.Error())
			}

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
