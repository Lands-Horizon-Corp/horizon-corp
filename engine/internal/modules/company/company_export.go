package company

import (
	"net/http"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/managers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/gin-gonic/gin"
)

type CompanyExport struct {
	db *providers.DatabaseService
}

func NewCompanyExport(db *providers.DatabaseService) *CompanyExport {
	return &CompanyExport{
		db: db,
	}
}

func (ce *CompanyExport) ExportAll(c *gin.Context) {
	csvManager := managers.NewCSVManager()
	csvManager.SetFileName("sample.csv")
	csvManager.SetHeaders([]string{"Name", "Age", "Email", "Country"})
	records := [][]string{
		{"John Doe", "30", "john.doe@example.com", "USA"},
		{"Jane Smith", "25", "jane.smith@example.com", "Canada"},
		{"Bob Johnson", "40", "bob.johnson@example.com", "UK"},
		{"Alice Williams", "28", "alice.williams@example.com", "Australia"},
		{"Michael Brown", "35", "michael.brown@example.com", "New Zealand"},
	}
	csvManager.AddRecords(records)
	if err := csvManager.WriteCSV(c); err != nil {
		c.String(http.StatusInternalServerError, "Failed to generate CSV: %v", err)
		return
	}
}

func (ce *CompanyExport) ExportAllFiltered() {

}

func (ce *CompanyExport) ExportSelected() {

}

func (ce *CompanyExport) ExportCurrentPage() {

}
