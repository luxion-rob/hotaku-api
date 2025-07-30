package entities

import "time"

type Author struct {
	AuthorID   string    `json:"author_id" gorm:"type:char(36);primaryKey"`
	ExternalID string    `json:"external_id" gorm:"type:char(36);not null"`
	AuthorName string    `json:"author_name" gorm:"type:char(50);not null"`
	AuthorBio  *string   `json:"author_bio,omitempty" gorm:"type:text"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
