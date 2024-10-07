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
func (h *MediaController) Create(ctx *gin.Context) {
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
	if err := h.repo.Create(media); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, resources.ToResourceMedia(*media))
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
	uid, err := helpers.ParseIDParam(ctx, "id")
	if err != nil {
		return
	}
	media, err := c.repo.GetByID(uid)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Media not found"})
		return
	}
	err = storage.DeleteFile(media.StorageKey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unable to delete the file"})
		return
	}
	if err := c.repo.Delete(uid); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Media not found"})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

func MediaRoutes(router *gin.RouterGroup, controller *MediaController) {
	group := router.Group("/medias")
	{

		group.POST("", controller.Create)
		group.GET("", controller.GetAll)
		group.GET("/:id", controller.GetByID)
		group.PUT("/:id", controller.Update)
		group.DELETE("/:id", controller.Delete)
	}
}
