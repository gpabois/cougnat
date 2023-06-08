package services

import (
	"github.com/gpabois/cougnat/core/result"
	reportingModels "github.com/gpabois/cougnat/reporting/models"
)

type IReportService interface {
	// Add the report to the pollution matrix
	HandleNewReport(newReport reportingModels.Report) result.Result[bool]
	// Remove the report from the pollution matrix
	HandleDeletedReport(deletedReport reportingModels.ReportID) result.Result[bool]
}
