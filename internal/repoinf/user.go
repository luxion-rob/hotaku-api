package repoinf

import (
	"hotaku-api/internal/domain/entities"
)

// UserRepository defines the interface for user data access
type UserRepository interface {
	Create(user *entities.User) error
	GetByID(id uint) (*entities.User, error)
	GetByEmail(email string) (*entities.User, error)
	Update(user *entities.User) error
	Delete(id uint) error
	List(offset, limit int) ([]entities.User, int64, error)
}
