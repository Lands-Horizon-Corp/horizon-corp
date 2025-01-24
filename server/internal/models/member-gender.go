package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MemberGender struct {
	ID        uuid.UUID      `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name        string                 `gorm:"size:255;unsigned"`
	Description string                 `gorm:"size:500"`
	History     []*MemberGenderHistory `gorm:"foreignKey:MemberProfileID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"history,omitempty"`
	CompanyID   uuid.UUID              `gorm:"unsigned" json:"company_id"`
	Company     *Company               `gorm:"foreignKey:CompanyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"company"`
}

func (v *MemberGender) BeforeCreate(tx *gorm.DB) (err error) {
	if v.ID == uuid.Nil {
		v.ID = uuid.New()
	}
	return
}

type MemberGenderResource struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`

	Name        string                         `json:"name"`
	Description string                         `json:"description"`
	History     []*MemberGenderHistoryResource `json:"history,omitempty"`
}

func (m *ModelTransformer) MemberGenderToResource(gender *MemberGender) *MemberGenderResource {
	if gender == nil {
		return nil
	}

	return &MemberGenderResource{
		ID:        gender.ID,
		CreatedAt: gender.CreatedAt.Format(time.RFC3339),
		UpdatedAt: gender.UpdatedAt.Format(time.RFC3339),
		DeletedAt: gender.DeletedAt.Time.Format(time.RFC3339),

		Name:        gender.Name,
		Description: gender.Description,
		History:     m.MemberGenderHistoryToResourceList(gender.History),
	}
}

func (m *ModelTransformer) MemberGenderToResourceList(genderList []*MemberGender) []*MemberGenderResource {
	if genderList == nil {
		return nil
	}

	var genderResources []*MemberGenderResource
	for _, gender := range genderList {
		genderResources = append(genderResources, m.MemberGenderToResource(gender))
	}
	return genderResources
}

func (m *ModelRepository) MemberGenderGetByID(id string, preloads ...string) (*MemberGender, error) {
	repo := NewGenericRepository[MemberGender](m.db.Client)
	return repo.GetByID(id, preloads...)
}

func (m *ModelRepository) MemberGenderCreate(membergender *MemberGender, preloads ...string) (*MemberGender, error) {
	repo := NewGenericRepository[MemberGender](m.db.Client)
	return repo.Create(membergender, preloads...)
}
func (m *ModelRepository) MemberGenderUpdate(membergender *MemberGender, preloads ...string) (*MemberGender, error) {
	repo := NewGenericRepository[MemberGender](m.db.Client)
	return repo.Update(membergender, preloads...)
}
func (m *ModelRepository) MemberGenderUpdateByID(id string, value *MemberGender, preloads ...string) (*MemberGender, error) {
	repo := NewGenericRepository[MemberGender](m.db.Client)
	return repo.UpdateByID(id, value, preloads...)
}
func (m *ModelRepository) MemberGenderDeleteByID(id string) error {
	repo := NewGenericRepository[MemberGender](m.db.Client)
	return repo.DeleteByID(id)
}
func (m *ModelRepository) MemberGenderGetAll(preloads ...string) ([]*MemberGender, error) {
	repo := NewGenericRepository[MemberGender](m.db.Client)
	return repo.GetAll(preloads...)
}
