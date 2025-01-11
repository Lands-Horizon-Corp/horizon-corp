package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Footstep struct {
	ID        uuid.UUID      `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Fields
	AccountType string    `gorm:"type:varchar(11);unsigned" json:"account_type"`
	Description string    `gorm:"type:varchar(1000)" json:"description,omitempty"`
	Activity    string    `gorm:"type:varchar(255);unsigned" json:"activity"`
	Latitude    *float64  `gorm:"type:decimal(10,7)" json:"latitude,omitempty"`
	Longitude   *float64  `gorm:"type:decimal(10,7)" json:"longitude,omitempty"`
	Timestamp   time.Time `gorm:"type:timestamp" json:"timestamp"`
	IsDeleted   bool      `gorm:"default:false" json:"is_deleted"`

	// Relationships
	AdminID    *uuid.UUID `gorm:"index" json:"admin_id,omitempty"`
	Admin      *Admin     `gorm:"foreignKey:AdminID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"admin,omitempty"`
	EmployeeID *uuid.UUID `gorm:"index" json:"employee_id,omitempty"`
	Employee   *Employee  `gorm:"foreignKey:EmployeeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"employee,omitempty"`
	OwnerID    *uuid.UUID `gorm:"index" json:"owner_id,omitempty"`
	Owner      *Owner     `gorm:"foreignKey:OwnerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"owner,omitempty"`
	MemberID   *uuid.UUID `gorm:"index" json:"member_id,omitempty"`
	Member     *Member    `gorm:"foreignKey:MemberID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"member,omitempty"`
}

type FootstepResource struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`

	AccountType string   `json:"accountType"`
	Description string   `json:"description"`
	Activity    string   `json:"activity"`
	Latitude    *float64 `json:"latitude,omitempty"`
	Longitude   *float64 `json:"longitude,omitempty"`
	Timestamp   string   `json:"timestamp"`
	IsDeleted   bool     `json:"isDeleted"`

	AdminID    *uuid.UUID        `json:"adminID"`
	Admin      *AdminResource    `json:"admin"`
	EmployeeID *uuid.UUID        `json:"employeeID"`
	Employee   *EmployeeResource `json:"employee"`
	OwnerID    *uuid.UUID        `json:"ownerID"`
	Owner      *OwnerResource    `json:"owner"`
	MemberID   *uuid.UUID        `json:"memberID"`
	Member     *MemberResource   `json:"member"`
}

func (m *ModelTransformer) FootstepToResource(footstep *Footstep) *FootstepResource {
	if footstep == nil {
		return nil
	}

	return &FootstepResource{

		ID:        footstep.ID,
		CreatedAt: footstep.CreatedAt.Format(time.RFC3339),
		UpdatedAt: footstep.UpdatedAt.Format(time.RFC3339),
		DeletedAt: footstep.DeletedAt.Time.Format(time.RFC3339),

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
