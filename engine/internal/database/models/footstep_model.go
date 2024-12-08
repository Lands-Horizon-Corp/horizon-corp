package models

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
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

type (
	FootstepResourceProvider interface {
		ToResource(footstep *Footstep) *FootstepResource
		ToResourceList(footstep []*Footstep) []*FootstepResource
	}
)

type FootstepModel struct {
	lc            *fx.Lifecycle
	db            *gorm.DB
	logger        *zap.Logger
	adminModel    AdminResourceProvider
	employeeModel EmployeeResourceProvider
	ownerModel    OwnerResourceProvider
	memberModel   MemberResourceProvider
}

func NewFootstepModel(
	lc *fx.Lifecycle,
	db *gorm.DB,
	logger *zap.Logger,
	adminModel AdminResourceProvider,
	employeeModel EmployeeResourceProvider,
	ownerModel OwnerResourceProvider,
	memberModel MemberResourceProvider,
) *FootstepModel {
	return &FootstepModel{
		lc:            lc,
		db:            db,
		logger:        logger,
		adminModel:    adminModel,
		employeeModel: employeeModel,
		ownerModel:    ownerModel,
		memberModel:   memberModel,
	}
}

func (fm *FootstepModel) SeedDatabase() {
}

func (fm *FootstepModel) ToResource(footstep *Footstep) *FootstepResource {
	if footstep == nil {
		return nil
	}
	return &FootstepResource{
		AccountType: footstep.AccountType,
		Description: footstep.Description,
		Activity:    footstep.Activity,
		AdminID:     footstep.AdminID,
		Admin:       fm.adminModel.ToResource(footstep.Admin),
		EmployeeID:  footstep.EmployeeID,
		Employee:    fm.employeeModel.ToResource(footstep.Employee),
		OwnerID:     footstep.OwnerID,
		Owner:       fm.ownerModel.ToResource(footstep.Owner),
		MemberID:    footstep.MemberID,
		Member:      fm.memberModel.ToResource(footstep.Member),
	}
}

func (fm *FootstepModel) ToResourceList(footsteps []*Footstep) []*FootstepResource {
	if footsteps == nil {
		return nil
	}
	var footstepResources []*FootstepResource
	for _, footstep := range footsteps {
		footstepResources = append(footstepResources, fm.ToResource(footstep))
	}
	return footstepResources
}
