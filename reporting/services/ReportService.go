package services

import (
	"context"
	"errors"

	auth_models "github.com/gpabois/cougnat/auth/models"
	auth_svcs "github.com/gpabois/cougnat/auth/services"
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

func (svc ImplReportService) Report(ctx context.Context, report models.Report) result.Result[models.Report] {
	ownerID := auth_models.ActorID_TryFromAny(ctx.Value("CurrentActor.ID"))
	report.Owner = ownerID

	// Chain operations
	res := result.FlatMap(
		// Create The Report
		result.FlatMap(
			svc.repo.Create(report),
			// Retrieve the report upon insertion
			func(reportID string) result.Result[option.Option[models.Report]] { return svc.repo.GetById(reportID) },
		),
		option.IntoResultFunc[models.Report](errors.New("created report not found")),
	)

	// If the result is successful, we proceed to send an event, and add permissions to the actor
	res.Then(func(report models.Report) {
		// Send an event
		svc.evrecv.OnNewReport(report.ID.Expect())

		// Add permissions read and write to the actor if he's bound (has an ID)
		if report.Owner.IsSome() && report.Owner.Expect().IsBound() {
			svc.authz.AddPermissions(
				report.Owner.Expect(),
				[]string{"read", "write"},
				report.ObjectID().Expect(),
			)
		}
	})

	return res
}

func GuardPermission() func(perm bool) result.Result[bool] {
	return func(perm bool) result.Result[bool] {
		if !perm {
			return result.Failed[bool](errors.New("not permitted"))
		}
		return result.Success[bool](true)
	}
}

func (svc ImplReportService) DeleteReport(ctx context.Context, reportID models.ReportID) result.Result[bool] {
	return result.ChainMap(
		// Execute the operation
		func(b bool) result.Result[bool] {
			return svc.repo.Delete(reportID)
		},
		// Check if the actor has the rights to delete the report
		result.ChainMap(
			func(currentActorID auth_models.ActorID) result.Result[bool] {
				objectID := models.ReportObjectID(reportID)
				return result.FlatMap(
					svc.authz.HasPermission(currentActorID, "write", objectID),
					GuardPermission(),
				)
			},
			// Check if the actor is authenticated
			auth_models.ActorID_TryFromAny(ctx.Value("CurrentActor.ID")).IntoResult(errors.New("not authenticated")),
		),
	).Then(func(_ bool) {
		svc.evrecv.OnDeletedReport(reportID)
	})
}

// Provide the report service
func ProvideReportService(repo repos.ReportRepository, authz auth_svcs.AuthorizationService, evrecv ev.ReportEventReceiver) ReportService {
	return &ImplReportService{
		repo,
		authz,
		evrecv,
	}
}
