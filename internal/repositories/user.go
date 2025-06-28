package repositories

import (
	"errors"
	"fmt"
	"hotaku-api/internal/domain/entities"
	"hotaku-api/internal/domain/interfaces"

	"gorm.io/gorm"
)

// UserRepositoryImpl implements the user repository interface
type UserRepositoryImpl struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepositoryImpl
func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &UserRepositoryImpl{db: db}
}

// Create saves a new user to the database
func (r *UserRepositoryImpl) Create(user *entities.User) error {
	return r.db.Create(user).Error
}

// GetByID retrieves a user by ID
func (r *UserRepositoryImpl) GetByID(id uint) (*entities.User, error) {
	var user entities.User
	if err := r.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to retrieve user by ID: %w", err)
	}
	return &user, nil
}

// GetByEmail retrieves a user by email
func (r *UserRepositoryImpl) GetByEmail(email string) (*entities.User, error) {
	var user entities.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to retrieve user by email: %w", err)
	}
	return &user, nil
}

// Update updates an existing user
func (r *UserRepositoryImpl) Update(user *entities.User) error {
	if err := r.db.Save(user).Error; err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

// Delete removes a user by ID
func (r *UserRepositoryImpl) Delete(id uint) error {
	if err := r.db.Delete(&entities.User{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

// List retrieves a paginated list of users
func (r *UserRepositoryImpl) List(offset, limit int) ([]entities.User, int64, error) {
	var users []entities.User
	var total int64

	// Get total count
	if err := r.db.Model(&entities.User{}).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count users: %w", err)
	}

	// Get paginated results
	if err := r.db.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to retrieve users list: %w", err)
	}

	return users, total, nil
}
