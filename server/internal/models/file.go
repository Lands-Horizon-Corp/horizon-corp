package models

import "gorm.io/gorm"

type File struct {
	gorm.Model
	FileName   string `gorm:"size:255;not null"`
	FileSize   int64  `gorm:"not null"`
	FileType   string `gorm:"size:50;not null"`
	StorageKey string `gorm:"size:255;not null;unique"`
	URL        string `gorm:"size:255"`
	BucketName string `gorm:"size:255"`
}
