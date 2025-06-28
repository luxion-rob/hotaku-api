package entities

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User represents the core user entity in the domain layer
type User struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// HashPassword hashes the user's password using bcrypt
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword verifies if the provided password matches the hashed password
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// IsValidEmail checks if the email is valid
func (u *User) IsValidEmail() bool {
	// Basic email validation - you might want to use a more robust validation library
	return len(u.Email) > 0 && len(u.Email) <= 254
}

// IsValidName checks if the name is valid
func (u *User) IsValidName() bool {
	return len(u.Name) >= 2 && len(u.Name) <= 100
}
