package models

type Role struct {
	RoleId   int    `gorm:"primaryKey;autoIncrement" json:"role_id"`
	RoleName string `gorm:"type:varchar(100)" json:"role_name"`
}
