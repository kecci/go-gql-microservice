// Code generated by MockGen. DO NOT EDIT.
// Source: resolver.go

// Package resolver is a generated GoMock package.
package resolver

import (
	gomock "github.com/golang/mock/gomock"
	health "github.com/kecci/go-gql-microservice/internal/model/health"
	reflect "reflect"
)

// MockhealthServiceInterface is a mock of healthServiceInterface interface
type MockhealthServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockhealthServiceInterfaceMockRecorder
}

// MockhealthServiceInterfaceMockRecorder is the mock recorder for MockhealthServiceInterface
type MockhealthServiceInterfaceMockRecorder struct {
	mock *MockhealthServiceInterface
}

// NewMockhealthServiceInterface creates a new mock instance
func NewMockhealthServiceInterface(ctrl *gomock.Controller) *MockhealthServiceInterface {
	mock := &MockhealthServiceInterface{ctrl: ctrl}
	mock.recorder = &MockhealthServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockhealthServiceInterface) EXPECT() *MockhealthServiceInterfaceMockRecorder {
	return m.recorder
}

// CheckHealth mocks base method
func (m *MockhealthServiceInterface) CheckHealth() (*health.Health, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckHealth")
	ret0, _ := ret[0].(*health.Health)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckHealth indicates an expected call of CheckHealth
func (mr *MockhealthServiceInterfaceMockRecorder) CheckHealth() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckHealth", reflect.TypeOf((*MockhealthServiceInterface)(nil).CheckHealth))
}