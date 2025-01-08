package transformers

import "time"

type TimesheetResource struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	EmployeeID uint              `json:"employeeID"`
	Employee   *EmployeeResource `json:"employee"`
	TimeIn     *time.Time        `json:"timeIn"`
	TimeOut    *time.Time        `json:"timeOut"`
	MediaInID  *uint             `json:"mediaInID"`
	MediaIn    *MediaResource    `json:"mediaIn"`
	MediaOutID *uint             `json:"mediaOutID"`
	MediaOut   *MediaResource    `json:"mediaOut"`
}
