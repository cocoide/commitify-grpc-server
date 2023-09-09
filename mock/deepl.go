// Code generated by MockGen. DO NOT EDIT.
// Source: deepl.go

// Package mock_gateway is a generated GoMock package.
package mock_gateway

import (
	reflect "reflect"

	gateway "github.com/cocoide/commitify-grpc-server/pkg/gateway"
	gomock "github.com/golang/mock/gomock"
)

// MockDeeplAPIGateway is a mock of DeeplAPIGateway interface.
type MockDeeplAPIGateway struct {
	ctrl     *gomock.Controller
	recorder *MockDeeplAPIGatewayMockRecorder
}

// MockDeeplAPIGatewayMockRecorder is the mock recorder for MockDeeplAPIGateway.
type MockDeeplAPIGatewayMockRecorder struct {
	mock *MockDeeplAPIGateway
}

// NewMockDeeplAPIGateway creates a new mock instance.
func NewMockDeeplAPIGateway(ctrl *gomock.Controller) *MockDeeplAPIGateway {
	mock := &MockDeeplAPIGateway{ctrl: ctrl}
	mock.recorder = &MockDeeplAPIGatewayMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeeplAPIGateway) EXPECT() *MockDeeplAPIGatewayMockRecorder {
	return m.recorder
}

// TranslateTexts mocks base method.
func (m *MockDeeplAPIGateway) TranslateTexts(texts []string, into gateway.Language) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TranslateTexts", texts, into)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TranslateTexts indicates an expected call of TranslateTexts.
func (mr *MockDeeplAPIGatewayMockRecorder) TranslateTexts(texts, into interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TranslateTexts", reflect.TypeOf((*MockDeeplAPIGateway)(nil).TranslateTexts), texts, into)
}

// TranslateTextsIntoJapanese mocks base method.
func (m *MockDeeplAPIGateway) TranslateTextsIntoJapanese(texts []string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TranslateTextsIntoJapanese", texts)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TranslateTextsIntoJapanese indicates an expected call of TranslateTextsIntoJapanese.
func (mr *MockDeeplAPIGatewayMockRecorder) TranslateTextsIntoJapanese(texts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TranslateTextsIntoJapanese", reflect.TypeOf((*MockDeeplAPIGateway)(nil).TranslateTextsIntoJapanese), texts)
}