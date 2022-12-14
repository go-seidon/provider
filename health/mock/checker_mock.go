// Code generated by MockGen. DO NOT EDIT.
// Source: health/checker/checker.go

// Package mock_health is a generated GoMock package.
package mock_health

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockChecker is a mock of Checker interface.
type MockChecker struct {
	ctrl     *gomock.Controller
	recorder *MockCheckerMockRecorder
}

// MockCheckerMockRecorder is the mock recorder for MockChecker.
type MockCheckerMockRecorder struct {
	mock *MockChecker
}

// NewMockChecker creates a new mock instance.
func NewMockChecker(ctrl *gomock.Controller) *MockChecker {
	mock := &MockChecker{ctrl: ctrl}
	mock.recorder = &MockCheckerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChecker) EXPECT() *MockCheckerMockRecorder {
	return m.recorder
}

// Status mocks base method.
func (m *MockChecker) Status() (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status.
func (mr *MockCheckerMockRecorder) Status() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockChecker)(nil).Status))
}

// MockDataSource is a mock of DataSource interface.
type MockDataSource struct {
	ctrl     *gomock.Controller
	recorder *MockDataSourceMockRecorder
}

// MockDataSourceMockRecorder is the mock recorder for MockDataSource.
type MockDataSourceMockRecorder struct {
	mock *MockDataSource
}

// NewMockDataSource creates a new mock instance.
func NewMockDataSource(ctrl *gomock.Controller) *MockDataSource {
	mock := &MockDataSource{ctrl: ctrl}
	mock.recorder = &MockDataSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataSource) EXPECT() *MockDataSourceMockRecorder {
	return m.recorder
}

// Ping mocks base method.
func (m *MockDataSource) Ping(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockDataSourceMockRecorder) Ping(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockDataSource)(nil).Ping), ctx)
}
