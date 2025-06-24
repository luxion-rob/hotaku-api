package models

type Author struct {
	AuthorID   int     `gorm:"column:author_id;primaryKey;autoIncrement" json:"author_id"`
	ExternalID string  `gorm:"column:external_id;type:char(36);uniqueIndex;not null" json:"external_id"`
	AuthorName string  `gorm:"column:author_name;type:varchar(255);not null" json:"author_name"`
	AuthorBio  *string `gorm:"column:author_bio;type:text" json:"author_bio,omitempty"`
}
