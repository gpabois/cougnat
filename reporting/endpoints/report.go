package reporting_endpoints

import (
	reporting_services "github.com/gpabois/cougnat/reporting/services"
	"github.com/gpabois/gostd/result"
)

type ReportRequest = reporting_services.ReportRequest
type ReportResponse = reporting_services.ReportResponse

type CreateReportEndpoint struct {
	svc reporting_services.IReportService
}

func ProvideReportEndpoint(report reporting_services.IReportService) *CreateReportEndpoint {
	return &CreateReportEndpoint{svc: report}
}

func (e *CreateReportEndpoint) Process(req ReportRequest) result.Result[ReportResponse] {
	return e.svc.Report(req)
}
