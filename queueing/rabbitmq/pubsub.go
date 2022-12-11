package rabbitmq

import (
	"context"

	"github.com/go-seidon/provider/queueing"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (que *rabbitQueue) Publish(ctx context.Context, p queueing.PublishParam) error {
	ch, err := que.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	currentTs := que.clock.Now().UTC()
	id, err := que.identifier.GenerateId()
	if err != nil {
		return err
	}

	err = ch.PublishWithContext(ctx, p.ExchangeName, "", true, false, amqp.Publishing{
		Body:         p.MessageBody,
		DeliveryMode: amqp.Persistent,
		Timestamp:    currentTs,
		MessageId:    id,
	})
	if err != nil {
		return err
	}

	return nil
}

func (que *rabbitQueue) Subscribe(ctx context.Context, p queueing.SubscribeParam) error {
	ch, err := que.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	delivery, err := ch.Consume(p.QueueName, "", false, false, false, false, nil)
	if err != nil {
		return err
	}

	forever := make(chan bool)
	go func() {
		for d := range delivery {
			p.Listener(ctx, &message{d: d})
		}
	}()
	<-forever

	return nil
}
