// Code generated by MockGen. DO NOT EDIT.
// Source: message_inputport.go

// Package mock_inputport is a generated GoMock package.
package mock_inputport

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entities "github.com/shima004/chat-server/entities"
	validation "github.com/shima004/chat-server/usecases/inputport/validation"
)

// MockMessageUsecase is a mock of MessageUsecase interface.
type MockMessageUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockMessageUsecaseMockRecorder
}

// MockMessageUsecaseMockRecorder is the mock recorder for MockMessageUsecase.
type MockMessageUsecaseMockRecorder struct {
	mock *MockMessageUsecase
}

// NewMockMessageUsecase creates a new mock instance.
func NewMockMessageUsecase(ctrl *gomock.Controller) *MockMessageUsecase {
	mock := &MockMessageUsecase{ctrl: ctrl}
	mock.recorder = &MockMessageUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessageUsecase) EXPECT() *MockMessageUsecaseMockRecorder {
	return m.recorder
}

// DeleteMessage mocks base method.
func (m *MockMessageUsecase) DeleteMessage(ctx context.Context, in *validation.DeleteMessageInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMessage", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMessage indicates an expected call of DeleteMessage.
func (mr *MockMessageUsecaseMockRecorder) DeleteMessage(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMessage", reflect.TypeOf((*MockMessageUsecase)(nil).DeleteMessage), ctx, in)
}

// FetchMessages mocks base method.
func (m *MockMessageUsecase) FetchMessages(ctx context.Context, in *validation.FatchMessagesInput) ([]*entities.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchMessages", ctx, in)
	ret0, _ := ret[0].([]*entities.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchMessages indicates an expected call of FetchMessages.
func (mr *MockMessageUsecaseMockRecorder) FetchMessages(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchMessages", reflect.TypeOf((*MockMessageUsecase)(nil).FetchMessages), ctx, in)
}

// PostMessage mocks base method.
func (m *MockMessageUsecase) PostMessage(ctx context.Context, in *validation.PostMessageInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostMessage", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostMessage indicates an expected call of PostMessage.
func (mr *MockMessageUsecaseMockRecorder) PostMessage(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostMessage", reflect.TypeOf((*MockMessageUsecase)(nil).PostMessage), ctx, in)
}

// UpdateMessage mocks base method.
func (m *MockMessageUsecase) UpdateMessage(ctx context.Context, in *validation.UpdateMessageInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMessage", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMessage indicates an expected call of UpdateMessage.
func (mr *MockMessageUsecaseMockRecorder) UpdateMessage(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMessage", reflect.TypeOf((*MockMessageUsecase)(nil).UpdateMessage), ctx, in)
}
