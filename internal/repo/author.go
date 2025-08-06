package repo

import (
	"errors"
	"fmt"
	"hotaku-api/internal/domain/apperrors"
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

	// Query the database for author with the specified ID
	if err := r.db.Where("author_id = ?", id).First(&author).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrAuthorNotFound
		}

		return nil, fmt.Errorf("failed to retrieve author by ID: %w", err)
	}

	return &author, nil
}

// Update updates an existing author
func (r *AuthorRepositoryImpl) Update(author *entities.Author) error {
	// Update the author record in the database
	res := r.db.Where("author_id = ?", author.AuthorID).Updates(author)

	if res.Error != nil {
		return fmt.Errorf("failed to update author: %w", res.Error)
	}
	if res.RowsAffected == 0 {
		return apperrors.ErrAuthorNotFound
	}
	return nil
}

// Delete removes an author by ID (hard delete)
func (r *AuthorRepositoryImpl) Delete(id string) error {
	// Delete the author record from the database
	res := r.db.Where("author_id = ?", id).Delete(&entities.Author{})

	if res.Error != nil {
		return fmt.Errorf("failed to delete author: %w", res.Error)
	}
	if res.RowsAffected == 0 {
		return apperrors.ErrAuthorNotFound
	}
	return nil
}

// List retrieves a paginated list of all authors
func (r *AuthorRepositoryImpl) List(offset, limit int) ([]entities.Author, int64, error) {
	var authors []entities.Author
	var total int64

	// Get total count of authors in the database
	if err := r.db.Model(&entities.Author{}).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count authors: %w", err)
	}

	// Get paginated results ordered by author name
	if err := r.db.Model(&entities.Author{}).Order("author_name").Offset(offset).Limit(limit).Find(&authors).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to retrieve authors list: %w", err)
	}

	return authors, total, nil
}
