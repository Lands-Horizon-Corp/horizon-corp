package transformers

type RoleResource struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	Name               string `json:"name"`
	Description        string `json:"description"`
	ApiKey             string `json:"apiKey"`
	Color              string `json:"color"`
	ReadRole           bool   `json:"readRole"`
	WriteRole          bool   `json:"writeRole"`
	UpdateRole         bool   `json:"updateRole"`
	DeleteRole         bool   `json:"deleteRole"`
	ReadErrorDetails   bool   `json:"readErrorDetails"`
	WriteErrorDetails  bool   `json:"writeErrorDetails"`
	UpdateErrorDetails bool   `json:"updateErrorDetails"`
	DeleteErrorDetails bool   `json:"deleteErrorDetails"`
	ReadGender         bool   `json:"readGender"`
	WriteGender        bool   `json:"writeGender"`
	UpdateGender       bool   `json:"updateGender"`
	DeleteGender       bool   `json:"deleteGender"`

	Admins    []*AdminResource    `json:"admins"`
	Owners    []*OwnerResource    `json:"owners"`
	Employees []*EmployeeResource `json:"employees"`
	Members   []*MemberResource   `json:"members"`
}
