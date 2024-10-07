package models

import "gorm.io/gorm"

type Feedback struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Email        string `json:"email" gorm:"not null"`
	Description  string `json:"description" gorm:"not null"`
	FeedbackType string `json:"feedback_type" gorm:"not null"`
	gorm.Model
}
