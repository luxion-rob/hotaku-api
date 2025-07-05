package entities

// Role represents the role entity in the domain layer
type Role struct {
	RoleID   string `json:"role_id" gorm:"type:char(36);primaryKey"`
	RoleName string `json:"role_name" gorm:"unique;not null"`
}
