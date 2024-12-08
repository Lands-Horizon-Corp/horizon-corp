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
	lc     *fx.Lifecycle
	db     *gorm.DB
	logger *zap.Logger
}

func NewGenderModel(
	lc *fx.Lifecycle,
	db *gorm.DB,
	logger *zap.Logger,
) *GenderModel {
	return &GenderModel{
		lc:     lc,
		db:     db,
		logger: logger,
	}
}

func (gm *GenderModel) SeedDatabase() {
}

func (gm *GenderModel) ToResource() {
}

func (gm *GenderModel) ToResourceList() {
}
