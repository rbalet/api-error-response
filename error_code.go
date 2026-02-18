package apierror

// ErrorCode represents specific error codes grouped by domain
type ErrorCode string

const (
	// AUTH
	AuthUnauthorized         ErrorCode = "AUTH_UNAUTHORIZED"
	AuthForbidden            ErrorCode = "AUTH_FORBIDDEN"
	AuthInvalidCredentials   ErrorCode = "AUTH_INVALID_CREDENTIALS"
	AuthTokenExpired         ErrorCode = "AUTH_TOKEN_EXPIRED"
	AuthTokenInvalid         ErrorCode = "AUTH_TOKEN_INVALID"
	AuthRefreshTokenInvalid  ErrorCode = "AUTH_REFRESH_TOKEN_INVALID"
	AuthAccountDisabled      ErrorCode = "AUTH_ACCOUNT_DISABLED"
	AuthAccountLocked        ErrorCode = "AUTH_ACCOUNT_LOCKED"
	AuthOAuthProviderError   ErrorCode = "AUTH_OAUTH_PROVIDER_ERROR"
	AuthSessionExpired       ErrorCode = "AUTH_SESSION_EXPIRED"

	// VALIDATION
	ValidationFailed            ErrorCode = "VALIDATION_FAILED"
	ValidationInvalidPayload    ErrorCode = "VALIDATION_INVALID_PAYLOAD"
	ValidationMissingField      ErrorCode = "VALIDATION_MISSING_FIELD"
	ValidationInvalidType       ErrorCode = "VALIDATION_INVALID_TYPE"
	ValidationFieldRequired     ErrorCode = "VALIDATION_FIELD_REQUIRED"
	ValidationFieldInvalidFormat ErrorCode = "VALIDATION_FIELD_INVALID_FORMAT"
	ValidationFieldTooShort     ErrorCode = "VALIDATION_FIELD_TOO_SHORT"
	ValidationFieldTooLong      ErrorCode = "VALIDATION_FIELD_TOO_LONG"
	ValidationFieldTooSmall     ErrorCode = "VALIDATION_FIELD_TOO_SMALL"
	ValidationFieldTooLarge     ErrorCode = "VALIDATION_FIELD_TOO_LARGE"
	ValidationFieldNotAllowed   ErrorCode = "VALIDATION_FIELD_NOT_ALLOWED"
	ValidationFieldNotUnique    ErrorCode = "VALIDATION_FIELD_NOT_UNIQUE"
	ValidationFieldOutOfRange   ErrorCode = "VALIDATION_FIELD_OUT_OF_RANGE"
	ValidationFieldEnumInvalid  ErrorCode = "VALIDATION_FIELD_ENUM_INVALID"

	// DOMAIN
	ResourceNotFound      ErrorCode = "RESOURCE_NOT_FOUND"
	ResourceAlreadyExists ErrorCode = "RESOURCE_ALREADY_EXISTS"
	ResourceConflict      ErrorCode = "RESOURCE_CONFLICT"
	ResourceLocked        ErrorCode = "RESOURCE_LOCKED"
	ResourceDeleted       ErrorCode = "RESOURCE_DELETED"
	UserNotFound          ErrorCode = "USER_NOT_FOUND"
	UserAlreadyExists     ErrorCode = "USER_ALREADY_EXISTS"
	UserEmailAlreadyUsed  ErrorCode = "USER_EMAIL_ALREADY_USED"
	UserUsernameAlreadyUsed ErrorCode = "USER_USERNAME_ALREADY_USED"
	UserInvalidState      ErrorCode = "USER_INVALID_STATE"
	OrderNotFound         ErrorCode = "ORDER_NOT_FOUND"
	OrderAlreadyPaid      ErrorCode = "ORDER_ALREADY_PAID"
	OrderOutOfStock       ErrorCode = "ORDER_OUT_OF_STOCK"
	PaymentFailed         ErrorCode = "PAYMENT_FAILED"
	PaymentDeclined       ErrorCode = "PAYMENT_DECLINED"
	PaymentProviderError  ErrorCode = "PAYMENT_PROVIDER_ERROR"

	// CONFLICT
	ConflictVersionMismatch        ErrorCode = "CONFLICT_VERSION_MISMATCH"
	ConflictDuplicateEntry         ErrorCode = "CONFLICT_DUPLICATE_ENTRY"
	ConflictInvalidStateTransition ErrorCode = "CONFLICT_INVALID_STATE_TRANSITION"

	// RATE LIMIT
	RateLimitExceeded ErrorCode = "RATE_LIMIT_EXCEEDED"
	QuotaExceeded     ErrorCode = "QUOTA_EXCEEDED"

	// SYSTEM
	SystemInternalError      ErrorCode = "SYSTEM_INTERNAL_ERROR"
	SystemDependencyFailure  ErrorCode = "SYSTEM_DEPENDENCY_FAILURE"
	SystemTimeout            ErrorCode = "SYSTEM_TIMEOUT"
	SystemDatabaseError      ErrorCode = "SYSTEM_DATABASE_ERROR"
	SystemCacheError         ErrorCode = "SYSTEM_CACHE_ERROR"
	SystemIOError            ErrorCode = "SYSTEM_IO_ERROR"
	SystemConfigurationError ErrorCode = "SYSTEM_CONFIGURATION_ERROR"

	// API
	APINotFound                ErrorCode = "API_NOT_FOUND"
	APIMethodNotAllowed        ErrorCode = "API_METHOD_NOT_ALLOWED"
	APIUnsupportedMediaType    ErrorCode = "API_UNSUPPORTED_MEDIA_TYPE"
	APIBadRequest              ErrorCode = "API_BAD_REQUEST"
	APIVersionNotSupported     ErrorCode = "API_VERSION_NOT_SUPPORTED"
)
