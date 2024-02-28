package messages

const (
	// Success represents a successful operation.
	Success string = "success"

	// Error represents a failed operation.
	Error string = "error"

	// NotFound represents a resource not found error.

	NotFound string = "not found"

	// Unauthorized represents an unauthorized access error.
	Unauthorized string = "unauthorized"

	// Forbidden represents a forbidden access error.
	Forbidden string = "forbidden"

	// BadRequest represents a bad request error.
	ValidationFailed string = "validation failed"
	BadRequest       string = "bad request"

	// InternalServerError represents an internal server error.
	InternalServerError string = "internal server error"

	// ServiceUnavailable represents a service unavailable error.
	ServiceUnavailable string = "service unavailable"

	// GatewayTimeout represents a gateway timeout error.
	GatewayTimeout string = "gateway timeout"

	// DuplicateResource represents a duplicate resource error.
	DuplicateResource string = "duplicate resource"

	// InvalidCredentials represents an invalid credentials error.
	InvalidCredentials string = "invalid credentials"
)
