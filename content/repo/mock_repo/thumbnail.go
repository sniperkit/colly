// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/urandom/readeef/content/repo (interfaces: Thumbnail)

// Package mock_repo is a generated GoMock package.
package mock_repo

import (
	gomock "github.com/golang/mock/gomock"
	content "github.com/urandom/readeef/content"
	reflect "reflect"
)

// MockThumbnail is a mock of Thumbnail interface
type MockThumbnail struct {
	ctrl     *gomock.Controller
	recorder *MockThumbnailMockRecorder
}

// MockThumbnailMockRecorder is the mock recorder for MockThumbnail
type MockThumbnailMockRecorder struct {
	mock *MockThumbnail
}

// NewMockThumbnail creates a new mock instance
func NewMockThumbnail(ctrl *gomock.Controller) *MockThumbnail {
	mock := &MockThumbnail{ctrl: ctrl}
	mock.recorder = &MockThumbnailMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockThumbnail) EXPECT() *MockThumbnailMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockThumbnail) Get(arg0 content.Article) (content.Thumbnail, error) {
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(content.Thumbnail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockThumbnailMockRecorder) Get(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockThumbnail)(nil).Get), arg0)
}

// Update mocks base method
func (m *MockThumbnail) Update(arg0 content.Thumbnail) error {
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockThumbnailMockRecorder) Update(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockThumbnail)(nil).Update), arg0)
}
