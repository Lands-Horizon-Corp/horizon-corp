package handlers

import (
	"horizon-server/internal/models"
	"horizon-server/internal/services"
	"horizon-server/pkg/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FileHandler struct {
	fileService services.FileService
}

func NewFileHandler(fileService services.FileService) *FileHandler {
	return &FileHandler{fileService: fileService}
}
func (h *FileHandler) UploadFile(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no file is received"})
		return
	}
	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to open the file"})
		return
	}
	defer file.Close()
	fileName := helpers.UniqueFileName(fileHeader.Filename)
	fileModel := &models.File{
		FileName: fileName,
	}
	if err := h.fileService.UploadFile(fileModel, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to upload the file"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "file uploaded successfully", "file": fileModel})
}

func (h *FileHandler) DeleteFile(c *gin.Context) {
	fileID, err := helpers.GetUintParam(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid file id"})
		return
	}
	if err := h.fileService.DeleteFile(fileID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete the file"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "file deleted successfully"})
}

func (h *FileHandler) DownloadFile(c *gin.Context) {
	fileID, err := helpers.GetUintParam(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file ID"})
		return
	}
	downloadURL, err := h.fileService.DownloadFile(fileID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate download URL"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"url": downloadURL})
}

func (h *FileHandler) GetFile(c *gin.Context) {
	fileID, err := helpers.GetUintParam(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid file id"})
		return
	}
	file, err := h.fileService.GetFileByID(fileID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "file not found"})
		return
	}
	c.JSON(http.StatusOK, file)
}
