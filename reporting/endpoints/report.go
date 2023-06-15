package endpoints

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/gpabois/cougnat/core/geojson"
	"github.com/gpabois/cougnat/reporting/models"
	"github.com/gpabois/cougnat/reporting/services"
)

type CreateReportRequest struct {
	Location   geojson.Feature   `serde:"location"`
	Type       models.ReportType `serde:"type"`
	Rate       int               `serde:"rate"`
	ReportedAt time.Time         `serde:"reported_at"`
}

type CreateReportResponse models.Report

type DeleteReportRequest struct {
	ReportID string
}

type Endpoints struct {
	Report       endpoint.Endpoint
	DeleteReport endpoint.Endpoint
}

func makeReportEndpoint(reportService services.IReportService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (response any, err error) {
		req := request.(CreateReportRequest)
		newReport := models.Report{}.From(req)
		res := reportService.Report(ctx, newReport)
		return res.Unwrap()
	}
}

func makeDeleteReportEndpoint(reportService services.IReportService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (response any, err error) {
		req := request.(DeleteReportRequest)
		res := reportService.DeleteReport(ctx, req.ReportID)
		return res.Unwrap()
	}
}

func ProvideEndpoints(reportService services.IReportService) Endpoints {
	return Endpoints{
		Report:       makeReportEndpoint(reportService),
		DeleteReport: makeDeleteReportEndpoint(reportService),
	}
}
