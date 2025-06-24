package models

import "time"

type UserFavoriteManga struct {
	FavoriteId int       `gorm:"column:favorite_id;primaryKey;autoIncrement" json:"favorite_id"`
	ExternalId string    `gorm:"column:external_id;type:char(36);not null;unique" json:"external_id"`
	UserId     uint64    `gorm:"column:user_id;not null;index:idx_user_manga,unique" json:"user_id"`
	MangaId    int       `gorm:"column:manga_id;not null;index:idx_user_manga,unique" json:"manga_id"`
	AddedAt    time.Time `gorm:"column:added_at;autoCreateTime" json:"added_at"`
}

func (UserFavoriteManga) TableName() string {
	return "user_favorite_mangas"
}
