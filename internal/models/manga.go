package models

import "time"

type Manga struct {
	MangaId     int       `gorm:"primaryKey;autoIncrement" json:"manga_id"`
	ExternalId  string    `gorm:"type:char(36)"         json:"external_id"`
	Title       string    `gorm:"type:varchar(255)"     json:"title"`
	CreatedAt   time.Time `gorm:"autoCreateTime"        json:"created_at"`
	Status      string    `gorm:"type:varchar(100)"     json:"status"`
	Description *string   `gorm:"type:text"             json:"description,omitempty"`
}
