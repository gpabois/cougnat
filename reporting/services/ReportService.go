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

// Implementation
type ReportService struct {
	repo   repos.IReportRepository
	authz  auth_svcs.IAuthorizationService
	events ev.IReportEventEmitter
}

// Report any annoyances.
func (svc *ReportService) Report(ctx context.Context, report models.Report) result.Result[models.Report] {
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
			return opt.IntoResult(errors.New("not found"))
		})).
		// Send an event, and create access controls for the owner
		Then(result.ThenFromAny(func(report models.Report) {
			// Send an event
			svc.events.OnNewReport(report)

			// Assert object id
			report.ObjectID().Expect()

			// Create an ACL for the Owner, if any
			if report.Owner.IsSome() {
				svc.authz.CreateAndAddRoleTo(
					ownerID.Expect(),
					"owner",
					report.ObjectID(),
					[]string{"read", "write"},
				)
			}
		})))
}

// Delete a report, if the actor has the right to do so.
func (svc *ReportService) DeleteReport(ctx context.Context, reportID models.ReportID) result.Result[models.ReportID] {
	// Check if the requester is authenticated
	currentActorIDRes := guards.IsAuthenticated(ctx)
	if currentActorIDRes.HasFailed() {
		return result.Result[string]{}.Failed(currentActorIDRes.UnwrapError())
	}
	currentActorID := currentActorIDRes.Expect()
	objectID := models.ReportObjectID(reportID)

	// Check if the requester has the permission to do so
	acl := auth_models.NewAccessControl(currentActorID, "write", option.Some(objectID))
	aclRes := guards.CheckAccessControl(svc.authz)(acl)
	if aclRes.HasFailed() {
		return result.Result[string]{}.Failed(aclRes.UnwrapError())
	}

	// Delete the report
	deleteRes := svc.repo.Delete(reportID)
	if deleteRes.HasFailed() {
		return result.Result[string]{}.Failed(deleteRes.UnwrapError())
	}
	// Clean the ACL
	svc.authz.RemoveByObjectID(objectID)

	// Send an event
	svc.events.OnDeletedReport(reportID)

	// Return successful
	return result.Success(reportID)
}

func ProvideReportService(repo repos.IReportRepository, authz auth_svcs.IAuthorizationService, events ev.IReportEventEmitter) IReportService {
	return &ReportService{
		repo,
		authz,
		events,
	}
}
