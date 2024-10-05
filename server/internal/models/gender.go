package models

// Gender represents the structure of the gender entity in the application.
type Gender struct {
	ID          uint   `json:"id" gorm:"primaryKey"`        // Unique identifier for the gender
	Name        string `json:"name" gorm:"unique;not null"` // Unique name of the gender
	Description string `json:"description,omitempty"`       // Description of the gender
}
