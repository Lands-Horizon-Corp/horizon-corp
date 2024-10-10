package controllers

import (
	"horizon/server/helpers"
	"horizon/server/internal/models"
	"horizon/server/internal/repositories"
	"horizon/server/internal/requests"
	"horizon/server/internal/resources"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ContactsController struct {
	repo *repositories.ContactsRepository
}

func NewContactsController(repo *repositories.ContactsRepository) *ContactsController {
	return &ContactsController{repo: repo}
}

func (c *ContactsController) Create(ctx *gin.Context) {
	var req requests.ContactRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contacts := models.Contact{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Description: req.Description,
	}

	if err := c.repo.Create(&contacts); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, resources.ToResourceContact(contacts))
}

func (c *ContactsController) GetAll(ctx *gin.Context) {
	contact, err := c.repo.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resources.ToResourceListContacts(contact))
}

func (c *ContactsController) GetByID(ctx *gin.Context) {
	uid, err := helpers.ParseIDParam(ctx, "id")
	if err != nil {
		return
	}

	contacts, err := c.repo.GetByID(uid)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Contacts not found"})
		return
	}
	ctx.JSON(http.StatusOK, resources.ToResourceContact(contacts))
}

func (c *ContactsController) Update(ctx *gin.Context) {

	id, err := helpers.ParseIDParam(ctx, "id")
	if err != nil {
		return
	}

	var req requests.ContactRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contacts := models.Contact{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Description: req.Description,
	}

	if err := c.repo.Update(id, &contacts); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resources.ToResourceContact(contacts))
}

func (c *ContactsController) Delete(ctx *gin.Context) {
	id, err := helpers.ParseIDParam(ctx, "id")
	if err != nil {
		return
	}
	if err := c.repo.Delete(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Contacts not found"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
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
