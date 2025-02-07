package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator"
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

type MediaRequest struct {
	ID *uint `json:"id"`

	FileName   string `json:"fileName" validate:"required,max=255"`
	FileSize   int64  `json:"fileSize" validate:"required,min=1"`
	FileType   string `json:"fileType" validate:"required,max=50"`
	StorageKey string `json:"storageKey" validate:"required,max=255"`
	URL        string `json:"url" validate:"required,url,max=255"`
	Key        string `json:"key,omitempty" validate:"max=255"`
	BucketName string `json:"bucketName,omitempty" validate:"max=255"`
}

func (m *ModelResource) MediaToResource(media *Media) *MediaResource {
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
func (m *ModelResource) MediaToResourceList(mediaList []*Media) []*MediaResource {
	if mediaList == nil {
		return nil
	}

	var mediaResources []*MediaResource
	for _, media := range mediaList {
		mediaResources = append(mediaResources, m.MediaToResource(media))
	}
	return mediaResources
}

// MediaToRecord converts a slice of Media pointers into CSV records and headers.
func (m *ModelResource) MediaToRecord(mediaList []*Media) ([][]string, []string) {
	// Convert Media structs to MediaResource structs
	resources := m.MediaToResourceList(mediaList)
	records := make([][]string, 0, len(resources))

	for _, media := range resources {
		// Basic Fields
		id := strconv.Itoa(int(media.ID))
		fileName := sanitizeCSVField(media.FileName)
		fileSize := strconv.FormatInt(media.FileSize, 10)
		fileType := sanitizeCSVField(media.FileType)
		storageKey := sanitizeCSVField(media.StorageKey)
		url := sanitizeCSVField(media.URL)
		key := sanitizeCSVField(media.Key)
		bucketName := sanitizeCSVField(media.BucketName)
		downloadURL := sanitizeCSVField(media.DownloadURL)
		createdAt := sanitizeCSVField(media.CreatedAt)
		updatedAt := sanitizeCSVField(media.UpdatedAt)

		// Handle Related Entities

		// Employees
		employees := "N/A"
		if len(media.Employees) > 0 {
			employeeNames := make([]string, 0, len(media.Employees))
			for _, emp := range media.Employees {
				fullName := fmt.Sprintf("%s %s", emp.FirstName, emp.LastName)
				employeeNames = append(employeeNames, sanitizeCSVField(fullName))
			}
			employees = strings.Join(employeeNames, "; ")
		}

		// Members
		members := "N/A"
		if len(media.Members) > 0 {
			memberNames := make([]string, 0, len(media.Members))
			for _, mem := range media.Members {
				fullName := fmt.Sprintf("%s %s", mem.FirstName, mem.LastName)
				memberNames = append(memberNames, sanitizeCSVField(fullName))
			}
			members = strings.Join(memberNames, "; ")
		}

		// Owners
		owners := "N/A"
		if len(media.Owners) > 0 {
			ownerNames := make([]string, 0, len(media.Owners))
			for _, own := range media.Owners {
				fullName := fmt.Sprintf("%s %s", own.FirstName, own.LastName)
				ownerNames = append(ownerNames, sanitizeCSVField(fullName))
			}
			owners = strings.Join(ownerNames, "; ")
		}

		// Admins
		admins := "N/A"
		if len(media.Admins) > 0 {
			adminNames := make([]string, 0, len(media.Admins))
			for _, adm := range media.Admins {
				fullName := fmt.Sprintf("%s %s", adm.FirstName, adm.LastName)
				adminNames = append(adminNames, sanitizeCSVField(fullName))
			}
			admins = strings.Join(adminNames, "; ")
		}

		// Companies
		companies := "N/A"
		if len(media.Companies) > 0 {
			companyNames := make([]string, 0, len(media.Companies))
			for _, comp := range media.Companies {
				companyNames = append(companyNames, sanitizeCSVField(comp.Name))
			}
			companies = strings.Join(companyNames, "; ")
		}

		// Branches
		branches := "N/A"
		if len(media.Branches) > 0 {
			branchNames := make([]string, 0, len(media.Branches))
			for _, branch := range media.Branches {
				branchNames = append(branchNames, sanitizeCSVField(branch.Name))
			}
			branches = strings.Join(branchNames, "; ")
		}

		// Assemble the record
		record := []string{
			id,
			fileName,
			fileSize,
			fileType,
			storageKey,
			url,
			key,
			bucketName,
			downloadURL,
			createdAt,
			updatedAt,
			employees,
			members,
			owners,
			admins,
			companies,
			branches,
		}
		records = append(records, record)
	}

	headers := []string{
		"ID",
		"File Name",
		"File Size",
		"File Type",
		"Storage Key",
		"URL",
		"Key",
		"Bucket Name",
		"Download URL",
		"Created At",
		"Updated At",
		"Employees",
		"Members",
		"Owners",
		"Admins",
		"Companies",
		"Branches",
	}

	return records, headers
}

func (m *ModelResource) ValidateMediaRequest(req *MediaRequest) error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return m.helpers.FormatValidationError(err)
	}
	return nil
}

func (m *ModelResource) MediaSeeders() error {
	m.logger.Info("Seeding Media")
	return nil
}
