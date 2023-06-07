package messaging

import (
	msg "github.com/gpabois/cougnat/core/messaging"
	"github.com/gpabois/cougnat/core/result"
	"github.com/gpabois/cougnat/core/serde"
	"github.com/gpabois/cougnat/reporting/events"
)

type ReportEventReceiver struct {
	messaging msg.IMessagingClient
}

func (receiver *ReportEventReceiver) OnNewReport(queueName string) result.Result[chan events.NewReportEvent] {
	return result.Chain(
		func(msgChannel chan msg.Message) result.Result[chan events.NewReportEvent] {
			newReportChannel := make(chan events.NewReportEvent)

			// Launch a go-routine to process incoming events
			go func() {
				for msg := range msgChannel {
					res := serde.UnMarshal[events.NewReportEvent](msg.Body, msg.ContentType)
					if res.IsSuccess() {
						newReportChannel <- res.Expect()
					}
				}
			}()

			return result.Success(newReportChannel)
		},
		(receiver.messaging).SubscribeFanout("cougnat.reports.created", queueName),
	)
}

func (receiver *ReportEventReceiver) OnDeletedReport(queueName string) result.Result[chan events.DeletedReportEvent] {
	return result.Chain(
		func(msgChannel chan msg.Message) result.Result[chan events.DeletedReportEvent] {
			deletedReportChannel := make(chan events.DeletedReportEvent)

			// Launch a go-routine to process incoming events
			go func() {
				for msg := range msgChannel {
					res := serde.UnMarshal[events.DeletedReportEvent](msg.Body, msg.ContentType)
					if res.IsSuccess() {
						deletedReportChannel <- res.Expect()
					}
				}
			}()

			return result.Success(deletedReportChannel)
		},
		(receiver.messaging).SubscribeFanout("cougnat.reports.deleted", queueName),
	)
}

func ProvideReportReceiver(messaging msg.IMessagingClient) events.IReportEventReceiver {
	return &ReportEventReceiver{messaging: messaging}
}
