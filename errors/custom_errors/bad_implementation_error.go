package customerrors

type BadImplementationError struct {
	*BaseError
}

// NewBadImplementationError creates a new BadImplementationError with the given message and inner error.
// It returns a pointer to the created BadImplementationError.
func NewBadImplementationError(msg string, inner error) *BadImplementationError {
	if msg == "" {
		msg = "Bad implementation"
	}

	return &BadImplementationError{
		BaseError: NewBaseError(msg, 400, inner),
	}
}
