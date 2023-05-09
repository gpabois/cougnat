package events

import "github.com/gpabois/cougnat/reporting/models"

//go:generate mockery --name ReportEventReceiver
type ReportEventReceiver interface {
	OnNewReport(report models.ReportID)
	OnDeletedReport(report models.ReportID)
}
