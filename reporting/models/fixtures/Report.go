package fixtures

import (
	"time"

	auth_fixtures "github.com/gpabois/cougnat/auth/models/fixtures"
	geojson "github.com/gpabois/cougnat/core/geojson"
	geo_fixtures "github.com/gpabois/cougnat/core/geojson/fixtures"
	"github.com/gpabois/cougnat/core/option"
	models "github.com/gpabois/cougnat/reporting/models"
)

func RandomReportType() models.ReportType {
	return models.ReportType{
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
