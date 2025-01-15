package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MemberAddress struct {
	ID        uuid.UUID      `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	MembersProfileID uuid.UUID      `gorm:"unsigned" json:"members_profile_id"`
	PostalCode       string         `gorm:"type:varchar(20)" json:"postal_code"`
	Province         string         `gorm:"type:varchar(255)" json:"province"`
	City             string         `gorm:"type:varchar(255)" json:"city"`
	Barangay         string         `gorm:"type:varchar(255)" json:"barangay"`
	Region           string         `gorm:"type:varchar(255)" json:"region"`
	Label            string         `gorm:"type:enum('work', 'home', 'province', 'business');default:'home'" json:"label"`
	MembersProfile   *MemberProfile `gorm:"foreignKey:MembersProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members_profile"`
}

func (v *MemberAddress) BeforeCreate(tx *gorm.DB) (err error) {
	if v.ID == uuid.Nil {
		v.ID = uuid.New()
	}
	return
}

type MemberAddressResource struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`

	MembersProfileID uuid.UUID              `json:"membersProfileID"`
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

		ID:        address.ID,
		CreatedAt: address.CreatedAt.Format(time.RFC3339),
		UpdatedAt: address.UpdatedAt.Format(time.RFC3339),
		DeletedAt: address.DeletedAt.Time.Format(time.RFC3339),

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

func (m *ModelRepository) MemberAddressGetByID(id string, preloads ...string) (*MemberAddress, error) {
	repo := NewGenericRepository[MemberAddress](m.db.Client)
	return repo.GetByID(id, preloads...)
}

func (m *ModelRepository) MemberAddressCreate(memberaddress *MemberAddress, preloads ...string) (*MemberAddress, error) {
	repo := NewGenericRepository[MemberAddress](m.db.Client)
	return repo.Create(memberaddress, preloads...)
}
func (m *ModelRepository) MemberAddressUpdate(memberaddress *MemberAddress, preloads ...string) (*MemberAddress, error) {
	repo := NewGenericRepository[MemberAddress](m.db.Client)
	return repo.Update(memberaddress, preloads...)
}
func (m *ModelRepository) MemberAddressUpdateByID(id string, column string, value interface{}, preloads ...string) (*MemberAddress, error) {
	repo := NewGenericRepository[MemberAddress](m.db.Client)
	return repo.UpdateByID(id, column, value, preloads...)
}
func (m *ModelRepository) MemberAddressDeleteByID(id string) error {
	repo := NewGenericRepository[MemberAddress](m.db.Client)
	return repo.DeleteByID(id)
}
func (m *ModelRepository) MemberAddressGetAll(preloads ...string) ([]*MemberAddress, error) {
	repo := NewGenericRepository[MemberAddress](m.db.Client)
	return repo.GetAll(preloads...)
}
