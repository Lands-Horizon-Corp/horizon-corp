package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MemberIncome struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	MembersProfileID uuid.UUID      `gorm:"unsigned" json:"members_profile_id"`
	Name             string         `gorm:"type:varchar(255);unsigned" json:"name"`
	Amount           float64        `gorm:"type:decimal(10,2);unsigned" json:"amount"`
	Date             time.Time      `gorm:"type:date" json:"date"`
	Description      string         `gorm:"type:text" json:"description"`
	MembersProfile   *MemberProfile `gorm:"foreignKey:MembersProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members_profile"`
}

type MemberIncomeResource struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`

	MembersProfileID uuid.UUID              `json:"membersProfileID"`
	Name             string                 `json:"name"`
	Amount           float64                `json:"amount"`
	Date             string                 `json:"date"`
	Description      string                 `json:"description"`
	MembersProfile   *MemberProfileResource `json:"membersProfile,omitempty"`
}

func (m *ModelTransformer) MemberIncomeToResource(income *MemberIncome) *MemberIncomeResource {
	if income == nil {
		return nil
	}

	return &MemberIncomeResource{

		ID:        income.ID,
		CreatedAt: income.CreatedAt.Format(time.RFC3339),
		UpdatedAt: income.UpdatedAt.Format(time.RFC3339),
		DeletedAt: income.DeletedAt.Time.Format(time.RFC3339),

		MembersProfileID: income.MembersProfileID,
		Name:             income.Name,
		Amount:           income.Amount,
		Date:             income.Date.Format("2006-01-02"),
		Description:      income.Description,
		MembersProfile:   m.MemberProfileToResource(income.MembersProfile),
	}
}

func (m *ModelTransformer) MemberIncomeToResourceList(incomeList []*MemberIncome) []*MemberIncomeResource {
	if incomeList == nil {
		return nil
	}

	var incomeResources []*MemberIncomeResource
	for _, income := range incomeList {
		incomeResources = append(incomeResources, m.MemberIncomeToResource(income))
	}
	return incomeResources
}
