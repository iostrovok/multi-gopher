// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/iostrovok/conveyor/faces (interfaces: ITestObject)

// Package mmock is a generated GoMock package.
package mmock

import (
	gomock "github.com/golang/mock/gomock"
	check "github.com/iostrovok/check"
	reflect "reflect"
)

// MockITestObject is a mock of ITestObject interface
type MockITestObject struct {
	ctrl     *gomock.Controller
	recorder *MockITestObjectMockRecorder
}

// MockITestObjectMockRecorder is the mock recorder for MockITestObject
type MockITestObjectMockRecorder struct {
	mock *MockITestObject
}

// NewMockITestObject creates a new mock instance
func NewMockITestObject(ctrl *gomock.Controller) *MockITestObject {
	mock := &MockITestObject{ctrl: ctrl}
	mock.recorder = &MockITestObjectMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockITestObject) EXPECT() *MockITestObjectMockRecorder {
	return m.recorder
}

// IsTestMode mocks base method
func (m *MockITestObject) IsTestMode() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsTestMode")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsTestMode indicates an expected call of IsTestMode
func (mr *MockITestObjectMockRecorder) IsTestMode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsTestMode", reflect.TypeOf((*MockITestObject)(nil).IsTestMode))
}

// Suffix mocks base method
func (m *MockITestObject) Suffix() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Suffix")
	ret0, _ := ret[0].(string)
	return ret0
}

// Suffix indicates an expected call of Suffix
func (mr *MockITestObjectMockRecorder) Suffix() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Suffix", reflect.TypeOf((*MockITestObject)(nil).Suffix))
}

// TestObject mocks base method
func (m *MockITestObject) TestObject() *check.C {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TestObject")
	ret0, _ := ret[0].(*check.C)
	return ret0
}

// TestObject indicates an expected call of TestObject
func (mr *MockITestObjectMockRecorder) TestObject() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TestObject", reflect.TypeOf((*MockITestObject)(nil).TestObject))
}
