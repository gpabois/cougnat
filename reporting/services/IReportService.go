package services

import (
	"context"

	"github.com/gpabois/cougnat/reporting/models"
	"github.com/gpabois/gostd/result"
)

type CreateReport struct {
}

//go:generate mockery
type IReportService interface {
	Report(ctx context.Context, report models.Report) result.Result[models.Report]
	DeleteReport(ctx context.Context, reportID models.ReportID) result.Result[models.ReportID]
}
