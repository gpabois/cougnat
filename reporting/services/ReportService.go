package services

import (
	"context"
	"errors"

	auth_models "github.com/gpabois/cougnat/auth/models"
	auth_svcs "github.com/gpabois/cougnat/auth/services"
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
	ownerID := ctx.Value("CurrentActorID")

	// Set the Owner of the Report
	if ownerID == nil {
		report.Owner = auth_models.AnonymousID("")
	} else {
		ownerID, ok := ownerID.(auth_models.ActorID)
		if ok {
			report.Owner = ownerID
		} else {
			report.Owner = auth_models.AnonymousID("")
		}
	}

	res := svc.repo.Create(report)

	if res.HasFailed() {
		return result.Failed[models.Report](res.UnwrapError())
	}

	res2 := svc.repo.GetById(res.Expect())

	if res2.HasFailed() {
		return result.Failed[models.Report](res2.UnwrapError())
	}

	res3 := res2.Expect()

	if res3.IsNone() {
		return result.Failed[models.Report](errors.New("created report not found"))
	}

	report = res3.Expect()

	// Send an event
	svc.evrecv.OnNewReport(report.ID)

	// Add permissions read and write to the user
	if report.Owner.IsBound() {
		svc.authz.AddPermissions(
			report.Owner,
			[]string{"read", "write"},
			report.ObjectID(),
		)
	}

	return result.Success(report)
}

func (svc ImplReportService) DeleteReport(ctx context.Context, reportID models.ReportID) result.Result[bool] {
	currentActorID := ctx.Value("CurrentActorID")
	objectID := models.ReportObjectID(reportID)
	canDelete := svc.authz.HasPermission(currentActorID, "write", objectID)

	if !canDelete {
		return result.Failed[bool](errors.New("not permitted"))
	}

	svc.evrecv.OnDeletedReport(reportID)
	return result.Success(true)
}

// Provide the report service
func ProvideReportService(repo repos.ReportRepository, authz auth_svcs.AuthorizationService, evrecv ev.ReportEventReceiver) ReportService {
	return &ImplReportService{
		repo,
		authz,
		evrecv,
	}
}
