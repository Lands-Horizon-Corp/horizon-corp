package models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model

	Name        string `gorm:"type:varchar(255);unique;not null" json:"name"`
	Description string `gorm:"type:text" json:"description"`
	ApiKey      string `gorm:"type:varchar(255);unique;not null" json:"api_key"`
	Color       string `gorm:"type:varchar(255)" json:"color"`

	ReadRole   bool `gorm:"default:false" json:"read_role"`
	WriteRole  bool `gorm:"default:false" json:"write_role"`
	UpdateRole bool `gorm:"default:false" json:"update_role"`
	DeleteRole bool `gorm:"default:false" json:"delete_role"`

	ReadErrorDetails   bool `gorm:"default:false" json:"read_error_details"`
	WriteErrorDetails  bool `gorm:"default:false" json:"write_error_details"`
	UpdateErrorDetails bool `gorm:"default:false" json:"update_error_details"`
	DeleteErrorDetails bool `gorm:"default:false" json:"delete_error_details"`

	ReadGender   bool `gorm:"default:false" json:"read_gender"`
	WriteGender  bool `gorm:"default:false" json:"write_gender"`
	UpdateGender bool `gorm:"default:false" json:"update_gender"`
	DeleteGender bool `gorm:"default:false" json:"delete_gender"`

	// Relationship 0 to many
	Admins    []*Admin    `gorm:"foreignKey:RoleID" json:"admins"`
	Owners    []*Owner    `gorm:"foreignKey:RoleID" json:"owners"`
	Employees []*Employee `gorm:"foreignKey:RoleID" json:"employees"`
	Members   []*Member   `gorm:"foreignKey:RoleID" json:"members"`
}

type RoleResource struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	Name               string `json:"name"`
	Description        string `json:"description"`
	ApiKey             string `json:"apiKey"`
	Color              string `json:"color"`
	ReadRole           bool   `json:"readRole"`
	WriteRole          bool   `json:"writeRole"`
	UpdateRole         bool   `json:"updateRole"`
	DeleteRole         bool   `json:"deleteRole"`
	ReadErrorDetails   bool   `json:"readErrorDetails"`
	WriteErrorDetails  bool   `json:"writeErrorDetails"`
	UpdateErrorDetails bool   `json:"updateErrorDetails"`
	DeleteErrorDetails bool   `json:"deleteErrorDetails"`
	ReadGender         bool   `json:"readGender"`
	WriteGender        bool   `json:"writeGender"`
	UpdateGender       bool   `json:"updateGender"`
	DeleteGender       bool   `json:"deleteGender"`

	Admins    []*AdminResource    `json:"admins"`
	Owners    []*OwnerResource    `json:"owners"`
	Employees []*EmployeeResource `json:"employees"`
	Members   []*MemberResource   `json:"members"`
}
