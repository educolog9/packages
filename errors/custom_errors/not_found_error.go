package customerrors

type NotFoundError struct {
	*BaseError
}

// NewNotFoundError creates a new NotFoundError with the given message and inner error.
// It returns a pointer to the created NotFoundError.
func NewNotFoundError(msg string, inner error) *NotFoundError {
	if msg == "" {
		msg = "Entity not found"
	}

	return &NotFoundError{
		BaseError: NewBaseError(msg, 404, inner),
	}
}
