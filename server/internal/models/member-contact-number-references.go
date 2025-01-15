package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MemberContactNumberReferences struct {
	ID        uuid.UUID      `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name          string `gorm:"type:varchar(255);unsigned" json:"name"`
	Description   string `gorm:"type:text" json:"description"`
	ContactNumber string `gorm:"type:varchar(255);unsigned" json:"contact_number"`

	MembersProfileID uuid.UUID `gorm:"unsigned" json:"members_profile_id"`
}

func (v *MemberContactNumberReferences) BeforeCreate(tx *gorm.DB) (err error) {
	if v.ID == uuid.Nil {
		v.ID = uuid.New()
	}
	return
}

type MemberContactNumberReferencesResource struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`

	Name          string `json:"name"`
	Description   string `json:"description"`
	ContactNumber string `json:"contactNumber"`
}

func (m *ModelTransformer) MemberContactNumberReferencesToResource(reference *MemberContactNumberReferences) *MemberContactNumberReferencesResource {
	if reference == nil {
		return nil
	}

	return &MemberContactNumberReferencesResource{

		ID:        reference.ID,
		CreatedAt: reference.CreatedAt.Format(time.RFC3339),
		UpdatedAt: reference.UpdatedAt.Format(time.RFC3339),
		DeletedAt: reference.DeletedAt.Time.Format(time.RFC3339),

		Name:          reference.Name,
		Description:   reference.Description,
		ContactNumber: reference.ContactNumber,
	}
}

func (m *ModelTransformer) MemberContactNumberReferencesToResourceList(referenceList []*MemberContactNumberReferences) []*MemberContactNumberReferencesResource {
	if referenceList == nil {
		return nil
	}

	var referenceResources []*MemberContactNumberReferencesResource
	for _, reference := range referenceList {
		referenceResources = append(referenceResources, m.MemberContactNumberReferencesToResource(reference))
	}
	return referenceResources
}

func (m *ModelRepository) MemberContactNumberReferencesGetByID(id string, preloads ...string) (*MemberContactNumberReferences, error) {
	repo := NewGenericRepository[MemberContactNumberReferences](m.db.Client)
	return repo.GetByID(id, preloads...)
}

func (m *ModelRepository) MemberContactNumberReferencesCreate(membercontactnumberreferences *MemberContactNumberReferences, preloads ...string) (*MemberContactNumberReferences, error) {
	repo := NewGenericRepository[MemberContactNumberReferences](m.db.Client)
	return repo.Create(membercontactnumberreferences, preloads...)
}
func (m *ModelRepository) MemberContactNumberReferencesUpdate(membercontactnumberreferences *MemberContactNumberReferences, preloads ...string) (*MemberContactNumberReferences, error) {
	repo := NewGenericRepository[MemberContactNumberReferences](m.db.Client)
	return repo.Update(membercontactnumberreferences, preloads...)
}
func (m *ModelRepository) MemberContactNumberReferencesUpdateByID(id string, column string, value interface{}, preloads ...string) (*MemberContactNumberReferences, error) {
	repo := NewGenericRepository[MemberContactNumberReferences](m.db.Client)
	return repo.UpdateByID(id, column, value, preloads...)
}
func (m *ModelRepository) MemberContactNumberReferencesDeleteByID(id string) error {
	repo := NewGenericRepository[MemberContactNumberReferences](m.db.Client)
	return repo.DeleteByID(id)
}
func (m *ModelRepository) MemberContactNumberReferencesGetAll(preloads ...string) ([]*MemberContactNumberReferences, error) {
	repo := NewGenericRepository[MemberContactNumberReferences](m.db.Client)
	return repo.GetAll(preloads...)
}
