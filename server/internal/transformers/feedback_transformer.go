package transformers

type FeedbackResource struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	Email        string `json:"email"`
	Description  string `json:"description"`
	FeedbackType string `json:"feedbackType"`
}
