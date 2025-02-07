package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model

	Name        string `gorm:"type:varchar(255);unique;unsigned" json:"name"`
	Description string `gorm:"type:text" json:"description"`
	ApiKey      string `gorm:"type:varchar(255);unique;unsigned" json:"api_key"`
	Color       string `gorm:"type:varchar(255)" json:"color"`

	ReadRole   bool `gorm:"default:false" json:"read_role"`
	WriteRole  bool `gorm:"default:false" json:"write_role"`
	UpdateRole bool `gorm:"default:false" json:"update_role"`
	DeleteRole bool `gorm:"default:false" json:"delete_role"`

	ReadErrorDetails   bool `gorm:"default:false" json:"read_error_details"`
	WriteErrorDetails  bool `gorm:"default:false" json:"write_error_details"`
	UpdateErrorDetails bool `gorm:"default:false" json:"update_error_details"`
	DeleteErrorDetails bool `gorm:"default:false" json:"delete_error_details"`

	ReadGender   bool `gorm:"default:false" json:"read_gender"`
	WriteGender  bool `gorm:"default:false" json:"write_gender"`
	UpdateGender bool `gorm:"default:false" json:"update_gender"`
	DeleteGender bool `gorm:"default:false" json:"delete_gender"`

	// Relationship 0 to many
	Admins    []*Admin    `gorm:"foreignKey:RoleID" json:"admins"`
	Owners    []*Owner    `gorm:"foreignKey:RoleID" json:"owners"`
	Employees []*Employee `gorm:"foreignKey:RoleID" json:"employees"`
	Members   []*Member   `gorm:"foreignKey:RoleID" json:"members"`
}

type RoleResource struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	Name               string `json:"name"`
	Description        string `json:"description"`
	ApiKey             string `json:"apiKey"`
	Color              string `json:"color"`
	ReadRole           bool   `json:"readRole"`
	WriteRole          bool   `json:"writeRole"`
	UpdateRole         bool   `json:"updateRole"`
	DeleteRole         bool   `json:"deleteRole"`
	ReadErrorDetails   bool   `json:"readErrorDetails"`
	WriteErrorDetails  bool   `json:"writeErrorDetails"`
	UpdateErrorDetails bool   `json:"updateErrorDetails"`
	DeleteErrorDetails bool   `json:"deleteErrorDetails"`
	ReadGender         bool   `json:"readGender"`
	WriteGender        bool   `json:"writeGender"`
	UpdateGender       bool   `json:"updateGender"`
	DeleteGender       bool   `json:"deleteGender"`

	Admins    []*AdminResource    `json:"admins"`
	Owners    []*OwnerResource    `json:"owners"`
	Employees []*EmployeeResource `json:"employees"`
	Members   []*MemberResource   `json:"members"`
}

type RoleRequest struct {
	Name               string `json:"name" validate:"required,max=255"`
	Description        string `json:"description,omitempty" validate:"max=1000"`
	ApiKey             string `json:"apiKey" validate:"required,max=255"`
	Color              string `json:"color,omitempty" validate:"max=255"`
	ReadRole           bool   `json:"readRole"`
	WriteRole          bool   `json:"writeRole"`
	UpdateRole         bool   `json:"updateRole"`
	DeleteRole         bool   `json:"deleteRole"`
	ReadErrorDetails   bool   `json:"readErrorDetails"`
	WriteErrorDetails  bool   `json:"writeErrorDetails"`
	UpdateErrorDetails bool   `json:"updateErrorDetails"`
	DeleteErrorDetails bool   `json:"deleteErrorDetails"`
	ReadGender         bool   `json:"readGender"`
	WriteGender        bool   `json:"writeGender"`
	UpdateGender       bool   `json:"updateGender"`
	DeleteGender       bool   `json:"deleteGender"`
}

// RoleToResource implements Models.
func (m *ModelResource) RoleToResource(role *Role) *RoleResource {
	if role == nil {
		return nil
	}

	return &RoleResource{
		ID:        role.ID,
		CreatedAt: role.CreatedAt.Format(time.RFC3339),
		UpdatedAt: role.UpdatedAt.Format(time.RFC3339),

		Name:               role.Name,
		Description:        role.Description,
		ApiKey:             role.ApiKey,
		Color:              role.Color,
		ReadRole:           role.ReadRole,
		WriteRole:          role.WriteRole,
		UpdateRole:         role.UpdateRole,
		DeleteRole:         role.DeleteRole,
		ReadErrorDetails:   role.ReadErrorDetails,
		WriteErrorDetails:  role.WriteErrorDetails,
		UpdateErrorDetails: role.UpdateErrorDetails,
		DeleteErrorDetails: role.DeleteErrorDetails,
		ReadGender:         role.ReadGender,
		WriteGender:        role.WriteGender,
		UpdateGender:       role.UpdateGender,
		DeleteGender:       role.DeleteGender,
		Admins:             m.AdminToResourceList(role.Admins),
		Owners:             m.OwnerToResourceList(role.Owners),
		Employees:          m.EmployeeToResourceList(role.Employees),
		Members:            m.MemberToResourceList(role.Members),
	}
}

func (m *ModelResource) RoleToResourceList(roles []*Role) []*RoleResource {
	if roles == nil {
		return nil
	}
	var roleResources []*RoleResource
	for _, role := range roles {
		roleResources = append(roleResources, m.RoleToResource(role))
	}
	return roleResources
}

// RoleToRecord converts a slice of Role pointers into CSV records and headers.
func (m *ModelResource) RoleToRecord(roles []*Role) ([][]string, []string) {
	// Convert Role structs to RoleResource structs
	resources := m.RoleToResourceList(roles)
	records := make([][]string, 0, len(resources))

	for _, role := range resources {
		// Basic Fields
		id := strconv.Itoa(int(role.ID))
		name := sanitizeCSVField(role.Name)
		description := sanitizeCSVField(role.Description)
		apiKey := sanitizeCSVField(role.ApiKey)
		color := sanitizeCSVField(role.Color)
		readRole := strconv.FormatBool(role.ReadRole)
		writeRole := strconv.FormatBool(role.WriteRole)
		updateRole := strconv.FormatBool(role.UpdateRole)
		deleteRole := strconv.FormatBool(role.DeleteRole)
		readErrorDetails := strconv.FormatBool(role.ReadErrorDetails)
		writeErrorDetails := strconv.FormatBool(role.WriteErrorDetails)
		updateErrorDetails := strconv.FormatBool(role.UpdateErrorDetails)
		deleteErrorDetails := strconv.FormatBool(role.DeleteErrorDetails)
		readGender := strconv.FormatBool(role.ReadGender)
		writeGender := strconv.FormatBool(role.WriteGender)
		updateGender := strconv.FormatBool(role.UpdateGender)
		deleteGender := strconv.FormatBool(role.DeleteGender)
		createdAt := sanitizeCSVField(role.CreatedAt)
		updatedAt := sanitizeCSVField(role.UpdatedAt)

		// Handle Related Entities

		// Admins
		admins := "N/A"
		if len(role.Admins) > 0 {
			adminNames := make([]string, 0, len(role.Admins))
			for _, admin := range role.Admins {
				fullName := fmt.Sprintf("%s %s", admin.FirstName, admin.LastName)
				adminNames = append(adminNames, sanitizeCSVField(fullName))
			}
			admins = strings.Join(adminNames, "; ")
		}

		// Owners
		owners := "N/A"
		if len(role.Owners) > 0 {
			ownerNames := make([]string, 0, len(role.Owners))
			for _, owner := range role.Owners {
				fullName := fmt.Sprintf("%s %s", owner.FirstName, owner.LastName)
				ownerNames = append(ownerNames, sanitizeCSVField(fullName))
			}
			owners = strings.Join(ownerNames, "; ")
		}

		// Employees
		employees := "N/A"
		if len(role.Employees) > 0 {
			employeeNames := make([]string, 0, len(role.Employees))
			for _, emp := range role.Employees {
				fullName := fmt.Sprintf("%s %s", emp.FirstName, emp.LastName)
				employeeNames = append(employeeNames, sanitizeCSVField(fullName))
			}
			employees = strings.Join(employeeNames, "; ")
		}

		// Members
		members := "N/A"
		if len(role.Members) > 0 {
			memberNames := make([]string, 0, len(role.Members))
			for _, mem := range role.Members {
				fullName := fmt.Sprintf("%s %s", mem.FirstName, mem.LastName)
				memberNames = append(memberNames, sanitizeCSVField(fullName))
			}
			members = strings.Join(memberNames, "; ")
		}

		// Assemble the record
		record := []string{
			id,
			name,
			description,
			apiKey,
			color,
			readRole,
			writeRole,
			updateRole,
			deleteRole,
			readErrorDetails,
			writeErrorDetails,
			updateErrorDetails,
			deleteErrorDetails,
			readGender,
			writeGender,
			updateGender,
			deleteGender,
			createdAt,
			updatedAt,
			admins,
			owners,
			employees,
			members,
		}
		records = append(records, record)
	}
	headers := []string{
		"ID",
		"Name",
		"Description",
		"API Key",
		"Color",
		"Read Role",
		"Write Role",
		"Update Role",
		"Delete Role",
		"Read Error Details",
		"Write Error Details",
		"Update Error Details",
		"Delete Error Details",
		"Read Gender",
		"Write Gender",
		"Update Gender",
		"Delete Gender",
		"Created At",
		"Updated At",
		"Admins",
		"Owners",
		"Employees",
		"Members",
	}

	return records, headers
}

func (m *ModelResource) ValidateRoleRequest(req *RoleRequest) error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return m.helpers.FormatValidationError(err)
	}
	return nil
}

func (m *ModelResource) RoleSeeders() error {
	m.logger.Info("Seeding Role")
	return nil
}
