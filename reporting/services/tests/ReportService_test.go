package services_tests

import (
	"context"
	"testing"

	auth_fixtures "github.com/gpabois/cougnat/auth/models/fixtures"
	auth_svcs "github.com/gpabois/cougnat/auth/services"
	auth_mocks "github.com/gpabois/cougnat/auth/services/mocks"
	auth_utils "github.com/gpabois/cougnat/auth/utils"
	"github.com/gpabois/cougnat/core/option"
	"github.com/gpabois/cougnat/core/rand"
	"github.com/gpabois/cougnat/core/result"
	ev "github.com/gpabois/cougnat/reporting/events"
	events_mocks "github.com/gpabois/cougnat/reporting/events/mocks"
	"github.com/gpabois/cougnat/reporting/models"
	fixtures "github.com/gpabois/cougnat/reporting/models/fixtures"
	repos "github.com/gpabois/cougnat/reporting/repositories"
	repo_mocks "github.com/gpabois/cougnat/reporting/repositories/mocks"
	"github.com/gpabois/cougnat/reporting/services"
	"github.com/stretchr/testify/assert"
	"go.uber.org/dig"
)

type mockedReportDependencies struct {
	authz              *auth_mocks.IAuthorizationService
	reportEventEmitter *events_mocks.IReportEventEmitter
	reportRepo         *repo_mocks.IReportRepository
}

func generateMockedReportServiceDependencies(t *testing.T, container *dig.Container) mockedReportDependencies {
	// Create mocked dependencies
	reportEventEmitter := events_mocks.NewIReportEventEmitter(t)
	reportRepo := repo_mocks.NewIReportRepository(t)
	authz := auth_mocks.NewIAuthorizationService(t)

	// Provide the report repository
	container.Provide(func() repos.IReportRepository {
		return reportRepo
	})
	container.Provide(func() ev.IReportEventEmitter {
		return reportEventEmitter
	})
	container.Provide(func() auth_svcs.IAuthorizationService {
		return authz
	})

	return mockedReportDependencies{
		authz, reportEventEmitter, reportRepo,
	}
}

// Test a successful use of Report Function of the Report Service
// The function must :
// 1 - Store the report, calls IReportRepository.Create
// 2 - Create an owner role with [read, write] permissions, and add the role to the requester (bind to CurrentActorID)
// 3 - Emits a NewReport event by calling the IReportEmitter.OnNewReport function
func Test_ReportService_Report_Success(t *testing.T) {
	container := dig.New()

	// Provide mocked deps for the report service
	deps := generateMockedReportServiceDependencies(t, container)

	// Provide the report service
	container.Provide(services.ProvideReportService)

	// Run the test
	err := container.Invoke(func(svc services.IReportService) {
		// Create fixtures
		report := fixtures.RandomAnonymousReport()
		reportID, _ := rand.RandomHex(20)
		ownerID := auth_fixtures.RandomAnonymousID()

		// Expected report
		expectedReport := report
		expectedReport.Owner = option.Some(ownerID)
		expectedReport.ID = option.Some(reportID)

		// Set the requester
		ctx := auth_utils.SetCurrentActorID(context.Background(), ownerID)

		// Mock deps
		deps.reportEventEmitter.EXPECT().OnNewReport(expectedReport).Return(result.Success(true))
		deps.reportRepo.EXPECT().Create(report).Return(result.Success(reportID))
		deps.reportRepo.EXPECT().GetById(reportID).Return(
			result.Success(
				option.Some(
					expectedReport,
				),
			),
		)
		deps.authz.EXPECT().CreateAndAddRoleTo(ownerID, "owner", option.Some(models.ReportObjectID(reportID)), []string{"read", "write"}).Return(result.Success(true))

		// Report
		res := svc.Report(ctx, report)

		// Should have successfuly created the report
		assert.True(t, res.IsSuccess(), res.UnwrapError())
		assert.Equal(t, expectedReport, res.Expect())

		// Should have called IAuthorizationService.CreateAndAddRoleTo
		// With ownerID as the ActorID, "owner" as the role name, reportID as the
		deps.authz.AssertCalled(t, "CreateAndAddRoleTo",
			ownerID,
			"owner",
			option.Some(models.ReportObjectID(reportID)),
			[]string{"read", "write"},
		)

		// Should have sent OnNewReport event to the receiver
		deps.reportEventEmitter.AssertCalled(t, "OnNewReport", expectedReport)
	})
	assert.Nil(t, err, err)
}
