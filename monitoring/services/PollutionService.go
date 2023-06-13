package services

import (
	"context"

	"github.com/gpabois/cougnat/core/result"
	"github.com/gpabois/cougnat/monitoring/models"
	"github.com/gpabois/cougnat/monitoring/repositories"
)

type PollutionService struct {
	monitoringRepo repositories.IMonitoringRepository
	pollutionRepo  repositories.IPollutionRepository
}

func (svc *PollutionService) AggregatePollutionMatrix(ctx context.Context, args AggregatePollutionTilesArgs) result.Result[models.PollutionMatrix] {
	orgMonRes := svc.monitoringRepo.GetOrganisationMonitoring(args.OrganisationID)
	if orgMonRes.HasFailed() {
		return result.Result[models.PollutionMatrix]{}.Failed(orgMonRes.UnwrapError())
	}
	orgMon := orgMonRes.Expect()

	svc.pollutionRepo.GetPollutionTimeSerie()
}
