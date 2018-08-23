// Code generated by MockGen. DO NOT EDIT.
// Source: namespace/namespace_processor.go

// Package namespace is a generated GoMock package.
package namespace

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// Mockprocessor is a mock of processor interface
type Mockprocessor struct {
	ctrl     *gomock.Controller
	recorder *MockprocessorMockRecorder
}

// MockprocessorMockRecorder is the mock recorder for Mockprocessor
type MockprocessorMockRecorder struct {
	mock *Mockprocessor
}

// NewMockprocessor creates a new mock instance
func NewMockprocessor(ctrl *gomock.Controller) *Mockprocessor {
	mock := &Mockprocessor{ctrl: ctrl}
	mock.recorder = &MockprocessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Mockprocessor) EXPECT() *MockprocessorMockRecorder {
	return m.recorder
}

// parseData mocks base method
func (m *Mockprocessor) parseData(organizationId int, serviceName string, view *namespaceView, data *namespaceStore) error {
	ret := m.ctrl.Call(m, "parseData", organizationId, serviceName, view, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// parseData indicates an expected call of parseData
func (mr *MockprocessorMockRecorder) parseData(organizationId, serviceName, view, data interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "parseData", reflect.TypeOf((*Mockprocessor)(nil).parseData), organizationId, serviceName, view, data)
}

// createNewNamespaceProcessor mocks base method
func (m *Mockprocessor) createNewNamespaceProcessor(organizationName, serviceName string, namespacev *namespaceView) error {
	ret := m.ctrl.Call(m, "createNewNamespaceProcessor", organizationName, serviceName, namespacev)
	ret0, _ := ret[0].(error)
	return ret0
}

// createNewNamespaceProcessor indicates an expected call of createNewNamespaceProcessor
func (mr *MockprocessorMockRecorder) createNewNamespaceProcessor(organizationName, serviceName, namespacev interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "createNewNamespaceProcessor", reflect.TypeOf((*Mockprocessor)(nil).createNewNamespaceProcessor), organizationName, serviceName, namespacev)
}

// retrieveAllNamespaceProcessor mocks base method
func (m *Mockprocessor) retrieveAllNamespaceProcessor(organizationName, serviceName string) ([]byte, error) {
	ret := m.ctrl.Call(m, "retrieveAllNamespaceProcessor", organizationName, serviceName)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// retrieveAllNamespaceProcessor indicates an expected call of retrieveAllNamespaceProcessor
func (mr *MockprocessorMockRecorder) retrieveAllNamespaceProcessor(organizationName, serviceName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "retrieveAllNamespaceProcessor", reflect.TypeOf((*Mockprocessor)(nil).retrieveAllNamespaceProcessor), organizationName, serviceName)
}