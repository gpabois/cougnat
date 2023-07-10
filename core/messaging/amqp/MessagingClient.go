package amqp

import (
	"context"
	"time"

	msg "github.com/gpabois/cougnat/core/messaging"
	"github.com/gpabois/gostd/result"
	amqp "github.com/rabbitmq/amqp091-go"
)

type MessagingClient struct {
	conn *amqp.Connection
}

func (client *MessagingClient) PublishFanout(msg []byte, exchangeName string) result.Result[bool] {
	ch, err := client.conn.Channel()

	if err != nil {
		return result.Failed[bool](err)
	}

	err = ch.ExchangeDeclare(exchangeName, "fanout", true, false, false, false, nil)

	if err != nil {
		return result.Failed[bool](err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx, exchangeName, "", true, true, amqp.Publishing{ContentType: "application/bson", Body: msg})

	if err != nil {
		return result.Failed[bool](err)
	}

	return result.Success(true)
}

func (client *MessagingClient) SubscribeFanout(exchangeName string, queueName string) result.Result[chan msg.Message] {
	ch, err := client.conn.Channel()

	if err != nil {
		return result.Failed[chan msg.Message](err)
	}

	if err = ch.ExchangeDeclare(exchangeName, "fanout", true, false, false, false, nil); err != nil {
		return result.Failed[chan msg.Message](err)
	}

	queue, err := ch.QueueDeclare(queueName, true, false, false, false, nil)

	if err != nil {
		return result.Failed[chan msg.Message](err)
	}

	if err = ch.QueueBind(
		queue.Name,   // name of the queue
		"",           // bindingKey
		exchangeName, // sourceExchange
		false,        // noWait
		nil,          // arguments
	); err != nil {
		return result.Failed[chan msg.Message](err)
	}

	rawChannel, err := ch.Consume(queue.Name, "", true, false, false, false, nil)

	if err != nil {
		return result.Failed[chan msg.Message](err)
	}

	msgChannel := make(chan msg.Message)

	go func() {
		for raw := range rawChannel {
			msgChannel <- msg.Message{
				ContentType: raw.ContentType,
				Body:        raw.Body,
			}
		}
	}()

	return result.Success(msgChannel)
}
