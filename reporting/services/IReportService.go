package reporting_services

import (
	auth_models "github.com/gpabois/cougnat/auth/models"
	reporting_models "github.com/gpabois/cougnat/reporting/models"
	"github.com/gpabois/gostd/result"
)

type ReportRequest = reporting_models.NewReport
type ReportResponse struct {
	ReportID reporting_models.ReportID `serde:"report_id"`
}
type ReportResult = result.Result[ReportResponse]

type DeleteReportRequest struct {
	Requester auth_models.ActorID       `serde:"requester"`
	ReportID  reporting_models.ReportID `serde:"report_id"`
}
type DeleteReportResponse struct {
	Result bool `serde:"result"`
}
type DeleteReportResult = result.Result[DeleteReportResponse]

//go:generate mockery
type IReportService interface {
	Report(request ReportRequest) ReportResult
	DeleteReport(request DeleteReportRequest) result.Result[DeleteReportResponse]
}
