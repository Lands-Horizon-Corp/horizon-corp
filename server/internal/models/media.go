package models

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"gorm.io/gorm"
)

type Media struct {
	gorm.Model

	// Fields
	FileName   string `gorm:"type:varchar(255);not null" json:"file_name"`
	FileSize   int64  `gorm:"not null" json:"file_size"`
	FileType   string `gorm:"type:varchar(50);not null" json:"file_type"`
	StorageKey string `gorm:"type:varchar(255);unique;not null" json:"storage_key"`
	URL        string `gorm:"type:varchar(255);not null" json:"url"`
	Key        string `gorm:"type:varchar(255)" json:"key"`
	BucketName string `gorm:"type:varchar(255)" json:"bucket_name"`

	// Relationship 0 to many
	Employees []*Employee `gorm:"foreignKey:MediaID" json:"employees"`
	Members   []*Member   `gorm:"foreignKey:MediaID" json:"members"`
	Owners    []*Owner    `gorm:"foreignKey:MediaID" json:"owners"`
	Admins    []*Admin    `gorm:"foreignKey:MediaID" json:"admins"`
	Companies []*Company  `gorm:"foreignKey:MediaID" json:"companies"`
	Branches  []*Branch   `gorm:"foreignKey:MediaID" json:"branches"`
}

type MediaResource struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	FileName    string `json:"fileName"`
	FileSize    int64  `json:"fileSize"`
	FileType    string `json:"fileType"`
	StorageKey  string `json:"storageKey"`
	URL         string `json:"uRL"`
	Key         string `json:"key"`
	DownloadURL string `json:"downloadURL"`
	BucketName  string `json:"bucketName"`

	Employees []*EmployeeResource `json:"employees"`
	Members   []*MemberResource   `json:"members"`
	Owners    []*OwnerResource    `json:"owners"`
	Admins    []*AdminResource    `json:"admins"`
	Companies []*CompanyResource  `json:"companies"`
	Branches  []*BranchResource   `json:"branches"`
}

type MediaRepository struct {
	*Repository[Media]
}

func NewMediaRepository(db *providers.DatabaseService) *MediaRepository {
	return &MediaRepository{
		Repository: NewRepository[Media](db),
	}
}
