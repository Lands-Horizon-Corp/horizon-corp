package transformers

type BranchResource struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	Name            string              `json:"name"`
	Address         string              `json:"address"`
	Longitude       float64             `json:"longitude"`
	Latitude        float64             `json:"latitude"`
	Email           string              `json:"email"`
	ContactNumber   string              `json:"contactNumber"`
	IsAdminVerified bool                `json:"isAdminVerified"`
	MediaID         *uint               `json:"mediaID"`
	Media           *MediaResource      `json:"media"`
	CompanyID       *uint               `json:"companyID"`
	Company         *CompanyResource    `json:"company"`
	Employees       []*EmployeeResource `json:"employees"`
	Members         []*MemberResource   `json:"members"`
}
