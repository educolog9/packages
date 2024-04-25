package middlewares

import (
	"net/http"

	"github.com/educolog9/packages/functions"
	"github.com/gin-gonic/gin"
)

// ParsePaginationParams is a middleware function that parses pagination parameters from the request.
// It extracts the pagination parameters from the request and sets them in the context for further processing.
// If there is an error while parsing the parameters, it returns a bad request error.
func ParsePaginationParams() gin.HandlerFunc {
	return func(c *gin.Context) {
		pagination, err := functions.ParsePaginationParams(c)
		if err != nil {
			functions.HandleErrorWithHttpStatus(c, http.StatusBadRequest, err)
			c.Abort()
			return
		}

		c.Set("pagination", pagination)
		c.Next()
	}
}
