package messaging

import (
	msg "github.com/gpabois/cougnat/core/messaging"
	"github.com/gpabois/cougnat/core/result"
	"github.com/gpabois/cougnat/core/serde"
	"github.com/gpabois/cougnat/reporting/events"
	"github.com/gpabois/cougnat/reporting/models"
)

type ReportEventEmitter struct {
	messaging msg.IMessagingClient
}

func (emitter *ReportEventEmitter) OnNewReport(report models.Report) result.Result[bool] {
	return result.Chain(func(msg []byte) result.Result[bool] {
		return (emitter.messaging).PublishFanout(msg, "cougnat.reports.created")
	}, serde.MarshalBson(events.NewReportEvent{Report: report}))
}

func (emitter *ReportEventEmitter) OnDeletedReport(report models.ReportID) result.Result[bool] {
	return result.Chain(func(msg []byte) result.Result[bool] {
		return (emitter.messaging).PublishFanout(msg, "cougnat.reports.deleted")
	}, serde.MarshalBson(events.DeletedReportEvent{ReportID: report}))
}

func ProvideReportEmitter(messaging msg.IMessagingClient) events.IReportEventEmitter {
	return &ReportEventEmitter{messaging: messaging}
}
