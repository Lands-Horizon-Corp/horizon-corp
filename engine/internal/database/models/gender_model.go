package models

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Gender struct {
	gorm.Model

	// Fields
	Name        string `gorm:"type:varchar(255);unique;not null" json:"name"`
	Description string `gorm:"type:text" json:"description"`

	// Relationship 0 to many
	Employees []*Employee `gorm:"foreignKey:GenderID" json:"employees"`
	Members   []*Member   `gorm:"foreignKey:GenderID" json:"members"`
	Owners    []*Owner    `gorm:"foreignKey:GenderID" json:"owners"`
	Admins    []*Admin    `gorm:"foreignKey:GenderID" json:"admins"`
}

type GenderResource struct {
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Employees   []*EmployeeResource `json:"employees"`
	Members     []*MemberResource   `json:"members"`
	Owners      []*OwnerResource    `json:"owners"`
	Admins      []*AdminResource    `json:"admins"`
}

type GenderModel struct {
	lc            *fx.Lifecycle
	db            *gorm.DB
	logger        *zap.Logger
	adminModel    *AdminModel
	employeeModel *EmployeeModel
	ownerModel    *OwnerModel
	memberModel   *MemberModel
}

func NewGenderModel(
	lc *fx.Lifecycle,
	db *gorm.DB,
	logger *zap.Logger,
	adminModel *AdminModel,
	employeeModel *EmployeeModel,
	ownerModel *OwnerModel,
	memberModel *MemberModel,
) *GenderModel {
	return &GenderModel{
		lc:            lc,
		db:            db,
		logger:        logger,
		adminModel:    adminModel,
		employeeModel: employeeModel,
		ownerModel:    ownerModel,
		memberModel:   memberModel,
	}
}

func (gm *GenderModel) SeedDatabase() {
}

func (gm *GenderModel) ToResource(gender *Gender) *GenderResource {
	if gender == nil {
		return nil
	}

	return &GenderResource{
		Name:        gender.Name,
		Description: gender.Description,
		Employees:   gm.employeeModel.ToResourceList(gender.Employees),
		Members:     gm.memberModel.ToResourceList(gender.Members),
		Owners:      gm.ownerModel.ToResourceList(gender.Owners),
		Admins:      gm.adminModel.ToResourceList(gender.Admins),
	}
}

func (gm *GenderModel) ToResourceList(genders []*Gender) []*GenderResource {
	if genders == nil {
		return nil
	}

	var genderResources []*GenderResource
	for _, gender := range genders {
		genderResources = append(genderResources, gm.ToResource(gender))
	}
	return genderResources
}
