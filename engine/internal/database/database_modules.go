package database

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database/models"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"database",

	fx.Provide(
		// Models
		models.NewAdminModel,
		models.NewBranchModel,
		models.NewCompanyModel,
		models.NewContactModel,
		models.NewEmployeeModel,
		models.NewFeedbackModel,
		models.NewFootstepModel,
		models.NewGenderModel,
		models.NewMediaModel,
		models.NewMemberModel,
		models.NewOwnerModel,
		models.NewRoleModel,
		models.NewTimesheetModel,

		// Database
	),
)
