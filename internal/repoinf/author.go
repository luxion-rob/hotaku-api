package repoinf

import "hotaku-api/internal/domain/entities"

type AuthorRepository interface {
	Create(author *entities.Author) error
	GetByID(id string) (*entities.Author, error)
	Update(author *entities.Author) error
	Delete(id string) error
	List(offset, limit int) (entities.AuthorList, int64, error)
}
