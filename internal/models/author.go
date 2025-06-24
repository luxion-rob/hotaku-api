package models

type Author struct {
	AuthorId   int     `gorm:"primaryKey;autoIncrement" json:"author_id"`
	ExternalId string  `gorm:"type:char(36)" json:"external_id"`
	AuthorName string  `gorm:"type:varchar(255)" json:"author_name"`
	AuthorBio  *string `gorm:"type:text" json:"author_bio,omitempty"`
}
