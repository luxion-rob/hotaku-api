package response

import "hotaku-api/internal/domain/dto"

// AuthResponse represents authentication response
type AuthResponse struct {
	Token string       `json:"token,omitempty"`
	User  *dto.UserDTO `json:"user,omitempty"`
}
