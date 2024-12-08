package database

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ProvideDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func ProvideModels() []interface{} {
	return []interface{}{
		&models.Admin{},
		&models.Branch{},
		&models.Company{},
		&models.Contact{},
		&models.Employee{},
		&models.Feedback{},
		&models.Footstep{},
		&models.Gender{},
		&models.Media{},
		&models.Member{},
		&models.Owner{},
		&models.Role{},
		&models.Timesheet{},
	}
}
