// Source: api/internal/storage (interfaces: RepositoryDB)

// Package mocks is a generated GoMock package.
package mocks

import (
	storage "api/internal/storage"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockRepositoryDB is a mock of RepositoryDB interface.
type MockRepositoryDB struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryDBMockRecorder
}

// MockRepositoryDBMockRecorder is the mock recorder for MockRepositoryDB.
type MockRepositoryDBMockRecorder struct {
	mock *MockRepositoryDB
}

// NewMockRepositoryDB creates a new mock instance.
func NewMockRepositoryDB(ctrl *gomock.Controller) *MockRepositoryDB {
	mock := &MockRepositoryDB{ctrl: ctrl}
	mock.recorder = &MockRepositoryDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryDB) EXPECT() *MockRepositoryDBMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockRepositoryDB) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockRepositoryDBMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRepositoryDB)(nil).Close))
}

// Create mocks base method.
func (m *MockRepositoryDB) Create(arg0 storage.User) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryDBMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepositoryDB)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockRepositoryDB) Delete(arg0 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRepositoryDBMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepositoryDB)(nil).Delete), arg0)
}

// Read mocks base method.
func (m *MockRepositoryDB) Read(arg0 uuid.UUID) (storage.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(storage.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockRepositoryDBMockRecorder) Read(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockRepositoryDB)(nil).Read), arg0)
}

// Update mocks base method.
func (m *MockRepositoryDB) Update(arg0 storage.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockRepositoryDBMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepositoryDB)(nil).Update), arg0)
}
