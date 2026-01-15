package middlewares

import (
	"context"
	"time"

	"github.com/educolog9/packages/types"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	// tenantDB is the database where white_labels collection is stored
	tenantDB = "user"
	// whiteLabelsColl is the collection name for white labels
	whiteLabelsColl = "white_labels"
	// xTenantDomainHeader is the header name for tenant domain
	xTenantDomainHeader = "X-Tenant-Domain"
)

// dbWhiteLabel represents the white label document in MongoDB
type dbWhiteLabel struct {
	ID             primitive.ObjectID `bson:"_id"`
	OrganizationID primitive.ObjectID `bson:"organizationId"`
	Domain         string             `bson:"domain"`
	Status         string             `bson:"status"`
	BackofficeUrl  string             `bson:"backofficeUrl"`
	WebUrl         string             `bson:"webUrl"`
}

// TenantResolver is a middleware that resolves the tenant from X-Tenant-Domain header.
// It queries the white_labels collection to get the organization information
// and injects a TenantContext into the request context.
//
// If no X-Tenant-Domain header is present, the request continues without tenant context.
// This allows the middleware to work with both public and authenticated endpoints.
func TenantResolver(client *mongo.Client) gin.HandlerFunc {
	collection := client.Database(tenantDB).Collection(whiteLabelsColl)

	return func(c *gin.Context) {
		span, _ := opentracing.StartSpanFromContext(c.Request.Context(), "TenantResolver")
		defer span.Finish()

		// Get domain from X-Tenant-Domain header
		domain := c.GetHeader(xTenantDomainHeader)
		if domain == "" {
			// No tenant header - continue without tenant context
			// This is valid for endpoints that don't require tenant resolution
			c.Next()
			return
		}

		// Query white_labels collection
		ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
		defer cancel()

		var whiteLabel dbWhiteLabel
		err := collection.FindOne(ctx, bson.M{"domain": domain}).Decode(&whiteLabel)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				// Domain not found - continue without tenant context
				// The handler can decide how to handle this case
				c.Set("tenantContext", &types.TenantContext{
					Domain:        domain,
					HasWhiteLabel: false,
				})
				c.Next()
				return
			}
			// Database error - log and continue without context
			span.SetTag("error", true)
			span.SetTag("error.message", err.Error())
			c.Next()
			return
		}

		// Create tenant context with resolved data
		tenantCtx := &types.TenantContext{
			Domain:         domain,
			OrganizationID: whiteLabel.OrganizationID.Hex(),
			WhiteLabelID:   whiteLabel.ID.Hex(),
			HasWhiteLabel:  true,
			Status:         whiteLabel.Status,
			BackofficeUrl:  whiteLabel.BackofficeUrl,
			WebUrl:         whiteLabel.WebUrl,
		}

		// Set tenant context in Gin context
		c.Set("tenantContext", tenantCtx)

		c.Next()
	}
}

// GetTenantContext retrieves the TenantContext from the Gin context.
// Returns nil if no tenant context is set.
func GetTenantContext(c *gin.Context) *types.TenantContext {
	if ctx, exists := c.Get("tenantContext"); exists {
		if tenantCtx, ok := ctx.(*types.TenantContext); ok {
			return tenantCtx
		}
	}
	return nil
}

// RequireTenant is a middleware that ensures a valid tenant context exists.
// Use this for endpoints that require a resolved tenant.
// Returns 400 Bad Request if X-Tenant-Domain header is missing.
// Returns 404 Not Found if the domain is not found in white_labels.
func RequireTenant(client *mongo.Client) gin.HandlerFunc {
	collection := client.Database(tenantDB).Collection(whiteLabelsColl)

	return func(c *gin.Context) {
		span, _ := opentracing.StartSpanFromContext(c.Request.Context(), "RequireTenant")
		defer span.Finish()

		// Get domain from X-Tenant-Domain header
		domain := c.GetHeader(xTenantDomainHeader)
		if domain == "" {
			c.JSON(400, types.ErrorResponse{
				Status:  400,
				Message: "Bad Request",
				Errors:  []string{"X-Tenant-Domain header is required"},
			})
			c.Abort()
			return
		}

		// Query white_labels collection
		ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
		defer cancel()

		var whiteLabel dbWhiteLabel
		err := collection.FindOne(ctx, bson.M{"domain": domain}).Decode(&whiteLabel)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(404, types.ErrorResponse{
					Status:  404,
					Message: "Not Found",
					Errors:  []string{"Domain not found"},
				})
				c.Abort()
				return
			}
			c.JSON(500, types.ErrorResponse{
				Status:  500,
				Message: "Internal Server Error",
				Errors:  []string{"Failed to resolve tenant"},
			})
			c.Abort()
			return
		}

		// Create tenant context with resolved data
		tenantCtx := &types.TenantContext{
			Domain:         domain,
			OrganizationID: whiteLabel.OrganizationID.Hex(),
			WhiteLabelID:   whiteLabel.ID.Hex(),
			HasWhiteLabel:  true,
			Status:         whiteLabel.Status,
			BackofficeUrl:  whiteLabel.BackofficeUrl,
			WebUrl:         whiteLabel.WebUrl,
		}

		// Set tenant context in Gin context
		c.Set("tenantContext", tenantCtx)

		c.Next()
	}
}
