package models

import (
	"strconv"
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/managers/filter"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type Feedback struct {
	gorm.Model

	// Fields
	Email        string `gorm:"type:varchar(255);not null" json:"email"`
	Description  string `gorm:"type:text;not null" json:"description"`
	FeedbackType string `gorm:"type:enum('bug', 'feature', 'general');not null" json:"feedback_type"`
}

type FeedbackResource struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	Email        string `json:"email"`
	Description  string `json:"description"`
	FeedbackType string `json:"feedbackType"`
}

type FeedbackRequest struct {
	Email        string `json:"email" validate:"required,email,max=255"`
	Description  string `json:"description" validate:"required"`
	FeedbackType string `json:"feedbackType" validate:"required,oneof=bug feature general"`
}

func (m *ModelResource) FeedbackToResource(feedback *Feedback) *FeedbackResource {
	if feedback == nil {
		return nil
	}
	return &FeedbackResource{
		ID:        feedback.ID,
		CreatedAt: feedback.CreatedAt.Format(time.RFC3339),
		UpdatedAt: feedback.UpdatedAt.Format(time.RFC3339),

		Email:        feedback.Email,
		Description:  feedback.Description,
		FeedbackType: feedback.FeedbackType,
	}
}

func (m *ModelResource) FeedbackToResourceList(feedbacks []*Feedback) []*FeedbackResource {
	if feedbacks == nil {
		return nil
	}
	var feedbackResources []*FeedbackResource
	for _, feedback := range feedbacks {
		feedbackResources = append(feedbackResources, m.FeedbackToResource(feedback))
	}
	return feedbackResources
}

// FeedbackToRecord converts a slice of Feedback pointers into CSV records and headers.
func (m *ModelResource) FeedbackToRecord(feedbacks []*Feedback) ([][]string, []string) {
	// Convert Feedback structs to FeedbackResource structs
	resource := m.FeedbackToResourceList(feedbacks)
	records := make([][]string, 0, len(resource))

	for _, feedback := range resource {
		// Basic Fields
		id := strconv.Itoa(int(feedback.ID))
		email := sanitizeCSVField(feedback.Email)
		description := sanitizeCSVField(feedback.Description)
		feedbackType := sanitizeCSVField(feedback.FeedbackType)
		createdAt := sanitizeCSVField(feedback.CreatedAt)
		updatedAt := sanitizeCSVField(feedback.UpdatedAt)

		// Assemble the record
		record := []string{
			id,
			email,
			description,
			feedbackType,
			createdAt,
			updatedAt,
		}
		records = append(records, record)
	}

	headers := []string{
		"ID",
		"Email",
		"Description",
		"Feedback Type",
		"Created At",
		"Updated At",
	}

	return records, headers
}

func (m *ModelResource) FeedbackFilterForAdmin(filters string) (filter.FilterPages[Feedback], error) {
	db := m.db.Client
	return m.FeedbackDB.GetPaginatedResult(db, filters)
}

func (m *ModelResource) FeedbackFilterForAdminRecord(filters string) ([]*Feedback, error) {
	db := m.db.Client
	return m.FeedbackDB.GetFilteredResults(db, filters)
}

func (m *ModelResource) ValidateFeedbackRequest(req *FeedbackRequest) error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return m.helpers.FormatValidationError(err)
	}
	return nil
}

func (m *ModelResource) FeedbackSeeders() error {
	feedbacks := []Feedback{
		{
			Email:        "user1@example.com",
			Description:  "Found a bug in the login page.",
			FeedbackType: "bug",
			Model: gorm.Model{
				CreatedAt: time.Now().AddDate(0, 0, -20),
				UpdatedAt: time.Now().AddDate(0, 0, -18),
			},
		},
		{
			Email:        "user2@example.com",
			Description:  "Can we have a dark mode feature?",
			FeedbackType: "feature",
			Model: gorm.Model{
				CreatedAt: time.Now().AddDate(0, 0, -19),
				UpdatedAt: time.Now().AddDate(0, 0, -17),
			},
		},
		{
			Email:        "user3@example.com",
			Description:  "The website is very user-friendly!",
			FeedbackType: "general",
			Model: gorm.Model{
				CreatedAt: time.Now().AddDate(0, 0, -18),
				UpdatedAt: time.Now().AddDate(0, 0, -16),
			},
		},
	}

	for i := 4; i <= 20; i++ {
		feedbacks = append(feedbacks, Feedback{
			Email:        "user" + strconv.Itoa(i) + "@example.com",
			Description:  "Sample feedback description for user " + strconv.Itoa(i) + ".",
			FeedbackType: []string{"bug", "feature", "general"}[i%3],
			Model: gorm.Model{
				CreatedAt: time.Now().AddDate(0, 0, -i),
				UpdatedAt: time.Now().AddDate(0, 0, -i+1),
			},
		})
	}

	if err := m.db.Client.Create(&feedbacks).Error; err != nil {
		return err
	}

	m.logger.Info("Feedback seeding completed successfully.")
	return nil
}
