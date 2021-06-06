// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ozoncp/ocp-project-api/internal/alarm (interfaces: Alarm)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockAlarm is a mock of Alarm interface.
type MockAlarm struct {
	ctrl     *gomock.Controller
	recorder *MockAlarmMockRecorder
}

// MockAlarmMockRecorder is the mock recorder for MockAlarm.
type MockAlarmMockRecorder struct {
	mock *MockAlarm
}

// NewMockAlarm creates a new mock instance.
func NewMockAlarm(ctrl *gomock.Controller) *MockAlarm {
	mock := &MockAlarm{ctrl: ctrl}
	mock.recorder = &MockAlarmMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAlarm) EXPECT() *MockAlarmMockRecorder {
	return m.recorder
}

// Alarms mocks base method.
func (m *MockAlarm) Alarms() <-chan struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Alarms")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// Alarms indicates an expected call of Alarms.
func (mr *MockAlarmMockRecorder) Alarms() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Alarms", reflect.TypeOf((*MockAlarm)(nil).Alarms))
}

// Close mocks base method.
func (m *MockAlarm) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockAlarmMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockAlarm)(nil).Close))
}

// NewTimeout mocks base method.
func (m *MockAlarm) ResetTimeout(arg0 time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ResetTimeout", arg0)
}

// NewTimeout indicates an expected call of NewTimeout.
func (mr *MockAlarmMockRecorder) NewTimeout(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetTimeout", reflect.TypeOf((*MockAlarm)(nil).ResetTimeout), arg0)
}
