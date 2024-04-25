package middlewares

import (
	"net/http"
	"strings"

	"github.com/educolog9/packages/functions"
	"github.com/educolog9/packages/types"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

type Key int

const UserClaimsKey Key = iota

// AdminMiddleware is a middleware that checks if the user is an admin
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, _ := opentracing.StartSpanFromContext(c.Request.Context(), "AdminMiddleware")
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
		// If the token is valid, check if the user is an admin
		// If the user is not an admin, return a 403 Forbidden
		// If the user is an admin, call c.Next()

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

		if !userClaims.IsAdmin() {
			response := types.ErrorResponse{
				Status:  http.StatusForbidden,
				Message: "Forbidden",
				Errors:  []string{"User is not an admin"},
			}
			c.JSON(http.StatusForbidden, response)
			c.Abort()
			return
		}

		c.Set("user_claims", userClaims)

		c.Next()
	}
}
