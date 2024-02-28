package middlewares

import (
	"net/http"
	"strings"

	"github.com/educolog7/packages/enums"
	"github.com/educolog7/packages/functions"
	"github.com/educolog7/packages/types"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

// RoleBasedAuthMiddleware is a middleware that checks if the user has one of the specified roles
func RoleBasedAuthMiddleware(allowedRoles []enums.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		span, _ := opentracing.StartSpanFromContext(c.Request.Context(), "RoleBasedAuthMiddleware")
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

		// Check if the user's role is in the allowedRoles array
		isAllowed := false
		for _, userRole := range userClaims.Roles {
			for _, allowedRole := range allowedRoles {
				if userRole == allowedRole {
					isAllowed = true
					break
				}
			}
			if isAllowed {
				break
			}
		}

		if !isAllowed {
			response := types.ErrorResponse{
				Status:  http.StatusForbidden,
				Message: "Forbidden",
				Errors:  []string{"User does not have the required role"},
			}
			c.JSON(http.StatusForbidden, response)
			c.Abort()
			return
		}

		c.Set("user_claims", userClaims)

		c.Next()
	}
}
