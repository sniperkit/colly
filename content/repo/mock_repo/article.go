// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/urandom/readeef/content/repo (interfaces: Article)

// Package mock_repo is a generated GoMock package.
package mock_repo

import (
	gomock "github.com/golang/mock/gomock"
	content "github.com/urandom/readeef/content"
	reflect "reflect"
)

// MockArticle is a mock of Article interface
type MockArticle struct {
	ctrl     *gomock.Controller
	recorder *MockArticleMockRecorder
}

// MockArticleMockRecorder is the mock recorder for MockArticle
type MockArticleMockRecorder struct {
	mock *MockArticle
}

// NewMockArticle creates a new mock instance
func NewMockArticle(ctrl *gomock.Controller) *MockArticle {
	mock := &MockArticle{ctrl: ctrl}
	mock.recorder = &MockArticleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockArticle) EXPECT() *MockArticleMockRecorder {
	return m.recorder
}

// All mocks base method
func (m *MockArticle) All(arg0 ...content.QueryOpt) ([]content.Article, error) {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "All", varargs...)
	ret0, _ := ret[0].([]content.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All
func (mr *MockArticleMockRecorder) All(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockArticle)(nil).All), arg0...)
}

// Count mocks base method
func (m *MockArticle) Count(arg0 content.User, arg1 ...content.QueryOpt) (int64, error) {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Count", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count
func (mr *MockArticleMockRecorder) Count(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockArticle)(nil).Count), varargs...)
}

// Favor mocks base method
func (m *MockArticle) Favor(arg0 bool, arg1 content.User, arg2 ...content.QueryOpt) error {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Favor", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Favor indicates an expected call of Favor
func (mr *MockArticleMockRecorder) Favor(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Favor", reflect.TypeOf((*MockArticle)(nil).Favor), varargs...)
}

// ForUser mocks base method
func (m *MockArticle) ForUser(arg0 content.User, arg1 ...content.QueryOpt) ([]content.Article, error) {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ForUser", varargs...)
	ret0, _ := ret[0].([]content.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ForUser indicates an expected call of ForUser
func (mr *MockArticleMockRecorder) ForUser(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForUser", reflect.TypeOf((*MockArticle)(nil).ForUser), varargs...)
}

// IDs mocks base method
func (m *MockArticle) IDs(arg0 content.User, arg1 ...content.QueryOpt) ([]content.ArticleID, error) {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IDs", varargs...)
	ret0, _ := ret[0].([]content.ArticleID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IDs indicates an expected call of IDs
func (mr *MockArticleMockRecorder) IDs(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IDs", reflect.TypeOf((*MockArticle)(nil).IDs), varargs...)
}

// Read mocks base method
func (m *MockArticle) Read(arg0 bool, arg1 content.User, arg2 ...content.QueryOpt) error {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Read", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Read indicates an expected call of Read
func (mr *MockArticleMockRecorder) Read(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockArticle)(nil).Read), varargs...)
}

// RemoveStaleUnreadRecords mocks base method
func (m *MockArticle) RemoveStaleUnreadRecords() error {
	ret := m.ctrl.Call(m, "RemoveStaleUnreadRecords")
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveStaleUnreadRecords indicates an expected call of RemoveStaleUnreadRecords
func (mr *MockArticleMockRecorder) RemoveStaleUnreadRecords() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveStaleUnreadRecords", reflect.TypeOf((*MockArticle)(nil).RemoveStaleUnreadRecords))
}
