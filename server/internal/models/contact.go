package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Contact struct {
	ID        uuid.UUID      `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Fields
	FirstName     string `gorm:"type:varchar(255);not null" json:"first_name"`
	LastName      string `gorm:"type:varchar(255);not null" json:"last_name"`
	Email         string `gorm:"type:varchar(255);unique;not null" json:"email"`
	ContactNumber string `gorm:"type:varchar(15);not null" json:"contact_number"`
	Description   string `gorm:"type:text" json:"description"`
}

func (v *Contact) BeforeCreate(tx *gorm.DB) (err error) {
	if v.ID == uuid.Nil {
		v.ID = uuid.New()
	}
	return
}

type ContactResource struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`

	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Email         string `json:"email"`
	ContactNumber string `json:"contactNumber"`
	Description   string `json:"description"`
}

func (m *ModelTransformer) ContactToResource(contact *Contact) *ContactResource {
	if contact == nil {
		return nil
	}

	return &ContactResource{
		ID:        contact.ID,
		CreatedAt: contact.CreatedAt.Format(time.RFC3339),
		UpdatedAt: contact.UpdatedAt.Format(time.RFC3339),
		DeletedAt: contact.DeletedAt.Time.Format(time.RFC3339),

		FirstName:     contact.FirstName,
		LastName:      contact.LastName,
		Email:         contact.Email,
		ContactNumber: contact.ContactNumber,
		Description:   contact.Description,
	}
}

func (m *ModelTransformer) ContactToResourceList(contactList []*Contact) []*ContactResource {
	if contactList == nil {
		return nil
	}

	var contactResources []*ContactResource
	for _, contact := range contactList {
		contactResources = append(contactResources, m.ContactToResource(contact))
	}
	return contactResources
}

func (m *ModelRepository) ContactGetByID(id string, preloads ...string) (*Contact, error) {
	repo := NewGenericRepository[Contact](m.db.Client)
	return repo.GetByID(id, preloads...)
}

func (m *ModelRepository) ContactCreate(contact *Contact, preloads ...string) (*Contact, error) {
	repo := NewGenericRepository[Contact](m.db.Client)
	return repo.Create(contact, preloads...)
}
func (m *ModelRepository) ContactUpdate(contact *Contact, preloads ...string) (*Contact, error) {
	repo := NewGenericRepository[Contact](m.db.Client)
	return repo.Update(contact, preloads...)
}
func (m *ModelRepository) ContactUpdateByID(id string, column string, value interface{}, preloads ...string) (*Contact, error) {
	repo := NewGenericRepository[Contact](m.db.Client)
	return repo.UpdateByID(id, column, value, preloads...)
}
func (m *ModelRepository) ContactDeleteByID(id string) error {
	repo := NewGenericRepository[Contact](m.db.Client)
	return repo.DeleteByID(id)
}
func (m *ModelRepository) ContactGetAll(preloads ...string) ([]*Contact, error) {
	repo := NewGenericRepository[Contact](m.db.Client)
	return repo.GetAll(preloads...)
}
