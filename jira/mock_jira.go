// Code generated by MockGen. DO NOT EDIT.
// Source: jira/jira.go

// Package jira is a generated GoMock package.
package jira

import (
	reflect "reflect"

	go_jira "github.com/andygrunwald/go-jira"
	gomock "github.com/golang/mock/gomock"
)

// MockAPI is a mock of API interface
type MockAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAPIMockRecorder
}

// MockAPIMockRecorder is the mock recorder for MockAPI
type MockAPIMockRecorder struct {
	mock *MockAPI
}

// NewMockAPI creates a new mock instance
func NewMockAPI(ctrl *gomock.Controller) *MockAPI {
	mock := &MockAPI{ctrl: ctrl}
	mock.recorder = &MockAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAPI) EXPECT() *MockAPIMockRecorder {
	return m.recorder
}

// GetBoards mocks base method
func (m *MockAPI) GetBoards() ([]go_jira.Board, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBoards")
	ret0, _ := ret[0].([]go_jira.Board)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBoards indicates an expected call of GetBoards
func (mr *MockAPIMockRecorder) GetBoards() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBoards", reflect.TypeOf((*MockAPI)(nil).GetBoards))
}

// GetSprints mocks base method
func (m *MockAPI) GetSprints(boardName string) ([]go_jira.Sprint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSprints", boardName)
	ret0, _ := ret[0].([]go_jira.Sprint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSprints indicates an expected call of GetSprints
func (mr *MockAPIMockRecorder) GetSprints(boardName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSprints", reflect.TypeOf((*MockAPI)(nil).GetSprints), boardName)
}