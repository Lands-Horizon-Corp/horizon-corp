package handlers

import (
	"horizon-server/config"
	"horizon-server/internal/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type FileHandler struct {
	FileService *services.FileService
	Config      *config.Config
}

func NewFileHandler(fileService *services.FileService, cfg *config.Config) *FileHandler {
	return &FileHandler{FileService: fileService, Config: cfg}
}

func (h *FileHandler) UploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file not provided"})
		return
	}
	defer file.Close()

	err = h.FileService.UploadFile(h.Config.Storage.BucketName, header.Filename, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to upload file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "file uploaded successfully"})
}

func (h *FileHandler) DeleteFile(c *gin.Context) {
	key := c.Param("key")

	err := h.FileService.DeleteFile(h.Config.Storage.BucketName, key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to delete file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "file deleted successfully"})
}

func (h *FileHandler) GeneratePresignedURL(c *gin.Context) {
	key := c.Param("key")

	url, err := h.FileService.GeneratePresignedURL(h.Config.Storage.BucketName, key, 15*time.Minute)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to generate URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"url": url})
}
