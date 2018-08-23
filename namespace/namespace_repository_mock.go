// Code generated by MockGen. DO NOT EDIT.
// Source: namespace/namespace_repository.go

// Package namespace is a generated GoMock package.
package namespace

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MocknamespaceRepositoryInterface is a mock of namespaceRepositoryInterface interface
type MocknamespaceRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MocknamespaceRepositoryInterfaceMockRecorder
}

// MocknamespaceRepositoryInterfaceMockRecorder is the mock recorder for MocknamespaceRepositoryInterface
type MocknamespaceRepositoryInterfaceMockRecorder struct {
	mock *MocknamespaceRepositoryInterface
}

// NewMocknamespaceRepositoryInterface creates a new mock instance
func NewMocknamespaceRepositoryInterface(ctrl *gomock.Controller) *MocknamespaceRepositoryInterface {
	mock := &MocknamespaceRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MocknamespaceRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MocknamespaceRepositoryInterface) EXPECT() *MocknamespaceRepositoryInterfaceMockRecorder {
	return m.recorder
}

// isNamespaceExist mocks base method
func (m *MocknamespaceRepositoryInterface) isNamespaceExist(organizationId int, serviceName, namespace string) (bool, error) {
	ret := m.ctrl.Call(m, "isNamespaceExist", organizationId, serviceName, namespace)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// isNamespaceExist indicates an expected call of isNamespaceExist
func (mr *MocknamespaceRepositoryInterfaceMockRecorder) isNamespaceExist(organizationId, serviceName, namespace interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isNamespaceExist", reflect.TypeOf((*MocknamespaceRepositoryInterface)(nil).isNamespaceExist), organizationId, serviceName, namespace)
}

// createConfiguration mocks base method
func (m *MocknamespaceRepositoryInterface) createConfiguration(organizationId int, serviceName, name string, configurations map[string]interface{}, createdBy string) error {
	ret := m.ctrl.Call(m, "createConfiguration", organizationId, serviceName, name, configurations, createdBy)
	ret0, _ := ret[0].(error)
	return ret0
}

// createConfiguration indicates an expected call of createConfiguration
func (mr *MocknamespaceRepositoryInterfaceMockRecorder) createConfiguration(organizationId, serviceName, name, configurations, createdBy interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "createConfiguration", reflect.TypeOf((*MocknamespaceRepositoryInterface)(nil).createConfiguration), organizationId, serviceName, name, configurations, createdBy)
}

// createNewNamespace mocks base method
func (m *MocknamespaceRepositoryInterface) createNewNamespace(namespaceStore *namespaceStore) error {
	ret := m.ctrl.Call(m, "createNewNamespace", namespaceStore)
	ret0, _ := ret[0].(error)
	return ret0
}

// createNewNamespace indicates an expected call of createNewNamespace
func (mr *MocknamespaceRepositoryInterfaceMockRecorder) createNewNamespace(namespaceStore interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "createNewNamespace", reflect.TypeOf((*MocknamespaceRepositoryInterface)(nil).createNewNamespace), namespaceStore)
}

// retrieveAllNamespace mocks base method
func (m *MocknamespaceRepositoryInterface) retrieveAllNamespace(organizationId int, serviceName string) ([]namespaceStore, error) {
	ret := m.ctrl.Call(m, "retrieveAllNamespace", organizationId, serviceName)
	ret0, _ := ret[0].([]namespaceStore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// retrieveAllNamespace indicates an expected call of retrieveAllNamespace
func (mr *MocknamespaceRepositoryInterfaceMockRecorder) retrieveAllNamespace(organizationId, serviceName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "retrieveAllNamespace", reflect.TypeOf((*MocknamespaceRepositoryInterface)(nil).retrieveAllNamespace), organizationId, serviceName)
}

// getOrganizationId mocks base method
func (m *MocknamespaceRepositoryInterface) getOrganizationId(organizationName string) (int, error) {
	ret := m.ctrl.Call(m, "getOrganizationId", organizationName)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getOrganizationId indicates an expected call of getOrganizationId
func (mr *MocknamespaceRepositoryInterfaceMockRecorder) getOrganizationId(organizationName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getOrganizationId", reflect.TypeOf((*MocknamespaceRepositoryInterface)(nil).getOrganizationId), organizationName)
}
