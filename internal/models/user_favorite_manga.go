package models

import "time"

type UserFavoriteManga struct {
	FavoriteId int       `gorm:"primaryKey;autoIncrement" json:"favorite_id"`
	ExternalId string    `gorm:"type:char(36)" json:"external_id"`
	UserId     uint64    `json:"user_id"`
	MangaId    int       `json:"manga_id"`
	AddedAt    time.Time `gorm:"autoCreateTime" json:"added_at"`
}
