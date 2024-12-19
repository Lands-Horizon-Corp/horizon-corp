package repositories

import (
	"horizon/server/internal/models"

	"gorm.io/gorm"
)

type ContactsRepository struct {
	*Repository[models.Contact]
}

func NewContactsRepository(db *gorm.DB) *ContactsRepository {
	return &ContactsRepository{
		Repository: NewRepository[models.Contact](db),
	}
}
