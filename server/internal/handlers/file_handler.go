package handlers

import (
	"encoding/json"
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

	err = h.FileService.UploadFile(h.Config.Storage.BucketName, header.Filename, file)
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

	fileURL := h.FileService.GetPublicURL(h.Config.Storage.BucketName, fileName)

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
	fileName := header.Filename

	progressCallback := func(fileName string, fileSize int64, progressBytes int64, progressPercentage float64) {
		select {
		case <-done:
			return
		default:
			progressData := gin.H{
				"file_name":      fileName,
				"file_size":      fileSize,
				"progress_bytes": progressBytes,
				"progress":       progressPercentage,
			}

			jsonData, err := json.Marshal(progressData)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to marshal progress data"})
				return
			}

			fmt.Fprintf(c.Writer, "data: %s\n\n", jsonData)
			c.Writer.Flush()
		}
	}

	err = h.FileService.UploadFileProgress(h.Config.Storage.BucketName, fileName, file, fileName, totalSize, progressCallback)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to upload file"})
		return
	}

	// Generate presigned URL for the uploaded file
	url, err := h.FileService.GeneratePresignedURL(h.Config.Storage.BucketName, fileName, 15*time.Minute)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to generate presigned URL"})
		return
	}

	// Generate a public download URL
	fileURL := h.FileService.GetPublicURL(h.Config.Storage.BucketName, fileName)

	// Final message after completion
	fmt.Fprintf(c.Writer, "data: {\"status\":\"done\"}\n\n")
	c.Writer.Flush()

	// Return file details in the final response
	c.JSON(http.StatusOK, gin.H{
		"message":     "file uploaded successfully",
		"file_name":   fileName,
		"file_size":   totalSize,
		"file_format": filepath.Ext(fileName),
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
