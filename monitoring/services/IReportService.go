package services

import (
	"github.com/gpabois/cougnat/core/result"
	"github.com/gpabois/cougnat/monitoring/repositories"
	reportingModels "github.com/gpabois/cougnat/reporting/models"
)

type IReportService interface {
	// Add the report to the pollution matrix
	HandleNewReport(newReport reportingModels.Report)
	// Remove the report from the pollution matrix
	HandleDeletedReport(deletedReport reportingModels.ReportID)
}

type ReportService struct {
	atomZoom   int
	polMatRepo repositories.IPolMatRepository
}

func (svc *ReportService) HandleNewReport(newReport reportingModels.Report) result.Result[bool] {
	// Expecting Geometry to be a Point[latLng]
	loc := newReport.Location.Geometry
	// Convert into TileIndex

}
