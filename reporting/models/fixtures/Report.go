package fixtures

import (
	"time"

	auth_fixtures "github.com/gpabois/cougnat/auth/models/fixtures"
	geo_fixtures "github.com/gpabois/cougnat/core/geo/fixtures"
	models "github.com/gpabois/cougnat/reporting/models"
)

func RandomReportNature() string {
	return "random"
}

func RandomAnonymousReport() models.Report {
	return models.Report{
		Owner:     auth_fixtures.RandomAnonymousID(),
		Location:  geo_fixtures.RandomPoint(),
		Nature:    RandomReportNature(),
		Rate:      5,
		CreatedAt: time.Now(),
	}
}
