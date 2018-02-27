// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Social-projects-Rivne/Rv-029.Go/backend/models (interfaces: SprintCRUD)

// Package mocks is a generated GoMock package.
package mocks

import (
	models "github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	gocql "github.com/gocql/gocql"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSprintCRUD is a mock of SprintCRUD interface
type MockSprintCRUD struct {
	ctrl     *gomock.Controller
	recorder *MockSprintCRUDMockRecorder
}

// MockSprintCRUDMockRecorder is the mock recorder for MockSprintCRUD
type MockSprintCRUDMockRecorder struct {
	mock *MockSprintCRUD
}

// NewMockSprintCRUD creates a new mock instance
func NewMockSprintCRUD(ctrl *gomock.Controller) *MockSprintCRUD {
	mock := &MockSprintCRUD{ctrl: ctrl}
	mock.recorder = &MockSprintCRUDMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSprintCRUD) EXPECT() *MockSprintCRUDMockRecorder {
	return m.recorder
}

// Delete mocks base method
func (m *MockSprintCRUD) Delete(arg0 *models.Sprint) error {
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockSprintCRUDMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSprintCRUD)(nil).Delete), arg0)
}

// FindByID mocks base method
func (m *MockSprintCRUD) FindByID(arg0 *models.Sprint) error {
	ret := m.ctrl.Call(m, "FindByID", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindByID indicates an expected call of FindByID
func (mr *MockSprintCRUDMockRecorder) FindByID(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockSprintCRUD)(nil).FindByID), arg0)
}

// GetSprintIssuesInProgress mocks base method
func (m *MockSprintCRUD) GetSprintIssuesInProgress(arg0 *models.Sprint) ([]models.Issue, error) {
	ret := m.ctrl.Call(m, "GetSprintIssuesInProgress", arg0)
	ret0, _ := ret[0].([]models.Issue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSprintIssuesInProgress indicates an expected call of GetSprintIssuesInProgress
func (mr *MockSprintCRUDMockRecorder) GetSprintIssuesInProgress(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSprintIssuesInProgress", reflect.TypeOf((*MockSprintCRUD)(nil).GetSprintIssuesInProgress), arg0)
}

// Insert mocks base method
func (m *MockSprintCRUD) Insert(arg0 *models.Sprint) error {
	ret := m.ctrl.Call(m, "Insert", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert
func (mr *MockSprintCRUDMockRecorder) Insert(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockSprintCRUD)(nil).Insert), arg0)
}

// List mocks base method
func (m *MockSprintCRUD) List(arg0 gocql.UUID) ([]map[string]interface{}, error) {
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].([]map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockSprintCRUDMockRecorder) List(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSprintCRUD)(nil).List), arg0)
}

// Update mocks base method
func (m *MockSprintCRUD) Update(arg0 *models.Sprint) error {
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockSprintCRUDMockRecorder) Update(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSprintCRUD)(nil).Update), arg0)
}
