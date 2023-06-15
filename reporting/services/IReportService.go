package services

import (
	"context"

	"github.com/gpabois/cougnat/core/result"
	"github.com/gpabois/cougnat/reporting/models"
)

//go:generate mockery
type IReportService interface {
	Report(ctx context.Context, report models.Report) result.Result[models.Report]
	DeleteReport(ctx context.Context, reportID models.ReportID) result.Result[bool]
}
