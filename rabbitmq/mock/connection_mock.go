// Code generated by MockGen. DO NOT EDIT.
// Source: rabbitmq/connection.go

// Package mock_rabbitmq is a generated GoMock package.
package mock_rabbitmq

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	amqp091_go "github.com/rabbitmq/amqp091-go"
)

// MockConnection is a mock of Connection interface.
type MockConnection struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionMockRecorder
}

// MockConnectionMockRecorder is the mock recorder for MockConnection.
type MockConnectionMockRecorder struct {
	mock *MockConnection
}

// NewMockConnection creates a new mock instance.
func NewMockConnection(ctrl *gomock.Controller) *MockConnection {
	mock := &MockConnection{ctrl: ctrl}
	mock.recorder = &MockConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnection) EXPECT() *MockConnectionMockRecorder {
	return m.recorder
}

// Channel mocks base method.
func (m *MockConnection) Channel() (*amqp091_go.Channel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Channel")
	ret0, _ := ret[0].(*amqp091_go.Channel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Channel indicates an expected call of Channel.
func (mr *MockConnectionMockRecorder) Channel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Channel", reflect.TypeOf((*MockConnection)(nil).Channel))
}

// Close mocks base method.
func (m *MockConnection) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockConnectionMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockConnection)(nil).Close))
}

// Init mocks base method.
func (m *MockConnection) Init(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockConnectionMockRecorder) Init(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockConnection)(nil).Init), ctx)
}

// IsClosed mocks base method.
func (m *MockConnection) IsClosed() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsClosed")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsClosed indicates an expected call of IsClosed.
func (mr *MockConnectionMockRecorder) IsClosed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsClosed", reflect.TypeOf((*MockConnection)(nil).IsClosed))
}

// NotifyBlocked mocks base method.
func (m *MockConnection) NotifyBlocked(receiver chan amqp091_go.Blocking) chan amqp091_go.Blocking {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotifyBlocked", receiver)
	ret0, _ := ret[0].(chan amqp091_go.Blocking)
	return ret0
}

// NotifyBlocked indicates an expected call of NotifyBlocked.
func (mr *MockConnectionMockRecorder) NotifyBlocked(receiver interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyBlocked", reflect.TypeOf((*MockConnection)(nil).NotifyBlocked), receiver)
}

// NotifyClose mocks base method.
func (m *MockConnection) NotifyClose(receiver chan *amqp091_go.Error) chan *amqp091_go.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotifyClose", receiver)
	ret0, _ := ret[0].(chan *amqp091_go.Error)
	return ret0
}

// NotifyClose indicates an expected call of NotifyClose.
func (mr *MockConnectionMockRecorder) NotifyClose(receiver interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyClose", reflect.TypeOf((*MockConnection)(nil).NotifyClose), receiver)
}

// MockChannel is a mock of Channel interface.
type MockChannel struct {
	ctrl     *gomock.Controller
	recorder *MockChannelMockRecorder
}

// MockChannelMockRecorder is the mock recorder for MockChannel.
type MockChannelMockRecorder struct {
	mock *MockChannel
}

// NewMockChannel creates a new mock instance.
func NewMockChannel(ctrl *gomock.Controller) *MockChannel {
	mock := &MockChannel{ctrl: ctrl}
	mock.recorder = &MockChannelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChannel) EXPECT() *MockChannelMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockChannel) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockChannelMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockChannel)(nil).Close))
}

// Consume mocks base method.
func (m *MockChannel) Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp091_go.Table) (<-chan amqp091_go.Delivery, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Consume", queue, consumer, autoAck, exclusive, noLocal, noWait, args)
	ret0, _ := ret[0].(<-chan amqp091_go.Delivery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Consume indicates an expected call of Consume.
func (mr *MockChannelMockRecorder) Consume(queue, consumer, autoAck, exclusive, noLocal, noWait, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consume", reflect.TypeOf((*MockChannel)(nil).Consume), queue, consumer, autoAck, exclusive, noLocal, noWait, args)
}

// ExchangeDeclare mocks base method.
func (m *MockChannel) ExchangeDeclare(name, kind string, durable, autoDelete, internal, noWait bool, args amqp091_go.Table) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExchangeDeclare", name, kind, durable, autoDelete, internal, noWait, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExchangeDeclare indicates an expected call of ExchangeDeclare.
func (mr *MockChannelMockRecorder) ExchangeDeclare(name, kind, durable, autoDelete, internal, noWait, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExchangeDeclare", reflect.TypeOf((*MockChannel)(nil).ExchangeDeclare), name, kind, durable, autoDelete, internal, noWait, args)
}

// PublishWithContext mocks base method.
func (m *MockChannel) PublishWithContext(ctx context.Context, exchange, key string, mandatory, immediate bool, msg amqp091_go.Publishing) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishWithContext", ctx, exchange, key, mandatory, immediate, msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishWithContext indicates an expected call of PublishWithContext.
func (mr *MockChannelMockRecorder) PublishWithContext(ctx, exchange, key, mandatory, immediate, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishWithContext", reflect.TypeOf((*MockChannel)(nil).PublishWithContext), ctx, exchange, key, mandatory, immediate, msg)
}

// QueueBind mocks base method.
func (m *MockChannel) QueueBind(name, key, exchange string, noWait bool, args amqp091_go.Table) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueueBind", name, key, exchange, noWait, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// QueueBind indicates an expected call of QueueBind.
func (mr *MockChannelMockRecorder) QueueBind(name, key, exchange, noWait, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueBind", reflect.TypeOf((*MockChannel)(nil).QueueBind), name, key, exchange, noWait, args)
}

// QueueDeclare mocks base method.
func (m *MockChannel) QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args amqp091_go.Table) (amqp091_go.Queue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueueDeclare", name, durable, autoDelete, exclusive, noWait, args)
	ret0, _ := ret[0].(amqp091_go.Queue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueueDeclare indicates an expected call of QueueDeclare.
func (mr *MockChannelMockRecorder) QueueDeclare(name, durable, autoDelete, exclusive, noWait, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueDeclare", reflect.TypeOf((*MockChannel)(nil).QueueDeclare), name, durable, autoDelete, exclusive, noWait, args)
}
