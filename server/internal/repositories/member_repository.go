package repositories

import (
	"horizon/server/internal/models"
	"regexp"

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

func (r *MemberRepository) FindByEmailUsernameOrContact(input string) (models.Member, error) {
	var member models.Member

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	phoneRegex := `^\+?[1-9]\d{1,14}$`

	isEmail, _ := regexp.MatchString(emailRegex, input)
	isPhone, _ := regexp.MatchString(phoneRegex, input)

	if isEmail {
		err := r.DB.Where("email = ?", input).First(&member).Error
		return member, handleDBError(err)
	} else if isPhone {
		err := r.DB.Where("contact_number = ?", input).First(&member).Error
		return member, handleDBError(err)
	} else {
		err := r.DB.Where("username = ?", input).First(&member).Error
		return member, handleDBError(err)
	}
}

func (r *MemberRepository) UpdateColumns(id uint, columns map[string]interface{}) (models.Member, error) {
	var member models.Member
	if err := r.DB.Model(&member).Where("id = ?", id).Updates(columns).Error; err != nil {
		return member, handleDBError(err)
	}
	if err := r.DB.First(&member, id).Error; err != nil {
		return member, handleDBError(err)
	}
	return member, nil
}
