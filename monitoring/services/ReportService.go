package services

import (
	"github.com/gpabois/cougnat/core/cfg"
	"github.com/gpabois/cougnat/core/result"
	"github.com/gpabois/cougnat/core/unit"
	"github.com/gpabois/cougnat/monitoring/repositories"
	reportingModels "github.com/gpabois/cougnat/reporting/models"
)

type ReportServiceConfiguration struct {
	TimeSampling unit.Sampling
	FloorZoom    int
	CeilZoom     int
}

type ReportService struct {
	cfg           ReportServiceConfiguration
	pollutionRepo repositories.IPollutionRepository
}

func (svc *ReportService) HandleNewReport(newReport reportingModels.Report) result.Result[bool] {
	// Convert into TileIndex
	return result.Chain(func(commands []repositories.IncPollutionCommand) result.Result[bool] {
		return svc.pollutionRepo.IncPollutionTileMany(commands)
	}, repositories.GenIncPollutionCommands(
		newReport,
		svc.cfg.CeilZoom,
		svc.cfg.FloorZoom,
		svc.cfg.TimeSampling,
	))
}

func (svc *ReportService) HandleDeletedReport(deletedReport reportingModels.ReportID) result.Result[bool] {
	return result.Success(true)
}

func ProvideReportService(config *cfg.ConfigMap, pollutionRepository repositories.IPollutionRepository) IReportService {
	return &ReportService{
		cfg: ReportServiceConfiguration{
			CeilZoom: cfg.GetInt(config, "Monitoring", "CeilZoom").UnwrapOr(func() int { return 16 }),
			TimeSampling: unit.Sampling{
				Unit:   cfg.Get(config, "Monitoring", "TimeSampling", "Unit").UnwrapOr(func() string { return unit.Minute }),
				Period: cfg.GetInt(config, "Monitoring", "TimeSampling", "Period").UnwrapOr(func() int { return 1 }),
			},
			FloorZoom: cfg.GetInt(config, "Monitoring", "FloorZoom").UnwrapOr(func() int { return 11 }),
		},
		pollutionRepo: pollutionRepository,
	}
}
