package middlewares

import (
	"net/http"
	"strings"

	"github.com/educolog9/packages/functions"
	"github.com/educolog9/packages/types"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

// AuthMiddleware is a middleware that checks if the user is authenticated
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, _ := opentracing.StartSpanFromContext(c.Request.Context(), "AuthMiddleware")
		defer span.Finish()

		authHeader := c.GetHeader("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) != 2 {
			response := types.ErrorResponse{
				Status:  http.StatusUnauthorized,
				Message: "Unauthorized",
				Errors:  []string{"Invalid authorization header format"},
			}
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		// Check if the token is valid
		// If the token is not valid, return a 401 Unauthorized
		// If the token is valid, call c.Next()
		userClaims, err := functions.ValidateToken(bearerToken[1])
		if err != nil {
			response := types.ErrorResponse{
				Status:  http.StatusUnauthorized,
				Message: "Unauthorized",
				Errors:  []string{err.Error()},
			}
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		c.Set("userClaims", userClaims)

		c.Next()
	}
}
