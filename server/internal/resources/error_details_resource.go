package resources

import (
	"horizon/server/internal/models"
	"time"
)

type ErrorDetailResource struct {
	ID        uint   `json:"id"`
	Message   string `json:"message"`
	Name      string `json:"name"`
	Stack     string `json:"stack,omitempty"`
	Response  string `json:"response,omitempty"`
	Status    int    `json:"status,omitempty"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func ToResourceErrorDetail(errorDetails models.ErrorDetail) ErrorDetailResource {
	return ErrorDetailResource{
		ID:        errorDetails.ID,
		Message:   errorDetails.Message,
		Name:      errorDetails.Name,
		Stack:     errorDetails.Stack,
		Response:  errorDetails.Response,
		Status:    errorDetails.Status,
		CreatedAt: errorDetails.CreatedAt.Format(time.RFC3339),
		UpdatedAt: errorDetails.UpdatedAt.Format(time.RFC3339),
	}
}

func ToResourceListErrorDetail(errorDetailsList []models.ErrorDetail) []ErrorDetailResource {
	var resources []ErrorDetailResource
	for _, errorDetails := range errorDetailsList {
		resources = append(resources, ToResourceErrorDetail(errorDetails))
	}
	return resources
}
