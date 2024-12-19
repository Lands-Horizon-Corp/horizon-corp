package repositories

import (
	"horizon/server/config"
	"horizon/server/helpers"
	"horizon/server/internal/models"

	"gorm.io/gorm"
)

type MemberRepository struct {
	*Repository[models.Member]
}

func NewMemberRepository(db *gorm.DB) *MemberRepository {
	return &MemberRepository{
		Repository: NewRepository[models.Member](db),
	}
}

func (r *MemberRepository) GetByEmail(email string) (*models.Member, error) {
	var member models.Member
	err := r.DB.Preload("Media").Where("email = ?", email).First(&member).Error
	return &member, handleDBError(err)
}

func (r *MemberRepository) GetByContactNumber(contactNumber string) (*models.Member, error) {
	var member models.Member
	err := r.DB.Preload("Media").Where("contact_number = ?", contactNumber).First(&member).Error
	return &member, handleDBError(err)
}

func (r *MemberRepository) GetByUsername(username string) (*models.Member, error) {
	var member models.Member
	err := r.DB.Preload("Media").Where("username = ?", username).First(&member).Error
	return &member, handleDBError(err)
}

func (r *MemberRepository) FindByEmailUsernameOrContact(input string) (*models.Member, error) {
	switch helpers.GetKeyType(input) {
	case "contact":
		return r.GetByContactNumber(input)
	case "email":
		return r.GetByEmail(input)
	default:
		return r.GetByUsername(input)
	}
}

func (r *MemberRepository) UpdatePassword(id uint, password string) error {
	newPassword, err := config.HashPassword(password)
	if err != nil {
		return err
	}
	updated := &models.Member{Password: newPassword}
	_, err = r.Repository.UpdateColumns(id, *updated, []string{})
	return err
}
