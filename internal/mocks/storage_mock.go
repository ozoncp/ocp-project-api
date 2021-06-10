// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ozoncp/ocp-project-api/internal/storage (interfaces: RepoStorage,ProjectStorage)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/ozoncp/ocp-project-api/internal/models"
)

// MockRepoStorage is a mock of RepoStorage interface.
type MockRepoStorage struct {
	ctrl     *gomock.Controller
	recorder *MockRepoStorageMockRecorder
}

// MockRepoStorageMockRecorder is the mock recorder for MockRepoStorage.
type MockRepoStorageMockRecorder struct {
	mock *MockRepoStorage
}

// NewMockRepoStorage creates a new mock instance.
func NewMockRepoStorage(ctrl *gomock.Controller) *MockRepoStorage {
	mock := &MockRepoStorage{ctrl: ctrl}
	mock.recorder = &MockRepoStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepoStorage) EXPECT() *MockRepoStorageMockRecorder {
	return m.recorder
}

// AddRepos mocks base method.
func (m *MockRepoStorage) AddRepos(arg0 []models.Repo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRepos", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRepos indicates an expected call of AddRepos.
func (mr *MockRepoStorageMockRecorder) AddRepos(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRepos", reflect.TypeOf((*MockRepoStorage)(nil).AddRepos), arg0)
}

// DescribeRepo mocks base method.
func (m *MockRepoStorage) DescribeRepo(arg0 uint64) (*models.Repo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeRepo", arg0)
	ret0, _ := ret[0].(*models.Repo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeRepo indicates an expected call of DescribeRepo.
func (mr *MockRepoStorageMockRecorder) DescribeRepo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeRepo", reflect.TypeOf((*MockRepoStorage)(nil).DescribeRepo), arg0)
}

// ListRepos mocks base method.
func (m *MockRepoStorage) ListRepos(arg0, arg1 uint64) ([]models.Repo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRepos", arg0, arg1)
	ret0, _ := ret[0].([]models.Repo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRepos indicates an expected call of ListRepos.
func (mr *MockRepoStorageMockRecorder) ListRepos(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRepos", reflect.TypeOf((*MockRepoStorage)(nil).ListRepos), arg0, arg1)
}

// RemoveRepo mocks base method.
func (m *MockRepoStorage) RemoveRepo(arg0 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveRepo", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveRepo indicates an expected call of RemoveRepo.
func (mr *MockRepoStorageMockRecorder) RemoveRepo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveRepo", reflect.TypeOf((*MockRepoStorage)(nil).RemoveRepo), arg0)
}

// MockProjectStorage is a mock of ProjectStorage interface.
type MockProjectStorage struct {
	ctrl     *gomock.Controller
	recorder *MockProjectStorageMockRecorder
}

// MockProjectStorageMockRecorder is the mock recorder for MockProjectStorage.
type MockProjectStorageMockRecorder struct {
	mock *MockProjectStorage
}

// NewMockProjectStorage creates a new mock instance.
func NewMockProjectStorage(ctrl *gomock.Controller) *MockProjectStorage {
	mock := &MockProjectStorage{ctrl: ctrl}
	mock.recorder = &MockProjectStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectStorage) EXPECT() *MockProjectStorageMockRecorder {
	return m.recorder
}

// AddProjects mocks base method.
func (m *MockProjectStorage) AddProjects(arg0 context.Context, arg1 []models.Project) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProjects", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddProjects indicates an expected call of AddProjects.
func (mr *MockProjectStorageMockRecorder) AddProjects(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProjects", reflect.TypeOf((*MockProjectStorage)(nil).AddProjects), arg0, arg1)
}

// DescribeProject mocks base method.
func (m *MockProjectStorage) DescribeProject(arg0 context.Context, arg1 uint64) (*models.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeProject", arg0, arg1)
	ret0, _ := ret[0].(*models.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeProject indicates an expected call of DescribeProject.
func (mr *MockProjectStorageMockRecorder) DescribeProject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeProject", reflect.TypeOf((*MockProjectStorage)(nil).DescribeProject), arg0, arg1)
}

// ListProjects mocks base method.
func (m *MockProjectStorage) ListProjects(arg0 context.Context, arg1, arg2 uint64) ([]models.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProjects", arg0, arg1, arg2)
	ret0, _ := ret[0].([]models.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProjects indicates an expected call of ListProjects.
func (mr *MockProjectStorageMockRecorder) ListProjects(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjects", reflect.TypeOf((*MockProjectStorage)(nil).ListProjects), arg0, arg1, arg2)
}

// RemoveProject mocks base method.
func (m *MockProjectStorage) RemoveProject(arg0 context.Context, arg1 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveProject", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveProject indicates an expected call of RemoveProject.
func (mr *MockProjectStorageMockRecorder) RemoveProject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveProject", reflect.TypeOf((*MockProjectStorage)(nil).RemoveProject), arg0, arg1)
}
