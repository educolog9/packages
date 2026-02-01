package middlewares

import (
	"net/http"
	"strings"

	"github.com/educolog9/packages/enums"
	"github.com/educolog9/packages/functions"
	"github.com/educolog9/packages/types"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

// OrgAdminMiddleware is a middleware that allows access to Admin and OrgAdmin users.
// - Admin: Full access without restrictions
// - OrgAdmin: Access allowed, but handlers MUST validate that operations only affect their organization
//
// The middleware sets "userClaims" in the context, and handlers should use:
// - userClaims.IsAdmin() to check if user has unrestricted access
// - userClaims.IsOrgAdmin() to check if user is an org admin (requires org validation)
// - userClaims.OrganizationID to get the user's organization for validation
func OrgAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, _ := opentracing.StartSpanFromContext(c.Request.Context(), "OrgAdminMiddleware")
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

		// Check if user is Admin or OrgAdmin
		isAllowed := false
		for _, role := range userClaims.Roles {
			if role == enums.Admin || role == enums.OrgAdmin {
				isAllowed = true
				break
			}
		}

		if !isAllowed {
			response := types.ErrorResponse{
				Status:  http.StatusForbidden,
				Message: "Forbidden",
				Errors:  []string{"User does not have admin or org_admin role"},
			}
			c.JSON(http.StatusForbidden, response)
			c.Abort()
			return
		}

		c.Set("userClaims", userClaims)

		c.Next()
	}
}
