package timesheet

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
)

type TimesheetService struct {
	db     *providers.DatabaseService
	engine *providers.EngineService
	models *models.ModelResource
}

func NewTimesheetService(
	db *providers.DatabaseService,
	engine *providers.EngineService,
	models *models.ModelResource,
) *TimesheetService {

	return &TimesheetService{

		db:     db,
		engine: engine,
		models: models,
	}
}

func (as *TimesheetService) RegisterRoutes() {
	as.engine.Client.Group("/api/v1/timesheet")
	{

	}
}
