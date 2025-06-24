package models

type Role struct {
	RoleID   int    `gorm:"column:role_id;primaryKey;autoIncrement" json:"role_id"`
	RoleName string `gorm:"column:role_name;type:varchar(100);not null" json:"role_name"`
}
