package controllers

import (
	"horizon/server/internal/models"
	"horizon/server/internal/repositories"
	"horizon/server/internal/requests"
	"horizon/server/internal/resources"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompaniesController struct {
	*BaseController[models.Company, *requests.CompanyRequest, resources.CompanyResource]
}

func NewCompaniesController(db *gorm.DB) *CompaniesController {
	repo := repositories.NewRepository[models.Company](db)
	baseController := NewBaseController(
		db,
		ToCompanyModel,
		resources.ToResourceCompany,
		resources.ToResourceListCompanies,
		UpdateCompanyModel,
	)

	baseController.Repo = repo
	return &CompaniesController{
		BaseController: baseController,
	}
}

func ToCompanyModel(req *requests.CompanyRequest) *models.Company {
	return &models.Company{
		Model:           gorm.Model{ID: req.ID},
		Name:            req.Name,
		Description:     req.Description,
		Address:         req.Address,
		Longitude:       req.Longitude,
		Latitude:        req.Latitude,
		ContactNumber:   req.ContactNumber,
		OwnerID:         req.OwnerID,
		MediaID:         req.MediaID,
		IsAdminVerified: req.IsAdminVerified,
	}
}

func UpdateCompanyModel(model *models.Company, req *requests.CompanyRequest) {
	model.Name = req.Name
	model.Description = req.Description
	model.Address = req.Address
	model.Longitude = req.Longitude
	model.Latitude = req.Latitude
	model.ContactNumber = req.ContactNumber
	model.OwnerID = req.OwnerID
	model.MediaID = req.MediaID
	model.IsAdminVerified = req.IsAdminVerified
}

func CompaniesRoutes(router *gin.RouterGroup, controller *CompaniesController) {
	group := router.Group("/companies")
	{
		group.POST("", controller.Create)
		group.GET("", controller.GetAll)
		group.GET("/:id", controller.GetByID)
		group.PUT("/:id", controller.Update)
		group.DELETE("/:id", controller.Delete)
		group.GET("/search", controller.Filter)

		group.GET("/export", controller.ExportAll)
		group.GET("/export-search", controller.ExportAllFiltered)
		group.GET("/export-selected", controller.ExportSelected)
		group.GET("/export-current-page/:page", controller.ExportCurrentPage)

	}
}
