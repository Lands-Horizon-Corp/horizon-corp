package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MemberGroup struct {
	ID        uuid.UUID      `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name        string                `gorm:"size:255;unsigned"`
	Description string                `gorm:"size:500"`
	History     []*MemberGroupHistory `gorm:"foreignKey:MemberGroupID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"history,omitempty"`
}

func (v *MemberGroup) BeforeCreate(tx *gorm.DB) (err error) {
	if v.ID == uuid.Nil {
		v.ID = uuid.New()
	}
	return
}

type MemberGroupResource struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`

	Name        string                        `json:"name"`
	Description string                        `json:"description"`
	History     []*MemberGroupHistoryResource `json:"history,omitempty"`
}

func (m *ModelTransformer) MemberGroupToResource(group *MemberGroup) *MemberGroupResource {
	if group == nil {
		return nil
	}

	return &MemberGroupResource{

		ID:        group.ID,
		CreatedAt: group.CreatedAt.Format(time.RFC3339),
		UpdatedAt: group.UpdatedAt.Format(time.RFC3339),
		DeletedAt: group.DeletedAt.Time.Format(time.RFC3339),

		Name:        group.Name,
		Description: group.Description,
		History:     m.MemberGroupHistoryToResourceList(group.History),
	}
}

func (m *ModelTransformer) MemberGroupToResourceList(groupList []*MemberGroup) []*MemberGroupResource {
	if groupList == nil {
		return nil
	}

	var groupResources []*MemberGroupResource
	for _, group := range groupList {
		groupResources = append(groupResources, m.MemberGroupToResource(group))
	}
	return groupResources
}

func (m *ModelRepository) MemberGroupGetByID(id string, preloads ...string) (*MemberGroup, error) {
	repo := NewGenericRepository[MemberGroup](m.db.Client)
	return repo.GetByID(id, preloads...)
}

func (m *ModelRepository) MemberGroupCreate(membergroup *MemberGroup, preloads ...string) (*MemberGroup, error) {
	repo := NewGenericRepository[MemberGroup](m.db.Client)
	return repo.Create(membergroup, preloads...)
}
func (m *ModelRepository) MemberGroupUpdate(membergroup *MemberGroup, preloads ...string) (*MemberGroup, error) {
	repo := NewGenericRepository[MemberGroup](m.db.Client)
	return repo.Update(membergroup, preloads...)
}
func (m *ModelRepository) MemberGroupUpdateByID(id string, column string, value interface{}, preloads ...string) (*MemberGroup, error) {
	repo := NewGenericRepository[MemberGroup](m.db.Client)
	return repo.UpdateByID(id, column, value, preloads...)
}
func (m *ModelRepository) MemberGroupDeleteByID(id string) error {
	repo := NewGenericRepository[MemberGroup](m.db.Client)
	return repo.DeleteByID(id)
}
func (m *ModelRepository) MemberGroupGetAll(preloads ...string) ([]*MemberGroup, error) {
	repo := NewGenericRepository[MemberGroup](m.db.Client)
	return repo.GetAll(preloads...)
}
