package models

import "time"

type Notification struct {
	NotificationId int       `gorm:"primaryKey;autoIncrement" json:"notification_id"`
	ExternalId     string    `gorm:"type:char(36)" json:"external_id"`
	Message        string    `gorm:"type:varchar(500)" json:"message"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
}
