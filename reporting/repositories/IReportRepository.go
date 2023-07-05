package reporting_repositories

import (
	"github.com/gpabois/cougnat/reporting/models"
	"github.com/gpabois/gostd/option"
	"github.com/gpabois/gostd/result"
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
