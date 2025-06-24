package models

import "time"

type MangaChapter struct {
	ChapterId     int       `gorm:"primaryKey;autoIncrement" json:"chapter_id"`
	ExternalId    string    `gorm:"type:char(36)" json:"external_id"`
	MangaId       *int      `json:"manga_id"`
	ChapterNumber int       `json:"chapter_number"`
	Title         string    `gorm:"type:varchar(255)" json:"title"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
}
