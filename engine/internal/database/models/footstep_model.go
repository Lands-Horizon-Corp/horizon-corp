package models

import (
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

// FootstepToResource implements Models.
func (m *ModelResource) FootstepToResource(footstep *Footstep) *FootstepResource {
	if footstep == nil {
		return nil
	}
	return &FootstepResource{
		AccountType: footstep.AccountType,
		Description: footstep.Description,
		Activity:    footstep.Activity,
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

// FootstepToResourceList implements Models.
func (m *ModelResource) FootstepToResourceList(footsteps []*Footstep) []*FootstepResource {
	if footsteps == nil {
		return nil
	}
	var footstepResources []*FootstepResource
	for _, footstep := range footsteps {
		footstepResources = append(footstepResources, m.FootstepToResource(footstep))
	}
	return footstepResources
}

func (m *ModelResource) FootstepSeeders() error {
	m.logger.Info("Seeding Footstep")
	return nil
}
