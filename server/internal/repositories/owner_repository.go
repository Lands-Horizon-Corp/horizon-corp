package repositories

import (
	"horizon/server/internal/models"
	"regexp"

	"gorm.io/gorm"
)

type OwnerRepository struct {
	DB *gorm.DB
}

func NewOwnerRepository(db *gorm.DB) *OwnerRepository {
	return &OwnerRepository{DB: db}
}

func (r *OwnerRepository) Create(owner *models.Owner) error {
	err := r.DB.Create(owner).Error
	return handleDBError(err)
}

func (r *OwnerRepository) GetAll() ([]models.Owner, error) {
	var owners []models.Owner
	err := r.DB.Preload("Media").Find(&owners).Error
	return owners, handleDBError(err)
}

func (r *OwnerRepository) GetByID(id uint) (models.Owner, error) {
	var owner models.Owner
	err := r.DB.Preload("Media").First(&owner, id).Error
	return owner, handleDBError(err)
}

func (r *OwnerRepository) Update(id uint, owner *models.Owner) error {
	owner.ID = id
	err := r.DB.Save(owner).Error
	return handleDBError(err)
}

func (r *OwnerRepository) Delete(id uint) error {
	err := r.DB.Delete(&models.Owner{}, id).Error
	return handleDBError(err)
}

func (r *OwnerRepository) FindByEmailUsernameOrContact(input string) (models.Owner, error) {
	var owner models.Owner

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	phoneRegex := `^\+?[1-9]\d{1,14}$`

	isEmail, _ := regexp.MatchString(emailRegex, input)
	isPhone, _ := regexp.MatchString(phoneRegex, input)

	if isEmail {
		err := r.DB.Preload("Media").Where("email = ?", input).First(&owner).Error
		return owner, handleDBError(err)
	} else if isPhone {
		err := r.DB.Preload("Media").Where("contact_number = ?", input).First(&owner).Error
		return owner, handleDBError(err)
	} else {
		err := r.DB.Preload("Media").Where("username = ?", input).First(&owner).Error
		return owner, handleDBError(err)
	}
}

func (r *OwnerRepository) UpdateColumns(id uint, columns map[string]interface{}) (models.Owner, error) {
	var owner models.Owner
	if err := r.DB.Model(&owner).Where("id = ?", id).Updates(columns).Error; err != nil {
		return owner, handleDBError(err)
	}
	if err := r.DB.Preload("Media").First(&owner, id).Error; err != nil {
		return owner, handleDBError(err)
	}
	return owner, nil
}
