package middlewares

import "github.com/gin-gonic/gin"

// CORSMiddleware is a middleware function that adds Cross-Origin Resource Sharing (CORS) headers to the response.
// It allows requests from any origin and includes common headers and methods used in HTTP requests.
// If the request method is OPTIONS, it responds with a 204 No Content status code.
// This middleware should be used to enable CORS support in your application.
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
