package repositories

import (
	"horizon/server/internal/models"

	"gorm.io/gorm"
)

type MemberRepository struct {
	DB *gorm.DB
}

func NewMemberRepository(db *gorm.DB) *MemberRepository {
	return &MemberRepository{DB: db}
}

func (r *MemberRepository) Create(member *models.Member) error {
	err := r.DB.Create(member).Error
	return handleDBError(err)
}

func (r *MemberRepository) GetAll() ([]models.Member, error) {
	var member []models.Member
	err := r.DB.Find(&member).Error
	return member, handleDBError(err)
}

func (r *MemberRepository) GetByID(id uint) (models.Member, error) {
	var member models.Member
	err := r.DB.First(&member, id).Error
	return member, handleDBError(err)
}

func (r *MemberRepository) Update(id uint, member *models.Member) error {
	member.ID = id
	err := r.DB.Save(member).Error
	return handleDBError(err)
}

func (r *MemberRepository) Delete(id uint) error {
	err := r.DB.Delete(&models.Member{}, id).Error
	return handleDBError(err)
}
