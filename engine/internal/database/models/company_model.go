package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/managers/filter"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model

	// Fields
	Name          string  `gorm:"type:varchar(255);not null" json:"name"`
	Description   string  `gorm:"type:text" json:"description"`
	Address       string  `gorm:"type:varchar(500)" json:"address"`
	Longitude     float64 `gorm:"type:decimal(10,7)" json:"longitude"`
	Latitude      float64 `gorm:"type:decimal(10,7)" json:"latitude"`
	ContactNumber string  `gorm:"type:varchar(255);not null" json:"contact_number"`

	// Relationship 0 to 1
	OwnerID *uint  `gorm:"type:bigint;unsigned" json:"owner_id"`
	Owner   *Owner `gorm:"foreignKey:OwnerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"owner"`

	// Relationship 0 to 1
	MediaID         *uint  `gorm:"type:bigint;unsigned" json:"media_id"`
	Media           *Media `gorm:"foreignKey:MediaID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"media"`
	IsAdminVerified bool   `gorm:"default:false" json:"is_admin_verified"`

	// Relationship 0 to many
	Branches []*Branch `gorm:"foreignKey:CompanyID" json:"branches"`
}

type CompanyResource struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	Name            string            `json:"name"`
	Description     string            `json:"description"`
	Address         string            `json:"address"`
	Longitude       float64           `json:"longitude"`
	Latitude        float64           `json:"latitude"`
	ContactNumber   string            `json:"contactNumber"`
	OwnerID         *uint             `json:"ownerID"`
	Owner           *OwnerResource    `json:"owner"`
	MediaID         *uint             `json:"mediaID"`
	Media           *MediaResource    `json:"media"`
	IsAdminVerified bool              `json:"isAdminVerified"`
	Branches        []*BranchResource `json:"branches"`
}

type CompanyRequest struct {
	Name            string  `json:"name" validate:"required,max=255"`
	Description     string  `json:"description,omitempty"`
	Address         string  `json:"address" validate:"required,max=500"`
	Longitude       float64 `json:"longitude" validate:"longitude"`
	Latitude        float64 `json:"latitude" validate:"latitude"`
	ContactNumber   string  `json:"contactNumber" validate:"required,max=255"`
	OwnerID         *uint   `json:"ownerID,omitempty"`
	MediaID         *uint   `json:"mediaID,omitempty"`
	IsAdminVerified bool    `json:"isAdminVerified"`
}

func (m *ModelResource) CompanyToResource(company *Company) *CompanyResource {
	if company == nil {
		return nil
	}

	return &CompanyResource{
		ID:        company.ID,
		CreatedAt: company.CreatedAt.Format(time.RFC3339),
		UpdatedAt: company.UpdatedAt.Format(time.RFC3339),

		Name:            company.Name,
		Description:     company.Description,
		Address:         company.Address,
		Longitude:       company.Longitude,
		Latitude:        company.Latitude,
		ContactNumber:   company.ContactNumber,
		OwnerID:         company.OwnerID,
		Owner:           m.OwnerToResource(company.Owner),
		MediaID:         company.MediaID,
		Media:           m.MediaToResource(company.Media),
		IsAdminVerified: company.IsAdminVerified,
		Branches:        m.BranchToResourceList(company.Branches),
	}
}

func (m *ModelResource) CompanyToResourceList(companies []*Company) []*CompanyResource {
	if companies == nil {
		return nil
	}
	var companyResources []*CompanyResource
	for _, company := range companies {
		fmt.Println("--")
		fmt.Println(company)
		fmt.Println("--")
		companyResources = append(companyResources, m.CompanyToResource(company))
	}
	return companyResources
}

func (m *ModelResource) CompanyToRecord(company []*Company) ([][]string, []string) {
	resource := m.CompanyToResourceList(company)
	records := make([][]string, 0, len(resource))

	for _, company := range resource {
		id := strconv.Itoa(int(company.ID))
		name := sanitizeCSVField(company.Name)
		description := sanitizeCSVField(company.Description)
		address := sanitizeCSVField(company.Address)
		longitude := fmt.Sprintf("%.6f", company.Longitude)
		latitude := fmt.Sprintf("%.6f", company.Latitude)
		contactNumber := sanitizeCSVField(company.ContactNumber)
		isAdminVerified := strconv.FormatBool(company.IsAdminVerified)
		createdAt := sanitizeCSVField(company.CreatedAt)
		updatedAt := sanitizeCSVField(company.UpdatedAt)

		// Handle Owner
		ownerName := "N/A"
		if company.Owner != nil {
			ownerName = sanitizeCSVField(company.Owner.FirstName)
		}

		// Handle Media
		mediaURL := "N/A"
		if company.Media != nil {
			mediaURL = sanitizeCSVField(company.Media.URL)
		}

		// Handle Branches
		branchAddresses := []string{}
		for _, branch := range company.Branches {
			branchAddresses = append(branchAddresses, sanitizeCSVField(branch.Address))
		}
		branches := strings.Join(branchAddresses, "; ")

		record := []string{
			id,
			name,
			description,
			address,
			longitude,
			latitude,
			contactNumber,
			ownerName,
			mediaURL,
			isAdminVerified,
			branches,
			createdAt,
			updatedAt,
		}
		records = append(records, record)
	}
	headers := []string{
		"ID",
		"Name",
		"Description",
		"Address",
		"Longitude",
		"Latitude",
		"Contact Number",
		"Owner Name",
		"Media URL",
		"Is Admin Verified",
		"Branches",
		"Created At",
		"Updated At",
	}

	return records, headers
}

func (m *ModelResource) ValidateCompanyRequest(req *CompanyRequest) error {
	validate := validator.New()
	validate.RegisterValidation("longitude", func(fl validator.FieldLevel) bool {
		lon := fl.Field().Float()
		return lon >= -180 && lon <= 180
	})
	validate.RegisterValidation("latitude", func(fl validator.FieldLevel) bool {
		lat := fl.Field().Float()
		return lat >= -90 && lat <= 90
	})
	err := validate.Struct(req)
	if err != nil {
		return m.helpers.FormatValidationError(err)
	}

	return nil
}

func (m *ModelResource) CompanyFilterForAdmin(filters string, pageSize, pageIndex int) (filter.FilterPages[Company], error) {

	db := m.db.Client.Model(&Company{}).
		Select("companies.*, COUNT(branches.id) as branches").
		Joins("LEFT JOIN branches ON branches.company_id = companies.id").
		Group("companies.id")
	return m.CompanyDB.GetPaginatedResult(db, filters, pageSize, pageIndex)
}
func (m *ModelResource) CompanyFilterForAdminRecord(filters string) ([]*Company, error) {
	db := m.db.Client
	return m.CompanyDB.GetFilteredResults(db, filters)
}

func (m *ModelResource) CompanySeeders() error {
	m.logger.Info("Seeding Company")
	return nil
}
