package models

type ChapterPage struct {
	PageId     int    `gorm:"primaryKey;autoIncrement" json:"page_id"`
	ExternalId string `gorm:"type:char(36)" json:"external_id"`
	ChapterId  *int   `json:"chapter_id"`
	ImageUrl   string `gorm:"type:varchar(500)" json:"image_url"`
	PageNumber *int   `json:"page_number"`
}
