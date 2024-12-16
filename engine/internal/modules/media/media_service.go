package media

import (
	"net/http"
	"strconv"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/managers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/gin-gonic/gin"
)

type MediaService struct {
	controller *managers.Controller[models.Media, models.MediaRequest, models.MediaResource]

	db             *providers.DatabaseService
	engine         *providers.EngineService
	storageService *providers.StorageProvider
	modelsResource *models.ModelResource
	models         *models.ModelResource
}

func NewMediaService(
	db *providers.DatabaseService,
	engine *providers.EngineService,
	storageService *providers.StorageProvider,
	modelsResource *models.ModelResource,
	models *models.ModelResource,
) *MediaService {
	controller := managers.NewController(
		models.MediaDB,
		models.ValidateMediaRequest,
		models.MediaToResource,
		models.MediaToResourceList,
	)

	return &MediaService{
		controller:     controller,
		db:             db,
		engine:         engine,
		storageService: storageService,
		modelsResource: modelsResource,
		models:         models,
	}
}

func (h *MediaService) Create(ctx *gin.Context) {
	mediaHeader, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no file is received"})
		return
	}
	media, err := h.storageService.UploadFile(mediaHeader)
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
	if err := h.modelsResource.MediaDB.Create(uploaded); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, h.modelsResource.MediaToResource(uploaded))
}

func (h *MediaService) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	media, err := h.modelsResource.MediaDB.FindByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Media not found"})
		return
	}

	err = h.storageService.DeleteFile(media.StorageKey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unable to delete the file"})
		return
	}
	if err := h.modelsResource.MediaDB.Delete(uint(id)); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Media not found"})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

func (as *MediaService) RegisterRoutes() {
	routes := as.engine.Client.Group("/api/v1/media")
	{
		// Common
		routes.POST("", as.Create)
		routes.DELETE("/:id", as.Delete)

		routes.GET("/", as.controller.GetAll)
		routes.GET("/:id", as.controller.GetByID)
		routes.PUT("/:id", as.controller.Update)
	}
}
