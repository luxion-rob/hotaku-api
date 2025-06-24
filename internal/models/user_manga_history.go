package models

import "time"

type UserMangaHistory struct {
	HistoryId      int       `gorm:"primaryKey;autoIncrement" json:"history_id"`
	ExternalId     string    `gorm:"type:char(36)" json:"external_id"`
	UserId         uint64    `json:"user_id"`
	MangaId        int       `json:"manga_id"`
	ReadChapterIds string    `gorm:"type:text" json:"read_chapter_ids"`
	ReadAt         time.Time `gorm:"autoCreateTime" json:"read_at"`
}
