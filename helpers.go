package apierror

import "time"

// Helper functions for creating common error responses

// NewValidationError creates a new validation error
func NewValidationError(message string, issues []ValidationIssue, traceID string) *ValidationError {
	code := ValidationFailed
	timestamp := time.Now().UTC().Format(time.RFC3339)
	return &ValidationError{
		Type:      ErrorTypeValidation,
		Code:      &code,
		Message:   &message,
		TraceID:   &traceID,
		Timestamp: &timestamp,
		Issues:    issues,
	}
}

// NewAuthError creates a new authentication error
func NewAuthError(code ErrorCode, message string, traceID string) *NonValidationError {
	timestamp := time.Now().UTC().Format(time.RFC3339)
	return &NonValidationError{
		Type:      ErrorTypeAuth,
		Code:      &code,
		Message:   &message,
		TraceID:   &traceID,
		Timestamp: &timestamp,
	}
}

// NewDomainError creates a new domain error
func NewDomainError(code ErrorCode, message string, traceID string) *NonValidationError {
	timestamp := time.Now().UTC().Format(time.RFC3339)
	return &NonValidationError{
		Type:      ErrorTypeDomain,
		Code:      &code,
		Message:   &message,
		TraceID:   &traceID,
		Timestamp: &timestamp,
	}
}

// NewSystemError creates a new system error
func NewSystemError(code ErrorCode, message string, traceID string) *NonValidationError {
	timestamp := time.Now().UTC().Format(time.RFC3339)
	return &NonValidationError{
		Type:      ErrorTypeSystem,
		Code:      &code,
		Message:   &message,
		TraceID:   &traceID,
		Timestamp: &timestamp,
	}
}

// NewNotFoundError creates a new not found error
func NewNotFoundError(message string, traceID string) *NonValidationError {
	code := ResourceNotFound
	timestamp := time.Now().UTC().Format(time.RFC3339)
	return &NonValidationError{
		Type:      ErrorTypeNotFound,
		Code:      &code,
		Message:   &message,
		TraceID:   &traceID,
		Timestamp: &timestamp,
	}
}

// NewRateLimitError creates a new rate limit error
func NewRateLimitError(message string, traceID string) *NonValidationError {
	code := RateLimitExceeded
	timestamp := time.Now().UTC().Format(time.RFC3339)
	return &NonValidationError{
		Type:      ErrorTypeRateLimit,
		Code:      &code,
		Message:   &message,
		TraceID:   &traceID,
		Timestamp: &timestamp,
	}
}

// NewConflictError creates a new conflict error
func NewConflictError(code ErrorCode, message string, traceID string) *NonValidationError {
	timestamp := time.Now().UTC().Format(time.RFC3339)
	return &NonValidationError{
		Type:      ErrorTypeConflict,
		Code:      &code,
		Message:   &message,
		TraceID:   &traceID,
		Timestamp: &timestamp,
	}
}

// NewAPIError creates a new API-level error
func NewAPIError(code ErrorCode, message string, traceID string) *NonValidationError {
	timestamp := time.Now().UTC().Format(time.RFC3339)
	return &NonValidationError{
		Type:      ErrorTypeAPI,
		Code:      &code,
		Message:   &message,
		TraceID:   &traceID,
		Timestamp: &timestamp,
	}
}
