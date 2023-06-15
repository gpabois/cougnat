package services_tests

import (
	"testing"
	"time"

	"github.com/gpabois/cougnat/core/cfg"
	"github.com/gpabois/cougnat/core/result"
	"github.com/gpabois/cougnat/core/unit"
	"github.com/gpabois/cougnat/monitoring/repositories"
	mockedRepos "github.com/gpabois/cougnat/monitoring/repositories/mocks"
	"github.com/gpabois/cougnat/monitoring/services"
	"github.com/gpabois/cougnat/reporting/models/fixtures"
	"github.com/stretchr/testify/assert"
	"go.uber.org/dig"
)

func Test_ReportService_HandleNewReport_Success(t *testing.T) {
	container := dig.New()

	// Provide a Mocked Pollution Repository
	pollutionRepo := mockedRepos.NewIPollutionRepository(t)
	container.Provide(func() repositories.IPollutionRepository {
		return pollutionRepo
	})

	// Provide Config Map
	const ceilZoom = 4
	const floorZoom = 5
	timeSampling := unit.Sampling{
		Period: 1,
		Unit:   unit.Minute,
	}

	container.Provide(func() *cfg.ConfigMap {
		return &cfg.ConfigMap{
			"Monitoring": cfg.ConfigMap{
				"CeilZoom":  "6",
				"FloorZoom": "4",
				"TimeSamplig": cfg.ConfigMap{
					"Unit":   unit.Minute,
					"Period": "1",
				},
			},
		}
	})

	// Provide the report service
	container.Provide(services.ProvideReportService)

	// Run the test
	err := container.Invoke(func(svc services.IReportService) {
		now := time.Now()

		report := fixtures.RandomAnonymousReport()
		report.ReportedAt = now

		// Prepare the mocked repository
		expectedCommands := repositories.GenIncPollutionCommands(report, ceilZoom, floorZoom, timeSampling).Expect()
		pollutionRepo.EXPECT().IncPollutionTileMany(expectedCommands).Return(result.Success(true))

		// Call the HandleNewReport function
		res := svc.HandleNewReport(report)

		// Should be true
		assert.True(t, res.IsSuccess(), res.UnwrapError())
	})

	assert.Nil(t, err)
}
