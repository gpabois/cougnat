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

	mockedPolMapRepo := mockedRepos.NewIPolMapRepository(t)
	// Provide the PolMap Repository
	container.Provide(func() repositories.IPolMapRepository {
		return mockedPolMapRepo
	})

	// Provide Config Map
	const clusterZoom = 4
	clusterTimeSampling := unit.Sampling{
		Period: 1,
		Unit:   unit.Year,
	}
	const tileSamplingZoom = 5
	tileTimeSampling := unit.Sampling{
		Period: 1,
		Unit:   unit.Minute,
	}

	container.Provide(func() *cfg.ConfigMap {
		return &cfg.ConfigMap{
			"Monitoring": cfg.ConfigMap{
				"TileSampling": cfg.ConfigMap{
					"Zoom":       "6",
					"TimeUnit":   unit.Minute,
					"TimePeriod": "1",
				},
				"Cluster": cfg.ConfigMap{
					"Zoom":       "4",
					"TimeUnit":   unit.Year,
					"TimePeriod": "1",
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
		expectedCommands := repositories.GenIncPollutionCommands(report, clusterZoom, clusterTimeSampling, tileSamplingZoom, tileTimeSampling).Expect()
		mockedPolMapRepo.EXPECT().IncPollutionTileMany(expectedCommands).Return(result.Success(true))

		// Call the HandleNewReport function
		res := svc.HandleNewReport(report)

		// Should be true
		assert.True(t, res.IsSuccess(), res.UnwrapError())
	})

	assert.Nil(t, err)
}