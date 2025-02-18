// Code generated by MockGen. DO NOT EDIT.
// Source: subject_black_list.go

// Package mock is a generated GoMock package.
package mock

import (
	dao "iam/pkg/database/dao"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sqlx "github.com/jmoiron/sqlx"
)

// MockSubjectBlackListManager is a mock of SubjectBlackListManager interface.
type MockSubjectBlackListManager struct {
	ctrl     *gomock.Controller
	recorder *MockSubjectBlackListManagerMockRecorder
}

// MockSubjectBlackListManagerMockRecorder is the mock recorder for MockSubjectBlackListManager.
type MockSubjectBlackListManagerMockRecorder struct {
	mock *MockSubjectBlackListManager
}

// NewMockSubjectBlackListManager creates a new mock instance.
func NewMockSubjectBlackListManager(ctrl *gomock.Controller) *MockSubjectBlackListManager {
	mock := &MockSubjectBlackListManager{ctrl: ctrl}
	mock.recorder = &MockSubjectBlackListManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubjectBlackListManager) EXPECT() *MockSubjectBlackListManagerMockRecorder {
	return m.recorder
}

// BulkCreate mocks base method.
func (m *MockSubjectBlackListManager) BulkCreate(subjectBlackList []dao.SubjectBlackList) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkCreate", subjectBlackList)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkCreate indicates an expected call of BulkCreate.
func (mr *MockSubjectBlackListManagerMockRecorder) BulkCreate(subjectBlackList interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkCreate", reflect.TypeOf((*MockSubjectBlackListManager)(nil).BulkCreate), subjectBlackList)
}

// BulkDelete mocks base method.
func (m *MockSubjectBlackListManager) BulkDelete(subjectPKs []int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkDelete", subjectPKs)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkDelete indicates an expected call of BulkDelete.
func (mr *MockSubjectBlackListManagerMockRecorder) BulkDelete(subjectPKs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkDelete", reflect.TypeOf((*MockSubjectBlackListManager)(nil).BulkDelete), subjectPKs)
}

// BulkDeleteWithTx mocks base method.
func (m *MockSubjectBlackListManager) BulkDeleteWithTx(tx *sqlx.Tx, subjectPKs []int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkDeleteWithTx", tx, subjectPKs)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkDeleteWithTx indicates an expected call of BulkDeleteWithTx.
func (mr *MockSubjectBlackListManagerMockRecorder) BulkDeleteWithTx(tx, subjectPKs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkDeleteWithTx", reflect.TypeOf((*MockSubjectBlackListManager)(nil).BulkDeleteWithTx), tx, subjectPKs)
}

// ListSubjectPK mocks base method.
func (m *MockSubjectBlackListManager) ListSubjectPK() ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSubjectPK")
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSubjectPK indicates an expected call of ListSubjectPK.
func (mr *MockSubjectBlackListManagerMockRecorder) ListSubjectPK() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSubjectPK", reflect.TypeOf((*MockSubjectBlackListManager)(nil).ListSubjectPK))
}
