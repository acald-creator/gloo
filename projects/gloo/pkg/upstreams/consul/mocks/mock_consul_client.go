// Code generated by MockGen. DO NOT EDIT.
// Source: consul_client.go

// Package mock_consul is a generated GoMock package.
package mock_consul

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "github.com/hashicorp/consul/api"
)

// MockClientWrapper is a mock of ClientWrapper interface.
type MockClientWrapper struct {
	ctrl     *gomock.Controller
	recorder *MockClientWrapperMockRecorder
}

// MockClientWrapperMockRecorder is the mock recorder for MockClientWrapper.
type MockClientWrapperMockRecorder struct {
	mock *MockClientWrapper
}

// NewMockClientWrapper creates a new mock instance.
func NewMockClientWrapper(ctrl *gomock.Controller) *MockClientWrapper {
	mock := &MockClientWrapper{ctrl: ctrl}
	mock.recorder = &MockClientWrapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientWrapper) EXPECT() *MockClientWrapperMockRecorder {
	return m.recorder
}

// Connect mocks base method.
func (m *MockClientWrapper) Connect(service, tag string, q *api.QueryOptions) ([]*api.CatalogService, *api.QueryMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect", service, tag, q)
	ret0, _ := ret[0].([]*api.CatalogService)
	ret1, _ := ret[1].(*api.QueryMeta)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Connect indicates an expected call of Connect.
func (mr *MockClientWrapperMockRecorder) Connect(service, tag, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockClientWrapper)(nil).Connect), service, tag, q)
}

// DataCenters mocks base method.
func (m *MockClientWrapper) DataCenters() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataCenters")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DataCenters indicates an expected call of DataCenters.
func (mr *MockClientWrapperMockRecorder) DataCenters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataCenters", reflect.TypeOf((*MockClientWrapper)(nil).DataCenters))
}

// Service mocks base method.
func (m *MockClientWrapper) Service(service, tag string, q *api.QueryOptions) ([]*api.CatalogService, *api.QueryMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Service", service, tag, q)
	ret0, _ := ret[0].([]*api.CatalogService)
	ret1, _ := ret[1].(*api.QueryMeta)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Service indicates an expected call of Service.
func (mr *MockClientWrapperMockRecorder) Service(service, tag, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Service", reflect.TypeOf((*MockClientWrapper)(nil).Service), service, tag, q)
}

// Services mocks base method.
func (m *MockClientWrapper) Services(q *api.QueryOptions) (map[string][]string, *api.QueryMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Services", q)
	ret0, _ := ret[0].(map[string][]string)
	ret1, _ := ret[1].(*api.QueryMeta)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Services indicates an expected call of Services.
func (mr *MockClientWrapperMockRecorder) Services(q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Services", reflect.TypeOf((*MockClientWrapper)(nil).Services), q)
}
