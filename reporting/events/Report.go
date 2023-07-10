package events

import (
	"github.com/gpabois/cougnat/core/events"
	reporting_models "github.com/gpabois/cougnat/reporting/models"
	"github.com/gpabois/gostd/result"
)

const NewReportEventName = "org.cougnat.events.reporting.ReportCreated"

type NewReportEvent struct {
	Report reporting_models.Report `serde:"report"`
}

func NotifyNewReport(events events.IEventService, report reporting_models.Report) result.Result[bool] {
	return events.Notify(NewReportEventName, NewReportEvent{Report: report})
}

const DeletedReportEventName = "org.cougnat.events.reporting.ReportDeleted"

type DeletedReportEvent struct {
	ReportID reporting_models.ReportID `serde:"report_id"`
}

func NotifyDeletedReport(events events.IEventService, reportID reporting_models.ReportID) result.Result[bool] {
	return events.Notify(DeletedReportEventName, DeletedReportEvent{ReportID: reportID})
}
