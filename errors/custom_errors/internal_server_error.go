package customerrors

type InternalServerError struct {
	*BaseError
}

// NewInternalServerError creates a new InternalServerError with the given message and inner error.
// It returns a pointer to the created InternalServerError.
func NewInternalServerError(msg string, inner error) *InternalServerError {
	if msg == "" {
		msg = "Internal Server Error"
	}
	return &InternalServerError{
		BaseError: NewBaseError(msg, 500, inner),
	}
}
