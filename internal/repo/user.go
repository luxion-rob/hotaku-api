package repo

import (
	"errors"
	"fmt"
	"hotaku-api/internal/domain/entities"
	"hotaku-api/internal/repoinf"
	"time"

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
	err := r.db.Model(&entities.User{}).Where("email = ? AND deleted_flag = ?", user.Email, false).Count(&count).Error
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
func (r *UserRepositoryImpl) GetByID(id string) (*entities.User, error) {
	var user entities.User
	if err := r.db.Where("user_id = ? AND deleted_flag = ?", id, false).Preload("Role").First(&user).Error; err != nil {
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
	if err := r.db.Where("email = ? AND deleted_flag = ?", email, false).Preload("Role").First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to retrieve user by email: %w", err)
	}
	return &user, nil
}

// Update updates an existing user
func (r *UserRepositoryImpl) Update(user *entities.User) error {
	res := r.db.Model(user).Where("user_id = ? AND deleted_flag = ?", user.UserID, false).Updates(user)
	if res.Error != nil {
		return fmt.Errorf("failed to update user: %w", res.Error)
	}
	if res.RowsAffected == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}

// Delete removes a user by ID (hard delete)
func (r *UserRepositoryImpl) Delete(id string) error {
	res := r.db.Where("user_id = ?", id).Delete(&entities.User{})
	if res.Error != nil {
		return fmt.Errorf("failed to delete user: %w", res.Error)
	}
	if res.RowsAffected == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}

// SoftDelete marks a user as deleted (soft delete)
func (r *UserRepositoryImpl) SoftDelete(id string) error {
	now := time.Now()
	res := r.db.Model(&entities.User{}).Where("user_id = ? AND deleted_flag = ?", id, false).Updates(map[string]interface{}{
		"deleted_flag": true,
		"deleted_at":   now,
	})
	if res.Error != nil {
		return fmt.Errorf("failed to soft delete user: %w", res.Error)
	}
	if res.RowsAffected == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}

// List retrieves a paginated list of all users (including deleted)
func (r *UserRepositoryImpl) List(offset, limit int) ([]entities.User, int64, error) {
	var users []entities.User
	var total int64

	// Get total count
	if err := r.db.Model(&entities.User{}).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count users: %w", err)
	}

	// Get paginated results
	if err := r.db.Preload("Role").Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to retrieve users list: %w", err)
	}

	return users, total, nil
}

// ListActive retrieves a paginated list of active users (not deleted)
func (r *UserRepositoryImpl) ListActive(offset, limit int) ([]entities.User, int64, error) {
	var users []entities.User
	var total int64

	// Get total count of active users
	if err := r.db.Model(&entities.User{}).Where("deleted_flag = ?", false).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count active users: %w", err)
	}

	// Get paginated results of active users
	if err := r.db.Where("deleted_flag = ?", false).Preload("Role").Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to retrieve active users list: %w", err)
	}

	return users, total, nil
}
