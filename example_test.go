package apierror_test

import (
	"encoding/json"
	"fmt"

	apierror "github.com/rbalet/api-error-response"
)

// User represents a user model
type User struct {
	UserID string `json:"userId"`
	Email  string `json:"email"`
}

func ExampleSuccessResponse() {
	// Create a success response
	user := User{
		UserID: "usr_1234567890",
		Email:  "user@example.com",
	}
	response := apierror.NewSuccessResponse(user)

	// Marshal to JSON
	jsonData, _ := json.MarshalIndent(response, "", "  ")
	fmt.Println(string(jsonData))

	// Output:
	// {
	//   "data": {
	//     "userId": "usr_1234567890",
	//     "email": "user@example.com"
	//   }
	// }
}

func ExampleNonValidationError() {
	// Create an authentication error
	message := "Invalid or expired authentication token"
	traceID := "trace-abc123"
	authError := apierror.NewAuthError(apierror.AuthUnauthorized, message, traceID)

	response := apierror.NewErrorResponse[User](authError)

	// Marshal to JSON
	jsonData, _ := json.MarshalIndent(response, "", "  ")
	fmt.Println(string(jsonData))

	// Output:
	// {
	//   "error": {
	//     "type": "AUTH",
	//     "code": "AUTH_UNAUTHORIZED",
	//     "message": "Invalid or expired authentication token",
	//     "traceId": "trace-abc123",
	//     "timestamp": "2026-02-18T12:30:00Z"
	//   }
	// }
}

func ExampleValidationError() {
	// Create a validation error with multiple issues
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

	validationError := apierror.NewValidationError("Request validation failed", issues, "trace-def456")
	response := apierror.NewErrorResponse[User](validationError)

	// Marshal to JSON
	jsonData, _ := json.MarshalIndent(response, "", "  ")
	fmt.Println(string(jsonData))

	// Output:
	// {
	//   "error": {
	//     "type": "VALIDATION",
	//     "code": "VALIDATION_FAILED",
	//     "message": "Request validation failed",
	//     "traceId": "trace-def456",
	//     "timestamp": "2026-02-18T12:35:00Z",
	//     "issues": [
	//       {
	//         "code": "VALIDATION_FIELD_REQUIRED",
	//         "path": ["user", "email"],
	//         "message": "Email is required"
	//       },
	//       {
	//         "code": "VALIDATION_FIELD_TOO_SHORT",
	//         "path": ["user", "password"],
	//         "message": "Password must be at least 8 characters",
	//         "meta": {
	//           "min": 8,
	//           "actual": 5
	//         }
	//       },
	//       {
	//         "code": "VALIDATION_FIELD_INVALID_FORMAT",
	//         "path": ["user", "phoneNumber"],
	//         "message": "Phone number format is invalid"
	//       }
	//     ]
	//   }
	// }
}

func ExampleTypeNarrowing() {
	// Demonstrating type checking
	user := User{UserID: "123", Email: "test@example.com"}
	response := apierror.NewSuccessResponse(user)

	if response.IsSuccess() {
		fmt.Printf("Success! User ID: %s\n", response.Data.UserID)
	} else if response.IsError() {
		fmt.Printf("Error: %s\n", *response.Error.GetMessage())
	}

	// Output:
	// Success! User ID: 123
}

func ExampleErrorTypeChecking() {
	// Create an error and check if it's a validation error
	message := "Request validation failed"
	issues := []apierror.ValidationIssue{}
	validationErr := apierror.NewValidationError(message, issues, "trace-001")

	if validationErr.IsValidationError() {
		fmt.Println("This is a validation error")
		if len(validationErr.Issues) > 0 {
			for _, issue := range validationErr.Issues {
				fmt.Printf("Field error: %s\n", *issue.Message)
			}
		}
	}

	// Output:
	// This is a validation error
}
