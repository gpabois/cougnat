package reporting_services

import (
	auth_svcs "github.com/gpabois/cougnat/auth/services"
	"github.com/gpabois/cougnat/core/events"
	"github.com/gpabois/cougnat/core/transaction"
	reporting_events "github.com/gpabois/cougnat/reporting/events"
	reporting_models "github.com/gpabois/cougnat/reporting/models"
	repos "github.com/gpabois/cougnat/reporting/repositories"
	authz "github.com/gpabois/goservice/authz"
	authz_services "github.com/gpabois/goservice/authz/services"
	"github.com/gpabois/gostd/option"
	"github.com/gpabois/gostd/result"
)

// Implementation
type ReportService struct {
	repo   repos.IReportRepository
	authz  auth_svcs.IAuthorizationService
	events events.IEventService
}

// Report any annoyances.
func (svc *ReportService) Report(report ReportRequest) ReportResult {
	return transaction.Begin(func(tx transaction.Transaction) ReportResult {
		res := transaction.With(svc.repo.Begin(option.Some(tx)),
			func(tx repos.IReportRepositoryTx) result.Result[reporting_models.ReportID] {
				return tx.Create(report)
			},
		)

		if res.HasFailed() {
			return result.Failed[ReportResponse](res.UnwrapError())
		}

		reportID := res.Expect()

		// Register ACL
		if report.Owner.IsSome() && report.Owner.Expect().IsBound() {
			reporter := report.Owner.Expect()
			objectID := reporting_models.ReportObjectID(reportID)
			res := svc.authz.CreateAndAddRoleTo(reporter, "owner", objectID, []string{"read", "write"})
			if res.HasFailed() {
				return result.Failed[ReportResponse](res.UnwrapError())
			}
		}

		reportRes := transaction.With(svc.repo.Begin(option.Some(tx)),
			func(tx repos.IReportRepositoryTx) result.Result[option.Option[reporting_models.Report]] {
				return tx.GetById(reportID)
			},
		)

		report := reportRes.Expect().Expect()

		// Notify events
		notifyRes := reporting_events.NotifyNewReport(svc.events, report)
		if notifyRes.HasFailed() {
			return result.Failed[ReportResponse](notifyRes.UnwrapError())
		}

		// Return result
		return result.Success(ReportResponse{ReportID: reportID})
	})
}

func (svc *ReportService) DeleteReport(request DeleteReportRequest) DeleteReportResult {
	objectID := reporting_models.ReportObjectID(request.ReportID)
	canRes := svc.authz.HasPermission(request.Requester, "write", option.Some(objectID))

	if canRes.HasFailed() {
		return result.Result[DeleteReportResponse]{}.Failed(canRes.UnwrapError())
	}

	can := canRes.Expect()
	if !can {
		return result.Result[DeleteReportResponse]{}.Failed(
			authz.NewNotAuthorizedError(authz_services.ACL{
				Subject:    request.Requester,
				Permission: "write",
				Object:     option.Some[any](objectID),
			}),
		)
	}

	return transaction.Begin(func(tx transaction.Transaction) DeleteReportResult {
		res := result.Map(transaction.With(svc.repo.Begin(option.Some(tx)),
			func(tx repos.IReportRepositoryTx) result.Result[bool] {
				return tx.Delete(request.ReportID)
			},
		), func(result bool) DeleteReportResponse {
			return DeleteReportResponse{Result: true}
		})

		if res.HasFailed() {
			return result.Result[DeleteReportResponse]{}.Failed(res.UnwrapError())
		}

		// Notify events
		notifyRes := reporting_events.NotifyDeletedReport(svc.events, request.ReportID)
		if notifyRes.HasFailed() {
			return result.Failed[DeleteReportResponse](notifyRes.UnwrapError())
		}

		return res
	})
}

func ProvideReportService(repo repos.IReportRepository, authz auth_svcs.IAuthorizationService, events events.IEventService) IReportService {
	return &ReportService{repo, authz, events}
}
