package reporting_endpoints

import (
	"time"

	auth_models "github.com/gpabois/cougnat/auth/models"
	"github.com/gpabois/cougnat/core/geojson"
	reporting_services "github.com/gpabois/cougnat/services"
	"github.com/gpabois/gostd/option"
	"github.com/gpabois/gostd/result"
)

type ReportRequest struct {
	Owner      option.Option[auth_models.ActorID]
	Location   geojson.Feature             `serde:"location"`
	Type       reporting_models.ReportType `serde:"type"`
	Rate       int                         `serde:"rate"`
	ReportedAt time.Time                   `serde:"reported_at"`
}

type ReportResponse int

type DeleteReportRequest struct {
	ReportID string
}

type CreateReportEndpoint struct {
	Report reporting_services.IReportService
}

func ProvideReportEndpoint(report reporting_services.IReportService) *CreateReportEndpoint {
	return &CreateReportEndpoint{Report: report}
}

func (e *CreateReportEndpoint) Process(req ReportRequest) result.Result[ReportResponse] {

}
