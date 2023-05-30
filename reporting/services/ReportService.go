package services

import (
	"context"
	"errors"

	"github.com/gpabois/cougnat/auth/guards"
	auth_models "github.com/gpabois/cougnat/auth/models"
	auth_svcs "github.com/gpabois/cougnat/auth/services"
	auth_utils "github.com/gpabois/cougnat/auth/utils"
	"github.com/gpabois/cougnat/core/option"
	"github.com/gpabois/cougnat/core/result"
	ev "github.com/gpabois/cougnat/reporting/events"
	"github.com/gpabois/cougnat/reporting/models"
	repos "github.com/gpabois/cougnat/reporting/repositories"
)

// Interface
type ReportService interface {
	Report(ctx context.Context, report models.Report) result.Result[models.Report]
	DeleteReport(ctx context.Context, reportID models.ReportID) result.Result[bool]
}

// Implementation
type ImplReportService struct {
	repo   repos.ReportRepository
	authz  auth_svcs.AuthorizationService
	evrecv ev.ReportEventReceiver
}

// Report any annoyances.
func (svc ImplReportService) Report(ctx context.Context, report models.Report) result.Result[models.Report] {
	ownerID := auth_utils.GetCurrentActorID(ctx)
	report.Owner = ownerID

	return result.Into[models.Report](svc.repo.
		// Create the report
		Create(report).ToAny().
		// Retrieve the new report from the repository
		Chain(result.ChainFromAny(func(reportID string) result.Result[option.Option[models.Report]] {
			return svc.repo.GetById(reportID)
		})).
		// Unwrap the option to get the report
		Chain(result.ChainFromAny(func(opt option.Option[models.Report]) result.Result[models.Report] {
			return opt.IntoResult(errors.New("created report not found"))
		})).
		// Send an event, and create access controls for the owner
		Then(result.ThenFromAny(func(report models.Report) {
			// Send an event
			svc.evrecv.OnNewReport(report.ID.Expect())

			// Create an ACL for the Owner, if any
			if report.Owner.IsSome() {
				svc.authz.CreateAndAddRoleTo(
					ownerID.Expect(),
					"owner",
					report.ObjectID().Expect(),
					[]string{"read", "write"},
				)
			}
		})))
}

// Delete a report, if the actor has the right to do so.
func (svc ImplReportService) DeleteReport(ctx context.Context, reportID models.ReportID) result.Result[bool] {
	return result.Into[bool](
		// Check if authenticated
		guards.IsAuthenticated(ctx).ToAny().
			// Return an access control to be checked
			Chain(result.ChainFromAny(func(currentActorID auth_models.ActorID) result.Result[auth_models.AccessControl] {
				objectID := models.ReportObjectID(reportID)
				return result.Success(auth_models.NewAccessControl(currentActorID, "write", objectID))
			})).
			// Check if has the permission
			Chain(result.ChainFromAny(guards.CheckAccessControl(svc.authz))).
			// Execute the deletion
			Chain(result.ChainFromAny(func(_ bool) result.Result[bool] {
				return svc.repo.Delete(reportID)
			})).
			// Send an event
			Then(result.ThenFromAny(func(deleted bool) {
				svc.evrecv.OnDeletedReport(reportID)
			})))
}

// Provide the report service
func ProvideReportService(repo repos.ReportRepository, authz auth_svcs.AuthorizationService, evrecv ev.ReportEventReceiver) ReportService {
	return &ImplReportService{
		repo,
		authz,
		evrecv,
	}
}