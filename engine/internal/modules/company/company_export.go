package company

import (
	"net/http"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/managers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/gin-gonic/gin"
)

type CompanyExport struct {
	db             *providers.DatabaseService
	modelsResource *models.ModelResource
}

func NewCompanyExport(
	db *providers.DatabaseService,
	modelsResource *models.ModelResource,
) *CompanyExport {
	return &CompanyExport{
		db:             db,
		modelsResource: modelsResource,
	}
}

func (ce *CompanyExport) ExportAll(ctx *gin.Context) {
	company, err := ce.modelsResource.CompanyDB.FindAll()
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to export CSV: %v", err)
		return
	}
	record, headers := ce.modelsResource.CompanyToRecord(company)
	csvManager := managers.NewCSVManager()
	csvManager.SetFileName("sample.csv")
	csvManager.SetHeaders(headers)
	csvManager.AddRecords(record)
	if err := csvManager.WriteCSV(ctx); err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to generate CSV: %v", err)
		return
	}
}

func (ce *CompanyExport) ExportAllFiltered() {

}

func (ce *CompanyExport) ExportSelected() {

}

func (ce *CompanyExport) ExportCurrentPage() {

}
