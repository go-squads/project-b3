// Code generated by MockGen. DO NOT EDIT.
// Source: organization/organization_processor.go

// Package organization is a generated GoMock package.
package organization

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

// createNewOrganizationProcessor mocks base method
func (m *Mockprocessor) createNewOrganizationProcessor(orginizationName string, userId int64) error {
	ret := m.ctrl.Call(m, "createNewOrganizationProcessor", orginizationName, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// createNewOrganizationProcessor indicates an expected call of createNewOrganizationProcessor
func (mr *MockprocessorMockRecorder) createNewOrganizationProcessor(orginizationName, userId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "createNewOrganizationProcessor", reflect.TypeOf((*Mockprocessor)(nil).createNewOrganizationProcessor), orginizationName, userId)
}

// addUserProcessor mocks base method
func (m *Mockprocessor) addUserProcessor(member *Member) error {
	ret := m.ctrl.Call(m, "addUserProcessor", member)
	ret0, _ := ret[0].(error)
	return ret0
}

// addUserProcessor indicates an expected call of addUserProcessor
func (mr *MockprocessorMockRecorder) addUserProcessor(member interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "addUserProcessor", reflect.TypeOf((*Mockprocessor)(nil).addUserProcessor), member)
}

// deleteUserFromGroupProcessor mocks base method
func (m *Mockprocessor) deleteUserFromGroupProcessor(organizationId, userId int64) error {
	ret := m.ctrl.Call(m, "deleteUserFromGroupProcessor", organizationId, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// deleteUserFromGroupProcessor indicates an expected call of deleteUserFromGroupProcessor
func (mr *MockprocessorMockRecorder) deleteUserFromGroupProcessor(organizationId, userId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "deleteUserFromGroupProcessor", reflect.TypeOf((*Mockprocessor)(nil).deleteUserFromGroupProcessor), organizationId, userId)
}

// updateRoleOfUserProcessor mocks base method
func (m *Mockprocessor) updateRoleOfUserProcessor(member *Member) error {
	ret := m.ctrl.Call(m, "updateRoleOfUserProcessor", member)
	ret0, _ := ret[0].(error)
	return ret0
}

// updateRoleOfUserProcessor indicates an expected call of updateRoleOfUserProcessor
func (mr *MockprocessorMockRecorder) updateRoleOfUserProcessor(member interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "updateRoleOfUserProcessor", reflect.TypeOf((*Mockprocessor)(nil).updateRoleOfUserProcessor), member)
}

// getAllMemberOfOrganizationProcessor mocks base method
func (m *Mockprocessor) getAllMemberOfOrganizationProcessor(organizationId int64) ([]map[string]interface{}, error) {
	ret := m.ctrl.Call(m, "getAllMemberOfOrganizationProcessor", organizationId)
	ret0, _ := ret[0].([]map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getAllMemberOfOrganizationProcessor indicates an expected call of getAllMemberOfOrganizationProcessor
func (mr *MockprocessorMockRecorder) getAllMemberOfOrganizationProcessor(organizationId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getAllMemberOfOrganizationProcessor", reflect.TypeOf((*Mockprocessor)(nil).getAllMemberOfOrganizationProcessor), organizationId)
}
