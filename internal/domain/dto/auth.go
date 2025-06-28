package dto

import (
	"net/http"
	"time"
)

// AuthRequest defines the interface for authentication-related requests
type AuthRequest interface {
	Validate() error
}

// RegisterRequest represents user registration request
type RegisterRequest struct {
	Name     string `json:"name" binding:"required,min=2,max=100"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=100"`
}

// LoginRequest represents user login request
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// UpdateProfileRequest represents profile update request
type UpdateProfileRequest struct {
	Name  string `json:"name,omitempty" binding:"omitempty,min=2,max=100"`
	Email string `json:"email,omitempty" binding:"omitempty,email"`
}

// ChangePasswordRequest represents password change request
type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required,min=6,max=100"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=NewPassword"`
}

// ValidationError represents a validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *ValidationError) Error() string {
	return e.Message
}

// GetField returns the field name
func (e *ValidationError) GetField() string {
	return e.Field
}

// GetMessage returns the error message
func (e *ValidationError) GetMessage() string {
	return e.Message
}

// APIResponse defines the interface for all API responses
type APIResponse interface {
	GetStatus() int
	GetMessage() string
}

// BaseResponse represents the base structure for all API responses
type BaseResponse struct {
	StatusCode int         `json:"statuscode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Error      interface{} `json:"error,omitempty"`
	Timestamp  int64       `json:"timestamp"`
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

// UserData represents user data in responses
type UserData struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewSuccessResponse creates a new success response
func NewSuccessResponse(message string, data interface{}) *BaseResponse {
	return &BaseResponse{
		StatusCode: http.StatusOK,
		Message:    message,
		Data:       data,
		Timestamp:  time.Now().Unix(),
	}
}

// NewErrorResponse creates a new error response
func NewErrorResponse(message string, err interface{}) *BaseResponse {
	return &BaseResponse{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
		Error:      err,
		Timestamp:  time.Now().Unix(),
	}
}

// NewAuthResponse creates a new authentication response
func NewAuthResponse(message, token string, user *UserData) *AuthResponse {
	return &AuthResponse{
		BaseResponse: BaseResponse{
			StatusCode: http.StatusOK,
			Message:    message,
			Timestamp:  time.Now().Unix(),
		},
		Token: token,
		User:  user,
	}
}

// NewUserResponse creates a new user response
func NewUserResponse(message string, user *UserData) *UserResponse {
	return &UserResponse{
		BaseResponse: BaseResponse{
			StatusCode: http.StatusOK,
			Message:    message,
			Timestamp:  time.Now().Unix(),
		},
		User: user,
	}
}

// NewValidationErrorResponse creates a new validation error response
func NewValidationErrorResponse(message string, errors []ValidationError) *ErrorResponse {
	return &ErrorResponse{
		BaseResponse: BaseResponse{
			StatusCode: http.StatusBadRequest,
			Message:    message,
			Timestamp:  time.Now().Unix(),
		},
		Details: errors,
	}
}

// GetStatus implements APIResponse interface
func (r *BaseResponse) GetStatus() int {
	return r.StatusCode
}

// GetMessage implements APIResponse interface
func (r *BaseResponse) GetMessage() string {
	return r.Message
}
