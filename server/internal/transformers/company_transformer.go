package transformers

type CompanyResource struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	Name            string            `json:"name"`
	Description     string            `json:"description"`
	Address         string            `json:"address"`
	Longitude       float64           `json:"longitude"`
	Latitude        float64           `json:"latitude"`
	ContactNumber   string            `json:"contactNumber"`
	OwnerID         *uint             `json:"ownerID"`
	Owner           *OwnerResource    `json:"owner"`
	MediaID         *uint             `json:"mediaID"`
	Media           *MediaResource    `json:"media"`
	IsAdminVerified bool              `json:"isAdminVerified"`
	Branches        []*BranchResource `json:"branches"`
}
