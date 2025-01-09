package models

import (
	"time"

	"gorm.io/gorm"
)

type MemberBranchRegistration struct {
	gorm.Model

	// Fields
	MemberID   uint       `gorm:"type:bigint;unsigned;not null" json:"member_id"`
	BranchID   uint       `gorm:"type:bigint;unsigned;not null" json:"branch_id"`
	Status     string     `gorm:"type:enum('Pending', 'Verified', 'Rejected');default:'Pending'" json:"status"`
	Remarks    string     `gorm:"type:text" json:"remarks"`
	VerifiedBy *uint      `gorm:"type:bigint;unsigned" json:"verified_by"`
	VerifiedAt *time.Time `gorm:"type:datetime" json:"verified_at"`

	// Relationships
	Member   *Member   `gorm:"foreignKey:MemberID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"member"`
	Branch   *Branch   `gorm:"foreignKey:BranchID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"branch"`
	Employee *Employee `gorm:"foreignKey:VerifiedBy;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"employee,omitempty"`
}

type MemberBranchRegistrationResource struct {
	ID                 uint              `json:"id"`
	CreatedAt          string            `json:"createdAt"`
	UpdatedAt          string            `json:"updatedAt"`
	MemberID           uint              `json:"memberID"`
	BranchID           uint              `json:"branchID"`
	Status             string            `json:"status"`
	Remarks            string            `json:"remarks"`
	VerifiedBy         *uint             `json:"verifiedBy,omitempty"`
	VerifiedAt         *string           `json:"verifiedAt,omitempty"`
	Member             *MemberResource   `json:"member,omitempty"`
	Branch             *BranchResource   `json:"branch,omitempty"`
	VerifiedByEmployee *EmployeeResource `json:"verifiedByEmployee,omitempty"`
}

func (m *ModelTransformer) MemberBranchRegistrationToResource(registration *MemberBranchRegistration) *MemberBranchRegistrationResource {
	if registration == nil {
		return nil
	}

	var verifiedAt *string
	if registration.VerifiedAt != nil {
		verified := registration.VerifiedAt.Format(time.RFC3339)
		verifiedAt = &verified
	}

	return &MemberBranchRegistrationResource{
		ID:                 registration.ID,
		CreatedAt:          registration.CreatedAt.Format(time.RFC3339),
		UpdatedAt:          registration.UpdatedAt.Format(time.RFC3339),
		MemberID:           registration.MemberID,
		BranchID:           registration.BranchID,
		Status:             registration.Status,
		Remarks:            registration.Remarks,
		VerifiedBy:         registration.VerifiedBy,
		VerifiedAt:         verifiedAt,
		Member:             m.MemberToResource(registration.Member),
		Branch:             m.BranchToResource(registration.Branch),
		VerifiedByEmployee: m.EmployeeToResource(registration.Employee),
	}
}

func (m *ModelTransformer) MemberBranchRegistrationToResourceList(registrationList []*MemberBranchRegistration) []*MemberBranchRegistrationResource {
	if registrationList == nil {
		return nil
	}

	var registrationResources []*MemberBranchRegistrationResource
	for _, registration := range registrationList {
		registrationResources = append(registrationResources, m.MemberBranchRegistrationToResource(registration))
	}
	return registrationResources
}
