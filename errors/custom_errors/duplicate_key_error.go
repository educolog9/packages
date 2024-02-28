package customerrors

type DuplicateKeyError struct {
	*BaseError
}

// NewDuplicateKeyError creates a new DuplicateKeyError with the given message and inner error.
// It returns a pointer to the created DuplicateKeyError.
func NewDuplicateKeyError(msg string, inner error) *DuplicateKeyError {
	if msg == "" {
		msg = "Duplicate key error"
	}

	return &DuplicateKeyError{
		BaseError: NewBaseError(msg, 409, inner),
	}
}
