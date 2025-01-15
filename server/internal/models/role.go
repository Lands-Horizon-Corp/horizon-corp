package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	ID        uuid.UUID      `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name        string `gorm:"type:varchar(255);unique;unsigned" json:"name"`
	Description string `gorm:"type:text" json:"description"`
	ApiKey      string `gorm:"type:varchar(255);unique;unsigned" json:"api_key"`
	Color       string `gorm:"type:varchar(255)" json:"color"`

	ReadRole   bool `gorm:"default:false" json:"read_role"`
	WriteRole  bool `gorm:"default:false" json:"write_role"`
	UpdateRole bool `gorm:"default:false" json:"update_role"`
	DeleteRole bool `gorm:"default:false" json:"delete_role"`

	// Relationship 0 to many
	Admins    []*Admin    `gorm:"foreignKey:RoleID" json:"admins"`
	Owners    []*Owner    `gorm:"foreignKey:RoleID" json:"owners"`
	Employees []*Employee `gorm:"foreignKey:RoleID" json:"employees"`
	Members   []*Member   `gorm:"foreignKey:RoleID" json:"members"`
}

func (v *Role) BeforeCreate(tx *gorm.DB) (err error) {
	if v.ID == uuid.Nil {
		v.ID = uuid.New()
	}
	return
}

type RoleResource struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`

	Name        string `json:"name"`
	Description string `json:"description"`
	ApiKey      string `json:"apiKey"`
	Color       string `json:"color"`

	ReadRole   bool `json:"readRole"`
	WriteRole  bool `json:"writeRole"`
	UpdateRole bool `json:"updateRole"`
	DeleteRole bool `json:"deleteRole"`

	Admins    []*AdminResource    `json:"admins"`
	Owners    []*OwnerResource    `json:"owners"`
	Employees []*EmployeeResource `json:"employees"`
	Members   []*MemberResource   `json:"members"`
}

func (m *ModelTransformer) RoleToResource(role *Role) *RoleResource {
	if role == nil {
		return nil
	}

	return &RoleResource{

		ID:        role.ID,
		CreatedAt: role.CreatedAt.Format(time.RFC3339),
		UpdatedAt: role.UpdatedAt.Format(time.RFC3339),
		DeletedAt: role.DeletedAt.Time.Format(time.RFC3339),

		Name:        role.Name,
		Description: role.Description,
		ApiKey:      role.ApiKey,
		Color:       role.Color,
		ReadRole:    role.ReadRole,
		WriteRole:   role.WriteRole,
		UpdateRole:  role.UpdateRole,
		DeleteRole:  role.DeleteRole,

		Admins:    m.AdminToResourceList(role.Admins),
		Owners:    m.OwnerToResourceList(role.Owners),
		Employees: m.EmployeeToResourceList(role.Employees),
		Members:   m.MemberToResourceList(role.Members),
	}
}

func (m *ModelTransformer) RoleToResourceList(roleList []*Role) []*RoleResource {
	if roleList == nil {
		return nil
	}

	var roleResources []*RoleResource
	for _, role := range roleList {
		roleResources = append(roleResources, m.RoleToResource(role))
	}
	return roleResources
}

func (m *ModelRepository) RoleGetByID(id string, preloads ...string) (*Role, error) {
	repo := NewGenericRepository[Role](m.db.Client)
	return repo.GetByID(id, preloads...)
}

func (m *ModelRepository) RoleCreate(role *Role, preloads ...string) (*Role, error) {
	repo := NewGenericRepository[Role](m.db.Client)
	return repo.Create(role, preloads...)
}
func (m *ModelRepository) RoleUpdate(role *Role, preloads ...string) (*Role, error) {
	repo := NewGenericRepository[Role](m.db.Client)
	return repo.Update(role, preloads...)
}
func (m *ModelRepository) RoleUpdateByID(id string, column string, value interface{}, preloads ...string) (*Role, error) {
	repo := NewGenericRepository[Role](m.db.Client)
	return repo.UpdateByID(id, column, value, preloads...)
}
func (m *ModelRepository) RoleDeleteByID(id string) error {
	repo := NewGenericRepository[Role](m.db.Client)
	return repo.DeleteByID(id)
}
func (m *ModelRepository) RoleGetAll(preloads ...string) ([]*Role, error) {
	repo := NewGenericRepository[Role](m.db.Client)
	return repo.GetAll(preloads...)
}
