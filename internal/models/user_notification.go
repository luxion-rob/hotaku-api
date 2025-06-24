package models

type UserNotification struct {
	UserId         uint64 `gorm:"primaryKey" json:"user_id"`
	NotificationId int    `gorm:"primaryKey" json:"notification_id"`
}
