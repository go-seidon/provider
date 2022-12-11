package queueing

import (
	"context"
	"time"
)

const (
	PROVIDER_RABBITMQ = "rabbitmq"
)

const (
	EXCHANGE_DIRECT  = "direct"
	EXCHANGE_FANOUT  = "fanout" //broadcasts all the messages it receives to all the queues it knows
	EXCHANGE_TOPIC   = "topic"
	EXCHANGE_HEADERS = "headers"
)

type Queuer interface {
	Manager
	Exchange
	Pubsub
	Queue
}

type Manager interface {
	Open(ctx context.Context) error
	Close(ctx context.Context) error
}

type Exchange interface {
	DeclareExchange(ctx context.Context, p DeclareExchangeParam) error
	BindQueue(ctx context.Context, p BindQueueParam) error
}

type DeclareExchangeParam struct {
	ExchangeName string
	ExchangeType string
}

type BindQueueParam struct {
	ExchangeName string
	QueueName    string
}

type Pubsub interface {
	Publish(ctx context.Context, p PublishParam) error
	Subscribe(ctx context.Context, p SubscribeParam) error
}

type PublishParam struct {
	ExchangeName string
	MessageBody  []byte
}

type SubscribeParam struct {
	QueueName string
	Listener  Listener
}

type Listener = func(ctx context.Context, message Message) error

type Message interface {
	GetId() string
	GetBody() []byte
	GetPublishedAt() time.Time
	Ack() error  // positively acknowledge the message being processed successfully
	Nack() error // negatively acknowledge the message can't be processed right now and requeue to different consumer if possible, if it's not possible requeue to dead-letter queue if any
	Drop() error // negatively acknowledge the message can't be processed at all and requeue to dead-letter queue if any
}

type Queue interface {
	DeclareQueue(ctx context.Context, p DeclareQueueParam) (*DeclareQueueResult, error)
}

type DeclareQueueParam struct {
	QueueName  string
	DeadLetter *DeclareQueueDeadLetter
}

type DeclareQueueResult struct {
	Name string
}

type DeclareQueueDeadLetter struct {
	ExchangeName string
	RoutingKey   string
}
