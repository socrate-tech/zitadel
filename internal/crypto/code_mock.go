// Code generated by MockGen. DO NOT EDIT.
// Source: code.go
//
// Generated by this command:
//
//	mockgen -source code.go -destination ./code_mock.go -package crypto
//

// Package crypto is a generated GoMock package.
package crypto

import (
	reflect "reflect"
	time "time"

	gomock "go.uber.org/mock/gomock"
)

// MockGenerator is a mock of Generator interface.
type MockGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockGeneratorMockRecorder
	isgomock struct{}
}

// MockGeneratorMockRecorder is the mock recorder for MockGenerator.
type MockGeneratorMockRecorder struct {
	mock *MockGenerator
}

// NewMockGenerator creates a new mock instance.
func NewMockGenerator(ctrl *gomock.Controller) *MockGenerator {
	mock := &MockGenerator{ctrl: ctrl}
	mock.recorder = &MockGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGenerator) EXPECT() *MockGeneratorMockRecorder {
	return m.recorder
}

// Alg mocks base method.
func (m *MockGenerator) Alg() EncryptionAlgorithm {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Alg")
	ret0, _ := ret[0].(EncryptionAlgorithm)
	return ret0
}

// Alg indicates an expected call of Alg.
func (mr *MockGeneratorMockRecorder) Alg() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Alg", reflect.TypeOf((*MockGenerator)(nil).Alg))
}

// Expiry mocks base method.
func (m *MockGenerator) Expiry() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Expiry")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// Expiry indicates an expected call of Expiry.
func (mr *MockGeneratorMockRecorder) Expiry() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Expiry", reflect.TypeOf((*MockGenerator)(nil).Expiry))
}

// Length mocks base method.
func (m *MockGenerator) Length() uint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(uint)
	return ret0
}

// Length indicates an expected call of Length.
func (mr *MockGeneratorMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockGenerator)(nil).Length))
}

// Runes mocks base method.
func (m *MockGenerator) Runes() []rune {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Runes")
	ret0, _ := ret[0].([]rune)
	return ret0
}

// Runes indicates an expected call of Runes.
func (mr *MockGeneratorMockRecorder) Runes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Runes", reflect.TypeOf((*MockGenerator)(nil).Runes))
}
