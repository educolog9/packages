package middlewares

import (
	languagues "github.com/educolog9/packages/languages"
	"github.com/gin-gonic/gin"
)

// LanguageMiddleware is a middleware that extracts the language from the Content-Language header
func LanguageMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Content-Language header
		lang := c.GetHeader("Content-Language")

		// If the header is not set, default to "es"
		if lang == "" {
			lang = languagues.Spanish.String()
		}

		// Set the language in the context
		c.Set("language", lang)

		c.Next()
	}
}
