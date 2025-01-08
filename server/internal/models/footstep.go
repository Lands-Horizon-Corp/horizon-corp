package models

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"gorm.io/gorm"
)

type Footstep struct {
	gorm.Model

	// Fields
	AccountType string `gorm:"type:varchar(11);not null" json:"account_type"`
	Description string `gorm:"type:varchar(1000)" json:"description,omitempty"`
	Activity    string `gorm:"type:varchar(255);not null" json:"activity"`

	// Relationship 0 to 1
	AdminID *uint  `gorm:"index" json:"admin_id,omitempty"`
	Admin   *Admin `gorm:"foreignKey:AdminID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"admin,omitempty"`

	// Relationship 0 to 1
	EmployeeID *uint     `gorm:"index" json:"employee_id,omitempty"`
	Employee   *Employee `gorm:"foreignKey:EmployeeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"employee,omitempty"`

	// Relationship 0 to 1
	OwnerID *uint  `gorm:"index" json:"owner_id,omitempty"`
	Owner   *Owner `gorm:"foreignKey:OwnerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"owner,omitempty"`

	// Relationship 0 to 1
	MemberID *uint   `gorm:"index" json:"member_id,omitempty"`
	Member   *Member `gorm:"foreignKey:MemberID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"member,omitempty"`
}

type FootstepResource struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	AccountType string            `json:"accountType"`
	Description string            `json:"description"`
	Activity    string            `json:"activity"`
	AdminID     *uint             `json:"adminID"`
	Admin       *AdminResource    `json:"admin"`
	EmployeeID  *uint             `json:"employeeID"`
	Employee    *EmployeeResource `json:"employee"`
	OwnerID     *uint             `json:"ownerID"`
	Owner       *OwnerResource    `json:"owner"`
	MemberID    *uint             `json:"memberID"`
	Member      *MemberResource   `json:"member"`
}

type FootstepRepository struct {
	*Repository[Footstep]
}

func NewFootstepRepository(db *providers.DatabaseService) *FootstepRepository {
	return &FootstepRepository{
		Repository: NewRepository[Footstep](db),
	}
}
