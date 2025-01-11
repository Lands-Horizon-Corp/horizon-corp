package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MemberBranchRegistration struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Fields
	MemberID   uuid.UUID  `gorm:"type:bigint;unsigned;unsigned" json:"member_id"`
	BranchID   uuid.UUID  `gorm:"type:bigint;unsigned;unsigned" json:"branch_id"`
	Status     string     `gorm:"type:enum('Pending', 'Verified', 'Rejected');default:'Pending'" json:"status"`
	Remarks    string     `gorm:"type:text" json:"remarks"`
	VerifiedBy *uuid.UUID `gorm:"type:bigint;unsigned" json:"verified_by"`
	VerifiedAt *time.Time `gorm:"type:datetime" json:"verified_at"`

	// Relationships
	Member   *Member   `gorm:"foreignKey:MemberID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"member"`
	Branch   *Branch   `gorm:"foreignKey:BranchID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"branch"`
	Employee *Employee `gorm:"foreignKey:VerifiedBy;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"employee,omitempty"`
}

type MemberBranchRegistrationResource struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`

	MemberID           uuid.UUID         `json:"memberID"`
	BranchID           uuid.UUID         `json:"branchID"`
	Status             string            `json:"status"`
	Remarks            string            `json:"remarks"`
	VerifiedBy         *uuid.UUID        `json:"verifiedBy,omitempty"`
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
		ID:        registration.ID,
		CreatedAt: registration.CreatedAt.Format(time.RFC3339),
		UpdatedAt: registration.UpdatedAt.Format(time.RFC3339),
		DeletedAt: registration.DeletedAt.Time.Format(time.RFC3339),

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
