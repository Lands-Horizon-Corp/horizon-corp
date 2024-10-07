package models

import "gorm.io/gorm"

type Feedback struct {
	Email        string `gorm:"type:varchar(255);not null;column:email"`
	Description  string `gorm:"type:text;not null;column:description"`
	FeedbackType string `gorm:"type:text;not null;column:feedback_type"`
	gorm.Model
}
