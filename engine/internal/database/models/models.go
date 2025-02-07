package models

import (
	"strings"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/helpers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/managers"
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
	cfg           *config.AppConfig
	db            *providers.DatabaseService
	storage       *providers.StorageProvider
	logger        *providers.LoggerService
	helpers       *helpers.HelpersFunction
	cryptoHelpers *helpers.HelpersCryptography
	Models        []MigrateItem

	AdminDB     *managers.Repository[Admin]
	BranchDB    *managers.Repository[Branch]
	CompanyDB   *managers.Repository[Company]
	ContactDB   *managers.Repository[Contact]
	EmployeeDB  *managers.Repository[Employee]
	FeedbackDB  *managers.Repository[Feedback]
	FootstepDB  *managers.Repository[Footstep]
	GenderDB    *managers.Repository[Gender]
	MediaDB     *managers.Repository[Media]
	MemberDB    *managers.Repository[Member]
	OwnerDB     *managers.Repository[Owner]
	RoleDB      *managers.Repository[Role]
	TimesheetDB *managers.Repository[Timesheet]
}

func NewModelResource(
	cfg *config.AppConfig,
	db *providers.DatabaseService,
	storage *providers.StorageProvider,
	logger *providers.LoggerService,
	helpers *helpers.HelpersFunction,
	cryptoHelpers *helpers.HelpersCryptography,

) (*ModelResource, error) {
	modelResource := &ModelResource{
		cfg:           cfg,
		db:            db,
		storage:       storage,
		logger:        logger,
		helpers:       helpers,
		cryptoHelpers: cryptoHelpers,

		AdminDB:     managers.NewRepository[Admin](db),
		BranchDB:    managers.NewRepository[Branch](db),
		CompanyDB:   managers.NewRepository[Company](db),
		ContactDB:   managers.NewRepository[Contact](db),
		EmployeeDB:  managers.NewRepository[Employee](db),
		FeedbackDB:  managers.NewRepository[Feedback](db),
		FootstepDB:  managers.NewRepository[Footstep](db),
		GenderDB:    managers.NewRepository[Gender](db),
		MediaDB:     managers.NewRepository[Media](db),
		MemberDB:    managers.NewRepository[Member](db),
		OwnerDB:     managers.NewRepository[Owner](db),
		RoleDB:      managers.NewRepository[Role](db),
		TimesheetDB: managers.NewRepository[Timesheet](db),
	}

	modelResource.Models = []MigrateItem{
		{Model: &Admin{}, Seeder: modelResource.AdminSeeders, ModelName: "Admin"},
		{Model: &Company{}, Seeder: modelResource.CompanySeeders, ModelName: "Company"},
		{Model: &Branch{}, Seeder: modelResource.BranchSeeders, ModelName: "Branch"},
		{Model: &Contact{}, Seeder: modelResource.ContactSeeders, ModelName: "Contact"},
		{Model: &Employee{}, Seeder: modelResource.EmployeeSeeders, ModelName: "Employee"},
		{Model: &Feedback{}, Seeder: modelResource.FeedbackSeeders, ModelName: "Feedback"},
		{Model: &Footstep{}, Seeder: modelResource.FootstepSeeders, ModelName: "Footstep"},
		{Model: &Gender{}, Seeder: modelResource.GenderSeeders, ModelName: "Gender"},
		{Model: &Media{}, Seeder: modelResource.MediaSeeders, ModelName: "Media"},
		{Model: &Member{}, Seeder: modelResource.MemberSeeders, ModelName: "Member"},
		{Model: &Owner{}, Seeder: modelResource.OwnerSeeders, ModelName: "Owner"},
		{Model: &Role{}, Seeder: modelResource.RoleSeeders, ModelName: "Role"},
		{Model: &Timesheet{}, Seeder: modelResource.TimesheetSeeders, ModelName: "Timesheet"},
	}

	return modelResource, nil
}

func sanitizeCSVField(field string) string {
	if strings.HasPrefix(field, "=") || strings.HasPrefix(field, "+") ||
		strings.HasPrefix(field, "-") || strings.HasPrefix(field, "@") {
		return "'" + field
	}
	return field
}
