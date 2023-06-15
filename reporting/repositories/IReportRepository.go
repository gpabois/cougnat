package repositories

import (
	"github.com/gpabois/cougnat/core/option"
	result "github.com/gpabois/cougnat/core/result"
	models "github.com/gpabois/cougnat/reporting/models"
)

//go:generate mockery
type IReportRepository interface {
	// Create a report
	Create(report models.Report) result.Result[models.ReportID]
	// Get a report
	GetById(id models.ReportID) result.Result[option.Option[models.Report]]
	// Delete a report
	Delete(reportID string) result.Result[bool]
}
