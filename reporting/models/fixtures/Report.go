package fixtures

import (
	"time"

	auth_fixtures "github.com/gpabois/cougnat/auth/models/fixtures"
	models "github.com/gpabois/cougnat/reporting/models"
	geojson "github.com/gpabois/gostd/geojson"
	geo_fixtures "github.com/gpabois/gostd/geojson/fixtures"
	"github.com/gpabois/gostd/option"
)

func RandomReportType() models.ReportType {
	return models.ReportType{
		ID:     10,
		Name:   "random_name",
		Label:  "Random name",
		Nature: "smell",
	}
}

func RandomAnonymousReport() models.Report {
	return models.Report{
		Owner:      option.Some(auth_fixtures.RandomAnonymousID()),
		Location:   geo_fixtures.RandomPoint(option.None[geojson.FeatureProperties]()),
		Type:       RandomReportType(),
		Rate:       5,
		ReportedAt: time.Now(),
	}
}

func AsNewReport(report models.Report) models.Report {
	return models.Report{
		Location:   report.Location,
		Type:       report.Type,
		Rate:       report.Rate,
		ReportedAt: report.ReportedAt,
	}
}
