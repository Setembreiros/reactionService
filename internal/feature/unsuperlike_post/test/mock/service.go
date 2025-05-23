// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_unsuperlike_post is a generated GoMock package.
package mock_unsuperlike_post

import (
	model "reactionservice/internal/model/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// DeleteSuperlikePost mocks base method.
func (m *MockRepository) DeleteSuperlikePost(superlike *model.SuperlikePost) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSuperlikePost", superlike)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSuperlikePost indicates an expected call of DeleteSuperlikePost.
func (mr *MockRepositoryMockRecorder) DeleteSuperlikePost(superlike interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSuperlikePost", reflect.TypeOf((*MockRepository)(nil).DeleteSuperlikePost), superlike)
}
