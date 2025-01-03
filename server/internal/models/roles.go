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

	Admins    []*Admin    `gorm:"foreignKey:RoleID" json:"admins"`
	Employees []*Employee `gorm:"foreignKey:RoleID" json:"employees"`
	Members   []*Member   `gorm:"foreignKey:RoleID" json:"members"`
}
