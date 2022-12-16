// Code generated by MockGen. DO NOT EDIT.
// Source: grpc/stream.go

// Package mock_grpc is a generated GoMock package.
package mock_grpc

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	metadata "google.golang.org/grpc/metadata"
)

// MockServerStream is a mock of ServerStream interface.
type MockServerStream struct {
	ctrl     *gomock.Controller
	recorder *MockServerStreamMockRecorder
}

// MockServerStreamMockRecorder is the mock recorder for MockServerStream.
type MockServerStreamMockRecorder struct {
	mock *MockServerStream
}

// NewMockServerStream creates a new mock instance.
func NewMockServerStream(ctrl *gomock.Controller) *MockServerStream {
	mock := &MockServerStream{ctrl: ctrl}
	mock.recorder = &MockServerStreamMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServerStream) EXPECT() *MockServerStreamMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockServerStream) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockServerStreamMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockServerStream)(nil).Context))
}

// RecvMsg mocks base method.
func (m_2 *MockServerStream) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockServerStreamMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockServerStream)(nil).RecvMsg), m)
}

// SendHeader mocks base method.
func (m *MockServerStream) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockServerStreamMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockServerStream)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockServerStream) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockServerStreamMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockServerStream)(nil).SendMsg), m)
}

// SetHeader mocks base method.
func (m *MockServerStream) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockServerStreamMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockServerStream)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockServerStream) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockServerStreamMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockServerStream)(nil).SetTrailer), arg0)
}
