package resources

import (
	"horizon/server/internal/models"
	"horizon/server/storage"
	"time"
)

type MediaResource struct {
	ID          uint                `json:"id"`
	FileName    string              `json:"fileName"`
	FileSize    int64               `json:"fileSize"`
	FileType    string              `json:"fileType"`
	StorageKey  string              `json:"storageKey"`
	URL         string              `json:"url"`
	Key         string              `json:"key,omitempty"`
	DownloadURL string              `json:"downloadURL"`
	BucketName  string              `json:"bucketName,omitempty"`
	Employees   []*EmployeeResource `json:"employees,omitempty"` // Updated to slice of pointers
	Members     []*MemberResource   `json:"members,omitempty"`   // Updated to slice of pointers
	Owners      []*OwnerResource    `json:"owners,omitempty"`    // Updated to slice of pointers
	Admins      []*AdminResource    `json:"admins,omitempty"`    // Updated to slice of pointers
	Companies   []*CompanyResource  `json:"companies,omitempty"` // Updated to slice of pointers
	Branches    []*BranchResource   `json:"branches,omitempty"`  // Updated to slice of pointers

	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// Convert models.Media to *MediaResource
func ToResourceMedia(media *models.Media) *MediaResource {
	temporaryURL, err := storage.GeneratePresignedURL(media.StorageKey)
	if err != nil {
		return &MediaResource{}
	}
	if media == nil {
		return nil
	}
	if media == nil {
		return nil
	}

	return &MediaResource{
		ID:          media.ID,
		FileName:    media.FileName,
		FileSize:    media.FileSize,
		FileType:    media.FileType,
		StorageKey:  media.StorageKey,
		URL:         media.URL,
		Key:         media.Key,
		BucketName:  media.BucketName,
		DownloadURL: temporaryURL,
		Employees:   ToResourceListEmployees(media.Employees),
		Members:     ToResourceListMembers(media.Members),
		Owners:      ToResourceListOwners(media.Owners),
		Admins:      ToResourceListAdmins(media.Admins),
		Companies:   ToResourceListCompanies(media.Companies),
		Branches:    ToResourceListBranch(media.Branches),
		CreatedAt:   media.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   media.UpdatedAt.Format(time.RFC3339),
	}
}

// Convert []*models.Media to []*MediaResource
func ToResourceListMedia(mediaList []*models.Media) []*MediaResource {
	var resources []*MediaResource
	for _, media := range mediaList {
		resources = append(resources, ToResourceMedia(media))
	}
	return resources
}
