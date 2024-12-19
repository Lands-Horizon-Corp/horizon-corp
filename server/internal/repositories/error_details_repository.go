package repositories

import (
	"horizon/server/internal/models"

	"gorm.io/gorm"
)

type ErrorDetailRepository struct {
	*Repository[models.ErrorDetail]
}

func NewErrorDetailRepository(db *gorm.DB) *ErrorDetailRepository {
	return &ErrorDetailRepository{
		Repository: NewRepository[models.ErrorDetail](db),
	}
}
