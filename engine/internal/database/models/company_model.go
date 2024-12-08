package models

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model

	// Fields
	Name          string  `gorm:"type:varchar(255);not null" json:"name"`
	Description   string  `gorm:"type:text" json:"description"`
	Address       string  `gorm:"type:varchar(500)" json:"address"`
	Longitude     float64 `gorm:"type:decimal(10,7)" json:"longitude"`
	Latitude      float64 `gorm:"type:decimal(10,7)" json:"latitude"`
	ContactNumber string  `gorm:"type:varchar(255);unique;not null" json:"contact_number"`
	OwnerID       *uint   `gorm:"type:bigint;unsigned" json:"owner_id"`

	// Relationship 0 to 1
	Owner   *Owner `gorm:"foreignKey:OwnerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"owner"`
	MediaID *uint  `gorm:"type:bigint;unsigned" json:"media_id"`

	// Relationship 0 to 1
	Media           *Media `gorm:"foreignKey:MediaID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"media"`
	IsAdminVerified bool   `gorm:"default:false" json:"is_admin_verified"`

	// Relationship 0 to many
	Branches []*Branch `gorm:"foreignKey:CompanyID" json:"branches"`
}

type CompanyResource struct {
	Name            string            `json:"name"`
	Description     string            `json:"description"`
	Address         string            `json:"address"`
	Longitude       float64           `json:"longitude"`
	Latitude        float64           `json:"latitude"`
	ContactNumber   string            `json:"contactNumber"`
	OwnerID         *uint             `json:"ownerID"`
	Owner           *OwnerResource    `json:"owner"`
	MediaID         *uint             `json:"mediaID"`
	Media           *MediaResource    `json:"media"`
	IsAdminVerified bool              `json:"isAdminVerified"`
	Branches        []*BranchResource `json:"branches"`
}

type CompanyModel struct {
	lc     *fx.Lifecycle
	db     *gorm.DB
	logger *zap.Logger
}

func NewCompanyModel(
	lc *fx.Lifecycle,
	db *gorm.DB,
	logger *zap.Logger,
) *CompanyModel {
	return &CompanyModel{
		lc:     lc,
		db:     db,
		logger: logger,
	}
}

func (cm *CompanyModel) SeedDatabase() {
}

func (cm *CompanyModel) ToResource() {
}

func (cm *CompanyModel) ToResourceList() {
}
