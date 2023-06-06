package messaging

import "github.com/gpabois/cougnat/core/result"

type IMessagingClient interface {
	Publish(msg []byte, exchangeName string, exchangeType string) result.Result[bool]
	PublishOnQueue(msg []byte, queueName string) result.Result[bool]
}
