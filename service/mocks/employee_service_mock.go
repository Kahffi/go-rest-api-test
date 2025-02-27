// Code generated by MockGen. DO NOT EDIT.
// Source: service/employee_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	web "github.com/Kahffi/go-rest-api-test/model/web"
	gomock "github.com/golang/mock/gomock"
)

// MockEmployeeService is a mock of EmployeeService interface.
type MockEmployeeService struct {
	ctrl     *gomock.Controller
	recorder *MockEmployeeServiceMockRecorder
}

// MockEmployeeServiceMockRecorder is the mock recorder for MockEmployeeService.
type MockEmployeeServiceMockRecorder struct {
	mock *MockEmployeeService
}

// NewMockEmployeeService creates a new mock instance.
func NewMockEmployeeService(ctrl *gomock.Controller) *MockEmployeeService {
	mock := &MockEmployeeService{ctrl: ctrl}
	mock.recorder = &MockEmployeeServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmployeeService) EXPECT() *MockEmployeeServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockEmployeeService) Create(ctx context.Context, request web.EmployeeCreateRequest) (web.EmployeeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, request)
	ret0, _ := ret[0].(web.EmployeeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockEmployeeServiceMockRecorder) Create(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockEmployeeService)(nil).Create), ctx, request)
}

// Delete mocks base method.
func (m *MockEmployeeService) Delete(ctx context.Context, employeeId uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, employeeId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockEmployeeServiceMockRecorder) Delete(ctx, employeeId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockEmployeeService)(nil).Delete), ctx, employeeId)
}

// FindAll mocks base method.
func (m *MockEmployeeService) FindAll(ctx context.Context) ([]web.EmployeeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx)
	ret0, _ := ret[0].([]web.EmployeeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockEmployeeServiceMockRecorder) FindAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockEmployeeService)(nil).FindAll), ctx)
}

// FindById mocks base method.
func (m *MockEmployeeService) FindById(ctx context.Context, employeeId uint64) (web.EmployeeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", ctx, employeeId)
	ret0, _ := ret[0].(web.EmployeeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockEmployeeServiceMockRecorder) FindById(ctx, employeeId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockEmployeeService)(nil).FindById), ctx, employeeId)
}

// Update mocks base method.
func (m *MockEmployeeService) Update(ctx context.Context, request web.EmployeeUpdateRequest) (web.EmployeeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, request)
	ret0, _ := ret[0].(web.EmployeeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockEmployeeServiceMockRecorder) Update(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockEmployeeService)(nil).Update), ctx, request)
}
