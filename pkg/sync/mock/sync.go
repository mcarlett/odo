// Code generated by MockGen. DO NOT EDIT.
// Source: sync.go

// Package mock is a generated GoMock package.
package mock

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	common "github.com/redhat-developer/odo/pkg/devfile/adapters/common"
)

// MockSyncClient is a mock of SyncClient interface
type MockSyncClient struct {
	ctrl     *gomock.Controller
	recorder *MockSyncClientMockRecorder
}

// MockSyncClientMockRecorder is the mock recorder for MockSyncClient
type MockSyncClientMockRecorder struct {
	mock *MockSyncClient
}

// NewMockSyncClient creates a new mock instance
func NewMockSyncClient(ctrl *gomock.Controller) *MockSyncClient {
	mock := &MockSyncClient{ctrl: ctrl}
	mock.recorder = &MockSyncClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSyncClient) EXPECT() *MockSyncClientMockRecorder {
	return m.recorder
}

// ExtractProjectToComponent mocks base method
func (m *MockSyncClient) ExtractProjectToComponent(arg0 common.ComponentInfo, arg1 string, arg2 io.Reader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtractProjectToComponent", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExtractProjectToComponent indicates an expected call of ExtractProjectToComponent
func (mr *MockSyncClientMockRecorder) ExtractProjectToComponent(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtractProjectToComponent", reflect.TypeOf((*MockSyncClient)(nil).ExtractProjectToComponent), arg0, arg1, arg2)
}
