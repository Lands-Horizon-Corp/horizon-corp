package models

import (
	"time"

	"gorm.io/gorm"
)

type MemberContactNumberReferences struct {
	gorm.Model
	Name          string `gorm:"type:varchar(255);unsigned" json:"name"`
	Description   string `gorm:"type:text" json:"description"`
	ContactNumber string `gorm:"type:varchar(255);unsigned" json:"contact_number"`

	MembersProfileID uint `gorm:"unsigned" json:"members_profile_id"`
}

type MemberContactNumberReferencesResource struct {
	ID            uint   `json:"id"`
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	ContactNumber string `json:"contactNumber"`
}

func (m *ModelTransformer) MemberContactNumberReferencesToResource(reference *MemberContactNumberReferences) *MemberContactNumberReferencesResource {
	if reference == nil {
		return nil
	}

	return &MemberContactNumberReferencesResource{
		ID:            reference.ID,
		CreatedAt:     reference.CreatedAt.Format(time.RFC3339),
		UpdatedAt:     reference.UpdatedAt.Format(time.RFC3339),
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
