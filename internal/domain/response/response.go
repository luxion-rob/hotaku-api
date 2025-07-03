package response

import (
	"net/http"
	"time"
)

type BaseResponse struct {
	StatusCode int         `json:"statuscode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Error      interface{} `json:"error,omitempty"`
	Timestamp  int64       `json:"timestamp"`
}

func SuccessResponse(message string, data interface{}) *BaseResponse {
	return &BaseResponse{
		StatusCode: http.StatusOK,
		Message:    message,
		Data:       data,
		Timestamp:  time.Now().Unix(),
	}
}

func ErrorResponse(message string, error interface{}) *BaseResponse {
	return &BaseResponse{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
		Error:      error,
		Timestamp:  time.Now().Unix(),
	}
}
