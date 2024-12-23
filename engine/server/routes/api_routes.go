package routes

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
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/gin-gonic/gin"
)

type APIRoutes struct {
	logger        *providers.LoggerService
	cache         *providers.CacheService
	engineService *providers.EngineService
	router        *gin.Engine

	// Services
	adminService     *admin.AdminService
	branchService    *branch.BranchService
	companyService   *company.CompanyService
	authService      *auth.AuthService
	contactService   *contact.ContactService
	employeeService  *employee.EmployeeService
	feedbackService  *feedback.FeedbackService
	footstepService  *footstep.FootstepService
	genderservice    *gender.GenderService
	mediaService     *media.MediaService
	memberService    *member.MemberService
	ownerService     *owner.OwnerService
	roleService      *role.RoleService
	timesheetService *timesheet.TimesheetService
}

func NewAPIRoutes(
	logger *providers.LoggerService,
	cache *providers.CacheService,
	engineService *providers.EngineService,

	// Services
	adminService *admin.AdminService,
	branchService *branch.BranchService,
	companyService *company.CompanyService,
	authService *auth.AuthService,
	contactService *contact.ContactService,
	employeeService *employee.EmployeeService,
	feedbackService *feedback.FeedbackService,
	footstepService *footstep.FootstepService,
	genderservice *gender.GenderService,
	mediaService *media.MediaService,
	memberService *member.MemberService,
	ownerService *owner.OwnerService,
	roleService *role.RoleService,
	timesheetService *timesheet.TimesheetService,

) *APIRoutes {
	return &APIRoutes{
		logger:        logger,
		cache:         cache,
		engineService: engineService,
		router:        engineService.Client,

		// Services
		adminService:     adminService,
		branchService:    branchService,
		companyService:   companyService,
		authService:      authService,
		contactService:   contactService,
		employeeService:  employeeService,
		feedbackService:  feedbackService,
		footstepService:  footstepService,
		genderservice:    genderservice,
		mediaService:     mediaService,
		memberService:    memberService,
		ownerService:     ownerService,
		roleService:      roleService,
		timesheetService: timesheetService,
	}
}

func (ar *APIRoutes) API() {
	ar.adminService.RegisterRoutes()
	ar.branchService.RegisterRoutes()
	ar.companyService.RegisterRoutes()
	ar.authService.RegisterRoutes()
	ar.contactService.RegisterRoutes()
	ar.employeeService.RegisterRoutes()
	ar.footstepService.RegisterRoutes()
	ar.genderservice.RegisterRoutes()
	ar.mediaService.RegisterRoutes()
	ar.memberService.RegisterRoutes()
	ar.ownerService.RegisterRoutes()
	ar.roleService.RegisterRoutes()
	ar.timesheetService.RegisterRoutes()
}
