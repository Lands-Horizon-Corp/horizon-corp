package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MemberIncome struct {
	ID        uuid.UUID      `gorm:"type:char(36);primary_key"`
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

func (v *MemberIncome) BeforeCreate(tx *gorm.DB) (err error) {
	if v.ID == uuid.Nil {
		v.ID = uuid.New()
	}
	return
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

func (m *ModelRepository) MemberIncomeGetByID(id string, preloads ...string) (*MemberIncome, error) {
	repo := NewGenericRepository[MemberIncome](m.db.Client)
	return repo.GetByID(id, preloads...)
}

func (m *ModelRepository) MemberIncomeCreate(memberincome *MemberIncome, preloads ...string) (*MemberIncome, error) {
	repo := NewGenericRepository[MemberIncome](m.db.Client)
	return repo.Create(memberincome, preloads...)
}
func (m *ModelRepository) MemberIncomeUpdate(memberincome *MemberIncome, preloads ...string) (*MemberIncome, error) {
	repo := NewGenericRepository[MemberIncome](m.db.Client)
	return repo.Update(memberincome, preloads...)
}
func (m *ModelRepository) MemberIncomeUpdateByID(id string, column string, value interface{}, preloads ...string) (*MemberIncome, error) {
	repo := NewGenericRepository[MemberIncome](m.db.Client)
	return repo.UpdateByID(id, column, value, preloads...)
}
func (m *ModelRepository) MemberIncomeDeleteByID(id string) error {
	repo := NewGenericRepository[MemberIncome](m.db.Client)
	return repo.DeleteByID(id)
}
func (m *ModelRepository) MemberIncomeGetAll(preloads ...string) ([]*MemberIncome, error) {
	repo := NewGenericRepository[MemberIncome](m.db.Client)
	return repo.GetAll(preloads...)
}
