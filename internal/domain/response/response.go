package response

import (
	"time"
)

// BaseResponse represents the base response for all API responses
type BaseResponse struct {
	StatusCode int    `json:"statuscode"`
	Message    string `json:"message"`
	Data       any    `json:"data,omitempty"`
	Error      any    `json:"error,omitempty"`
	Timestamp  int64  `json:"timestamp"`
}

// SuccessResponse returns a new success response
func SuccessResponse(code int, message string, data any) *BaseResponse {
	return &BaseResponse{
		StatusCode: code,
		Message:    message,
		Data:       data,
		Timestamp:  time.Now().Unix(),
	}
}

// ErrorResponse returns a new error response
func ErrorResponse(code int, message string, error any) *BaseResponse {
	return &BaseResponse{
		StatusCode: code,
		Message:    message,
		Error:      error,
		Timestamp:  time.Now().Unix(),
	}
}
