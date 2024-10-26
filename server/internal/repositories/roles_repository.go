package repositories

import (
	"horizon/server/internal/models"

	"gorm.io/gorm"
)

type RoleRepository struct {
	*Repository[models.Role]
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{
		Repository: NewRepository[models.Role](db),
	}
}
