package middlewares

import (
	"net/http"

	"github.com/educolog9/packages/types"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ParseMongoIDMiddleware is a middleware function that parses a MongoDB ObjectID from the request parameter and sets it in the context.
// If the ID is invalid, it returns a JSON response with a status code of 400 Bad Request and an error message.
// The middleware function is designed to be used with the Gin framework.
func ParseMongoIDMiddleware(paramName ...string) gin.HandlerFunc {
	param := "id"
	if len(paramName) > 0 {
		param = paramName[0]
	}

	return func(c *gin.Context) {
		idStr := c.Param(param)
		_, err := primitive.ObjectIDFromHex(idStr)
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

		c.Set(param, idStr)
		c.Next()
	}
}
