package models

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type Gender struct {
	gorm.Model

	// Fields
	Name        string `gorm:"type:varchar(255);unique;unsigned" json:"name"`
	Description string `gorm:"type:text" json:"description"`

	// Relationship 0 to many
	Employees []*Employee `gorm:"foreignKey:GenderID" json:"employees"`
	Members   []*Member   `gorm:"foreignKey:GenderID" json:"members"`
	Owners    []*Owner    `gorm:"foreignKey:GenderID" json:"owners"`
	Admins    []*Admin    `gorm:"foreignKey:GenderID" json:"admins"`
}

type GenderResource struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	Name        string              `json:"name"`
	Description string              `json:"description"`
	Employees   []*EmployeeResource `json:"employees"`
	Members     []*MemberResource   `json:"members"`
	Owners      []*OwnerResource    `json:"owners"`
	Admins      []*AdminResource    `json:"admins"`
}

type GenderRequest struct {
	Name        string `json:"name" validate:"required,max=255"`
	Description string `json:"description,omitempty" validate:"max=1000"`
}

func (m *ModelResource) GenderToResource(gender *Gender) *GenderResource {
	if gender == nil {
		return nil
	}
	return &GenderResource{
		ID:        gender.ID,
		CreatedAt: gender.CreatedAt.Format(time.RFC3339),
		UpdatedAt: gender.UpdatedAt.Format(time.RFC3339),

		Name:        gender.Name,
		Description: gender.Description,
		Employees:   m.EmployeeToResourceList(gender.Employees),
		Members:     m.MemberToResourceList(gender.Members),
		Owners:      m.OwnerToResourceList(gender.Owners),
		Admins:      m.AdminToResourceList(gender.Admins),
	}
}

func (m *ModelResource) GenderToResourceList(genders []*Gender) []*GenderResource {
	if genders == nil {
		return nil
	}

	var genderResources []*GenderResource
	for _, gender := range genders {
		genderResources = append(genderResources, m.GenderToResource(gender))
	}
	return genderResources
}

// GenderToRecord converts a slice of Gender pointers into CSV records and headers.
func (m *ModelResource) GenderToRecord(genders []*Gender) ([][]string, []string) {
	// Convert Gender structs to GenderResource structs
	resource := m.GenderToResourceList(genders)
	records := make([][]string, 0, len(resource))

	for _, gender := range resource {
		// Basic Fields
		id := strconv.Itoa(int(gender.ID))
		name := sanitizeCSVField(gender.Name)
		description := sanitizeCSVField(gender.Description)
		createdAt := sanitizeCSVField(gender.CreatedAt)
		updatedAt := sanitizeCSVField(gender.UpdatedAt)

		// Handle Related Entities
		employees := "N/A"
		if len(gender.Employees) > 0 {
			employeeNames := make([]string, 0, len(gender.Employees))
			for _, emp := range gender.Employees {
				fullName := fmt.Sprintf("%s %s", emp.FirstName, emp.LastName)
				employeeNames = append(employeeNames, sanitizeCSVField(fullName))
			}
			employees = strings.Join(employeeNames, "; ")
		}

		members := "N/A"
		if len(gender.Members) > 0 {
			memberNames := make([]string, 0, len(gender.Members))
			for _, mem := range gender.Members {
				fullName := fmt.Sprintf("%s %s", mem.FirstName, mem.LastName)
				memberNames = append(memberNames, sanitizeCSVField(fullName))
			}
			members = strings.Join(memberNames, "; ")
		}

		owners := "N/A"
		if len(gender.Owners) > 0 {
			ownerNames := make([]string, 0, len(gender.Owners))
			for _, own := range gender.Owners {
				fullName := fmt.Sprintf("%s %s", own.FirstName, own.LastName)
				ownerNames = append(ownerNames, sanitizeCSVField(fullName))
			}
			owners = strings.Join(ownerNames, "; ")
		}

		admins := "N/A"
		if len(gender.Admins) > 0 {
			adminNames := make([]string, 0, len(gender.Admins))
			for _, adm := range gender.Admins {
				fullName := fmt.Sprintf("%s %s", adm.FirstName, adm.LastName)
				adminNames = append(adminNames, sanitizeCSVField(fullName))
			}
			admins = strings.Join(adminNames, "; ")
		}

		// Assemble the record
		record := []string{
			id,
			name,
			description,
			createdAt,
			updatedAt,
			employees,
			members,
			owners,
			admins,
		}
		records = append(records, record)
	}

	headers := []string{
		"ID",
		"Name",
		"Description",
		"Created At",
		"Updated At",
		"Employees",
		"Members",
		"Owners",
		"Admins",
	}

	return records, headers
}

func (m *ModelResource) ValidateGenderRequest(req *GenderRequest) error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return m.helpers.FormatValidationError(err)
	}
	return nil
}

func (m *ModelResource) GenderSeeders() error {
	m.logger.Info("Seeding Genders")

	// Define predefined genders
	predefinedGenders := []GenderRequest{
		{
			Name:        "Male",
			Description: "Male gender",
		},
		{
			Name:        "Female",
			Description: "Female gender",
		},
		{
			Name:        "Other",
			Description: "Other gender identities",
		},
		{
			Name:        "Prefer Not to Say",
			Description: "Prefer not to disclose gender",
		},
	}

	for _, genderReq := range predefinedGenders {

		gender := &Gender{
			Name:        genderReq.Name,
			Description: genderReq.Description,
		}

		err := m.GenderDB.Create(gender)
		if err != nil {
			log.Printf("Error seeding gender %s: %v", gender.Name, err)
		}
	}

	m.logger.Info("Gender seeding completed")
	return nil
}
