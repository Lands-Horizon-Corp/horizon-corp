package models

import (
	"time"

	"gorm.io/gorm"
)

type Footstep struct {
	gorm.Model

	// Fields
	AccountType string    `gorm:"type:varchar(11);unsigned" json:"account_type"`
	Description string    `gorm:"type:varchar(1000)" json:"description,omitempty"`
	Activity    string    `gorm:"type:varchar(255);unsigned" json:"activity"`
	Latitude    *float64  `gorm:"type:decimal(10,7)" json:"latitude,omitempty"`
	Longitude   *float64  `gorm:"type:decimal(10,7)" json:"longitude,omitempty"`
	Timestamp   time.Time `gorm:"type:timestamp" json:"timestamp"`
	IsDeleted   bool      `gorm:"default:false" json:"is_deleted"`

	// Relationships
	AdminID    *uint     `gorm:"index" json:"admin_id,omitempty"`
	Admin      *Admin    `gorm:"foreignKey:AdminID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"admin,omitempty"`
	EmployeeID *uint     `gorm:"index" json:"employee_id,omitempty"`
	Employee   *Employee `gorm:"foreignKey:EmployeeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"employee,omitempty"`
	OwnerID    *uint     `gorm:"index" json:"owner_id,omitempty"`
	Owner      *Owner    `gorm:"foreignKey:OwnerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"owner,omitempty"`
	MemberID   *uint     `gorm:"index" json:"member_id,omitempty"`
	Member     *Member   `gorm:"foreignKey:MemberID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"member,omitempty"`
}

type FootstepResource struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	AccountType string   `json:"accountType"`
	Description string   `json:"description"`
	Activity    string   `json:"activity"`
	Latitude    *float64 `json:"latitude,omitempty"`
	Longitude   *float64 `json:"longitude,omitempty"`
	Timestamp   string   `json:"timestamp"`
	IsDeleted   bool     `json:"isDeleted"`

	AdminID    *uint             `json:"adminID"`
	Admin      *AdminResource    `json:"admin"`
	EmployeeID *uint             `json:"employeeID"`
	Employee   *EmployeeResource `json:"employee"`
	OwnerID    *uint             `json:"ownerID"`
	Owner      *OwnerResource    `json:"owner"`
	MemberID   *uint             `json:"memberID"`
	Member     *MemberResource   `json:"member"`
}

func (m *ModelTransformer) FootstepToResource(footstep *Footstep) *FootstepResource {
	if footstep == nil {
		return nil
	}

	return &FootstepResource{
		ID:          footstep.ID,
		CreatedAt:   footstep.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   footstep.UpdatedAt.Format(time.RFC3339),
		AccountType: footstep.AccountType,
		Description: footstep.Description,
		Activity:    footstep.Activity,
		Latitude:    footstep.Latitude,
		Longitude:   footstep.Longitude,
		Timestamp:   footstep.Timestamp.Format(time.RFC3339),
		IsDeleted:   footstep.IsDeleted,
		AdminID:     footstep.AdminID,
		Admin:       m.AdminToResource(footstep.Admin),
		EmployeeID:  footstep.EmployeeID,
		Employee:    m.EmployeeToResource(footstep.Employee),
		OwnerID:     footstep.OwnerID,
		Owner:       m.OwnerToResource(footstep.Owner),
		MemberID:    footstep.MemberID,
		Member:      m.MemberToResource(footstep.Member),
	}
}

func (m *ModelTransformer) FootstepToResourceList(footstepList []*Footstep) []*FootstepResource {
	if footstepList == nil {
		return nil
	}

	var footstepResources []*FootstepResource
	for _, footstep := range footstepList {
		footstepResources = append(footstepResources, m.FootstepToResource(footstep))
	}
	return footstepResources
}
