// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/task/task.go

// Package mock_task is a generated GoMock package.
package mock_task

import (
	context "context"
	reflect "reflect"
	domain "sampleProject/internal/domain"
	params "sampleProject/internal/params"

	gomock "github.com/golang/mock/gomock"
)

// MockTaskRepository is a mock of TaskRepository interface.
type MockTaskRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTaskRepositoryMockRecorder
}

// MockTaskRepositoryMockRecorder is the mock recorder for MockTaskRepository.
type MockTaskRepositoryMockRecorder struct {
	mock *MockTaskRepository
}

// NewMockTaskRepository creates a new mock instance.
func NewMockTaskRepository(ctrl *gomock.Controller) *MockTaskRepository {
	mock := &MockTaskRepository{ctrl: ctrl}
	mock.recorder = &MockTaskRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskRepository) EXPECT() *MockTaskRepositoryMockRecorder {
	return m.recorder
}

// Begin mocks base method.
func (m *MockTaskRepository) Begin(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Begin", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Begin indicates an expected call of Begin.
func (mr *MockTaskRepositoryMockRecorder) Begin(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*MockTaskRepository)(nil).Begin), ctx)
}

// Commit mocks base method.
func (m *MockTaskRepository) Commit() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockTaskRepositoryMockRecorder) Commit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockTaskRepository)(nil).Commit))
}

// Create mocks base method.
func (m *MockTaskRepository) Create(ctx context.Context, input params.CreateParams) (*domain.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, input)
	ret0, _ := ret[0].(*domain.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockTaskRepositoryMockRecorder) Create(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTaskRepository)(nil).Create), ctx, input)
}

// CreateLogTable mocks base method.
func (m *MockTaskRepository) CreateLogTable(ctx context.Context, tableName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLogTable", ctx, tableName)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateLogTable indicates an expected call of CreateLogTable.
func (mr *MockTaskRepositoryMockRecorder) CreateLogTable(ctx, tableName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLogTable", reflect.TypeOf((*MockTaskRepository)(nil).CreateLogTable), ctx, tableName)
}

// Get mocks base method.
func (m *MockTaskRepository) Get(ctx context.Context, id int) (*domain.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(*domain.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockTaskRepositoryMockRecorder) Get(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTaskRepository)(nil).Get), ctx, id)
}

// GetList mocks base method.
func (m *MockTaskRepository) GetList(ctx context.Context) ([]*domain.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetList", ctx)
	ret0, _ := ret[0].([]*domain.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetList indicates an expected call of GetList.
func (mr *MockTaskRepositoryMockRecorder) GetList(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetList", reflect.TypeOf((*MockTaskRepository)(nil).GetList), ctx)
}

// Rollback mocks base method.
func (m *MockTaskRepository) Rollback() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback")
	ret0, _ := ret[0].(error)
	return ret0
}

// Rollback indicates an expected call of Rollback.
func (mr *MockTaskRepositoryMockRecorder) Rollback() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockTaskRepository)(nil).Rollback))
}

// Update mocks base method.
func (m *MockTaskRepository) Update(ctx context.Context, dao *domain.Task) (*domain.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, dao)
	ret0, _ := ret[0].(*domain.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockTaskRepositoryMockRecorder) Update(ctx, dao interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTaskRepository)(nil).Update), ctx, dao)
}

// MockTaskCache is a mock of TaskCache interface.
type MockTaskCache struct {
	ctrl     *gomock.Controller
	recorder *MockTaskCacheMockRecorder
}

// MockTaskCacheMockRecorder is the mock recorder for MockTaskCache.
type MockTaskCacheMockRecorder struct {
	mock *MockTaskCache
}

// NewMockTaskCache creates a new mock instance.
func NewMockTaskCache(ctrl *gomock.Controller) *MockTaskCache {
	mock := &MockTaskCache{ctrl: ctrl}
	mock.recorder = &MockTaskCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskCache) EXPECT() *MockTaskCacheMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockTaskCache) Get(ctx context.Context, id int) (*domain.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(*domain.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockTaskCacheMockRecorder) Get(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTaskCache)(nil).Get), ctx, id)
}

// Set mocks base method.
func (m *MockTaskCache) Set(ctx context.Context, data *domain.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockTaskCacheMockRecorder) Set(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockTaskCache)(nil).Set), ctx, data)
}

// MockTaskQueue is a mock of TaskQueue interface.
type MockTaskQueue struct {
	ctrl     *gomock.Controller
	recorder *MockTaskQueueMockRecorder
}

// MockTaskQueueMockRecorder is the mock recorder for MockTaskQueue.
type MockTaskQueueMockRecorder struct {
	mock *MockTaskQueue
}

// NewMockTaskQueue creates a new mock instance.
func NewMockTaskQueue(ctrl *gomock.Controller) *MockTaskQueue {
	mock := &MockTaskQueue{ctrl: ctrl}
	mock.recorder = &MockTaskQueueMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskQueue) EXPECT() *MockTaskQueueMockRecorder {
	return m.recorder
}

// SendMessage mocks base method.
func (m *MockTaskQueue) SendMessage(msg string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMessage", msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMessage indicates an expected call of SendMessage.
func (mr *MockTaskQueueMockRecorder) SendMessage(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockTaskQueue)(nil).SendMessage), msg)
}
