package middlewares

import (
	"net/http"
	"strconv"

	"github.com/educolog9/packages/types"
	"github.com/gin-gonic/gin"
)

// ParseIntIDMiddlware is a middleware function that parses an integer ID from the request parameter and sets it in the context.
// If the ID is invalid, it returns a JSON response with a status code of 400 Bad Request and an error message.
// The middleware function is designed to be used with the Gin framework.
func ParseIntIDMiddlware() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			response := types.ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "Invalid ID",
				Errors:  []string{err.Error()},
			}
			c.JSON(http.StatusBadRequest, response)
			c.Abort()
			return
		}

		c.Set("intID", id)
		c.Next()
	}
}
