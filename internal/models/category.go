package models

type Category struct {
	CategoryId   int    `gorm:"primaryKey;autoIncrement" json:"category_id"`
	ExternalId   string `gorm:"type:char(36)" json:"external_id"`
	CategoryName string `gorm:"type:varchar(255)" json:"category_name"`
}
