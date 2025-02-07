package controllers

import (
	"fmt"
	"net/http"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api/handlers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/helpers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type MediaController struct {
	repository      *models.ModelRepository
	transformer     *models.ModelTransformer
	footstep        *handlers.FootstepHandler
	currentUser     *handlers.CurrentUser
	storageProvider *providers.StorageProvider
	helpers         *helpers.HelpersFunction
}

func NewMediaController(
	repository *models.ModelRepository,
	transformer *models.ModelTransformer,
	footstep *handlers.FootstepHandler,
	currentUser *handlers.CurrentUser,
	storageProvider *providers.StorageProvider,
	helpers *helpers.HelpersFunction,
) *MediaController {
	return &MediaController{
		repository:      repository,
		transformer:     transformer,
		footstep:        footstep,
		currentUser:     currentUser,
		storageProvider: storageProvider,
		helpers:         helpers,
	}
}

func (c *MediaController) Index(ctx *gin.Context) {

}

func (c *MediaController) Team(ctx *gin.Context) {}

func (c *MediaController) Show(ctx *gin.Context) {
	id := ctx.Param("id")
	media, err := c.repository.MediaGetByID(id, c.helpers.GetPreload(ctx)...)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Media not found"})
		return
	}
	ctx.JSON(http.StatusOK, c.transformer.MediaToResource(media))
}

type MediaStoreRequest struct {
	ID *uuid.UUID `json:"id"`

	FileName   string `json:"fileName" validate:"required,max=255"`
	FileSize   int64  `json:"fileSize" validate:"required,min=1"`
	FileType   string `json:"fileType" validate:"required,max=50"`
	StorageKey string `json:"storageKey" validate:"required,max=255"`
	URL        string `json:"url" validate:"required,url,max=255"`
	Key        string `json:"key,omitempty" validate:"max=255"`
	BucketName string `json:"bucketName,omitempty" validate:"max=255"`
}

func (c *MediaController) Store(ctx *gin.Context) {
	var req MediaStoreRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	uploaded := &models.Media{
		FileName:   req.FileName,
		FileSize:   req.FileSize,
		FileType:   req.FileType,
		StorageKey: req.StorageKey,
		URL:        req.URL,
		BucketName: req.BucketName,
	}
	mediaUpload, err := c.repository.MediaCreate(uploaded)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, c.transformer.MediaToResource(mediaUpload))

}

type MediaUpdateRequest struct {
	FileName    string `json:"fileName" validate:"required,max=255"`
	Description string `json:"description" validate:"required,max=1024"`
}

func (c *MediaController) Update(ctx *gin.Context) {
	var req MediaUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
		return
	}
	if err := validator.New().Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data", "details": err.Error()})
		return
	}
	id := ctx.Param("id")
	updatedMedia, err := c.repository.MediaUpdateByID(id, &models.Media{
		FileName: req.FileName, Description: req.Description,
	}, c.helpers.GetPreload(ctx)...)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
		return
	}
	_, err = c.footstep.Create(ctx, "Media", "Update", "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to log activity"})
		return
	}
	ctx.JSON(http.StatusOK, c.transformer.MediaToResource(updatedMedia))
}

func (c *MediaController) Destroy(ctx *gin.Context) {
	id := ctx.Param("id")
	media, err := c.repository.MediaGetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Media not found"})
		return
	}
	err = c.storageProvider.DeleteFile(media.StorageKey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unable to delete the file"})
		return
	}
	if err := c.repository.MediaDeleteByID(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Media not found"})
		return
	}
	description := fmt.Sprintf("Deleted media file: %s (%s, %d bytes)", media.FileName, media.FileType, media.FileSize)
	_, err = c.footstep.Create(ctx, "Media", "Uploade", description)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to log activity"})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

func (c *MediaController) Upload(ctx *gin.Context) {
	mediaHeader, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no file is received"})
		return
	}
	media, err := c.storageProvider.UploadFile(mediaHeader)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unable to upload the file"})
		return
	}
	uploaded := &models.Media{
		FileName:   media.FileName,
		FileSize:   media.FileSize,
		FileType:   media.FileType,
		StorageKey: media.StorageKey,
		URL:        media.URL,
		BucketName: media.BucketName,
	}
	mediaUpload, err := c.repository.MediaCreate(uploaded, c.helpers.GetPreload(ctx)...)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	description := fmt.Sprintf("Uploaded media file: %s (%s, %d bytes)", media.FileName, media.FileType, media.FileSize)
	_, err = c.footstep.Create(ctx, "Media", "Uploade", description)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to log activity"})
		return
	}
	ctx.JSON(http.StatusCreated, c.transformer.MediaToResource(mediaUpload))
}
