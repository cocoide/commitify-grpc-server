// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	entity "github.com/cocoide/commitify-grpc-server/internal/domain/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockLangService is a mock of LangService interface.
type MockLangService struct {
	ctrl     *gomock.Controller
	recorder *MockLangServiceMockRecorder
}

// MockLangServiceMockRecorder is the mock recorder for MockLangService.
type MockLangServiceMockRecorder struct {
	mock *MockLangService
}

// NewMockLangService creates a new mock instance.
func NewMockLangService(ctrl *gomock.Controller) *MockLangService {
	mock := &MockLangService{ctrl: ctrl}
	mock.recorder = &MockLangServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLangService) EXPECT() *MockLangServiceMockRecorder {
	return m.recorder
}

// TranslateTexts mocks base method.
func (m *MockLangService) TranslateTexts(texts []string, into entity.LanguageType) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TranslateTexts", texts, into)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TranslateTexts indicates an expected call of TranslateTexts.
func (mr *MockLangServiceMockRecorder) TranslateTexts(texts, into interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TranslateTexts", reflect.TypeOf((*MockLangService)(nil).TranslateTexts), texts, into)
}

// TranslateTextsIntoJapanese mocks base method.
func (m *MockLangService) TranslateTextsIntoJapanese(texts []string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TranslateTextsIntoJapanese", texts)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TranslateTextsIntoJapanese indicates an expected call of TranslateTextsIntoJapanese.
func (mr *MockLangServiceMockRecorder) TranslateTextsIntoJapanese(texts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TranslateTextsIntoJapanese", reflect.TypeOf((*MockLangService)(nil).TranslateTextsIntoJapanese), texts)
}

// MockNLPService is a mock of NLPService interface.
type MockNLPService struct {
	ctrl     *gomock.Controller
	recorder *MockNLPServiceMockRecorder
}

// MockNLPServiceMockRecorder is the mock recorder for MockNLPService.
type MockNLPServiceMockRecorder struct {
	mock *MockNLPService
}

// NewMockNLPService creates a new mock instance.
func NewMockNLPService(ctrl *gomock.Controller) *MockNLPService {
	mock := &MockNLPService{ctrl: ctrl}
	mock.recorder = &MockNLPServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNLPService) EXPECT() *MockNLPServiceMockRecorder {
	return m.recorder
}

// GetAnswerFromPrompt mocks base method.
func (m *MockNLPService) GetAnswerFromPrompt(prompt string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAnswerFromPrompt", prompt)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAnswerFromPrompt indicates an expected call of GetAnswerFromPrompt.
func (mr *MockNLPServiceMockRecorder) GetAnswerFromPrompt(prompt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAnswerFromPrompt", reflect.TypeOf((*MockNLPService)(nil).GetAnswerFromPrompt), prompt)
}