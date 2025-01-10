package models

import (
	"time"

	"gorm.io/gorm"
)

type Media struct {
	gorm.Model

	// Fields
	FileName   string `gorm:"type:varchar(255);unsigned" json:"file_name"`
	FileSize   int64  `gorm:"unsigned" json:"file_size"`
	FileType   string `gorm:"type:varchar(50);unsigned" json:"file_type"`
	StorageKey string `gorm:"type:varchar(255);unique;unsigned" json:"storage_key"`
	URL        string `gorm:"type:varchar(255);unsigned" json:"url"`
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

func (m *ModelTransformer) MediaToResource(media *Media) *MediaResource {
	if media == nil {
		return nil
	}
	temporaryURL, err := m.storage.GeneratePresignedURL(media.StorageKey)
	if err != nil {
		return nil
	}
	return &MediaResource{
		ID:        media.ID,
		CreatedAt: media.CreatedAt.Format(time.RFC3339),
		UpdatedAt: media.UpdatedAt.Format(time.RFC3339),

		FileName:    media.FileName,
		FileSize:    media.FileSize,
		FileType:    media.FileType,
		StorageKey:  media.StorageKey,
		URL:         media.URL,
		Key:         media.Key,
		BucketName:  media.BucketName,
		DownloadURL: temporaryURL,
		Employees:   m.EmployeeToResourceList(media.Employees),
		Members:     m.MemberToResourceList(media.Members),
		Owners:      m.OwnerToResourceList(media.Owners),
		Admins:      m.AdminToResourceList(media.Admins),
		Companies:   m.CompanyToResourceList(media.Companies),
		Branches:    m.BranchToResourceList(media.Branches),
	}
}

// MediaToResourceList implements Models.
func (m *ModelTransformer) MediaToResourceList(mediaList []*Media) []*MediaResource {
	if mediaList == nil {
		return nil
	}

	var mediaResources []*MediaResource
	for _, media := range mediaList {
		mediaResources = append(mediaResources, m.MediaToResource(media))
	}
	return mediaResources
}
