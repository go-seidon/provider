package rabbitmq

import (
	"github.com/go-seidon/provider/datetime"
	"github.com/go-seidon/provider/identifier"
	"github.com/go-seidon/provider/rabbitmq"
)

type RabbitParam struct {
	Connection rabbitmq.Connection
	Clock      datetime.Clock
	Identifier identifier.Identifier
}

type RabbitOption = func(*RabbitParam)

func WithConnection(conn rabbitmq.Connection) RabbitOption {
	return func(p *RabbitParam) {
		p.Connection = conn
	}
}

func WithClock(c datetime.Clock) RabbitOption {
	return func(p *RabbitParam) {
		p.Clock = c
	}
}

func WithIdentifier(i identifier.Identifier) RabbitOption {
	return func(p *RabbitParam) {
		p.Identifier = i
	}
}
