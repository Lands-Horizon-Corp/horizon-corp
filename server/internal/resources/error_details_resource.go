package resources

import (
	"horizon/server/helpers"
	"horizon/server/internal/models"
	"time"
)

// ErrorDetailsResource represents the structure of the error details resource for API responses.
type ErrorDetailsResource struct {
	ID        uint   `json:"id"`
	Message   string `json:"message"`
	Name      string `json:"name"`
	Stack     string `json:"stack,omitempty"`
	Response  string `json:"response,omitempty"`
	Status    int    `json:"status,omitempty"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// ToResource converts an ErrorDetails model to an ErrorDetailsResource.
func ToResourceErrorDetails(errorDetails models.ErrorDetails) ErrorDetailsResource {
	return ErrorDetailsResource{
		ID:        errorDetails.ID,
		Message:   errorDetails.Message,
		Name:      errorDetails.Name,
		Stack:     helpers.SafeString(errorDetails.Stack),
		Response:  helpers.SafeString(errorDetails.Response),
		Status:    helpers.SafeInt(errorDetails.Status),
		CreatedAt: errorDetails.CreatedAt.Format(time.RFC3339),
		UpdatedAt: errorDetails.UpdatedAt.Format(time.RFC3339),
	}
}

// ToResourceList converts a slice of ErrorDetails models to a slice of ErrorDetailsResource.
func ToResourceListErrorDetails(errorDetailsList []models.ErrorDetails) []ErrorDetailsResource {
	var resources []ErrorDetailsResource
	for _, errorDetails := range errorDetailsList {
		resources = append(resources, ToResourceErrorDetails(errorDetails))
	}
	return resources
}
