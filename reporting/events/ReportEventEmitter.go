package events

import (
	"github.com/gpabois/cougnat/core/result"
	"github.com/gpabois/cougnat/reporting/models"
)

type NewReportEvent struct {
	Report models.Report `serde:"report"`
}

type DeletedReportEvent struct {
	ReportID models.ReportID `serde:"report_id"`
}

//go:generate mockery
type IReportEventEmitter interface {
	OnNewReport(report models.Report) result.Result[bool]
	OnDeletedReport(report models.ReportID) result.Result[bool]
}
