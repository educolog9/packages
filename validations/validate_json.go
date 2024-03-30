package validations

import (
	"net/http"

	"github.com/educolog7/packages/errors/messages"
	"github.com/educolog7/packages/types"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func ValidateJSON(obj interface{}, trans ut.Translator) gin.HandlerFunc {
	return func(c *gin.Context) {
		var errors []string

		if trans == nil {
			en := en.New()
			uni := ut.New(en, nil)
			trans, _ = uni.GetTranslator("en")
		}

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
