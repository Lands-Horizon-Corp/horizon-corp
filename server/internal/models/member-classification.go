package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MemberClassification struct {
	ID        uuid.UUID      `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name        string                         `gorm:"size:255;unsigned"`
	Description string                         `gorm:"size:500"`
	History     []*MemberClassificationHistory `gorm:"foreignKey:MemberClassificationID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"history,omitempty"`

	CompanyID uuid.UUID `gorm:"unsigned" json:"company_id"`
	Company   *Company  `gorm:"foreignKey:CompanyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"company"`
}

func (v *MemberClassification) BeforeCreate(tx *gorm.DB) (err error) {
	if v.ID == uuid.Nil {
		v.ID = uuid.New()
	}
	return
}

type MemberClassificationResource struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`

	Name        string                                 `json:"name"`
	Description string                                 `json:"description"`
	History     []*MemberClassificationHistoryResource `json:"history,omitempty"`
}

func (m *ModelTransformer) MemberClassificationToResource(classification *MemberClassification) *MemberClassificationResource {
	if classification == nil {
		return nil
	}

	return &MemberClassificationResource{

		ID:        classification.ID,
		CreatedAt: classification.CreatedAt.Format(time.RFC3339),
		UpdatedAt: classification.UpdatedAt.Format(time.RFC3339),
		DeletedAt: classification.DeletedAt.Time.Format(time.RFC3339),

		Name:        classification.Name,
		Description: classification.Description,
		History:     m.MemberClassificationHistoryToResourceList(classification.History),
	}
}

func (m *ModelTransformer) MemberClassificationToResourceList(classificationList []*MemberClassification) []*MemberClassificationResource {
	if classificationList == nil {
		return nil
	}

	var classificationResources []*MemberClassificationResource
	for _, classification := range classificationList {
		classificationResources = append(classificationResources, m.MemberClassificationToResource(classification))
	}
	return classificationResources
}

func (m *ModelRepository) MemberClassificationGetByID(id string, preloads ...string) (*MemberClassification, error) {
	repo := NewGenericRepository[MemberClassification](m.db.Client)
	return repo.GetByID(id, preloads...)
}

func (m *ModelRepository) MemberClassificationCreate(memberclassification *MemberClassification, preloads ...string) (*MemberClassification, error) {
	repo := NewGenericRepository[MemberClassification](m.db.Client)
	return repo.Create(memberclassification, preloads...)
}
func (m *ModelRepository) MemberClassificationUpdate(memberclassification *MemberClassification, preloads ...string) (*MemberClassification, error) {
	repo := NewGenericRepository[MemberClassification](m.db.Client)
	return repo.Update(memberclassification, preloads...)
}
func (m *ModelRepository) MemberClassificationUpdateByID(id string, value *MemberClassification, preloads ...string) (*MemberClassification, error) {
	repo := NewGenericRepository[MemberClassification](m.db.Client)
	return repo.UpdateByID(id, value, preloads...)
}
func (m *ModelRepository) MemberClassificationDeleteByID(id string) error {
	repo := NewGenericRepository[MemberClassification](m.db.Client)
	return repo.DeleteByID(id)
}
func (m *ModelRepository) MemberClassificationGetAll(preloads ...string) ([]*MemberClassification, error) {
	repo := NewGenericRepository[MemberClassification](m.db.Client)
	return repo.GetAll(preloads...)
}
