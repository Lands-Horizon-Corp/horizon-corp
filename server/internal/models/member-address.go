package models

import (
	"time"

	"gorm.io/gorm"
)

type MemberAddress struct {
	gorm.Model
	MembersProfileID uint           `gorm:"unsigned" json:"members_profile_id"`
	PostalCode       string         `gorm:"type:varchar(20)" json:"postal_code"`
	Province         string         `gorm:"type:varchar(255)" json:"province"`
	City             string         `gorm:"type:varchar(255)" json:"city"`
	Barangay         string         `gorm:"type:varchar(255)" json:"barangay"`
	Region           string         `gorm:"type:varchar(255)" json:"region"`
	Label            string         `gorm:"type:enum('work', 'home', 'province', 'business');default:'home'" json:"label"`
	MembersProfile   *MemberProfile `gorm:"foreignKey:MembersProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members_profile"`
}

type MemberAddressResource struct {
	ID               uint                   `json:"id"`
	CreatedAt        string                 `json:"createdAt"`
	UpdatedAt        string                 `json:"updatedAt"`
	MembersProfileID uint                   `json:"membersProfileID"`
	PostalCode       string                 `json:"postalCode"`
	Province         string                 `json:"province"`
	City             string                 `json:"city"`
	Barangay         string                 `json:"barangay"`
	Region           string                 `json:"region"`
	Label            string                 `json:"label"`
	MembersProfile   *MemberProfileResource `json:"membersProfile,omitempty"`
}

func (m *ModelTransformer) MemberAddressToResource(address *MemberAddress) *MemberAddressResource {
	if address == nil {
		return nil
	}

	return &MemberAddressResource{
		ID:               address.ID,
		CreatedAt:        address.CreatedAt.Format(time.RFC3339),
		UpdatedAt:        address.UpdatedAt.Format(time.RFC3339),
		MembersProfileID: address.MembersProfileID,
		PostalCode:       address.PostalCode,
		Province:         address.Province,
		City:             address.City,
		Barangay:         address.Barangay,
		Region:           address.Region,
		Label:            address.Label,
		MembersProfile:   m.MemberProfileToResource(address.MembersProfile),
	}
}

func (m *ModelTransformer) MemberAddressToResourceList(addressList []*MemberAddress) []*MemberAddressResource {
	if addressList == nil {
		return nil
	}

	var addressResources []*MemberAddressResource
	for _, address := range addressList {
		addressResources = append(addressResources, m.MemberAddressToResource(address))
	}
	return addressResources
}
