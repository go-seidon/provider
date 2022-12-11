package rabbitmq

import (
	"context"
	"time"

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
			startedAt := que.clock.Now().UTC()

			err := p.Listener(ctx, &message{d: d})

			finishedAt := que.clock.Now().UTC()

			log := que.logger.WithFields(map[string]interface{}{
				"queue_name":   p.QueueName,
				"message_id":   d.MessageId,
				"published_at": d.Timestamp.UTC().Format(time.RFC3339),
				"started_at":   startedAt.Format(time.RFC3339),
				"finished_at":  finishedAt.Format(time.RFC3339),
			})

			if err != nil {
				log.WithError(err).Errorf("Failed processing %s at %s", d.MessageId, p.QueueName)
			} else {
				log.Infof("Processing %s at %s", d.MessageId, p.QueueName)
			}

		}
	}()
	<-forever

	return nil
}
