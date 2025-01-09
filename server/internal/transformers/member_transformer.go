package transformers

import "time"

type MemberResource struct {
	ID                 uint               `json:"id"`
	CreatedAt          string             `json:"createdAt"`
	UpdatedAt          string             `json:"updatedAt"`
	FirstName          string             `json:"firstName"`
	LastName           string             `json:"lastName"`
	MiddleName         string             `json:"middleName,omitempty"`
	PermanentAddress   string             `json:"permanentAddress,omitempty"`
	Description        string             `json:"description,omitempty"`
	BirthDate          time.Time          `json:"birthDate"`
	Username           string             `json:"username"`
	Email              string             `json:"email"`
	IsEmailVerified    bool               `json:"isEmailVerified"`
	IsContactVerified  bool               `json:"isContactVerified"`
	IsSkipVerification bool               `json:"isSkipVerification"`
	ContactNumber      string             `json:"contactNumber"`
	Status             string             `json:"status"`
	Media              *MediaResource     `json:"media,omitempty"`
	Branch             *BranchResource    `json:"branch,omitempty"`
	Longitude          *float64           `json:"longitude,omitempty"`
	Latitude           *float64           `json:"latitude,omitempty"`
	Role               *RoleResource      `json:"role,omitempty"`
	Gender             *GenderResource    `json:"gender,omitempty"`
	Footsteps          []FootstepResource `json:"footsteps,omitempty"`
}

type MemberProfileResource struct {
	ID                    uint                                 `json:"id"`
	CreatedAt             string                               `json:"createdAt"`
	UpdatedAt             string                               `json:"updatedAt"`
	Description           string                               `json:"description"`
	Notes                 string                               `json:"notes,omitempty"`
	ContactNumber         string                               `json:"contactNumber"`
	OldReferenceID        string                               `json:"oldReferenceId,omitempty"`
	Status                string                               `json:"status"`
	PassbookNumber        string                               `json:"passbookNumber,omitempty"`
	IsClosed              bool                                 `json:"isClosed"`
	Occupation            string                               `json:"occupation,omitempty"`
	BusinessAddress       string                               `json:"businessAddress,omitempty"`
	BusinessContact       string                               `json:"businessContact,omitempty"`
	TinNumber             string                               `json:"tinNumber,omitempty"`
	CivilStatus           string                               `json:"civilStatus"`
	SSSNumber             string                               `json:"sssNumber,omitempty"`
	PagibigNumber         string                               `json:"pagibigNumber,omitempty"`
	PhilhealthNumber      string                               `json:"philhealthNumber,omitempty"`
	IsMutualFundMember    bool                                 `json:"isMutualFundMember"`
	IsMicroFinanceMember  bool                                 `json:"isMicroFinanceMember"`
	IsDosriMember         bool                                 `json:"isDosriMember"`
	MemberType            *MemberTypeResource                  `json:"memberType,omitempty"`
	Classification        *MemberClassificationResource        `json:"classification,omitempty"`
	EducationalAttainment *MemberEducationalAttainmentResource `json:"educationalAttainment,omitempty"`
	Group                 *MemberGroupResource                 `json:"group,omitempty"`
	Center                *MemberCenterResource                `json:"center,omitempty"`
	OccupationDetail      *MemberOccupationResource            `json:"occupationDetail,omitempty"`
	Addresses             []MemberAddressResource              `json:"addresses,omitempty"`
	Media                 *MediaResource                       `json:"media,omitempty"`
}

type MemberClassificationResource struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	Name        string `json:"name"`
	Description string `json:"description"`
}

type MemberEducationalAttainmentResource struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	Name        string `json:"name"`
	Description string `json:"description"`
}

type MemberGroupResource struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	Name        string `json:"name"`
	Description string `json:"description"`
}

type MemberCenterResource struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	Name        string `json:"name"`
	Description string `json:"description"`
}

type MemberOccupationResource struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	Name        string `json:"name"`
	Description string `json:"description"`
}

type MemberAddressResource struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	PostalCode string `json:"postalCode"`
	Province   string `json:"province"`
	City       string `json:"city"`
	Barangay   string `json:"barangay"`
	Region     string `json:"region"`
	Label      string `json:"label"`
}

type MemberBranchRegistrationResource struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	Member     *MemberResource   `json:"member"`
	Branch     *BranchResource   `json:"branch"`
	Status     string            `json:"status"`
	Remarks    string            `json:"remarks"`
	VerifiedBy *EmployeeResource `json:"verifiedBy,omitempty"`
	VerifiedAt time.Time         `json:"verifiedAt,omitempty"`
}

type MemberHistoryResource struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	MemberID    uint   `json:"memberId"`
	Description string `json:"description"`
}

type MemberIncomeResource struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	Name        string    `json:"name"`
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
}

type MemberExpensesResource struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	Name        string    `json:"name"`
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
}

// MemberCloseRemarks
type MemberCloseRemarksResource struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type MemberWalletResource struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	Debit       float64   `json:"debit"`
	Credit      float64   `json:"credit"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
}

type MemberAssetsResource struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	Name        string    `json:"name"`
	EntryDate   time.Time `json:"entryDate"`
	Description string    `json:"description"`
}

type MemberTypeResource struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	Name        string `json:"name"`
	Description string `json:"description"`
	Prefix      string `json:"prefix"`
}
