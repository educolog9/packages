package middlewares

import (
	"github.com/educolog7/packages/countries"
	"github.com/gin-gonic/gin"
)

// CountryMiddleware is a middleware that extracts the country from the X-Country header
func CountryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the X-Country header
		country := c.GetHeader("X-Country")

		// If the header is not set, default to "US"
		if country == "" {
			country = countries.DominicanRepublic.String()
		}

		// Set the country in the context
		c.Set("country", country)

		c.Next()
	}
}
