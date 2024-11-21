package controllers

import (
	"horizon/server/internal/models"
	"horizon/server/internal/repositories"
	"horizon/server/internal/requests"
	"horizon/server/internal/resources"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ContactsController struct {
	*BaseController[models.Contact, *requests.ContactRequest, resources.ContactResource]
}

func NewContactsController(db *gorm.DB) *ContactsController {
	repo := repositories.NewRepository[models.Contact](db)
	baseController := NewBaseController[models.Contact, *requests.ContactRequest, resources.ContactResource](
		db,
		ToContactModel,
		resources.ToResourceContact,
		resources.ToResourceListContacts,
		UpdateContactModel,
	)

	baseController.Repo = repo
	return &ContactsController{
		BaseController: baseController,
	}
}

func ToContactModel(req *requests.ContactRequest) *models.Contact {
	return &models.Contact{
		Model:         gorm.Model{ID: req.ID},
		FirstName:     req.FirstName,
		LastName:      req.LastName,
		Email:         req.Email,
		Description:   req.Description,
		ContactNumber: req.ContactNumber,
	}
}

func UpdateContactModel(model *models.Contact, req *requests.ContactRequest) {
	model.FirstName = req.FirstName
	model.LastName = req.LastName
	model.Email = req.Email
	model.Description = req.Description
	model.ContactNumber = req.ContactNumber
}

func ContactsRoutes(router *gin.RouterGroup, controller *ContactsController) {
	group := router.Group("/contacts")
	{
		group.POST("", controller.Create)
		group.GET("", controller.GetAll)
		group.GET("/:id", controller.GetByID)
		group.PUT("/:id", controller.Update)
		group.DELETE("/:id", controller.Delete)
	}
}
