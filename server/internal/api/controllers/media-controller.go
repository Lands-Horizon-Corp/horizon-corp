package controllers

import (
	"net/http"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api/handlers"
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
}

func NewMediaController(
	repository *models.ModelRepository,
	transformer *models.ModelTransformer,
	footstep *handlers.FootstepHandler,
	currentUser *handlers.CurrentUser,
	storageProvider *providers.StorageProvider,
) *MediaController {
	return &MediaController{
		repository:      repository,
		transformer:     transformer,
		footstep:        footstep,
		currentUser:     currentUser,
		storageProvider: storageProvider,
	}
}

// GET: /
// Retrieve all media files for the logged-in user along with the total size of all media files.
// Supports filtering with or without pagination. Results can be converted to records.
//
// Enhancements:
// - Apply pagination and filtering to prevent performance bottlenecks when handling large datasets.
// - Implement audit logging to track media file access (e.g., who accessed which files and when).
// - Include metadata in the response (e.g., upload date, file type, size, download count) for better usability.
// - Add support for sorting (e.g., by file name, upload date, or size) to enhance user experience.
// Endpoint: GET /api/v1/branch
func (c *MediaController) Index(ctx *gin.Context) {}

// GET: /:id
// Retrieve a specific media file by its ID.
//
// Enhancements:
// - Validate user permissions to ensure the requester is authorized to access the specific media file.
// - Include detailed metadata in the response (e.g., file size, type, upload date, download count).
// - Log access attempts for the media file to ensure auditability.
// - Handle edge cases, such as file not found or unauthorized access, with clear error messages.
func (c *MediaController) Show(ctx *gin.Context) {}

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

// PUT: /:id
// Update the name and description of a media file.
//
// Enhancements:
// - Validate that the logged-in user owns the file before allowing updates.
// - Log all updates with details such as who made the changes and what was updated.
// - Ensure atomic updates to avoid partial updates in case of errors.
// - Provide detailed error messages for invalid or unauthorized updates.
// - Add support for versioning metadata updates to allow rollback or tracking of changes.
func (c *MediaController) Update(ctx *gin.Context) {}

// DELETE: /:id
// Delete a specific media file by its ID, only if it belongs to the logged-in user.
//
// Enhancements:
// - Implement a soft delete mechanism, marking the file as "deleted" while allowing recovery if accidental.
// - Log all delete actions with details such as who deleted the file and when.
// - Validate file ownership before allowing deletion to prevent unauthorized actions.
// - Provide clear error messages for failed deletions (e.g., file not found, unauthorized).
// - Consider adding admin-level functionality to restore soft-deleted files if necessary.
func (c *MediaController) Destroy(ctx *gin.Context) {}

// GET: /team
// Retrieve media files based on user role:
//
//	Member: Access not allowed.
//	Admin: Can retrieve all media files, including those uploaded by branches, owners, employees, and members.
//	Owner: Can retrieve all media files from their owned branches, employees, and members.
//	Employee: Can retrieve all media files uploaded by employees and members within their branch.
//
// Supports filtering with or without pagination. Results can be converted to records.
// Enhancements:
// - Enforce role-based access control to prevent unauthorized access.
// - Apply pagination and filtering to manage large datasets efficiently.
// - Add sorting capabilities (e.g., by upload date, size, or name) for better usability.
// - Implement audit logging to track who accessed media files and under what context.
// - Ensure sensitive data is excluded from the response based on role-specific access.
func (c *MediaController) Team(ctx *gin.Context) {}

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
	mediaUpload, err := c.repository.MediaCreate(uploaded)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, c.transformer.MediaToResource(mediaUpload))
}
