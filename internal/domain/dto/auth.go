package dto

import (
	"time"
)

// AuthResponse represents authentication response
type AuthResponse struct {
	Token string   `json:"token,omitempty"`
	User  *UserDTO `json:"user,omitempty"`
}

// UserDTO represents user data in responses
type UserDTO struct {
	UserID    string    `json:"user_id"`
	RoleID    string    `json:"role_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AuthDTO struct {
	UserID    string    `json:"user_id"`
	RoleID    string    `json:"role_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
