package models

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"go.uber.org/fx"
	"go.uber.org/zap"
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

type (
	MediaResourceProvider interface {
		SeedDatabase()
		ToResource(media *Media) *MediaResource
		ToResourceList(media []*Media) []*MediaResource
	}
)

type MediaModel struct {
	lc              *fx.Lifecycle
	db              *gorm.DB
	logger          *zap.Logger
	adminModel      AdminResourceProvider
	employeeModel   EmployeeResourceProvider
	ownerModel      OwnerResourceProvider
	memberModel     MemberResourceProvider
	companyModel    CompanyResourceProvider
	branchModel     BranchResourceProvider
	storageProvider *providers.StorageProvider
}

func NewMediaModel(
	lc *fx.Lifecycle,
	db *gorm.DB,
	logger *zap.Logger,
	adminModel AdminResourceProvider,
	employeeModel EmployeeResourceProvider,
	ownerModel OwnerResourceProvider,
	memberModel MemberResourceProvider,
	companyModel CompanyResourceProvider,
	branchModel BranchResourceProvider,
	storageProvider *providers.StorageProvider,
) *MediaModel {
	return &MediaModel{
		lc:              lc,
		db:              db,
		logger:          logger,
		adminModel:      adminModel,
		employeeModel:   employeeModel,
		ownerModel:      ownerModel,
		memberModel:     memberModel,
		companyModel:    companyModel,
		branchModel:     branchModel,
		storageProvider: storageProvider,
	}
}

func (mm *MediaModel) SeedDatabase() {
}

func (mm *MediaModel) ToResource(media *Media) *MediaResource {
	if media == nil {
		return nil
	}

	temporaryURL, err := mm.storageProvider.GeneratePresignedURL(media.StorageKey)
	if err != nil {
		return nil
	}

	return &MediaResource{
		FileName:    media.FileName,
		FileSize:    media.FileSize,
		FileType:    media.FileType,
		StorageKey:  media.StorageKey,
		URL:         media.URL,
		Key:         media.Key,
		BucketName:  media.BucketName,
		DownloadURL: temporaryURL,
		Employees:   mm.employeeModel.ToResourceList(media.Employees),
		Members:     mm.memberModel.ToResourceList(media.Members),
		Owners:      mm.ownerModel.ToResourceList(media.Owners),
		Admins:      mm.adminModel.ToResourceList(media.Admins),
		Companies:   mm.companyModel.ToResourceList(media.Companies),
		Branches:    mm.branchModel.ToResourceList(media.Branches),
	}
}

func (mm *MediaModel) ToResourceList(mediaList []*Media) []*MediaResource {
	if mediaList == nil {
		return nil
	}

	var mediaResources []*MediaResource
	for _, media := range mediaList {
		mediaResources = append(mediaResources, mm.ToResource(media))
	}
	return mediaResources
}
