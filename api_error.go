package apierror

// ErrorType represents high-level error categories
type ErrorType string

const (
	ErrorTypeAuth      ErrorType = "AUTH"
	ErrorTypeValidation ErrorType = "VALIDATION"
	ErrorTypeDomain    ErrorType = "DOMAIN"
	ErrorTypeConflict  ErrorType = "CONFLICT"
	ErrorTypeNotFound  ErrorType = "NOT_FOUND"
	ErrorTypeRateLimit ErrorType = "RATE_LIMIT"
	ErrorTypeSystem    ErrorType = "SYSTEM"
	ErrorTypeAPI       ErrorType = "API"
)

// ValidationIssue represents a single validation problem, typically associated with a specific field or path
type ValidationIssue struct {
	Code    *ErrorCode             `json:"code,omitempty"`
	Path    []interface{}          `json:"path,omitempty"`
	Message *string                `json:"message,omitempty"`
	Meta    map[string]interface{} `json:"meta,omitempty"`
}

// ApiErrorBase is the base structure shared by all errors
type ApiErrorBase struct {
	Type      *ErrorType `json:"type,omitempty"`
	Code      *ErrorCode `json:"code,omitempty"`
	Message   *string    `json:"message,omitempty"`
	TraceID   *string    `json:"traceId,omitempty"`
	Timestamp *string    `json:"timestamp,omitempty"`
}

// ValidationError is used when request input fails validation
type ValidationError struct {
	Type      ErrorType          `json:"type"`
	Code      *ErrorCode         `json:"code,omitempty"`
	Message   *string            `json:"message,omitempty"`
	TraceID   *string            `json:"traceId,omitempty"`
	Timestamp *string            `json:"timestamp,omitempty"`
	Issues    []ValidationIssue  `json:"issues,omitempty"`
}

// NonValidationError is used for all other error types
type NonValidationError struct {
	Type      ErrorType  `json:"type"`
	Code      *ErrorCode `json:"code,omitempty"`
	Message   *string    `json:"message,omitempty"`
	TraceID   *string    `json:"traceId,omitempty"`
	Timestamp *string    `json:"timestamp,omitempty"`
}

// ApiError represents any API error (validation or non-validation)
// In Go, this is implemented as an interface that both ValidationError and NonValidationError satisfy
type ApiError interface {
	GetType() ErrorType
	GetCode() *ErrorCode
	GetMessage() *string
	GetTraceID() *string
	GetTimestamp() *string
	IsValidationError() bool
}

// GetType returns the error type
func (e *ValidationError) GetType() ErrorType {
	return e.Type
}

// GetCode returns the error code
func (e *ValidationError) GetCode() *ErrorCode {
	return e.Code
}

// GetMessage returns the error message
func (e *ValidationError) GetMessage() *string {
	return e.Message
}

// GetTraceID returns the trace ID
func (e *ValidationError) GetTraceID() *string {
	return e.TraceID
}

// GetTimestamp returns the timestamp
func (e *ValidationError) GetTimestamp() *string {
	return e.Timestamp
}

// IsValidationError returns true for ValidationError
func (e *ValidationError) IsValidationError() bool {
	return true
}

// GetType returns the error type
func (e *NonValidationError) GetType() ErrorType {
	return e.Type
}

// GetCode returns the error code
func (e *NonValidationError) GetCode() *ErrorCode {
	return e.Code
}

// GetMessage returns the error message
func (e *NonValidationError) GetMessage() *string {
	return e.Message
}

// GetTraceID returns the trace ID
func (e *NonValidationError) GetTraceID() *string {
	return e.TraceID
}

// GetTimestamp returns the timestamp
func (e *NonValidationError) GetTimestamp() *string {
	return e.Timestamp
}

// IsValidationError returns false for NonValidationError
func (e *NonValidationError) IsValidationError() bool {
	return false
}
