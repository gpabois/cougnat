package services

import (
	"context"

	"github.com/gpabois/cougnat/core/result"
	"github.com/gpabois/cougnat/monitoring/models"
	"github.com/gpabois/cougnat/monitoring/repositories"
)

type PollutionService struct {
	monitoringRepo repositories.IMonitoringRepository
	polMapRepo     repositories.IPolMapRepository
}

func (svc *PollutionService) GetPollutionTiles(ctx context.Context, args GetPollutionTilesArgs) result.Result[models.PolTileCollection] {
	orgMonRes := svc.monitoringRepo.GetOrganisationMonitoring(args.OrganisationID)
	if orgMonRes.HasFailed() {
		return result.Result[[]models.PolTile]{}.Failed(orgMonRes.UnwrapError())
	}
	orgMon := orgMonRes.Expect()
}
