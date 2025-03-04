package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MemberGovernmentBenefits struct {
	ID        uuid.UUID      `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	MembersProfileID uuid.UUID      `gorm:"unsigned" json:"members_profile_id"`
	Country          string         `gorm:"type:varchar(255);unsigned" json:"country"`
	Name             string         `gorm:"type:varchar(255);unsigned" json:"name"`
	Description      string         `gorm:"type:text" json:"description"`
	Value            float64        `gorm:"type:decimal(10,2)" json:"value"`
	FrontMediaID     *uuid.UUID     `gorm:"type:bigint;unsigned" json:"front_media_id"`
	BackMediaID      *uuid.UUID     `gorm:"type:bigint;unsigned" json:"back_media_id"`
	MembersProfile   *MemberProfile `gorm:"foreignKey:MembersProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members_profile"`
	FrontMedia       *Media         `gorm:"foreignKey:FrontMediaID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"front_media,omitempty"`
	BackMedia        *Media         `gorm:"foreignKey:BackMediaID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"back_media,omitempty"`
}

func (v *MemberGovernmentBenefits) BeforeCreate(tx *gorm.DB) (err error) {
	if v.ID == uuid.Nil {
		v.ID = uuid.New()
	}
	return
}

type MemberGovernmentBenefitsResource struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`

	MembersProfileID uuid.UUID              `json:"membersProfileID"`
	Country          string                 `json:"country"`
	Name             string                 `json:"name"`
	Description      string                 `json:"description"`
	Value            float64                `json:"value"`
	FrontMediaID     *uuid.UUID             `json:"frontMediaID,omitempty"`
	BackMediaID      *uuid.UUID             `json:"backMediaID,omitempty"`
	MembersProfile   *MemberProfileResource `json:"membersProfile,omitempty"`
	FrontMedia       *MediaResource         `json:"frontMedia,omitempty"`
	BackMedia        *MediaResource         `json:"backMedia,omitempty"`
}

func (m *ModelTransformer) MemberGovernmentBenefitsToResource(benefit *MemberGovernmentBenefits) *MemberGovernmentBenefitsResource {
	if benefit == nil {
		return nil
	}

	return &MemberGovernmentBenefitsResource{

		ID:        benefit.ID,
		CreatedAt: benefit.CreatedAt.Format(time.RFC3339),
		UpdatedAt: benefit.UpdatedAt.Format(time.RFC3339),
		DeletedAt: benefit.DeletedAt.Time.Format(time.RFC3339),

		MembersProfileID: benefit.MembersProfileID,
		Country:          benefit.Country,
		Name:             benefit.Name,
		Description:      benefit.Description,
		Value:            benefit.Value,
		FrontMediaID:     benefit.FrontMediaID,
		BackMediaID:      benefit.BackMediaID,
		MembersProfile:   m.MemberProfileToResource(benefit.MembersProfile),
		FrontMedia:       m.MediaToResource(benefit.FrontMedia),
		BackMedia:        m.MediaToResource(benefit.BackMedia),
	}
}

func (m *ModelTransformer) MemberGovernmentBenefitsToResourceList(benefitList []*MemberGovernmentBenefits) []*MemberGovernmentBenefitsResource {
	if benefitList == nil {
		return nil
	}

	var benefitResources []*MemberGovernmentBenefitsResource
	for _, benefit := range benefitList {
		benefitResources = append(benefitResources, m.MemberGovernmentBenefitsToResource(benefit))
	}
	return benefitResources
}

func (m *ModelRepository) MemberGovernmentBenefitsGetByID(id string, preloads ...string) (*MemberGovernmentBenefits, error) {
	repo := NewGenericRepository[MemberGovernmentBenefits](m.db.Client)
	return repo.GetByID(id, preloads...)
}

func (m *ModelRepository) MemberGovernmentBenefitsCreate(membergovernmentbenefits *MemberGovernmentBenefits, preloads ...string) (*MemberGovernmentBenefits, error) {
	repo := NewGenericRepository[MemberGovernmentBenefits](m.db.Client)
	return repo.Create(membergovernmentbenefits, preloads...)
}
func (m *ModelRepository) MemberGovernmentBenefitsUpdate(membergovernmentbenefits *MemberGovernmentBenefits, preloads ...string) (*MemberGovernmentBenefits, error) {
	repo := NewGenericRepository[MemberGovernmentBenefits](m.db.Client)
	return repo.Update(membergovernmentbenefits, preloads...)
}
func (m *ModelRepository) MemberGovernmentBenefitsUpdateByID(id string, value *MemberGovernmentBenefits, preloads ...string) (*MemberGovernmentBenefits, error) {
	repo := NewGenericRepository[MemberGovernmentBenefits](m.db.Client)
	return repo.UpdateByID(id, value, preloads...)
}
func (m *ModelRepository) MemberGovernmentBenefitsDeleteByID(id string) error {
	repo := NewGenericRepository[MemberGovernmentBenefits](m.db.Client)
	return repo.DeleteByID(id)
}
func (m *ModelRepository) MemberGovernmentBenefitsGetAll(preloads ...string) ([]*MemberGovernmentBenefits, error) {
	repo := NewGenericRepository[MemberGovernmentBenefits](m.db.Client)
	return repo.GetAll(preloads...)
}
