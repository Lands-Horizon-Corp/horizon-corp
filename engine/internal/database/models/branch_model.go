package models

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Branch struct {
	gorm.Model

	// Fields
	Name            string  `gorm:"type:varchar(255);not null" json:"name"`
	Address         string  `gorm:"type:varchar(500)" json:"address"`
	Longitude       float64 `gorm:"type:decimal(10,7)" json:"longitude"`
	Latitude        float64 `gorm:"type:decimal(10,7)" json:"latitude"`
	Email           string  `gorm:"type:varchar(255);unique;not null" json:"email"`
	ContactNumber   string  `gorm:"type:varchar(15);unique;not null" json:"contact_number"`
	IsAdminVerified bool    `gorm:"default:false" json:"is_admin_verified"`

	// Relationship 0 to 1
	MediaID *uint  `gorm:"type:bigint;unsigned" json:"media_id"`
	Media   *Media `gorm:"foreignKey:MediaID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"media"`

	// Relationship 0 to 1
	CompanyID *uint    `gorm:"type:bigint;unsigned;not null" json:"company_id"`
	Company   *Company `gorm:"foreignKey:CompanyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"company"`

	// Relationship 0 to many
	Employees []*Employee `gorm:"foreignKey:BranchID" json:"employees"`

	// Relationship 0 to many
	Members []*Member `gorm:"foreignKey:BranchID" json:"members"`
}

type BranchResource struct {
	Name            string              `json:"name"`
	Address         string              `json:"address"`
	Longitude       float64             `json:"longitude"`
	Latitude        float64             `json:"latitude"`
	Email           string              `json:"email"`
	ContactNumber   string              `json:"contactNumber"`
	IsAdminVerified bool                `json:"isAdminVerified"`
	MediaID         *uint               `json:"mediaID"`
	Media           *MediaResource      `json:"media"`
	CompanyID       *uint               `json:"companyID"`
	Company         *CompanyResource    `json:"company"`
	Employees       []*EmployeeResource `json:"employees"`
	Members         []*MemberResource   `json:"members"`
}

type BranchModel struct {
	lc            *fx.Lifecycle
	db            *gorm.DB
	logger        *zap.Logger
	mediaModel    *MediaModel
	companyModel  *CompanyModel
	employeeModel *EmployeeModel
	memberModel   *MemberModel
}

func NewBranchModel(
	lc *fx.Lifecycle,
	db *gorm.DB,
	logger *zap.Logger,
	mediaModel *MediaModel,
	companyModel *CompanyModel,
	employeeModel *EmployeeModel,
	memberModel *MemberModel,
) *BranchModel {
	return &BranchModel{
		lc:            lc,
		db:            db,
		logger:        logger,
		mediaModel:    mediaModel,
		companyModel:  companyModel,
		employeeModel: employeeModel,
		memberModel:   memberModel,
	}
}

func (bm *BranchModel) SeedDatabase() {

}

func (bm *BranchModel) ToResource(branch *Branch) *BranchResource {
	if branch == nil {
		return nil
	}

	return &BranchResource{
		Name:            branch.Name,
		Address:         branch.Address,
		Longitude:       branch.Longitude,
		Latitude:        branch.Latitude,
		Email:           branch.Email,
		ContactNumber:   branch.ContactNumber,
		IsAdminVerified: branch.IsAdminVerified,
		MediaID:         branch.MediaID,
		Media:           bm.mediaModel.ToResource(branch.Media),
		CompanyID:       branch.CompanyID,
		Company:         bm.companyModel.ToResource(branch.Company),
		Employees:       bm.employeeModel.ToResourceList(branch.Employees),
		Members:         bm.memberModel.ToResourceList(branch.Members),
	}
}

func (bm *BranchModel) ToResourceList(branches []*Branch) []*BranchResource {
	if branches == nil {
		return nil
	}
	var branchResources []*BranchResource
	for _, branch := range branches {
		branchResources = append(branchResources, bm.ToResource(branch))
	}
	return branchResources
}
