package service

import (
	"horizon-core/internal/models"
	"horizon-core/internal/repository"

	"gorm.io/gorm"
)

type MemberRepository struct {
	*repository.ModelRepository[models.Member]
}

func NewMemberRepository(db *gorm.DB) *MemberRepository {
	return &MemberRepository{
		ModelRepository: repository.NewModelRepository[models.Member](db),
	}
}
