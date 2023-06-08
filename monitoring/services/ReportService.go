package services

import (
	"github.com/gpabois/cougnat/core/cfg"
	"github.com/gpabois/cougnat/core/result"
	"github.com/gpabois/cougnat/core/unit"
	"github.com/gpabois/cougnat/monitoring/repositories"
	reportingModels "github.com/gpabois/cougnat/reporting/models"
)

type ReportServiceConfiguration struct {
	// Tile sampling
	TileTimeSampling unit.Sampling
	TileZoomSampling int
	// Cluster sampling
	ClusterTimeSampling unit.Sampling
	ClusterZoom         int
}

type ReportService struct {
	cfg        ReportServiceConfiguration
	polMapRepo repositories.IPolMapRepository
}

func (svc *ReportService) HandleNewReport(newReport reportingModels.Report) result.Result[bool] {
	// Convert into TileIndex
	return result.Chain(func(commands []repositories.IncPollutionCommand) result.Result[bool] {
		return svc.polMapRepo.IncPollutionTileMany(commands)
	}, repositories.GenIncPollutionCommands(
		newReport,
		svc.cfg.ClusterZoom,
		svc.cfg.ClusterTimeSampling,
		svc.cfg.TileZoomSampling,
		svc.cfg.TileTimeSampling,
	))
}

func (svc *ReportService) HandleDeletedReport(deletedReport reportingModels.ReportID) result.Result[bool] {
	return result.Success(true)
}

func ProvideReportService(config *cfg.ConfigMap, polMapRepo repositories.IPolMapRepository) IReportService {
	return &ReportService{
		cfg: ReportServiceConfiguration{
			TileZoomSampling: cfg.GetInt(config, "Monitoring", "TileSampling", "Zoom").UnwrapOr(func() int { return 16 }),
			TileTimeSampling: unit.Sampling{
				Unit:   cfg.Get(config, "Monitoring", "TileSampling", "TimeUnit").UnwrapOr(func() string { return unit.Minute }),
				Period: cfg.GetInt(config, "Monitoring", "TileSampling", "TimePeriod").UnwrapOr(func() int { return 5 }),
			},
			ClusterZoom: cfg.GetInt(config, "Monitoring", "Cluster", "Zoom").UnwrapOr(func() int { return 4 }),
			ClusterTimeSampling: unit.Sampling{
				Unit:   cfg.Get(config, "Monitoring", "Cluster", "TimeUnit").UnwrapOr(func() string { return unit.Year }),
				Period: cfg.GetInt(config, "Monitoring", "Cluster", "TimePeriod").UnwrapOr(func() int { return 1 }),
			},
		},
		polMapRepo: polMapRepo,
	}
}
