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
	AccountType    string    `gorm:"type:varchar(11);unsigned" json:"account_type"`
	Module         string    `gorm:"type:varchar(255);unsigned" json:"module"`
	Description    string    `gorm:"type:varchar(1000)" json:"description,omitempty"`
	Activity       string    `gorm:"type:varchar(255);unsigned" json:"activity"`
	Latitude       *float64  `gorm:"type:decimal(10,7)" json:"latitude,omitempty"`
	Longitude      *float64  `gorm:"type:decimal(10,7)" json:"longitude,omitempty"`
	Timestamp      time.Time `gorm:"type:timestamp" json:"timestamp"`
	IsDeleted      bool      `gorm:"default:false" json:"is_deleted"`
	IPAddress      string    `gorm:"type:varchar(45)" json:"ip_address"`
	UserAgent      string    `gorm:"type:varchar(1000)" json:"user_agent"`
	Referer        string    `gorm:"type:varchar(1000)" json:"referer"`
	Location       string    `gorm:"type:varchar(255)" json:"location"`
	AcceptLanguage string    `gorm:"type:varchar(255)" json:"accept_language"`

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

func (v *Footstep) BeforeCreate(tx *gorm.DB) (err error) {
	if v.ID == uuid.Nil {
		v.ID = uuid.New()
	}
	return
}

type FootstepResource struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`

	AccountType    string   `json:"accountType"`
	Module         string   `json:"module"`
	Description    string   `json:"description"`
	Activity       string   `json:"activity"`
	Latitude       *float64 `json:"latitude,omitempty"`
	Longitude      *float64 `json:"longitude,omitempty"`
	Timestamp      string   `json:"timestamp"`
	IsDeleted      bool     `json:"isDeleted"`
	IPAddress      string   `json:"ipAddress"`
	UserAgent      string   `json:"userAgent"`
	Referer        string   `json:"referer"`
	Location       string   `json:"location"`
	AcceptLanguage string   `json:"acceptLanguage"`

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
		ID:             footstep.ID,
		CreatedAt:      footstep.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      footstep.UpdatedAt.Format(time.RFC3339),
		DeletedAt:      footstep.DeletedAt.Time.Format(time.RFC3339),
		Module:         footstep.Module,
		AccountType:    footstep.AccountType,
		Description:    footstep.Description,
		Activity:       footstep.Activity,
		Latitude:       footstep.Latitude,
		Longitude:      footstep.Longitude,
		Timestamp:      footstep.Timestamp.Format(time.RFC3339),
		IsDeleted:      footstep.IsDeleted,
		IPAddress:      footstep.IPAddress,
		UserAgent:      footstep.UserAgent,
		Referer:        footstep.Referer,
		Location:       footstep.Location,
		AcceptLanguage: footstep.AcceptLanguage,
		AdminID:        footstep.AdminID,
		Admin:          m.AdminToResource(footstep.Admin),
		EmployeeID:     footstep.EmployeeID,
		Employee:       m.EmployeeToResource(footstep.Employee),
		OwnerID:        footstep.OwnerID,
		Owner:          m.OwnerToResource(footstep.Owner),
		MemberID:       footstep.MemberID,
		Member:         m.MemberToResource(footstep.Member),
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

func (m *ModelRepository) FootstepGetByID(id string, preloads ...string) (*Footstep, error) {
	repo := NewGenericRepository[Footstep](m.db.Client)
	return repo.GetByID(id, preloads...)
}

func (m *ModelRepository) FootstepCreate(footstep *Footstep, preloads ...string) (*Footstep, error) {
	repo := NewGenericRepository[Footstep](m.db.Client)
	return repo.Create(footstep, preloads...)
}
func (m *ModelRepository) FootstepUpdate(footstep *Footstep, preloads ...string) (*Footstep, error) {
	repo := NewGenericRepository[Footstep](m.db.Client)
	return repo.Update(footstep, preloads...)
}
func (m *ModelRepository) FootstepUpdateByID(id string, value *Footstep, preloads ...string) (*Footstep, error) {
	repo := NewGenericRepository[Footstep](m.db.Client)
	return repo.UpdateByID(id, value, preloads...)
}
func (m *ModelRepository) FootstepDeleteByID(id string) error {
	repo := NewGenericRepository[Footstep](m.db.Client)
	return repo.DeleteByID(id)
}
func (m *ModelRepository) FootstepGetAll(preloads ...string) ([]*Footstep, error) {
	repo := NewGenericRepository[Footstep](m.db.Client)
	return repo.GetAll(preloads...)
}
