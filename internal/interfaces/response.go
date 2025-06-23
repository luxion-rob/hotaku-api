package interfaces

import (
	"os"
	"time"
)

// APIResponse defines the interface for all API responses
type APIResponse interface {
	GetStatus() int
	GetMessage() string
}

// BaseResponse represents the base structure for all API responses
type BaseResponse struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	Error     interface{} `json:"error,omitempty"`
	Timestamp int64       `json:"timestamp"`
}

// AuthResponse represents authentication response
type AuthResponse struct {
	BaseResponse
	Token string    `json:"token,omitempty"`
	User  *UserData `json:"user,omitempty"`
}

// UserResponse represents user-related responses
type UserResponse struct {
	BaseResponse
	User *UserData `json:"user,omitempty"`
}

// ErrorResponse represents error responses
type ErrorResponse struct {
	BaseResponse
	Details []ValidationError `json:"details,omitempty"`
}

// SuccessResponse represents success responses
type SuccessResponse struct {
	BaseResponse
}

// UserData represents user data in responses
type UserData struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// HealthResponse represents health check response
type HealthResponse struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
	Version   string `json:"version,omitempty"`
}

// PaginationMeta represents pagination metadata
type PaginationMeta struct {
	Page       int   `json:"page"`
	PerPage    int   `json:"per_page"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

// PaginatedResponse represents paginated responses
type PaginatedResponse struct {
	BaseResponse
	Meta PaginationMeta `json:"meta"`
}

// NewSuccessResponse returns a BaseResponse indicating a successful operation with the provided message and data. The response includes the current Unix timestamp.
func NewSuccessResponse(message string, data interface{}) *BaseResponse {
	return &BaseResponse{
		Success:   true,
		Message:   message,
		Data:      data,
		Timestamp: time.Now().Unix(),
	}
}

// NewErrorResponse returns a BaseResponse representing an error, with the provided message, error details, and the current timestamp.
func NewErrorResponse(message string, err interface{}) *BaseResponse {
	return &BaseResponse{
		Success:   false,
		Message:   message,
		Error:     err,
		Timestamp: time.Now().Unix(),
	}
}

// NewAuthResponse returns an AuthResponse indicating successful authentication, including a message, token, user data, and the current timestamp.
func NewAuthResponse(message, token string, user *UserData) *AuthResponse {
	return &AuthResponse{
		BaseResponse: BaseResponse{
			Success:   true,
			Message:   message,
			Timestamp: time.Now().Unix(),
		},
		Token: token,
		User:  user,
	}
}

// NewUserResponse returns a UserResponse indicating success, containing a message and user data with the current timestamp.
func NewUserResponse(message string, user *UserData) *UserResponse {
	return &UserResponse{
		BaseResponse: BaseResponse{
			Success:   true,
			Message:   message,
			Timestamp: time.Now().Unix(),
		},
		User: user,
	}
}

// NewValidationErrorResponse returns an ErrorResponse representing a failed operation with validation error details and a timestamp.
func NewValidationErrorResponse(message string, errors []ValidationError) *ErrorResponse {
	return &ErrorResponse{
		BaseResponse: BaseResponse{
			Success:   false,
			Message:   message,
			Timestamp: time.Now().Unix(),
		},
		Details: errors,
	}
}

// NewHealthResponse returns a HealthResponse indicating the API's health status, including a version string from the APP_VERSION environment variable or a default value.
func NewHealthResponse() *HealthResponse {
	var version string
	if version = os.Getenv("APP_VERSION"); version == "" {
		version = "1.0.0"
	}

	return &HealthResponse{
		Status:    "healthy",
		Message:   "API is running smoothly",
		Timestamp: time.Now().Unix(),
		Version:   version,
	}
}

// GetStatus implements APIResponse interface
func (r *BaseResponse) GetStatus() int {
	if r.Success {
		return 200
	}
	return 500
}

// GetMessage implements APIResponse interface
func (r *BaseResponse) GetMessage() string {
	return r.Message
}
