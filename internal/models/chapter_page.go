package models

type ChapterPage struct {
	PageID     int    `gorm:"column:page_id;primaryKey;autoIncrement" json:"page_id"`
	ExternalID string `gorm:"column:external_id;type:char(36);uniqueIndex;not null" json:"external_id"`
	ChapterID  *int   `gorm:"column:chapter_id;index" json:"chapter_id"`
	ImageURL   string `gorm:"column:image_url;type:varchar(500);not null" json:"image_url"`
	PageNumber *int   `gorm:"column:page_number" json:"page_number"`
}
