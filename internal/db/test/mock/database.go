// Code generated by MockGen. DO NOT EDIT.
// Source: database.go

// Package mock_database is a generated GoMock package.
package mock_database

import (
	model "reactionservice/internal/model/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDatabaseClient is a mock of DatabaseClient interface.
type MockDatabaseClient struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseClientMockRecorder
}

// MockDatabaseClientMockRecorder is the mock recorder for MockDatabaseClient.
type MockDatabaseClientMockRecorder struct {
	mock *MockDatabaseClient
}

// NewMockDatabaseClient creates a new mock instance.
func NewMockDatabaseClient(ctrl *gomock.Controller) *MockDatabaseClient {
	mock := &MockDatabaseClient{ctrl: ctrl}
	mock.recorder = &MockDatabaseClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabaseClient) EXPECT() *MockDatabaseClientMockRecorder {
	return m.recorder
}

// Clean mocks base method.
func (m *MockDatabaseClient) Clean() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Clean")
}

// Clean indicates an expected call of Clean.
func (mr *MockDatabaseClientMockRecorder) Clean() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clean", reflect.TypeOf((*MockDatabaseClient)(nil).Clean))
}

// CreateLikePost mocks base method.
func (m *MockDatabaseClient) CreateLikePost(data *model.LikePost) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLikePost", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateLikePost indicates an expected call of CreateLikePost.
func (mr *MockDatabaseClientMockRecorder) CreateLikePost(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLikePost", reflect.TypeOf((*MockDatabaseClient)(nil).CreateLikePost), data)
}

// CreateReview mocks base method.
func (m *MockDatabaseClient) CreateReview(data *model.Review) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateReview", data)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateReview indicates an expected call of CreateReview.
func (mr *MockDatabaseClientMockRecorder) CreateReview(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReview", reflect.TypeOf((*MockDatabaseClient)(nil).CreateReview), data)
}

// CreateSuperlikePost mocks base method.
func (m *MockDatabaseClient) CreateSuperlikePost(data *model.SuperlikePost) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSuperlikePost", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSuperlikePost indicates an expected call of CreateSuperlikePost.
func (mr *MockDatabaseClientMockRecorder) CreateSuperlikePost(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSuperlikePost", reflect.TypeOf((*MockDatabaseClient)(nil).CreateSuperlikePost), data)
}

// DeleteLikePost mocks base method.
func (m *MockDatabaseClient) DeleteLikePost(data *model.LikePost) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLikePost", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteLikePost indicates an expected call of DeleteLikePost.
func (mr *MockDatabaseClientMockRecorder) DeleteLikePost(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLikePost", reflect.TypeOf((*MockDatabaseClient)(nil).DeleteLikePost), data)
}

// DeleteSuperlikePost mocks base method.
func (m *MockDatabaseClient) DeleteSuperlikePost(data *model.SuperlikePost) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSuperlikePost", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSuperlikePost indicates an expected call of DeleteSuperlikePost.
func (mr *MockDatabaseClientMockRecorder) DeleteSuperlikePost(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSuperlikePost", reflect.TypeOf((*MockDatabaseClient)(nil).DeleteSuperlikePost), data)
}

// GetLikePost mocks base method.
func (m *MockDatabaseClient) GetLikePost(postId, username string) (*model.LikePost, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLikePost", postId, username)
	ret0, _ := ret[0].(*model.LikePost)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLikePost indicates an expected call of GetLikePost.
func (mr *MockDatabaseClientMockRecorder) GetLikePost(postId, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLikePost", reflect.TypeOf((*MockDatabaseClient)(nil).GetLikePost), postId, username)
}

// GetReviewById mocks base method.
func (m *MockDatabaseClient) GetReviewById(id uint64) (*model.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReviewById", id)
	ret0, _ := ret[0].(*model.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReviewById indicates an expected call of GetReviewById.
func (mr *MockDatabaseClientMockRecorder) GetReviewById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReviewById", reflect.TypeOf((*MockDatabaseClient)(nil).GetReviewById), id)
}

// GetSuperlikePost mocks base method.
func (m *MockDatabaseClient) GetSuperlikePost(postId, username string) (*model.SuperlikePost, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSuperlikePost", postId, username)
	ret0, _ := ret[0].(*model.SuperlikePost)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSuperlikePost indicates an expected call of GetSuperlikePost.
func (mr *MockDatabaseClientMockRecorder) GetSuperlikePost(postId, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSuperlikePost", reflect.TypeOf((*MockDatabaseClient)(nil).GetSuperlikePost), postId, username)
}
