package company

import (
	"net/http"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/helpers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/managers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/gin-gonic/gin"
)

type CompanyExport struct {
	db            *providers.DatabaseService
	modelResource *models.ModelResource
	helpers       *helpers.HelpersFunction
}

func NewCompanyExport(
	db *providers.DatabaseService,
	modelResource *models.ModelResource,
	helpers *helpers.HelpersFunction,
) *CompanyExport {
	return &CompanyExport{
		db:            db,
		modelResource: modelResource,
		helpers:       helpers,
	}
}

func (ce *CompanyExport) ExportAll(ctx *gin.Context) {
	company, err := ce.modelResource.CompanyDB.FindAll()
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to export CSV: %v", err)
		return
	}
	record, headers := ce.modelResource.CompanyToRecord(company)
	csvManager := managers.NewCSVManager()
	csvManager.SetFileName("company-export-all.csv")
	csvManager.SetHeaders(headers)
	csvManager.AddRecords(record)
	if err := csvManager.WriteCSV(ctx); err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to generate CSV: %v", err)
		return
	}
}

func (ce *CompanyExport) ExportAllFiltered(ctx *gin.Context) {
	filterParam := ctx.Query("filter")
	if filterParam == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "filter parameter is required"})
		return
	}
	company, err := ce.modelResource.CompanyFilterForAdminRecord(filterParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Companies not found."})
		return
	}
	record, headers := ce.modelResource.CompanyToRecord(company)
	csvManager := managers.NewCSVManager()
	csvManager.SetFileName("company-export-all-filtered.csv")
	csvManager.SetHeaders(headers)
	csvManager.AddRecords(record)
	if err := csvManager.WriteCSV(ctx); err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to generate CSV: %v", err)
		return
	}
}

func (ce *CompanyExport) ExportSelected(ctx *gin.Context) {
	ids, err := ce.helpers.ParseIDsFromQuery(ctx, "ids")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	company, err := ce.modelResource.CompanyDB.GetAllByIDs(ids)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve companies."})
		return
	}

	record, headers := ce.modelResource.CompanyToRecord(company)
	csvManager := managers.NewCSVManager()
	csvManager.SetFileName("company-export-all-filtered.csv")
	csvManager.SetHeaders(headers)
	csvManager.AddRecords(record)
	if err := csvManager.WriteCSV(ctx); err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to generate CSV: %v", err)
		return
	}
}
