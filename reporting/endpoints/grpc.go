package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	geo "github.com/gpabois/cougnat/core/geojson"
	"github.com/gpabois/cougnat/reporting/models"
	"github.com/gpabois/cougnat/reporting/services"
)

type Endpoints struct {
	Report endpoint.Endpoint
}

type ReportRequest struct {
	Nature   string
	Rate     uint8
	Location struct {
		P0 float64
		P1 float64
	}
}

type ReportResponse struct {
	ReportID string
}

func MakeEndpoints(s services.ReportService) Endpoints {
	return Endpoints{
		Report: makeReportEndpoint(s),
	}
}

func makeReportEndpoint(svc services.ReportService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (response any, err error) {
		req := request.(ReportRequest)

		ret := svc.Report(ctx, models.Report{
			Nature:   req.Nature,
			Rate:     int(req.Rate),
			Location: geo.NewPoint(req.Location.P0, req.Location.P1),
		})

		return ret.IntoAnyTuple()
	}
}
