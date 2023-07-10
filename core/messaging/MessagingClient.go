package messaging

import "github.com/gpabois/gostd/result"

type Message struct {
	ContentType string
	Body        []byte
}

type IMessagingClient interface {
	PublishFanout(msg []byte, exchangeName string) result.Result[bool]
	SubscribeFanout(exchange string, queueName string) result.Result[chan Message]
}
