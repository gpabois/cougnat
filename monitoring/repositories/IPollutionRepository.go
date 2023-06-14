package repositories

import (
	"github.com/gpabois/cougnat/core/result"
	slippy_map "github.com/gpabois/cougnat/core/slippy-map"
	time_serie "github.com/gpabois/cougnat/core/time-serie"
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
	// The tile coordinates
	tileIndex models.TimeTileIndex
	// The pollution data
	report reportingModels.Report
}

func NewIncPollutionCommand(tile models.TimeTileIndex, report reportingModels.Report) IncPollutionCommand {
	return IncPollutionCommand{
		TileIndex: tile,
		Report:    report,
	}
}

func GenIncPollutionCommands(
	report reportingModels.Report,
	ceilZoom int,
	floorZoom int,
	timeSampling unit.Sampling,
) result.Result[[]IncPollutionCommand] {
	commands := []IncPollutionCommand{}
	loc := report.Location.Geometry

	for zoom := floorZoom; zoom >= ceilZoom; zoom-- {
		tileIndexResult := slippy_map.TileIndexFromGeometry(loc, zoom)

		if tileIndexResult.HasFailed() {
			return result.Failed[[]IncPollutionCommand](tileIndexResult.UnwrapError())
		}

		tileIndex := tileIndexResult.Expect()
		timeTileIndex := models.GetTimeTileIndex(tileIndex, report.ReportedAt, timeSampling)

		commands = append(commands, NewIncPollutionCommand(timeTileIndex, report))
	}

	return result.Success(commands)
}

//go:generate mockery
type IPollutionRepository interface {
	IncPollutionTileMany(commands []IncPollutionCommand) result.Result[bool]
	DecPollutionTileMany(index models.TimeTileIndex, reportType reportingModels.ReportType) result.Result[bool]
	GetPollutionTiles(tileBounds slippy_map.TileBounds, timeBounds time_serie.TimeInterval) result.Result[models.PollutionTileCollection]
}
