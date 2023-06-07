package repositories

import (
	"github.com/gpabois/cougnat/core/result"
	"github.com/gpabois/cougnat/monitoring/models"
	reportingModels "github.com/gpabois/cougnat/reporting/models"
)

type IPolMapRepository interface {
	IncPollution(index models.TimeTileIndex, reportType reportingModels.ReportType) result.Result[bool]
	DecPollution(index models.TimeTileIndex, reportType reportingModels.ReportType) result.Result[bool]
}
