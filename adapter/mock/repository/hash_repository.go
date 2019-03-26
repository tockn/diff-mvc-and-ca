// Code generated by MockGen. DO NOT EDIT.
// Source: domain/repository/hash_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockHash is a mock of Hash interface
type MockHash struct {
	ctrl     *gomock.Controller
	recorder *MockHashMockRecorder
}

// MockHashMockRecorder is the mock recorder for MockHash
type MockHashMockRecorder struct {
	mock *MockHash
}

// NewMockHash creates a new mock instance
func NewMockHash(ctrl *gomock.Controller) *MockHash {
	mock := &MockHash{ctrl: ctrl}
	mock.recorder = &MockHashMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHash) EXPECT() *MockHashMockRecorder {
	return m.recorder
}

// Encode mocks base method
func (m *MockHash) Encode(id int64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Encode", id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Encode indicates an expected call of Encode
func (mr *MockHashMockRecorder) Encode(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encode", reflect.TypeOf((*MockHash)(nil).Encode), id)
}

// Decode mocks base method
func (m *MockHash) Decode(idStr string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode", idStr)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decode indicates an expected call of Decode
func (mr *MockHashMockRecorder) Decode(idStr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockHash)(nil).Decode), idStr)
}
