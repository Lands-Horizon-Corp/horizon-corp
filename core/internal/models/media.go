package models

import (
	"time"

	"gorm.io/gorm"
)

type Media struct {
	ID           string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	URL          string    `gorm:"type:varchar(255);not null"`
	FileName     string    `gorm:"type:varchar(255);not null"`
	FileType     string    `gorm:"type:varchar(50);not null"`
	FileSize     int       `gorm:"not null"`
	UploadTime   time.Time `gorm:"not null"`
	Description  string    `gorm:"type:text"`
	BucketName   string    `gorm:"type:varchar(255);not null"`
	TemporaryURL string    `gorm:"type:varchar(255)"`
	gorm.Model
}
