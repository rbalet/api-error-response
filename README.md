# API Error Response

A standardized contract for API error and success responses, available in **TypeScript** and **Go**. Provides type-safe structures that make it easy to handle both success and error cases with full type safety.

## Table of Contents

- [Overview](#overview)
- [The Problem](#the-problem)
- [The Solution](#the-solution)
- [Language Support](#language-support)
  - [TypeScript](#typescript)
  - [Go](#go)
- [Type Definitions](#type-definitions)
  - [ApiResponse](#apiresponse)
  - [ApiError](#apierror)
  - [ValidationError vs NonValidationError](#validationerror-vs-nonvalidationerror)
  - [ValidationIssue](#validationissue)
  - [ErrorType](#errortype)
  - [ErrorCode](#errorcode)
- [Examples](#examples)
  - [TypeScript Examples](#typescript-examples)
  - [Go Examples](#go-examples)
- [Conventions & Guidelines](#conventions--guidelines)
  - [TraceId and Timestamp](#traceid-and-timestamp)
  - [HTTP Status Code Mapping](#http-status-code-mapping)
  - [Backward Compatibility](#backward-compatibility)
- [Installation & Usage](#installation--usage)
  - [TypeScript Installation](#typescript-installation)
  - [Go Installation](#go-installation)
- [License](#license)

## Overview

This library defines a standard shape for API responses that cleanly separates success and error cases. Every response is either a success with `data` or an error with `error`, but never both.

Available in both **TypeScript** (with discriminated unions) and **Go** (with generics and interfaces).

## The Problem

APIs often return inconsistent error shapes:
- Some nest errors in a `message` field, others use `error`, `errors`, or `errorMessage`
- Validation errors are structured differently than other errors
- Clients can't easily distinguish error types without runtime checks
- Type safety is lost when handling different error scenarios

## The Solution

**ApiResponse** provides a single, predictable contract.

### TypeScript

```typescript
type ApiResponse<T> = 
  | { data: T; error?: never }       // Success case
  | { data?: never; error: ApiError } // Error case
```

### Go

```go
type ApiResponse[T any] struct {
    Data  *T       `json:"data,omitempty"`
    Error ApiError `json:"error,omitempty"`
}
```

This ensures:
- Type-safe handling with language-specific type systems
- Clear separation between success and error states
- Consistent error structure across all endpoints
- Detailed validation errors with field-level issues

## Language Support

### TypeScript

The TypeScript implementation uses discriminated unions for compile-time type safety. See the [TypeScript Examples](#typescript-examples) section for usage.

### Go

The Go implementation uses generics (Go 1.18+) and interfaces. The `ApiError` interface is implemented by both `ValidationError` and `NonValidationError`. See the [Go Examples](#go-examples) section for usage.

## Type Definitions

### ApiResponse

The root type for all API responses.

```typescript
export type ApiResponse<T, Codes extends string = ErrorCodeBase> =
  | { data: T; error?: never }
  | { data?: never; error: ApiError<Codes> };
```

**Generic Parameters:**
- `T` - The type of the success response data
- `Codes` - Optional custom error code type (defaults to `ErrorCodeBase`)

**Properties:**
- `data` - Present only on success, contains the response payload of type `T`
- `error` - Present only on error, contains detailed error information

### ApiError

A union type representing all possible error shapes.

```typescript
export type ApiError<Codes extends string = ErrorCodeBase> = 
  | ValidationError<Codes> 
  | NonValidationError<Codes>;
```

All errors share a common base structure:

```typescript
interface ApiErrorBase<Codes extends string = ErrorCode> {
  type?: ErrorType;
  code?: Codes;
  message?: string;
  traceId?: string;
  timestamp?: string;
}
```

**Properties:**
- `type` - Broad error category (e.g., `"AUTH"`, `"VALIDATION"`, `"SYSTEM"`)
- `code` - Specific error code from the `ErrorCode` enum
- `message` - Human-readable error description
- `traceId` - Unique identifier for request tracing and debugging
- `timestamp` - ISO 8601 timestamp when the error occurred

### ValidationError vs NonValidationError

#### ValidationError

Used when request input fails validation. Includes an `issues` array with field-level details.

```typescript
interface ValidationError<Codes extends string = ErrorCodeBase> extends ApiErrorBase<Codes> {
  type: "VALIDATION";
  issues?: ValidationIssue<Codes>[];
}
```

**When to use:**
- Request body validation fails
- Query parameters are invalid
- Path parameters don't meet constraints
- Form data has errors in multiple fields

#### NonValidationError

Used for all other error types: authentication, authorization, domain logic, system failures, etc.

```typescript
interface NonValidationError<Codes extends string = ErrorCodeBase> extends ApiErrorBase<Codes> {
  type: Exclude<ErrorType, "VALIDATION">;
  issues?: never;
}
```

**When to use:**
- Authentication/authorization failures
- Resource not found
- Business rule violations
- Rate limiting
- System/infrastructure errors

### ValidationIssue

Represents a single validation problem, typically associated with a specific field or path.

```typescript
interface ValidationIssue<Codes extends string = ErrorCode> {
  code?: Codes;
  path?: (string | number)[];
  message?: string;
  meta?: Record<string, unknown>;
}
```

**Properties:**
- `code` - Specific validation error code (e.g., `VALIDATION_FIELD_TOO_SHORT`)
- `path` - JSON path to the invalid field (e.g., `["user", "email"]`)
- `message` - Human-readable description of the validation issue
- `meta` - Additional context (e.g., `{ min: 8, max: 100, actual: 5 }`)

### ErrorType

High-level error categories:

```typescript
type ErrorType =
  | "AUTH"        // Authentication/authorization
  | "VALIDATION"  // Input validation
  | "DOMAIN"      // Business logic/domain rules
  | "CONFLICT"    // State conflicts
  | "NOT_FOUND"   // Resource not found
  | "RATE_LIMIT"  // Rate limiting/quotas
  | "SYSTEM"      // Infrastructure/system errors
  | "API";        // API-level errors (e.g., unsupported media type)
```

### ErrorCode

An enum with specific, actionable error codes grouped by domain. Examples:

```typescript
enum ErrorCode {
  // AUTH
  AUTH_UNAUTHORIZED = "AUTH_UNAUTHORIZED",
  AUTH_FORBIDDEN = "AUTH_FORBIDDEN",
  AUTH_TOKEN_EXPIRED = "AUTH_TOKEN_EXPIRED",
  
  // VALIDATION
  VALIDATION_FAILED = "VALIDATION_FAILED",
  VALIDATION_FIELD_REQUIRED = "VALIDATION_FIELD_REQUIRED",
  VALIDATION_FIELD_TOO_SHORT = "VALIDATION_FIELD_TOO_SHORT",
  
  // DOMAIN
  RESOURCE_NOT_FOUND = "RESOURCE_NOT_FOUND",
  USER_EMAIL_ALREADY_USED = "USER_EMAIL_ALREADY_USED",
  ORDER_OUT_OF_STOCK = "ORDER_OUT_OF_STOCK",
  
  // SYSTEM
  SYSTEM_INTERNAL_ERROR = "SYSTEM_INTERNAL_ERROR",
  SYSTEM_DATABASE_ERROR = "SYSTEM_DATABASE_ERROR",
  
  // ... and many more
}
```

See `error-code.enum.ts` (TypeScript) or `error_code.go` (Go) for the complete list.

## Examples

### TypeScript Examples

#### Success Response

```typescript
const response: ApiResponse<{ userId: string; email: string }> = {
  data: {
    userId: "usr_1234567890",
    email: "user@example.com"
  }
};
```

#### Non-Validation Error

```typescript
const response: ApiResponse<never> = {
  error: {
    type: "AUTH",
    code: "AUTH_UNAUTHORIZED",
    message: "Invalid or expired authentication token",
    traceId: "trace-abc123",
    timestamp: "2026-02-16T12:30:00Z"
  }
};
```

#### Validation Error with Multiple Issues

```typescript
const response: ApiResponse<never> = {
  error: {
    type: "VALIDATION",
    code: "VALIDATION_FAILED",
    message: "Request validation failed",
    traceId: "trace-def456",
    timestamp: "2026-02-16T12:35:00Z",
    issues: [
      {
        code: "VALIDATION_FIELD_REQUIRED",
        path: ["user", "email"],
        message: "Email is required"
      },
      {
        code: "VALIDATION_FIELD_TOO_SHORT",
        path: ["user", "password"],
        message: "Password must be at least 8 characters",
        meta: { min: 8, actual: 5 }
      },
      {
        code: "VALIDATION_FIELD_INVALID_FORMAT",
        path: ["user", "phoneNumber"],
        message: "Phone number format is invalid"
      }
    ]
  }
};
```

#### Server-Side Helper

Here's a minimal TypeScript helper for building responses:

```typescript
import { ApiResponse, ApiError, ErrorCode, ValidationIssue } from './api-error-response';

// Success helper
function successResponse<T>(data: T): ApiResponse<T> {
  return { data };
}

// Error helpers
function errorResponse(error: ApiError): ApiResponse<never> {
  return { error };
}

function authError(message: string, traceId?: string): ApiResponse<never> {
  return {
    error: {
      type: "AUTH",
      code: ErrorCode.AUTH_UNAUTHORIZED,
      message,
      traceId,
      timestamp: new Date().toISOString()
    }
  };
}

function validationError(
  issues: ValidationIssue[],
  message = "Validation failed",
  traceId?: string
): ApiResponse<never> {
  return {
    error: {
      type: "VALIDATION",
      code: ErrorCode.VALIDATION_FAILED,
      message,
      issues,
      traceId,
      timestamp: new Date().toISOString()
    }
  };
}

function notFoundError(message: string, traceId?: string): ApiResponse<never> {
  return {
    error: {
      type: "DOMAIN",
      code: ErrorCode.RESOURCE_NOT_FOUND,
      message,
      traceId,
      timestamp: new Date().toISOString()
    }
  };
}

// Usage
const userResponse = successResponse({ userId: "123", name: "Alice" });

const authFailure = authError("Token expired", "trace-001");

const validationFailure = validationError([
  { code: ErrorCode.VALIDATION_FIELD_REQUIRED, path: ["email"], message: "Email required" }
]);
```

#### Client-Side Type Narrowing

Use type guards to safely handle success vs. error cases:

```typescript
import { ApiResponse } from './api-error-response';

async function fetchUser(userId: string): Promise<ApiResponse<{ name: string; email: string }>> {
  const response = await fetch(`/api/users/${userId}`);
  return response.json();
}

// Type guard
function isSuccess<T>(response: ApiResponse<T>): response is { data: T; error?: never } {
  return 'data' in response && response.data !== undefined;
}

function isError<T>(response: ApiResponse<T>): response is { data?: never; error: ApiError } {
  return 'error' in response && response.error !== undefined;
}

// Usage with type narrowing
const response = await fetchUser("usr_123");

if (isSuccess(response)) {
  // TypeScript knows response.data exists
  console.log("User name:", response.data.name);
  console.log("User email:", response.data.email);
} else if (isError(response)) {
  // TypeScript knows response.error exists
  console.error("Error type:", response.error.type);
  console.error("Error code:", response.error.code);
  console.error("Message:", response.error.message);
  
  if (response.error.type === "VALIDATION" && response.error.issues) {
    response.error.issues.forEach(issue => {
      console.error(`Field ${issue.path?.join('.')}: ${issue.message}`);
    });
  }
}

// Alternative: using 'data' presence directly
if (response.data) {
  console.log(response.data.name);
} else if (response.error) {
  console.error(response.error.message);
}
```

### Go Examples

#### Success Response

```go
package main

import (
    "encoding/json"
    "fmt"
    apierror "github.com/rbalet/api-error-response"
)

type User struct {
    UserID string `json:"userId"`
    Email  string `json:"email"`
}

func main() {
    user := User{
        UserID: "usr_1234567890",
        Email:  "user@example.com",
    }
    response := apierror.NewSuccessResponse(user)
    
    jsonData, _ := json.MarshalIndent(response, "", "  ")
    fmt.Println(string(jsonData))
}
```

#### Non-Validation Error

```go
package main

import (
    "encoding/json"
    "fmt"
    apierror "github.com/rbalet/api-error-response"
)

func main() {
    authError := apierror.NewAuthError(
        apierror.AuthUnauthorized,
        "Invalid or expired authentication token",
        "trace-abc123",
    )
    
    response := apierror.NewErrorResponse[interface{}](authError)
    
    jsonData, _ := json.MarshalIndent(response, "", "  ")
    fmt.Println(string(jsonData))
}
```

#### Validation Error with Multiple Issues

```go
package main

import (
    "encoding/json"
    "fmt"
    apierror "github.com/rbalet/api-error-response"
)

func main() {
    emailMsg := "Email is required"
    passwordMsg := "Password must be at least 8 characters"
    phoneMsg := "Phone number format is invalid"
    
    emailCode := apierror.ValidationFieldRequired
    passwordCode := apierror.ValidationFieldTooShort
    phoneCode := apierror.ValidationFieldInvalidFormat
    
    issues := []apierror.ValidationIssue{
        {
            Code:    &emailCode,
            Path:    []interface{}{"user", "email"},
            Message: &emailMsg,
        },
        {
            Code:    &passwordCode,
            Path:    []interface{}{"user", "password"},
            Message: &passwordMsg,
            Meta: map[string]interface{}{
                "min":    8,
                "actual": 5,
            },
        },
        {
            Code:    &phoneCode,
            Path:    []interface{}{"user", "phoneNumber"},
            Message: &phoneMsg,
        },
    }
    
    validationError := apierror.NewValidationError(
        "Request validation failed",
        issues,
        "trace-def456",
    )
    response := apierror.NewErrorResponse[interface{}](validationError)
    
    jsonData, _ := json.MarshalIndent(response, "", "  ")
    fmt.Println(string(jsonData))
}
```

#### Type Checking in Go

```go
package main

import (
    "fmt"
    apierror "github.com/rbalet/api-error-response"
)

type User struct {
    UserID string `json:"userId"`
    Name   string `json:"name"`
}

func main() {
    user := User{UserID: "123", Name: "Alice"}
    response := apierror.NewSuccessResponse(user)
    
    if response.IsSuccess() {
        fmt.Printf("Success! User ID: %s\n", response.Data.UserID)
    } else if response.IsError() {
        fmt.Printf("Error: %s\n", *response.Error.GetMessage())
    }
    
    // Check if error is a validation error
    if response.IsError() && response.Error.IsValidationError() {
        if validationErr, ok := response.Error.(*apierror.ValidationError); ok {
            for _, issue := range validationErr.Issues {
                fmt.Printf("Field error: %s\n", *issue.Message)
            }
        }
    }
}
```

#### Helper Functions

The Go package includes helper functions for creating common error types:

```go
package main

import apierror "github.com/rbalet/api-error-response"

func main() {
    // Authentication error
    authErr := apierror.NewAuthError(
        apierror.AuthUnauthorized,
        "Token expired",
        "trace-001",
    )
    
    // Domain error
    domainErr := apierror.NewDomainError(
        apierror.UserNotFound,
        "User not found",
        "trace-002",
    )
    
    // Not found error
    notFoundErr := apierror.NewNotFoundError(
        "Resource not found",
        "trace-003",
    )
    
    // System error
    systemErr := apierror.NewSystemError(
        apierror.SystemDatabaseError,
        "Database connection failed",
        "trace-004",
    )
    
    // Rate limit error
    rateLimitErr := apierror.NewRateLimitError(
        "Too many requests",
        "trace-005",
    )
    
    // Use these errors in responses
    response := apierror.NewErrorResponse[interface{}](authErr)
}
```

## Conventions & Guidelines

### TraceId and Timestamp

**TraceId:**
- Generate a unique trace ID per request (e.g., UUID or request ID from your framework)
- Include it in all error responses for debugging and log correlation
- Pass it through to downstream services to maintain request context
- Log it server-side to link errors back to specific requests

**Timestamp:**
- Use ISO 8601 format (e.g., `2026-02-16T12:35:00Z`)
- Set when the error occurs, not when the request started
- Helpful for debugging time-sensitive issues and understanding error timing

### HTTP Status Code Mapping

Match HTTP status codes to error types for RESTful conventions:

| Error Type | HTTP Status | Examples |
|------------|-------------|----------|
| `AUTH` | 401 Unauthorized / 403 Forbidden | `AUTH_UNAUTHORIZED`, `AUTH_FORBIDDEN` |
| `VALIDATION` | 400 Bad Request / 422 Unprocessable Entity | All `VALIDATION_*` codes |
| `NOT_FOUND` / `DOMAIN` (resource) | 404 Not Found | `RESOURCE_NOT_FOUND`, `USER_NOT_FOUND` |
| `CONFLICT` | 409 Conflict | `CONFLICT_DUPLICATE_ENTRY` |
| `RATE_LIMIT` | 429 Too Many Requests | `RATE_LIMIT_EXCEEDED` |
| `SYSTEM` | 500 Internal Server Error / 503 Service Unavailable | `SYSTEM_INTERNAL_ERROR`, `SYSTEM_DATABASE_ERROR` |
| `API` | 405 Method Not Allowed / 415 Unsupported Media | `API_METHOD_NOT_ALLOWED` |
| `DOMAIN` (business logic) | 400 Bad Request / 422 Unprocessable Entity | `ORDER_OUT_OF_STOCK`, `PAYMENT_FAILED` |

**Note:** This is guidance, not a strict requirement. Adjust based on your API conventions.

### Backward Compatibility

- **Add, don't remove**: Add new error codes instead of changing existing ones
- **Optional fields**: Keep all fields optional to allow gradual adoption
- **New issues fields**: When adding new properties to `ValidationIssue`, make them optional
- **Version error codes**: If breaking changes are needed, consider prefixing (e.g., `V2_*`)

## Installation & Usage

### TypeScript Installation

Since this library doesn't have an npm package yet, you can consume it in two ways:

#### Option 1: Copy Types Directly

Copy the type definition files into your project:

```bash
# Copy files to your project
cp api-response.d.ts your-project/src/types/
cp api-error.d.ts your-project/src/types/
cp error-code.enum.ts your-project/src/types/
```

Then import:

```typescript
import { ApiResponse, ApiError, ErrorCode } from './types/api-response';
```

#### Option 2: Git Dependency

Install directly from the GitHub repository:

```bash
npm install rbalet/api-error-response
# or
yarn add rbalet/api-error-response
```

Then import:

```typescript
import { ApiResponse, ApiError, ErrorCode } from 'api-error-response';
```

**Note:** If importing from the Git repo, you may need to configure your TypeScript paths or bundler to resolve the `.d.ts` and `.ts` files correctly.

### Go Installation

#### Option 1: Go Module (Recommended)

Install using Go modules:

```bash
go get github.com/rbalet/api-error-response
```

Then import in your code:

```go
import apierror "github.com/rbalet/api-error-response"
```

#### Option 2: Copy Go Files

Copy the Go files into your project:

```bash
# Copy files to your project
cp error_code.go your-project/pkg/apierror/
cp api_error.go your-project/pkg/apierror/
cp api_response.go your-project/pkg/apierror/
cp helpers.go your-project/pkg/apierror/
```

**Note:** Requires Go 1.22 or later (for generics support).

## License

MIT License - see [LICENSE](LICENSE) file for details.

Copyright (c) 2026 Raphaël Balet
