package dto

import "time"

// UserDTO represents user data in responses
type UserDTO struct {
	UserID    string    `json:"user_id"`
	RoleID    string    `json:"role_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
