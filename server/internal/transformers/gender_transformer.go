package transformers

type GenderResource struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	Name        string              `json:"name"`
	Description string              `json:"description"`
	Employees   []*EmployeeResource `json:"employees"`
	Members     []*MemberResource   `json:"members"`
	Owners      []*OwnerResource    `json:"owners"`
	Admins      []*AdminResource    `json:"admins"`
}
