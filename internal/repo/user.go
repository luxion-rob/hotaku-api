package repo

import (
	"errors"
	"fmt"
	"hotaku-api/internal/domain/entities"
	"hotaku-api/internal/repoinf"

	"gorm.io/gorm"
)

// UserRepositoryImpl implements the user repository interface
type UserRepositoryImpl struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepositoryImpl
func NewUserRepository(db *gorm.DB) repoinf.UserRepository {
	return &UserRepositoryImpl{db: db}
}

// Create saves a new user to the database
func (r *UserRepositoryImpl) Create(user *entities.User) error {
	var count int64
	err := r.db.Model(&entities.User{}).Where("email = ?", user.Email).Count(&count).Error
	if err != nil {
		return fmt.Errorf("failed to find user: %w", err)
	}
	if count > 0 {
		return fmt.Errorf("user already exists")
	}

	err = r.db.Create(user).Error
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
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
	res := r.db.Model(user).Where("id = ?", user.ID).Updates(user)
	if res.Error != nil {
		return fmt.Errorf("failed to update user: %w", res.Error)
	}
	if res.RowsAffected == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}

// Delete removes a user by ID
func (r *UserRepositoryImpl) Delete(id uint) error {
	res := r.db.Delete(&entities.User{}, id)
	if res.Error != nil {
		return fmt.Errorf("failed to delete user: %w", res.Error)
	}
	if res.RowsAffected == 0 {
		return fmt.Errorf("user not found")
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
