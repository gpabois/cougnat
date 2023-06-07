package events

import (
	"github.com/gpabois/cougnat/core/result"
)

//go:generate mockery --name ReportEventEmitter
type IReportEventReceiver interface {
	OnNewReport(queueName string) result.Result[chan NewReportEvent]
	OnDeletedReport(queueName string) result.Result[chan DeletedReportEvent]
}
