package models

type Group struct {
	GroupId    int    `gorm:"primaryKey;autoIncrement" json:"group_id"`
	ExternalId string `gorm:"type:char(36)" json:"external_id"`
	GroupName  string `gorm:"type:varchar(255)" json:"group_name"`
}
