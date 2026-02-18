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

func Example_successResponse() {
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

func Example_nonValidationError() {
	// Create an authentication error
	message := "Invalid or expired authentication token"
	traceID := "trace-abc123"
	authError := apierror.NewAuthError(apierror.AuthUnauthorized, message, traceID)

	response := apierror.NewErrorResponse[User](authError)

	// Check if it's an error response
	if response.IsError() {
		fmt.Println("Error type:", response.Error.GetType())
		fmt.Println("Error code:", *response.Error.GetCode())
		fmt.Println("Error message:", *response.Error.GetMessage())
	}

	// Output:
	// Error type: AUTH
	// Error code: AUTH_UNAUTHORIZED
	// Error message: Invalid or expired authentication token
}

func Example_validationError() {
	// Create a validation error with multiple issues
	emailMsg := "Email is required"
	passwordMsg := "Password must be at least 8 characters"

	emailCode := apierror.ValidationFieldRequired
	passwordCode := apierror.ValidationFieldTooShort

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
	}

	validationError := apierror.NewValidationError("Request validation failed", issues, "trace-def456")
	response := apierror.NewErrorResponse[User](validationError)

	// Check validation error details
	if response.IsError() && response.Error.IsValidationError() {
		if validationErr, ok := response.Error.(*apierror.ValidationError); ok {
			fmt.Println("Validation error with", len(validationErr.Issues), "issues")
			for i, issue := range validationErr.Issues {
				fmt.Printf("Issue %d: %s\n", i+1, *issue.Message)
			}
		}
	}

	// Output:
	// Validation error with 2 issues
	// Issue 1: Email is required
	// Issue 2: Password must be at least 8 characters
}

func Example_typeNarrowing() {
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

func Example_errorTypeChecking() {
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
