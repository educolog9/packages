package functions

import (
	"net/http"

	customerrors "github.com/educolog7/packages/errors/custom_errors"
	"github.com/gin-gonic/gin"
)

func HandleErrorWithHttpStatus(c *gin.Context, httpStatus int, err error) {
	status := httpStatus
	message := ""
	errors := []string{err.Error()}

	switch {
	case httpStatus == http.StatusInternalServerError:
		message = "Internal Server Error"
	case httpStatus == http.StatusNotFound:
		message = "Your requested item is not found"
	case httpStatus == http.StatusConflict:
		message = "This item already exists"
	case httpStatus == http.StatusBadRequest:
		message = "Given param is not valid"
	default:
		status = http.StatusInternalServerError
		message = "An unexpected error occurred"
	}

	c.JSON(status, gin.H{
		"status":  status,
		"message": message,
		"errors":  errors,
	})
}

// HandleError handles the given error and sends a JSON response with the error details.
func HandleError(c *gin.Context, err customerrors.BaseErrorInterface) {
	var errors []string

	if err.GetInner() != nil {
		errors = append(errors, err.GetInner().Error())
	}

	c.JSON(err.GetCode(), gin.H{
		"status":  err.GetCode(),
		"message": err.GetMessage(),
		"errors":  errors,
	})
}
