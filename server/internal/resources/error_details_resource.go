package resources

import (
	"horizon/server/internal/models"
	"time"
)

type ErrorDetailResource struct {
	ID       uint   `json:"id"`
	Message  string `json:"message"`
	Name     string `json:"name"`
	Stack    string `json:"stack"`
	Response string `json:"response"`
	Status   int    `json:"status"`

	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func ToResourceErrorDetail(errorDetail models.ErrorDetail) ErrorDetailResource {
	return ErrorDetailResource{
		ID:       errorDetail.ID,
		Message:  errorDetail.Message,
		Name:     errorDetail.Name,
		Stack:    errorDetail.Stack,
		Response: errorDetail.Response,
		Status:   errorDetail.Status,

		CreatedAt: errorDetail.CreatedAt.Format(time.RFC3339),
		UpdatedAt: errorDetail.UpdatedAt.Format(time.RFC3339),
	}
}

func ToResourceListErrorDetail(errorDetails []models.ErrorDetail) []ErrorDetailResource {
	resourceList := make([]ErrorDetailResource, len(errorDetails))
	for i, errorDetail := range errorDetails {
		resourceList[i] = ToResourceErrorDetail(errorDetail)
	}
	return resourceList
}
