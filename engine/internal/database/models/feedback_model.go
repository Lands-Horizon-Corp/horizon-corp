package models

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
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
	Email        string `json:"email"`
	Description  string `json:"description"`
	FeedbackType string `json:"feedbackType"`
}

type FeedbackModel struct {
	lc     *fx.Lifecycle
	db     *gorm.DB
	logger *zap.Logger
}

func NewFeedbackModel(
	lc *fx.Lifecycle,
	db *gorm.DB,
	logger *zap.Logger,
) *FeedbackModel {
	return &FeedbackModel{
		lc:     lc,
		db:     db,
		logger: logger,
	}
}

func (fm *FeedbackModel) SeedDatabase() {
}

func (fm *FeedbackModel) ToResource() {
}

func (fm *FeedbackModel) ToResourceList() {
}
