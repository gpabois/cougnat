package repositories

import (
	"github.com/gpabois/cougnat/core/option"
	result "github.com/gpabois/cougnat/core/result"
	models "github.com/gpabois/cougnat/reporting/models"
)

// Interface of the ReportRepository
//
//go:generate mockery --name ReportRepository
type ReportRepository interface {
	// Create a report
	Create(report models.Report) result.Result[models.ReportID]
	// Get a report
	GetById(id models.ReportID) result.Result[option.Option[models.Report]]
	// Delete a report
	Delete(reportID string) result.Result[bool]
	// Find reports by actor id, sorted from desc
	// FindByActorID(actorID auth.ActorID, page repos.PageCursor) iter.Iterator[result.Result[models.Report]]
	// Find near, and within a time frame
	// FindNearWithinTimeFrame(near geo.Near, frame time.TimeFrame, page repos.PageCursor) iter.Iterator[result.Result[models.Report]]
}
