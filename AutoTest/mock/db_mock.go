// Code generated by MockGen. DO NOT EDIT.
// Source: D:\code\project_based_learning\AutoTest\mock\db.go

// Package mock_main is a generated GoMock package.
package main

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockOrderDBI is a mock of OrderDBI interface.
type MockOrderDBI struct {
	ctrl     *gomock.Controller
	recorder *MockOrderDBIMockRecorder
}

// MockOrderDBIMockRecorder is the mock recorder for MockOrderDBI.
type MockOrderDBIMockRecorder struct {
	mock *MockOrderDBI
}

// NewMockOrderDBI creates a new mock instance.
func NewMockOrderDBI(ctrl *gomock.Controller) *MockOrderDBI {
	mock := &MockOrderDBI{ctrl: ctrl}
	mock.recorder = &MockOrderDBIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderDBI) EXPECT() *MockOrderDBIMockRecorder {
	return m.recorder
}

// GetName mocks base method.
func (m *MockOrderDBI) GetName(orderid int) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName", orderid)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockOrderDBIMockRecorder) GetName(orderid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockOrderDBI)(nil).GetName), orderid)
}
