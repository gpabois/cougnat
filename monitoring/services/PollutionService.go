package services

import (
	"time"

	"github.com/gpabois/cougnat/core/result"
	"github.com/gpabois/cougnat/monitoring/models"
	"github.com/gpabois/cougnat/monitoring/repositories"
)

type PollutionService struct {
	monitoringRepo repositories.IMonitoringRepository
	polMapRepo     repositories.IPolMapRepository
}

func (svc *PollutionService) GetPollutionTiles(orgID string, sectionID []string, zoom int, begin time.Time, end time.Time) result.Result[models.PolTileCollection] {
	orgMonRes := svc.monitoringRepo.GetOrganisationMonitoring(orgID)
}
