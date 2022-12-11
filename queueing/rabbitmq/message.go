package rabbitmq

import (
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type message struct {
	d amqp.Delivery
}

func (m *message) GetId() string {
	return m.d.MessageId
}

func (m *message) GetBody() []byte {
	return m.d.Body
}

func (m *message) GetPublishedAt() time.Time {
	return m.d.Timestamp
}

func (m *message) Ack() error {
	return m.d.Ack(false)
}

func (m *message) Nack() error {
	return m.d.Nack(false, true)
}

func (m *message) Drop() error {
	return m.d.Nack(false, false)
}
