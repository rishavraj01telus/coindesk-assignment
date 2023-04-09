// Code generated by MockGen. DO NOT EDIT.
// Source: coindesk/client (interfaces: ICryptoClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	models "coindesk/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockICryptoClient is a mock of ICryptoClient interface.
type MockICryptoClient struct {
	ctrl     *gomock.Controller
	recorder *MockICryptoClientMockRecorder
}

// MockICryptoClientMockRecorder is the mock recorder for MockICryptoClient.
type MockICryptoClientMockRecorder struct {
	mock *MockICryptoClient
}

// NewMockICryptoClient creates a new mock instance.
func NewMockICryptoClient(ctrl *gomock.Controller) *MockICryptoClient {
	mock := &MockICryptoClient{ctrl: ctrl}
	mock.recorder = &MockICryptoClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICryptoClient) EXPECT() *MockICryptoClientMockRecorder {
	return m.recorder
}

// GetCurrentPrice mocks base method.
func (m *MockICryptoClient) GetCurrentPrice() (models.CoinDeskResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentPrice")
	ret0, _ := ret[0].(models.CoinDeskResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentPrice indicates an expected call of GetCurrentPrice.
func (mr *MockICryptoClientMockRecorder) GetCurrentPrice() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentPrice", reflect.TypeOf((*MockICryptoClient)(nil).GetCurrentPrice))
}
