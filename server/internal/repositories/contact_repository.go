package repositories

import (
	"horizon/server/internal/models"

	"gorm.io/gorm"
)

type ContactsRepository struct {
	DB *gorm.DB
}

func NewContactsRepository(db *gorm.DB) *ContactsRepository {
	return &ContactsRepository{DB: db}
}

func (r *ContactsRepository) Create(contacts *models.Contacts) error {
	return r.DB.Create(contacts).Error
}

func (r *ContactsRepository) GetAll() ([]models.Contacts, error) {
	var contactss []models.Contacts
	err := r.DB.Find(&contactss).Error
	return contactss, err
}

func (r *ContactsRepository) GetByID(id uint) (models.Contacts, error) {
	var contacts models.Contacts
	err := r.DB.First(&contacts, id).Error
	return contacts, err
}

func (r *ContactsRepository) Update(contacts *models.Contacts) error {
	return r.DB.Save(contacts).Error
}

func (r *ContactsRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Contacts{}, id).Error
}
