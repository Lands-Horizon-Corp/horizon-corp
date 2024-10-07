package models

import "gorm.io/gorm"

type Media struct {
	gorm.Model
	FileName   string `json:"file_name" gorm:"size:255;not null"`
	FileSize   int64  `json:"file_size" gorm:"not null"`
	FileType   string `json:"file_type" gorm:"size:50;not null"`
	StorageKey string `json:"storage_key" gorm:"size:255;not null;unique"`
	URL        string `json:"url" gorm:"size:255"`
	BucketName string `json:"bucket_name" gorm:"size:255"`
}
