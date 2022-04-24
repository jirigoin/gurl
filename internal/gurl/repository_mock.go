// Code generated by MockGen. DO NOT EDIT.
// Source: internal/gurl/repository.go

// Package gurl is a generated GoMock package.
package gurl

import (
	os "os"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepositoryImpl is a mock of RepositoryImpl interface.
type MockRepositoryImpl struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryImplMockRecorder
}

// MockRepositoryImplMockRecorder is the mock recorder for MockRepositoryImpl.
type MockRepositoryImplMockRecorder struct {
	mock *MockRepositoryImpl
}

// NewMockRepositoryImpl creates a new mock instance.
func NewMockRepositoryImpl(ctrl *gomock.Controller) *MockRepositoryImpl {
	mock := &MockRepositoryImpl{ctrl: ctrl}
	mock.recorder = &MockRepositoryImplMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryImpl) EXPECT() *MockRepositoryImplMockRecorder {
	return m.recorder
}

// Store mocks base method.
func (m *MockRepositoryImpl) Store(url string, storeFile *os.File) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", url, storeFile)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockRepositoryImplMockRecorder) Store(url, storeFile interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockRepositoryImpl)(nil).Store), url, storeFile)
}