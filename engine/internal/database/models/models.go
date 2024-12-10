package models

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
)

type UserStatus string

const (
	AdminPending    UserStatus = "Pending"
	AdminVerified   UserStatus = "Verified"
	AdminNotAllowed UserStatus = "Not allowed"
)

type MigrateItem struct {
	Model     interface{}
	Seeder    func() error
	ModelName string
}

type ModelResource struct {
	db      *providers.DatabaseService
	storage *providers.StorageProvider
	logger  *providers.LoggerService
	Models  []MigrateItem
}

func NewModelResource(
	db *providers.DatabaseService,
	storage *providers.StorageProvider,
	logger *providers.LoggerService,
) (*ModelResource, error) {

	modelResource := &ModelResource{
		db:      db,
		storage: storage,
		logger:  logger,
	}

	migrateItems := []MigrateItem{
		{
			Model:     &Admin{},
			Seeder:    modelResource.AdminSeeders,
			ModelName: "Admin",
		},
		{
			Model:     &Branch{},
			Seeder:    modelResource.BranchSeeders,
			ModelName: "Branch",
		},
		{
			Model:     &Company{},
			Seeder:    modelResource.CompanySeeders,
			ModelName: "Company",
		},
		{
			Model:     &Contact{},
			Seeder:    modelResource.ContactSeeders,
			ModelName: "Contact",
		},
		{
			Model:     &Employee{},
			Seeder:    modelResource.EmployeeSeeders,
			ModelName: "Employee",
		},
		{
			Model:     &Feedback{},
			Seeder:    modelResource.FeedbackSeeders,
			ModelName: "Feedback",
		},
		{
			Model:     &Footstep{},
			Seeder:    modelResource.FootstepSeeders,
			ModelName: "Footstep",
		},
		{
			Model:     &Gender{},
			Seeder:    modelResource.GenderSeeders,
			ModelName: "Gender",
		},
		{
			Model:     &Media{},
			Seeder:    modelResource.MediaSeeders,
			ModelName: "Media",
		},
		{
			Model:     &Member{},
			Seeder:    modelResource.MemberSeeders,
			ModelName: "Member",
		},
		{
			Model:     &Owner{},
			Seeder:    modelResource.OwnerSeeders,
			ModelName: "Owner",
		},
		{
			Model:     &Role{},
			Seeder:    modelResource.RoleSeeders,
			ModelName: "Role",
		},
		{
			Model:     &Timesheet{},
			Seeder:    modelResource.TimesheetSeeders,
			ModelName: "Timesheet",
		},
	}

	modelResource.Models = migrateItems

	return modelResource, nil
}
