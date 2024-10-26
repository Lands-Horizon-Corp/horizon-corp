package resources

import (
	"horizon/server/internal/models"
	"time"
)

type MediaResource struct {
	ID         uint               `json:"id"`
	FileName   string             `json:"fileName"`
	FileSize   int64              `json:"fileSize"`
	FileType   string             `json:"fileType"`
	StorageKey string             `json:"storageKey"`
	URL        string             `json:"url"`
	Key        string             `json:"key,omitempty"`
	BucketName string             `json:"bucketName,omitempty"`
	Employees  []EmployeeResource `json:"employees,omitempty"`
	Members    []MemberResource   `json:"members,omitempty"`
	Owners     []OwnerResource    `json:"owners,omitempty"`
	Admins     []AdminResource    `json:"admins,omitempty"`
	Companies  []CompanyResource  `json:"companies,omitempty"`
	Branches   []BranchResource   `json:"branches,omitempty"`

	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func ToResourceMedia(media models.Media) MediaResource {
	return MediaResource{
		ID:         media.ID,
		FileName:   media.FileName,
		FileSize:   media.FileSize,
		FileType:   media.FileType,
		StorageKey: media.StorageKey,
		URL:        media.URL,
		Key:        media.Key,
		BucketName: media.BucketName,
		Employees:  ToResourceListEmployees(media.Employees),
		Members:    ToResourceListMembers(media.Members),
		Owners:     ToResourceListOwners(media.Owners),
		Admins:     ToResourceListAdmins(media.Admins),
		Companies:  ToResourceListCompanies(media.Companies),
		Branches:   ToResourceListBranch(media.Branches),

		CreatedAt: media.CreatedAt.Format(time.RFC3339),
		UpdatedAt: media.UpdatedAt.Format(time.RFC3339),
	}
}

func ToResourceListMedia(mediaList []models.Media) []MediaResource {
	var resources []MediaResource
	for _, media := range mediaList {
		resources = append(resources, ToResourceMedia(media))
	}
	return resources
}
