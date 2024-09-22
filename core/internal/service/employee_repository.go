package service

import (
	"horizon-core/internal/models"
	"horizon-core/internal/repository"

	"gorm.io/gorm"
)

type EmployeeRepository struct {
	*repository.ModelRepository[models.Employee]
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{
		ModelRepository: repository.NewModelRepository[models.Employee](db),
	}
}
