package entities

import (
	"hotaku-api/internal/domain/dto"
	"time"
)

type Author struct {
	AuthorID   string    `json:"author_id" gorm:"type:char(36);primaryKey"`
	ExternalID string    `json:"external_id" gorm:"type:char(36);not null,unique"`
	AuthorName string    `json:"author_name" gorm:"type:char(50);not null"`
	AuthorBio  *string   `json:"author_bio,omitempty" gorm:"type:text"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type AuthorList []Author

func (a *Author) ToDTO() *dto.AuthorDTO {
	return &dto.AuthorDTO{
		AuthorID:   a.AuthorID,
		AuthorName: a.AuthorName,
		AuthorBio:  a.AuthorBio,
	}
}

func (a *Author) ToAuthorResponse() *dto.AuthorResponse {
	authorDTO := a.ToDTO()

	return &dto.AuthorResponse{
		AuthorDTO:  *authorDTO,
		ExternalID: a.ExternalID,
		CreatedAt:  a.CreatedAt,
		UpdatedAt:  a.UpdatedAt,
	}
}

func (al AuthorList) ToAuthorListResponse(total int64, offset, limit int) *dto.AuthorListResponse {
	authorDTOs := make([]dto.AuthorDTO, len(al))
	for i := range al {
		authorDTOs[i] = *al[i].ToDTO()
	}

	return &dto.AuthorListResponse{
		Authors: authorDTOs,
		Total:   total,
		Offset:  offset,
		Limit:   limit,
	}
}
