package routes

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/modules/admin"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/modules/auth"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/modules/branch"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/modules/company"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/gin-gonic/gin"
)

type APIRoutes struct {
	logger        *providers.LoggerService
	cache         *providers.CacheService
	engineService *providers.EngineService
	router        *gin.Engine

	// Services
	adminService   *admin.AdminService
	branchService  *branch.BranchService
	companyService *company.CompanyService
	authService    *auth.AuthService
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

) *APIRoutes {
	return &APIRoutes{
		logger:        logger,
		cache:         cache,
		engineService: engineService,
		router:        engineService.Client,

		// Services
		adminService:   adminService,
		branchService:  branchService,
		companyService: companyService,
		authService:    authService,
	}
}

func (ar *APIRoutes) API() {
	ar.adminService.RegisterRoutes()
	ar.branchService.RegisterRoutes()
	ar.companyService.RegisterRoutes()
	ar.authService.RegisterRoutes()
}
