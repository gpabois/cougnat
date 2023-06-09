package repositories

import (
	"github.com/gpabois/cougnat/core/result"
	slippy_map "github.com/gpabois/cougnat/core/slippy-map"
	"github.com/gpabois/cougnat/core/unit"
	"github.com/gpabois/cougnat/monitoring/models"
	reportingModels "github.com/gpabois/cougnat/reporting/models"
)

type IncPollutionCommand struct {
	// The area (sharding key)
	ClusterIndex models.TimeTileIndex
	// The tile
	TileIndex models.TimeTileIndex
	// The pollution data
	Report reportingModels.Report
}

type DecPollutionCommand struct {
	// The area (sharding key)
	clusterIndex models.TimeTileIndex
	// The tile
	tileIndex models.TimeTileIndex
	// The pollution data
	report reportingModels.Report
}

func NewIncPollutionCommand(cluster models.TimeTileIndex, tile models.TimeTileIndex, report reportingModels.Report) IncPollutionCommand {
	return IncPollutionCommand{
		ClusterIndex: cluster,
		TileIndex:    tile,
		Report:       report,
	}
}

func GenIncPollutionCommands(
	report reportingModels.Report,
	clusterZoom int,
	clusterTimeSampling unit.Sampling,
	tileSamplingZoom int,
	tileTimeSampling unit.Sampling,
) result.Result[[]IncPollutionCommand] {
	commands := []IncPollutionCommand{}
	loc := report.Location.Geometry

	areaIndexRes := slippy_map.TileIndexFromGeometry(loc, clusterZoom)
	if areaIndexRes.HasFailed() {
		return result.Failed[[]IncPollutionCommand](areaIndexRes.UnwrapError())
	}
	areaIndex := areaIndexRes.Expect()
	clusterIndex := models.GetTimeTileIndex(areaIndex, report.ReportedAt, clusterTimeSampling)

	for zoom := tileSamplingZoom; zoom >= clusterZoom; zoom-- {
		tileIndexRes := slippy_map.TileIndexFromGeometry(loc, zoom)
		if tileIndexRes.HasFailed() {
			return result.Failed[[]IncPollutionCommand](tileIndexRes.UnwrapError())
		}
		tileIndex := tileIndexRes.Expect()
		timeTileIndex := models.GetTimeTileIndex(tileIndex, report.ReportedAt, tileTimeSampling)

		commands = append(commands, NewIncPollutionCommand(clusterIndex, timeTileIndex, report))
	}

	return result.Success(commands)
}

//go:generate mockery
type IPolMapRepository interface {
	IncPollutionTileMany(commands []IncPollutionCommand) result.Result[bool]
	DecPollutionTileMany(index models.TimeTileIndex, reportType reportingModels.ReportType) result.Result[bool]
	GetPollutionTimeSerie(upperLeft slippy_map.TileIndex, lowerRight slippy_map.TileIndex, beginTime int, endTime int) result.Result[models.PollutionTime]
}
