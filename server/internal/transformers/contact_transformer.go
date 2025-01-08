package transformers

type ContactResource struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Email         string `json:"email"`
	ContactNumber string `json:"contactNumber"`
	Description   string `json:"description"`
}
