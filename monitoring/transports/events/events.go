package events

import (
	"github.com/gpabois/cougnat/core/result"
	reporting_events "github.com/gpabois/cougnat/reporting/events"
)

type EventTransport struct {
	reports reporting_events.IReportEventReceiver
}

func (transport EventTransport) HandleNewReport(endpoint func(newReportEvent reporting_events.NewReportEvent)) result.Result[bool] {
	res := transport.reports.OnNewReport("cougnat.monitoring.OnNewReport")

	if res.HasFailed() {
		return result.Failed[bool](res.UnwrapError())
	}

	ch := res.Expect()
	go func() {
		for newReportEvent := range ch {
			endpoint(newReportEvent)
		}
	}()

	return result.Success(true)
}

func (transport EventTransport) HandleDeletedReport(endpoint func(deletedReportEvent reporting_events.DeletedReportEvent)) result.Result[bool] {
	res := transport.reports.OnDeletedReport("cougnat.monitoring.OnDeletedReport")

	if res.HasFailed() {
		return result.Failed[bool](res.UnwrapError())
	}

	ch := res.Expect()

	go func() {
		for deletedReportEvent := range ch {
			endpoint(deletedReportEvent)
		}
	}()

	return result.Success(true)
}
