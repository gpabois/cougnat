package services

import (
	"context"
	"testing"

	auth_mocks "github.com/gpabois/cougnat/auth/mocks/services"
	auth_fixtures "github.com/gpabois/cougnat/auth/models/fixtures"
	auth_svcs "github.com/gpabois/cougnat/auth/services"
	"github.com/gpabois/cougnat/core/option"
	"github.com/gpabois/cougnat/core/rand"
	"github.com/gpabois/cougnat/core/result"
	ev "github.com/gpabois/cougnat/reporting/events"
	ev_mocks "github.com/gpabois/cougnat/reporting/mocks/events"
	repo_mocks "github.com/gpabois/cougnat/reporting/mocks/repositories"
	fixtures "github.com/gpabois/cougnat/reporting/models/fixtures"
	repos "github.com/gpabois/cougnat/reporting/repositories"
	"github.com/stretchr/testify/assert"
	"go.uber.org/dig"
)

func Test_ReportService_Report_Success(t *testing.T) {
	container := dig.New()

	// Mocks deps
	reportEvRecv := ev_mocks.NewReportEventReceiver(t)
	reportRepo := repo_mocks.NewReportRepository(t)
	authz := auth_mocks.NewAuthorizationService(t)

	// Provide the report repository
	container.Provide(func() repos.ReportRepository {
		return reportRepo
	})
	container.Provide(func() ev.ReportEventReceiver {
		return reportEvRecv
	})
	container.Provide(func() auth_svcs.AuthorizationService {
		return authz
	})

	// Provide the report service
	container.Provide(ProvideReportService)

	err := container.Invoke(func(svc ReportService) {
		// Create fixtures
		report := fixtures.RandomAnonymousReport()
		reportID, _ := rand.RandomHex(20)
		owner := auth_fixtures.RandomAnonymousID()

		// Expected report
		expectedReport := report
		expectedReport.Owner = owner
		expectedReport.ID = reportID

		// Set the ActorID
		ctx := context.WithValue(context.Background(), "CurrentActorID", owner)

		// Mock deps
		reportEvRecv.EXPECT().OnNewReport(reportID).Return()
		reportRepo.EXPECT().Create(report).Return(result.Success(reportID))
		reportRepo.EXPECT().GetById(reportID).Return(
			result.Success(
				option.Some(
					expectedReport,
				),
			),
		)

		// Report
		res := svc.Report(ctx, report)

		// Should have successfuly created the report
		assert.True(t, res.IsSuccess(), res.UnwrapError())
		assert.Equal(t, expectedReport, res.Expect())
		// Should have sent OnNewReport event to the receiver
		reportEvRecv.AssertCalled(t, "OnNewReport", reportID)
	})

	assert.Nil(t, err)
}
