package repo

import (
	"errors"
	"fmt"
	"hotaku-api/internal/domain/entities"
	"hotaku-api/internal/repoinf"

	"gorm.io/gorm"
)

// AuthorRepositoryImpl implements the author repository interface
type AuthorRepositoryImpl struct {
	db *gorm.DB
}

// NewAuthorRepository creates a new instance of AuthorRepositoryImpl
func NewAuthorRepository(db *gorm.DB) repoinf.AuthorRepository {
	return &AuthorRepositoryImpl{db: db}
}

// Create saves a new author to the database
func (r *AuthorRepositoryImpl) Create(author *entities.Author) error {
	// Create the author
	if err := r.db.Create(author).Error; err != nil {
		return fmt.Errorf("failed to create author: %w", err)
	}
	return nil
}

// GetByID retrieves an author by ID
func (r *AuthorRepositoryImpl) GetByID(id string) (*entities.Author, error) {
	var author entities.Author

	if err := r.db.Where("author_id = ?", id).First(&author).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("author not found")
		}

		return nil, fmt.Errorf("failed to retrieve author by ID: %w", err)
	}

	return &author, nil
}

// Update updates an existing author
func (r *AuthorRepositoryImpl) Update(author *entities.Author) error {
	res := r.db.Where("author_id = ?", author.AuthorID).Updates(author)

	if res.Error != nil {
		return fmt.Errorf("failed to update author: %w", res.Error)
	}
	if res.RowsAffected == 0 {
		return fmt.Errorf("author not found")
	}
	return nil
}

// Delete removes an author by ID (hard delete)
func (r *AuthorRepositoryImpl) Delete(id string) error {
	res := r.db.Where("author_id = ?", id).Delete(&entities.Author{})

	if res.Error != nil {
		return fmt.Errorf("failed to delete author: %w", res.Error)
	}
	if res.RowsAffected == 0 {
		return fmt.Errorf("author not found")
	}
	return nil
}
