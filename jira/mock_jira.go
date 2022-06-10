// Code generated by MockGen. DO NOT EDIT.
// Source: jira/jira.go

// Package mock_jira is a generated GoMock package.
package jira

import (
	reflect "reflect"

	jira "github.com/andygrunwald/go-jira"
	gomock "github.com/golang/mock/gomock"
)

// MockAPI is a mock of API interface.
type MockAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAPIMockRecorder
}

// MockAPIMockRecorder is the mock recorder for MockAPI.
type MockAPIMockRecorder struct {
	mock *MockAPI
}

// NewMockAPI creates a new mock instance.
func NewMockAPI(ctrl *gomock.Controller) *MockAPI {
	mock := &MockAPI{ctrl: ctrl}
	mock.recorder = &MockAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPI) EXPECT() *MockAPIMockRecorder {
	return m.recorder
}

// GetBoard mocks base method.
func (m *MockAPI) GetBoard(name string) (*jira.Board, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBoard", name)
	ret0, _ := ret[0].(*jira.Board)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBoard indicates an expected call of GetBoard.
func (mr *MockAPIMockRecorder) GetBoard(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBoard", reflect.TypeOf((*MockAPI)(nil).GetBoard), name)
}

// GetBoardLayout mocks base method.
func (m *MockAPI) GetBoardLayout(name string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBoardLayout", name)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBoardLayout indicates an expected call of GetBoardLayout.
func (mr *MockAPIMockRecorder) GetBoardLayout(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBoardLayout", reflect.TypeOf((*MockAPI)(nil).GetBoardLayout), name)
}

// GetBoards mocks base method.
func (m *MockAPI) GetBoards() ([]jira.Board, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBoards")
	ret0, _ := ret[0].([]jira.Board)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBoards indicates an expected call of GetBoards.
func (mr *MockAPIMockRecorder) GetBoards() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBoards", reflect.TypeOf((*MockAPI)(nil).GetBoards))
}

// GetIssues mocks base method.
func (m *MockAPI) GetIssues(boardName, sprintName string) ([]jira.Issue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIssues", boardName, sprintName)
	ret0, _ := ret[0].([]jira.Issue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIssues indicates an expected call of GetIssues.
func (mr *MockAPIMockRecorder) GetIssues(boardName, sprintName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssues", reflect.TypeOf((*MockAPI)(nil).GetIssues), boardName, sprintName)
}

// GetIssuesForBoard mocks base method.
func (m *MockAPI) GetIssuesForBoard(boardName string) ([]jira.Issue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIssuesForBoard", boardName)
	ret0, _ := ret[0].([]jira.Issue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIssuesForBoard indicates an expected call of GetIssuesForBoard.
func (mr *MockAPIMockRecorder) GetIssuesForBoard(boardName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssuesForBoard", reflect.TypeOf((*MockAPI)(nil).GetIssuesForBoard), boardName)
}

// GetSprints mocks base method.
func (m *MockAPI) GetSprints(boardName string) ([]jira.Sprint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSprints", boardName)
	ret0, _ := ret[0].([]jira.Sprint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSprints indicates an expected call of GetSprints.
func (mr *MockAPIMockRecorder) GetSprints(boardName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSprints", reflect.TypeOf((*MockAPI)(nil).GetSprints), boardName)
}

// IssueSearch mocks base method.
func (m *MockAPI) IssueSearch(query string, opts *jira.SearchOptions) ([]jira.Issue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IssueSearch", query, opts)
	ret0, _ := ret[0].([]jira.Issue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IssueSearch indicates an expected call of IssueSearch.
func (mr *MockAPIMockRecorder) IssueSearch(query, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IssueSearch", reflect.TypeOf((*MockAPI)(nil).IssueSearch), query, opts)
}
