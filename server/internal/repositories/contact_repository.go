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

func (r *ContactsRepository) Create(contact *models.Contact) error {
	return r.DB.Create(contact).Error
}

func (r *ContactsRepository) GetAll() ([]models.Contact, error) {
	var contact []models.Contact
	err := r.DB.Find(&contact).Error
	return contact, err
}

func (r *ContactsRepository) GetByID(id uint) (models.Contact, error) {
	var contact models.Contact
	err := r.DB.First(&contact, id).Error
	return contact, err
}

func (r *ContactsRepository) Update(contact *models.Contact) error {
	return r.DB.Save(contact).Error
}

func (r *ContactsRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Contact{}, id).Error
}
