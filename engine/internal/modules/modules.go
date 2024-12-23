package modules

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/modules/admin"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/modules/auth"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/modules/branch"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/modules/company"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/modules/contact"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/modules/employee"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/modules/feedback"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/modules/footstep"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/modules/gender"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/modules/media"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/modules/member"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/modules/owner"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/modules/role"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/modules/timesheet"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"modules",
	admin.Module,
	auth.Module,
	branch.Module,
	company.Module,
	contact.Module,
	employee.Module,
	feedback.Module,
	footstep.Module,
	gender.Module,
	media.Module,
	member.Module,
	owner.Module,
	role.Module,
	timesheet.Module,
)
