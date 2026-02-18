package apierror

// ApiResponse represents a standardized API response that is either a success with data or an error
// In Go, we use a generic struct with pointers to ensure only one field is populated
type ApiResponse[T any] struct {
	Data  *T        `json:"data,omitempty"`
	Error ApiError  `json:"error,omitempty"`
}

// NewSuccessResponse creates a success response with data
func NewSuccessResponse[T any](data T) *ApiResponse[T] {
	return &ApiResponse[T]{
		Data: &data,
	}
}

// NewErrorResponse creates an error response
func NewErrorResponse[T any](err ApiError) *ApiResponse[T] {
	return &ApiResponse[T]{
		Error: err,
	}
}

// IsSuccess returns true if the response contains data
func (r *ApiResponse[T]) IsSuccess() bool {
	return r.Data != nil
}

// IsError returns true if the response contains an error
func (r *ApiResponse[T]) IsError() bool {
	return r.Error != nil
}
