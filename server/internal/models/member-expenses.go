package models

import (
	"time"

	"gorm.io/gorm"
)

type MemberExpenses struct {
	gorm.Model
	MembersProfileID uint           `gorm:"unsigned" json:"members_profile_id"`
	Name             string         `gorm:"type:varchar(255);unsigned" json:"name"`
	Amount           float64        `gorm:"type:decimal(10,2);unsigned" json:"amount"`
	Date             time.Time      `gorm:"type:date" json:"date"`
	Description      string         `gorm:"type:text" json:"description"`
	MembersProfile   *MemberProfile `gorm:"foreignKey:MembersProfileID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members_profile"`
}

type MemberExpensesResource struct {
	ID               uint                   `json:"id"`
	CreatedAt        string                 `json:"createdAt"`
	UpdatedAt        string                 `json:"updatedAt"`
	MembersProfileID uint                   `json:"membersProfileID"`
	Name             string                 `json:"name"`
	Amount           float64                `json:"amount"`
	Date             string                 `json:"date"`
	Description      string                 `json:"description"`
	MembersProfile   *MemberProfileResource `json:"membersProfile,omitempty"`
}

func (m *ModelTransformer) MemberExpensesToResource(expense *MemberExpenses) *MemberExpensesResource {
	if expense == nil {
		return nil
	}

	return &MemberExpensesResource{
		ID:               expense.ID,
		CreatedAt:        expense.CreatedAt.Format(time.RFC3339),
		UpdatedAt:        expense.UpdatedAt.Format(time.RFC3339),
		MembersProfileID: expense.MembersProfileID,
		Name:             expense.Name,
		Amount:           expense.Amount,
		Date:             expense.Date.Format("2006-01-02"),
		Description:      expense.Description,
		MembersProfile:   m.MemberProfileToResource(expense.MembersProfile),
	}
}

func (m *ModelTransformer) MemberExpensesToResourceList(expensesList []*MemberExpenses) []*MemberExpensesResource {
	if expensesList == nil {
		return nil
	}

	var expensesResources []*MemberExpensesResource
	for _, expense := range expensesList {
		expensesResources = append(expensesResources, m.MemberExpensesToResource(expense))
	}
	return expensesResources
}
