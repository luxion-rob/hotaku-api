package entities

import (
	"hotaku-api/internal/domain/dto"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User represents the core user entity in the domain layer
type User struct {
	UserID      string     `json:"user_id" gorm:"type:char(36);primaryKey"`
	RoleID      string     `json:"role_id" gorm:"type:char(36);not null"`
	Email       string     `json:"email" gorm:"unique;not null"`
	Password    string     `json:"-" gorm:"not null"`
	Name        string     `json:"name" gorm:"not null"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedFlag bool       `json:"deleted_flag" gorm:"not null;default:false"`
	DeletedAt   *time.Time `json:"deleted_at"`

	// Relationships
	Role *Role `json:"role,omitempty" gorm:"foreignKey:RoleID;references:RoleID"`
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

// IsDeleted checks if the user is soft deleted
func (u *User) IsDeleted() bool {
	return u.DeletedFlag
}

func (u *User) ToDTO() *dto.UserDTO {
	return &dto.UserDTO{
		UserID: u.UserID,
		RoleID: u.RoleID,
		Name:   u.Name,
		Email:  u.Email,
	}
}

func (u *User) ToAuthResponse(token string) *dto.AuthResponse {
	return &dto.AuthResponse{
		Token: token,
		User:  u.ToDTO(),
	}
}
