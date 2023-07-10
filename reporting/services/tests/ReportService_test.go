package services_tests

import (
	"context"
	"testing"

	auth_fixtures "github.com/gpabois/cougnat/auth/models/fixtures"
	auth_svcs "github.com/gpabois/cougnat/auth/services"
	auth_mocks "github.com/gpabois/cougnat/auth/services/mocks"
	auth_utils "github.com/gpabois/cougnat/auth/utils"
	"github.com/gpabois/cougnat/core/events"
	"github.com/gpabois/cougnat/monitoring/services"
	events_mocks "github.com/gpabois/cougnat/reporting/events/mocks"
	reporting_models "github.com/gpabois/cougnat/reporting/models"
	fixtures "github.com/gpabois/cougnat/reporting/models/fixtures"
	repos "github.com/gpabois/cougnat/reporting/repositories"
	repo_mocks "github.com/gpabois/cougnat/reporting/repositories/mocks"
	reporting_services "github.com/gpabois/cougnat/reporting/services"
	"github.com/gpabois/gostd/option"
	"github.com/gpabois/gostd/rand"
	"github.com/gpabois/gostd/result"
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
	container.Provide(func() events.IEventService {
		return reportEventEmitter
	})
	container.Provide(func() auth_svcs.IAuthorizationService {
		return authz
	})

	return mockedReportDependencies{
		authz, reportEventEmitter, reportRepo,
	}
}

func setupReportServiceTest(t *testing.T) *dig.Container {
	container := dig.New()
	// Provide mocked deps for the report service
	deps := generateMockedReportServiceDependencies(t, container)
	// Provide the report service
	container.Provide(func() mockedReportDependencies { return deps })
	container.Provide(services.ProvideReportService)
	//
	return container
}

// Test a successful use of Report Function of the Report Service
// The function must :
// 1 - Store the report, calls IReportRepository.Create
// 2 - Create an owner role with [read, write] permissions, and add the role to the owner
// 3 - Emits a ReportCreated event
func Test_ReportService_Report_Success(t *testing.T) {
	// Setup the test
	container := setupReportServiceTest(t)

	// Run the test
	err := container.Invoke(func(svc reporting_services.IReportService, deps mockedReportDependencies) {
		// Create fixtures
		report := fixtures.RandomAnonymousReport()
		reportID := 120
		ownerID := auth_fixtures.RandomAnonymousID()

		// Expected report
		expectedReport := report
		expectedReport.Owner = option.Some(ownerID)
		expectedReport.ID = option.Some(reportID)

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
		deps.authz.EXPECT().CreateAndAddRoleTo(
			ownerID,
			"owner",
			option.Some(reporting_models.ReportObjectID(reportID)), []string{"read", "write"}).Return(result.Success(true))

		// services
		res := svc.Report(reporting_services.ReportRequest{
			Owner:    option.Some(ownerID),
			TypeID:   expectedReport.Type.ID,
			Location: expectedReport.Location,
			Rate:     expectedReport.Rate,
		})

		// Should have successfuly created the report
		assert.True(t, res.IsSuccess(), res.UnwrapError())
		assert.Equal(t, expectedReport, res.Expect())

		// Should have called IAuthorizationService.CreateAndAddRoleTo
		// With ownerID as the ActorID, "owner" as the role name, reportID as the
		deps.authz.AssertCalled(t, "CreateAndAddRoleTo",
			ownerID,
			"owner",
			option.Some(reporting_models.ReportObjectID(reportID)),
			[]string{"read", "write"},
		)

		// Should have sent OnNewReport event to the receiver
		deps.reportEventEmitter.AssertCalled(t, "OnNewReport", expectedReport)
	})
	assert.Nil(t, err, err)
}

func Test_ReportService_DeleteReport_Success(t *testing.T) {
	// Setup the test
	container := setupReportServiceTest(t)
	// Run the test
	err := container.Invoke(func(svc services.IReportService, deps mockedReportDependencies) {
		// Create a report ID
		reportID, _ := rand.RandomHex(20)
		objectID := models.ReportObjectID(reportID)

		// Requester
		requesterID := auth_fixtures.RandomAnonymousID()

		// Set the requester
		ctx := auth_utils.SetCurrentActorID(context.Background(), requesterID)

		// Setup the mocked functions
		deps.authz.EXPECT().HasPermission(requesterID, "write", option.Some(objectID)).Return(result.Success(true))
		deps.authz.EXPECT().RemoveByObjectID(objectID).Return(result.Success(true))
		deps.reportEventEmitter.EXPECT().OnDeletedReport(reportID).Return(result.Success(true))
		deps.reportRepo.EXPECT().Delete(reportID).Return(result.Success(true))

		res := svc.DeleteReport(ctx, reportID)
		assert.True(t, res.IsSuccess(), res.UnwrapError())
		assert.Equal(t, reportID, res.Expect())

		deps.authz.AssertCalled(t, "HasPermission", requesterID, "write", option.Some(objectID))
		deps.authz.AssertCalled(t, "RemoveByObjectID", objectID)
		deps.reportEventEmitter.AssertCalled(t, "OnDeletedReport", reportID)
		deps.reportRepo.AssertCalled(t, "Delete", reportID)
	})
	assert.Nil(t, err, err)
}
