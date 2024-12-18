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

	Employees []*Employee `gorm:"foreignKey:MediaID" json:"employees"`
	Members   []*Member   `gorm:"foreignKey:MediaID" json:"members"`
	Owners    []*Owner    `gorm:"foreignKey:MediaID" json:"owners"`
	Admins    []*Admin    `gorm:"foreignKey:MediaID" json:"admins"`
	Companies []*Company  `gorm:"foreignKey:MediaID" json:"companies"`
	Branches  []*Branch   `gorm:"foreignKey:MediaID" json:"branches"`
}
