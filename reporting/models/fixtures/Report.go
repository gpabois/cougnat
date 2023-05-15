package fixtures

import (
	"time"

	auth_fixtures "github.com/gpabois/cougnat/auth/models/fixtures"
	geo_fixtures "github.com/gpabois/cougnat/core/geo/fixtures"
	"github.com/gpabois/cougnat/core/option"
	models "github.com/gpabois/cougnat/reporting/models"
)

func RandomReportNature() string {
	return "random"
}

func RandomAnonymousReport() models.Report {
	return models.Report{
		Owner:     option.Some(auth_fixtures.RandomAnonymousID()),
		Location:  geo_fixtures.RandomPoint(),
		Nature:    RandomReportNature(),
		Rate:      5,
		CreatedAt: time.Now(),
	}
}
