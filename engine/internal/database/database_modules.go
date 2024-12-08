package database

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database/models"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
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
	),
	fx.Invoke(func(lc fx.Lifecycle, db *gorm.DB, logger *zap.Logger) {
		Migrate(lc, db, logger)
	}),
)
