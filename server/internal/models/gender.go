package models

import "gorm.io/gorm"

type Gender struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255);unique;not null" json:"name"`
	Description string `gorm:"type:text" json:"description"`

	Employees []Employee `gorm:"foreignKey:MediaID" json:"employees"`
	Members   []Member   `gorm:"foreignKey:MediaID" json:"members"`
	Owners    []Owner    `gorm:"foreignKey:MediaID" json:"owners"`
	Admins    []Admin    `gorm:"foreignKey:MediaID" json:"admins"`
}
