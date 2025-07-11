package repoinf

import (
	"hotaku-api/internal/domain/entities"
)

// UserRepository defines the interface for user data access
type UserRepository interface {
	Create(user *entities.User) error
	GetByID(id string) (*entities.User, error)
	GetByEmail(email string) (*entities.User, error)
	Update(user *entities.User) error
	Delete(id string) error
	SoftDelete(id string) error
	List(offset, limit int) ([]entities.User, int64, error)
	ListActive(offset, limit int) ([]entities.User, int64, error)
}
