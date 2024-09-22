package models

type Permission struct {
	ID                string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	RoleID            string `gorm:"type:uuid;not null"`
	Name              string `gorm:"type:varchar(255);not null"`
	Description       string `gorm:"type:text"`
	Read              bool   `gorm:"default:false"`
	ReadDescription   string `gorm:"type:text"`
	Update            bool   `gorm:"default:false"`
	UpdateDescription string `gorm:"type:text"`
	Create            bool   `gorm:"default:false"`
	CreateDescription string `gorm:"type:text"`
}
