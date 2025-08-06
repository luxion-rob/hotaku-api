package dto

import (
	"time"
)

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

type ChangePasswordDTO struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}
