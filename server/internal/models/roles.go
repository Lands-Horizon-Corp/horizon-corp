package models

import "gorm.io/gorm"

type Roles struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"unique;not null"`
	Description string `json:"description,omitempty"`
	gorm.Model

	ApiKey string `json:"api_key,omitempty" gorm:"unique;not null"`

	// Role Module
	ReadRole   bool `json:"read_role,omitempty" gorm:"default:false"`
	WriteRole  bool `json:"write_role,omitempty" gorm:"default:false"`
	UpdateRole bool `json:"update_role,omitempty" gorm:"default:false"`
	DeleteRole bool `json:"delete_role,omitempty" gorm:"default:false"`

	// Error Details Module
	ReadErrorDetails   bool `json:"read_error_details,omitempty" gorm:"default:false"`
	WriteErrorDetails  bool `json:"write_error_details,omitempty" gorm:"default:false"`
	UpdateErrorDetails bool `json:"update_error_details,omitempty" gorm:"default:false"`
	DeleteErrorDetails bool `json:"delete_error_details,omitempty" gorm:"default:false"`

	// Gender Module
	ReadGender   bool `json:"read_gender,omitempty" gorm:"default:false"`
	WriteGender  bool `json:"write_gender,omitempty" gorm:"default:false"`
	UpdateGender bool `json:"update_gender,omitempty" gorm:"default:false"`
	DeleteGender bool `json:"delete_gender,omitempty" gorm:"default:false"`
}
