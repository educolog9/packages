package customerrors

type UnauthorizedError struct {
	*BaseError
}

// NewUnauthorizedError creates a new UnauthorizedError with the given message and inner error.
// It returns a pointer to the created UnauthorizedError.
func NewUnauthorizedError(msg string, inner error) *UnauthorizedError {
	if msg == "" {
		msg = "Unauthorized"
	}
	return &UnauthorizedError{
		BaseError: NewBaseError(msg, 401, inner),
	}
}
