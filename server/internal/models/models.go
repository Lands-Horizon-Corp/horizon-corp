package models

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"go.uber.org/fx"
)

type UserStatus string

const (
	PendingStatus    UserStatus = "Pending"
	VerifiedStatus   UserStatus = "Verified"
	NotAllowedStatus UserStatus = "Not allowed"
)

type UserAccountType string

const (
	AdminrAccountType    UserAccountType = "Admin"
	EmployeerAccountType UserAccountType = "Employee"
	MemberAccountType    UserAccountType = "Member"
	OwnerrAccountType    UserAccountType = "Owner"
)

var Module = fx.Module(
	"models",
	fx.Provide(
	// NewAdminRepository,
	// NewBranchRepository,
	// NewCompanyRepository,
	// NewContactRepository,
	// NewEmployeeRepository,
	// NewFeedbackRepository,
	// NewFootstepRepository,
	// NewGenderRepository,
	// NewMediaRepository,
	// NewMemberRepository,
	// NewOwnerRepository,
	// NewRoleRepository,
	),
	fx.Invoke(func(
		db *providers.DatabaseService,
		logger *providers.LoggerService,
	) {
		err := db.Client.AutoMigrate(
			&Admin{},
			&Branch{},
			&Company{},
			&Contact{},
			&Employee{},
			&Feedback{},
			&Footstep{},
			&Gender{},
			&Media{},
			&Member{},
			&Owner{},
			&Role{},
			&MemberProfile{},
			&MemberClassification{},
			&MemberClassificationHistory{},
			&MemberGender{},
			&MemberGenderHistory{},
			&MemberCenter{},
			&MemberCenterHistory{},
			&MemberGroup{},
			&MemberGroupHistory{},
			&MemberEducationalAttainment{},
			&MemberEducationalAttainmentHistory{},
			&MemberOccupation{},
			&MemberOccupationHistory{},
			&MemberType{},
			&MemberTypeHistory{},
			&MemberDescription{},
			&MemberRecruits{},
			&ContactNumberReferences{},
			&MemberWallet{},
			&MemberIncome{},
			&MemberExpenses{},
			&MemberCloseRemarks{},
			&MemberJointAccounts{},
			&MemberAddress{},
			&MemberGovernmentBenefits{},
			&MutualFundsHistory{},
			&MemberAssets{},
			&MemberRelativeAccounts{},
		)
		if err != nil {
			logger.Fatal("failed to migrate database")
		}
		logger.Info("Database migration completed successfully")
	}),
)
