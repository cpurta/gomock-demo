// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStorageClient is a mock of StorageClient interface.
type MockStorageClient struct {
	ctrl     *gomock.Controller
	recorder *MockStorageClientMockRecorder
}

// MockStorageClientMockRecorder is the mock recorder for MockStorageClient.
type MockStorageClientMockRecorder struct {
	mock *MockStorageClient
}

// NewMockStorageClient creates a new mock instance.
func NewMockStorageClient(ctrl *gomock.Controller) *MockStorageClient {
	mock := &MockStorageClient{ctrl: ctrl}
	mock.recorder = &MockStorageClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageClient) EXPECT() *MockStorageClientMockRecorder {
	return m.recorder
}

// GetFile mocks base method.
func (m *MockStorageClient) GetFile(ctx context.Context, path string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFile", ctx, path)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFile indicates an expected call of GetFile.
func (mr *MockStorageClientMockRecorder) GetFile(ctx, path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFile", reflect.TypeOf((*MockStorageClient)(nil).GetFile), ctx, path)
}
