package company

import (
	"fmt"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/managers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/xuri/excelize/v2"
)

type CompanyExport struct {
	db *providers.DatabaseService
}

func NewCompanyExport(db *providers.DatabaseService) *CompanyExport {
	return &CompanyExport{
		db: db,
	}
}

func (ce *CompanyExport) ExportAll() ([]byte, error) {
	users := []map[string]interface{}{
		{
			"ID":    1,
			"Name":  "Alice",
			"Email": "alice@example.com",
			"Age":   30,
		},
		{
			"ID":    2,
			"Name":  "Bob",
			"Email": "bob@example.com",
			"Age":   25,
		},
		{
			"ID":    3,
			"Name":  "Charlie",
			"Email": "charlie@example.com",
			"Age":   35,
		},
	}
	builder := managers.NewExportExcelBuilder().
		AddSheet("Users", users).
		SetHeaders("Users", []string{"ID", "Name", "Email", "Age"}).
		SetFilters("Users", func(data interface{}) bool {
			user, ok := data.(map[string]interface{})
			if !ok {
				return false
			}
			age, ok := user["Age"].(int)
			if !ok {
				return false
			}
			return age >= 30
		}).
		EnableAutoSize("Users", true)
	f := excelize.NewFile()
	headerStyle, err := builder.CreateStyle(f, &excelize.Style{
		Font: &excelize.Font{
			Bold:  true,
			Color: "FFFFFF",
			Size:  12,
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#4F81BD"},
			Pattern: 1,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
		},
	})
	if err != nil {
		return nil, fmt.Errorf("error creating style: %v", err)
	}
	builder.SetStyles("Users", map[string]int{
		"A1": headerStyle,
		"B1": headerStyle,
		"C1": headerStyle,
		"D1": headerStyle,
	})
	excelBytes, err := builder.Build()
	if err != nil {
		return nil, fmt.Errorf("error building Excel file: %w", err)
	}
	return excelBytes, nil
}

func (ce *CompanyExport) ExportAllFiltered() {

}

func (ce *CompanyExport) ExportSelected() {

}

func (ce *CompanyExport) ExportCurrentPage() {

}
