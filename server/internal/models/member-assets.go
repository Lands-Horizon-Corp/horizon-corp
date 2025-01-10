package models

import (
	"time"

	"gorm.io/gorm"
)

type MemberAssets struct {
	gorm.Model
	MembersProfileID uint           `gorm:"unsigned" json:"members_profile_id"`
	EntryDate        time.Time      `gorm:"type:date;unsigned" json:"entry_date"`
	Description      string         `gorm:"type:text" json:"description"`
	Name             string         `gorm:"type:varchar(255);unsigned" json:"name"`
	MembersProfile   *MemberProfile `gorm:"foreignKey:MembersProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members_profile"`
}

type MemberAssetsResource struct {
	ID               uint                   `json:"id"`
	CreatedAt        string                 `json:"createdAt"`
	UpdatedAt        string                 `json:"updatedAt"`
	MembersProfileID uint                   `json:"membersProfileID"`
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
		ID:               asset.ID,
		CreatedAt:        asset.CreatedAt.Format(time.RFC3339),
		UpdatedAt:        asset.UpdatedAt.Format(time.RFC3339),
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
