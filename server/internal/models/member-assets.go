package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MemberAssets struct {
	ID        uuid.UUID      `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	MembersProfileID uuid.UUID      `gorm:"unsigned" json:"members_profile_id"`
	EntryDate        time.Time      `gorm:"type:date;unsigned" json:"entry_date"`
	Description      string         `gorm:"type:text" json:"description"`
	Name             string         `gorm:"type:varchar(255);unsigned" json:"name"`
	MembersProfile   *MemberProfile `gorm:"foreignKey:MembersProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members_profile"`
}

func (v *MemberAssets) BeforeCreate(tx *gorm.DB) (err error) {
	if v.ID == uuid.Nil {
		v.ID = uuid.New()
	}
	return
}

type MemberAssetsResource struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`

	MembersProfileID uuid.UUID              `json:"membersProfileID"`
	EntryDate        string                 `json:"entryDate"`
	Description      string                 `json:"description"`
	Name             string                 `json:"name"`
	MembersProfile   *MemberProfileResource `json:"membersProfile,omitempty"`
}

func (m *ModelTransformer) MemberAssetsToResource(asset *MemberAssets) *MemberAssetsResource {
	if asset == nil {
		return nil
	}

	return &MemberAssetsResource{
		ID:        asset.ID,
		CreatedAt: asset.CreatedAt.Format(time.RFC3339),
		UpdatedAt: asset.UpdatedAt.Format(time.RFC3339),
		DeletedAt: asset.DeletedAt.Time.Format(time.RFC3339),

		MembersProfileID: asset.MembersProfileID,
		EntryDate:        asset.EntryDate.Format("2006-01-02"),
		Description:      asset.Description,
		Name:             asset.Name,
		MembersProfile:   m.MemberProfileToResource(asset.MembersProfile),
	}
}

func (m *ModelTransformer) MemberAssetsToResourceList(assetList []*MemberAssets) []*MemberAssetsResource {
	if assetList == nil {
		return nil
	}

	var assetResources []*MemberAssetsResource
	for _, asset := range assetList {
		assetResources = append(assetResources, m.MemberAssetsToResource(asset))
	}
	return assetResources
}

func (m *ModelRepository) MemberAssetsGetByID(id string, preloads ...string) (*MemberAssets, error) {
	repo := NewGenericRepository[MemberAssets](m.db.Client)
	return repo.GetByID(id, preloads...)
}

func (m *ModelRepository) MemberAssetsCreate(memberassets *MemberAssets, preloads ...string) (*MemberAssets, error) {
	repo := NewGenericRepository[MemberAssets](m.db.Client)
	return repo.Create(memberassets, preloads...)
}
func (m *ModelRepository) MemberAssetsUpdate(memberassets *MemberAssets, preloads ...string) (*MemberAssets, error) {
	repo := NewGenericRepository[MemberAssets](m.db.Client)
	return repo.Update(memberassets, preloads...)
}
func (m *ModelRepository) MemberAssetsUpdateByID(id string, column string, value *MemberAssets, preloads ...string) (*MemberAssets, error) {
	repo := NewGenericRepository[MemberAssets](m.db.Client)
	return repo.UpdateByID(id, value, preloads...)
}
func (m *ModelRepository) MemberAssetsDeleteByID(id string) error {
	repo := NewGenericRepository[MemberAssets](m.db.Client)
	return repo.DeleteByID(id)
}
func (m *ModelRepository) MemberAssetsGetAll(preloads ...string) ([]*MemberAssets, error) {
	repo := NewGenericRepository[MemberAssets](m.db.Client)
	return repo.GetAll(preloads...)
}
