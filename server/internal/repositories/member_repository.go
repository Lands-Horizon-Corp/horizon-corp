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
	return r.DB.Create(member).Error
}

func (r *MemberRepository) GetAll() ([]models.Member, error) {
	var member []models.Member
	err := r.DB.Find(&member).Error
	return member, err
}

func (r *MemberRepository) GetByID(id uint) (models.Member, error) {
	var member models.Member
	err := r.DB.First(&member, id).Error
	return member, err
}

func (r *MemberRepository) Update(member *models.Member) error {
	return r.DB.Save(member).Error
}

func (r *MemberRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Member{}, id).Error
}
