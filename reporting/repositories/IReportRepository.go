package reporting_repositories

import (
	reporting_models "github.com/gpabois/cougnat/reporting/models"
	"github.com/gpabois/gostd/option"
	"github.com/gpabois/gostd/result"
)

//go:generate mockery
type IReportRepository interface {
	// Create a report
	Create(report reporting_models.Report) result.Result[reporting_models.ReportID]
	// Get a report
	GetById(id reporting_models.ReportID) result.Result[option.Option[reporting_models.Report]]
	// Delete a report
	Delete(reportID reporting_models.ReportID) result.Result[bool]
}
