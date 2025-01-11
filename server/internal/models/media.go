package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Media struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Fields
	FileName    string `gorm:"type:varchar(255);unsigned" json:"file_name"`
	FileSize    int64  `gorm:"unsigned" json:"file_size"`
	FileType    string `gorm:"type:varchar(50);unsigned" json:"file_type"`
	StorageKey  string `gorm:"type:varchar(255);unique;unsigned" json:"storage_key"`
	URL         string `gorm:"type:varchar(255);unsigned" json:"url"`
	Key         string `gorm:"type:varchar(255)" json:"key"`
	BucketName  string `gorm:"type:varchar(255)" json:"bucket_name"`
	Description string `json:"description" gorm:"type:text"`

	// Enhancements
	IsPublic  bool       `gorm:"default:false" json:"is_public"`
	ExpiresAt *time.Time `gorm:"type:timestamp" json:"expires_at"`

	// Metadata (grouping related fields)
	Metadata Metadata `gorm:"embedded" json:"metadata"`

	// Relationship 0 to many
	Employees []*Employee `gorm:"foreignKey:MediaID" json:"employees"`
	Members   []*Member   `gorm:"foreignKey:MediaID" json:"members"`
	Owners    []*Owner    `gorm:"foreignKey:MediaID" json:"owners"`
	Admins    []*Admin    `gorm:"foreignKey:MediaID" json:"admins"`
	Companies []*Company  `gorm:"foreignKey:MediaID" json:"companies"`
	Branches  []*Branch   `gorm:"foreignKey:MediaID" json:"branches"`
}

// Metadata structure
type Metadata struct {
	UploadedBy    string `json:"uploaded_by"`
	UploadedAt    string `json:"uploaded_at"`
	LastAccessed  string `json:"last_accessed"`
	DownloadCount int    `json:"download_count"`
	FileChecksum  string `json:"file_checksum"`
}

type MediaResource struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`

	Description string              `json:"description"`
	FileName    string              `json:"fileName"`
	FileSize    int64               `json:"fileSize"`
	FileType    string              `json:"fileType"`
	StorageKey  string              `json:"storageKey"`
	URL         string              `json:"uRL"`
	Key         string              `json:"key"`
	DownloadURL string              `json:"downloadURL"`
	BucketName  string              `json:"bucketName"`
	IsPublic    bool                `json:"isPublic"`
	ExpiresAt   string              `json:"expiresAt,omitempty"`
	Metadata    Metadata            `json:"metadata"`
	Employees   []*EmployeeResource `json:"employees"`
	Members     []*MemberResource   `json:"members"`
	Owners      []*OwnerResource    `json:"owners"`
	Admins      []*AdminResource    `json:"admins"`
	Companies   []*CompanyResource  `json:"companies"`
	Branches    []*BranchResource   `json:"branches"`
}

func (m *ModelTransformer) MediaToResource(media *Media) *MediaResource {
	if media == nil {
		return nil
	}
	temporaryURL, err := m.storage.GeneratePresignedURL(media.StorageKey)
	if err != nil {
		return nil
	}

	expiresAt := ""
	if media.ExpiresAt != nil {
		expiresAt = media.ExpiresAt.Format(time.RFC3339)
	}

	return &MediaResource{

		ID:        media.ID,
		CreatedAt: media.CreatedAt.Format(time.RFC3339),
		UpdatedAt: media.UpdatedAt.Format(time.RFC3339),
		DeletedAt: media.DeletedAt.Time.Format(time.RFC3339),

		Description: media.Description,
		FileName:    media.FileName,
		FileSize:    media.FileSize,
		FileType:    media.FileType,
		StorageKey:  media.StorageKey,
		URL:         media.URL,
		Key:         media.Key,
		BucketName:  media.BucketName,
		DownloadURL: temporaryURL,
		IsPublic:    media.IsPublic,
		ExpiresAt:   expiresAt,
		Metadata:    media.Metadata,
		Employees:   m.EmployeeToResourceList(media.Employees),
		Members:     m.MemberToResourceList(media.Members),
		Owners:      m.OwnerToResourceList(media.Owners),
		Admins:      m.AdminToResourceList(media.Admins),
		Companies:   m.CompanyToResourceList(media.Companies),
		Branches:    m.BranchToResourceList(media.Branches),
	}
}

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
