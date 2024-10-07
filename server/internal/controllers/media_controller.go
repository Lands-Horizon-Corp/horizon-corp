package controllers

import (
	"horizon/server/helpers"
	"horizon/server/internal/models"
	"horizon/server/internal/repositories"
	"horizon/server/internal/requests"
	"horizon/server/internal/resources"
	"horizon/server/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MediaController struct {
	repo *repositories.MediaRepository
}

func NewMediaController(repo *repositories.MediaRepository) *MediaController {
	return &MediaController{repo: repo}
}
func (h *MediaController) UploadMedia(ctx *gin.Context) {
	mediaHeader, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no file is received"})
		return
	}
	media, err := storage.UploadFile(mediaHeader)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unable to upload the file"})
		return
	}
	ctx.JSON(http.StatusCreated, resources.ToResourceMedia(*media))
}
func (h *MediaController) DeleteMedia(ctx *gin.Context) {
	storageKey := ctx.Param("id")
	err := storage.DeleteFile(storageKey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unable to delete the file"})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

func (c *MediaController) Create(ctx *gin.Context) {
	var req requests.MediaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	media := models.Media{
		FileName:   req.FileName,
		FileSize:   req.FileSize,
		FileType:   req.FileType,
		StorageKey: req.StorageKey,
		URL:        req.URL,
		BucketName: req.BucketName,
	}

	if err := c.repo.Create(&media); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, resources.ToResourceMedia(media))
}

func (c *MediaController) GetAll(ctx *gin.Context) {
	medias, err := c.repo.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resources.ToResourceListMedia(medias))
}

func (c *MediaController) GetByID(ctx *gin.Context) {
	uid, err := helpers.ParseIDParam(ctx, "id")
	if err != nil {
		return
	}

	media, err := c.repo.GetByID(uid)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Media not found"})
		return
	}
	ctx.JSON(http.StatusOK, resources.ToResourceMedia(media))
}

func (c *MediaController) Update(ctx *gin.Context) {
	var req requests.MediaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	media := models.Media{
		FileName:   req.FileName,
		FileSize:   req.FileSize,
		FileType:   req.FileType,
		StorageKey: req.StorageKey,
		URL:        req.URL,
		BucketName: req.BucketName,
	}

	if err := c.repo.Update(&media); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resources.ToResourceMedia(media))
}

func (c *MediaController) Delete(ctx *gin.Context) {
	id, err := helpers.ParseIDParam(ctx, "id")
	if err != nil {
		return
	}
	if err := c.repo.Delete(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Media not found"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func MediaRoutes(router *gin.RouterGroup, controller *MediaController) {
	group := router.Group("/medias")
	{
		group.POST("/upload", controller.UploadMedia)
		group.POST("/delete/:id", controller.UploadMedia)

		group.POST("/", controller.Create)
		group.GET("/", controller.GetAll)
		group.GET("/:id", controller.GetByID)
		group.PUT("/:id", controller.Update)
		group.DELETE("/:id", controller.Delete)
	}
}
