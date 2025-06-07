package interfaces

import "time"

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

// NewSuccessResponse creates a new success response
func NewSuccessResponse(message string, data interface{}) *BaseResponse {
	return &BaseResponse{
		Success:   true,
		Message:   message,
		Data:      data,
		Timestamp: time.Now().Unix(),
	}
}

// NewErrorResponse creates a new error response
func NewErrorResponse(message string, err interface{}) *BaseResponse {
	return &BaseResponse{
		Success:   false,
		Message:   message,
		Error:     err,
		Timestamp: time.Now().Unix(),
	}
}

// NewAuthResponse creates a new authentication response
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

// NewUserResponse creates a new user response
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

// NewValidationErrorResponse creates a new validation error response
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

// NewHealthResponse creates a new health response
func NewHealthResponse() *HealthResponse {
	return &HealthResponse{
		Status:    "healthy",
		Message:   "API is running smoothly",
		Timestamp: time.Now().Unix(),
		Version:   "1.0.0",
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