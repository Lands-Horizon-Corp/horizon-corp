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
	err := r.DB.Create(contact).Error
	return handleDBError(err)
}

func (r *ContactsRepository) GetAll() ([]models.Contact, error) {
	var contact []models.Contact
	err := r.DB.Find(&contact).Error
	return contact, handleDBError(err)
}

func (r *ContactsRepository) GetByID(id uint) (models.Contact, error) {
	var contact models.Contact
	err := r.DB.First(&contact, id).Error
	return contact, handleDBError(err)
}

func (r *ContactsRepository) Update(id uint, contact *models.Contact) error {
	contact.ID = id
	err := r.DB.Save(contact).Error
	return handleDBError(err)
}

func (r *ContactsRepository) Delete(id uint) error {
	err := r.DB.Delete(&models.Contact{}, id).Error
	return handleDBError(err)
}
