package models

import "gorm.io/gorm"

type Media struct {
	gorm.Model
	FileName   string `gorm:"type:varchar(255);not null" json:"file_name"`
	FileSize   int64  `gorm:"not null" json:"file_size"`
	FileType   string `gorm:"type:varchar(50);not null" json:"file_type"`
	StorageKey string `gorm:"type:varchar(255);unique;not null" json:"storage_key"`
	URL        string `gorm:"type:varchar(255);not null" json:"url"`
	Key        string `gorm:"type:varchar(255)" json:"key"`
	BucketName string `gorm:"type:varchar(255)" json:"bucket_name"`
}
