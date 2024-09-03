package handlers

import (
	"fmt"
	"horizon-server/config"
	"horizon-server/internal/services"
	"net/http"
	"path/filepath"
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

	// Upload the file
	err = h.FileService.UploadFile(h.Config.Storage.BucketName, header.Filename, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to upload file"})
		return
	}

	// Get file details
	fileSize := header.Size
	fileFormat := filepath.Ext(header.Filename)
	fileName := header.Filename

	// Generate a presigned URL for download (optional)
	url, err := h.FileService.GeneratePresignedURL(h.Config.Storage.BucketName, fileName, 15*time.Minute)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to generate presigned URL"})
		return
	}

	// Generate a public download URL
	fileURL := h.FileService.GetPublicURL(h.Config.Storage.BucketName, fileName)

	// Return file details
	c.JSON(http.StatusOK, gin.H{
		"message":     "file uploaded successfully",
		"file_name":   fileName,
		"file_size":   fileSize,
		"file_format": fileFormat,
		"file_url":    fileURL,
		"temp_url":    url,
	})
}

func (h *FileHandler) UploadFileProgress(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file not provided"})
		return
	}
	defer file.Close()

	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	done := make(chan struct{})
	go func() {
		<-c.Request.Context().Done()
		close(done)
	}()

	var totalSize int64 = header.Size
	var uploadedSize int64

	progressCallback := func(progress int64) {
		uploadedSize += progress
		percentage := float64(uploadedSize) / float64(totalSize) * 100
		select {
		case <-done:
			return
		default:
			fmt.Fprintf(c.Writer, "data: %d\n\n", int(percentage))
			c.Writer.Flush()
		}
	}

	// Upload the file with progress tracking
	err = h.FileService.UploadFileProgress(h.Config.Storage.BucketName, header.Filename, file, progressCallback)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to upload file"})
		return
	}

	fileSize := header.Size
	fileFormat := filepath.Ext(header.Filename)
	fileName := header.Filename

	url, err := h.FileService.GeneratePresignedURL(h.Config.Storage.BucketName, fileName, 15*time.Minute)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to generate presigned URL"})
		return
	}

	// Generate a public download URL
	fileURL := h.FileService.GetPublicURL(h.Config.Storage.BucketName, fileName)

	// Final message after completion
	fmt.Fprintf(c.Writer, "data: done\n\n")
	c.Writer.Flush()

	// Return file details
	c.JSON(http.StatusOK, gin.H{
		"message":     "file uploaded successfully",
		"file_name":   fileName,
		"file_size":   fileSize,
		"file_format": fileFormat,
		"file_url":    fileURL,
		"temp_url":    url,
	})
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

	// Return the presigned URL along with key (file name)
	c.JSON(http.StatusOK, gin.H{
		"url":       url,
		"file_name": key,
	})
}
