package rabbitmq

import (
	"context"

	"github.com/go-seidon/provider/datetime"
	"github.com/go-seidon/provider/identifier"
	"github.com/go-seidon/provider/identifier/ksuid"
	"github.com/go-seidon/provider/queueing"
	"github.com/go-seidon/provider/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

type rabbitQueue struct {
	conn       rabbitmq.Connection
	clock      datetime.Clock
	identifier identifier.Identifier
}

func (que *rabbitQueue) DeclareQueue(ctx context.Context, p queueing.DeclareQueueParam) (*queueing.DeclareQueueResult, error) {
	ch, err := que.conn.Channel()
	if err != nil {
		return nil, err
	}
	defer ch.Close()

	args := amqp.Table{}
	if p.DeadLetter != nil {
		if p.DeadLetter.ExchangeName != "" {
			args["x-dead-letter-exchange"] = p.DeadLetter.ExchangeName
		}
		if p.DeadLetter.RoutingKey != "" {
			args["x-dead-letter-routing-key"] = p.DeadLetter.RoutingKey
		}
	}

	q, err := ch.QueueDeclare(p.QueueName, true, false, false, false, args)
	if err != nil {
		return nil, err
	}

	res := &queueing.DeclareQueueResult{
		Name: q.Name,
	}
	return res, nil
}

func NewQueueing(opts ...RabbitOption) *rabbitQueue {
	p := RabbitParam{}
	for _, opt := range opts {
		opt(&p)
	}

	clock := p.Clock
	if clock == nil {
		clock = datetime.NewClock()
	}

	identifier := p.Identifier
	if identifier == nil {
		identifier = ksuid.NewIdentifier()
	}

	return &rabbitQueue{
		conn:       p.Connection,
		clock:      clock,
		identifier: identifier,
	}
}
