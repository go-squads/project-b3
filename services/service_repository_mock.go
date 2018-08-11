// Code generated by MockGen. DO NOT EDIT.
// Source: services/service_repository.go

// Package services is a generated GoMock package.
package services

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockserviceRepositoryInterface is a mock of serviceRepositoryInterface interface
type MockserviceRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockserviceRepositoryInterfaceMockRecorder
}

// MockserviceRepositoryInterfaceMockRecorder is the mock recorder for MockserviceRepositoryInterface
type MockserviceRepositoryInterfaceMockRecorder struct {
	mock *MockserviceRepositoryInterface
}

// NewMockserviceRepositoryInterface creates a new mock instance
func NewMockserviceRepositoryInterface(ctrl *gomock.Controller) *MockserviceRepositoryInterface {
	mock := &MockserviceRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockserviceRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockserviceRepositoryInterface) EXPECT() *MockserviceRepositoryInterfaceMockRecorder {
	return m.recorder
}

// getAll mocks base method
func (m *MockserviceRepositoryInterface) getAll(organizationId int) ([]service, error) {
	ret := m.ctrl.Call(m, "getAll", organizationId)
	ret0, _ := ret[0].([]service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getAll indicates an expected call of getAll
func (mr *MockserviceRepositoryInterfaceMockRecorder) getAll(organizationId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getAll", reflect.TypeOf((*MockserviceRepositoryInterface)(nil).getAll), organizationId)
}

// createService mocks base method
func (m *MockserviceRepositoryInterface) createService(servicestore service) error {
	ret := m.ctrl.Call(m, "createService", servicestore)
	ret0, _ := ret[0].(error)
	return ret0
}

// createService indicates an expected call of createService
func (mr *MockserviceRepositoryInterfaceMockRecorder) createService(servicestore interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "createService", reflect.TypeOf((*MockserviceRepositoryInterface)(nil).createService), servicestore)
}

// deleteService mocks base method
func (m *MockserviceRepositoryInterface) deleteService(servicestore service) error {
	ret := m.ctrl.Call(m, "deleteService", servicestore)
	ret0, _ := ret[0].(error)
	return ret0
}

// deleteService indicates an expected call of deleteService
func (mr *MockserviceRepositoryInterfaceMockRecorder) deleteService(servicestore interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "deleteService", reflect.TypeOf((*MockserviceRepositoryInterface)(nil).deleteService), servicestore)
}

// getServiceToken mocks base method
func (m *MockserviceRepositoryInterface) getServiceToken(name string) (*serviceToken, error) {
	ret := m.ctrl.Call(m, "getServiceToken", name)
	ret0, _ := ret[0].(*serviceToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getServiceToken indicates an expected call of getServiceToken
func (mr *MockserviceRepositoryInterfaceMockRecorder) getServiceToken(name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getServiceToken", reflect.TypeOf((*MockserviceRepositoryInterface)(nil).getServiceToken), name)
}

// findOneServiceByName mocks base method
func (m *MockserviceRepositoryInterface) findOneServiceByName(name string) (*service, error) {
	ret := m.ctrl.Call(m, "findOneServiceByName", name)
	ret0, _ := ret[0].(*service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// findOneServiceByName indicates an expected call of findOneServiceByName
func (mr *MockserviceRepositoryInterfaceMockRecorder) findOneServiceByName(name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "findOneServiceByName", reflect.TypeOf((*MockserviceRepositoryInterface)(nil).findOneServiceByName), name)
}

// translateNameToIdRepository mocks base method
func (m *MockserviceRepositoryInterface) translateNameToIdRepository(organizationName string) (int, error) {
	ret := m.ctrl.Call(m, "translateNameToIdRepository", organizationName)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// translateNameToIdRepository indicates an expected call of translateNameToIdRepository
func (mr *MockserviceRepositoryInterfaceMockRecorder) translateNameToIdRepository(organizationName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "translateNameToIdRepository", reflect.TypeOf((*MockserviceRepositoryInterface)(nil).translateNameToIdRepository), organizationName)
}

// generateToken mocks base method
func (m *MockserviceRepositoryInterface) generateToken() string {
	ret := m.ctrl.Call(m, "generateToken")
	ret0, _ := ret[0].(string)
	return ret0
}

// generateToken indicates an expected call of generateToken
func (mr *MockserviceRepositoryInterfaceMockRecorder) generateToken() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "generateToken", reflect.TypeOf((*MockserviceRepositoryInterface)(nil).generateToken))
}
