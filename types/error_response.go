package types

// ErrorResponse represents an error response returned by the API.
type ErrorResponse struct {
	Status  int      `json:"status"`  // The HTTP status code of the error response.
	Message string   `json:"message"` // The error message.
	Errors  []string `json:"errors"`  // A list of specific error messages.
}
