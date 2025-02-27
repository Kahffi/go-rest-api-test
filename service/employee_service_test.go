package service

import (
	"context"
	"errors"
	"github.com/Kahffi/go-rest-api-test/model/domain"
	"github.com/Kahffi/go-rest-api-test/model/web"
	"github.com/Kahffi/go-rest-api-test/repository/mocks"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var employeeResponseTpl = web.EmployeeResponse{
	Id:        1,
	Name:      "Harun maskiu",
	Role:      "Admin",
	Email:     "gone@away.com",
	Phone:     "72346782364",
	DateHired: "24/10/2010",
}

var employeeModelTpl = domain.Employee{
	EmployeeID: 1,
	Name:       "Harun maskiu",
	Role:       "Admin",
	Email:      "gone@away.com",
	Phone:      "72346782364",
	DateHired:  "24/10/2010",
}

func TestCreateEmployee(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockEmployeeRepository(ctrl)
	mockValidator := validator.New()
	employeeService := NewEmployeeService(mockRepo, mockValidator)

	employeeCreateReq := web.EmployeeCreateRequest{
		Name:      "Harun maskiu",
		Role:      "Admin",
		Email:     "gone@away.com",
		Phone:     "72346782364",
		DateHired: "24/10/2010",
	}

	tests := []struct {
		name      string
		input     web.EmployeeCreateRequest
		mock      func()
		expect    web.EmployeeResponse
		expectErr bool
	}{
		{
			name:  "success",
			input: employeeCreateReq,
			mock: func() {
				mockRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(employeeModelTpl, nil)
			},
			expect:    employeeResponseTpl,
			expectErr: false,
		},
		{
			name:      "validation error",
			input:     web.EmployeeCreateRequest{Name: ""},
			mock:      func() {},
			expect:    web.EmployeeResponse{},
			expectErr: true,
		},
		{
			name:  "repository error",
			input: employeeCreateReq,
			mock: func() {
				mockRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(domain.Employee{}, errors.New("database error"))
			},
			expect:    web.EmployeeResponse{},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			resp, err := employeeService.Create(context.Background(), tt.input)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expect, resp)
			}
		})
	}
}

func TestDeleteEmployee(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockEmployeeRepository(ctrl)
	employeeService := NewEmployeeService(mockRepo, validator.New())

	tests := []struct {
		name       string
		employeeId uint64
		mock       func()
		expectErr  bool
	}{
		{
			name:       "success",
			employeeId: 1,
			mock: func() {
				mockRepo.EXPECT().FindById(gomock.Any(), employeeModelTpl.EmployeeID).Return(employeeModelTpl, nil)
				mockRepo.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectErr: false,
		},
		{
			name:       "not found",
			employeeId: 99,
			mock: func() {
				mockRepo.EXPECT().FindById(gomock.Any(), uint64(99)).Return(domain.Employee{}, errors.New("not found"))
			},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := employeeService.Delete(context.Background(), tt.employeeId)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateEmployee(t *testing.T) {
	employeeUpdateReqTpl := web.EmployeeUpdateRequest{
		Id:        1,
		Name:      "Harun maskiu",
		Role:      "Admin",
		Email:     "gone@away.com",
		Phone:     "72346782364",
		DateHired: "24/10/2010",
	}

	tests := []struct {
		name    string
		mock    func(mockEmployeeRepo *mocks.MockEmployeeRepository)
		input   web.EmployeeUpdateRequest
		expects error
	}{
		{
			name: "Success",
			mock: func(mockEmployeeRepo *mocks.MockEmployeeRepository) {
				mockEmployeeRepo.EXPECT().FindById(gomock.Any(), uint64(1)).Return(employeeModelTpl, nil)
				mockEmployeeRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(employeeModelTpl, nil)
			},
			input:   employeeUpdateReqTpl,
			expects: nil,
		},
		{
			name: "Employee Not Found",
			mock: func(mockEmployeeRepo *mocks.MockEmployeeRepository) {
				mockEmployeeRepo.EXPECT().FindById(gomock.Any(), employeeModelTpl.EmployeeID).Return(domain.Employee{}, errors.New("employee not found"))
			},
			input:   employeeUpdateReqTpl,
			expects: errors.New("employee not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockEmployeeRepo := mocks.NewMockEmployeeRepository(ctrl)
			tt.mock(mockEmployeeRepo)

			service := NewEmployeeService(mockEmployeeRepo, validator.New())
			_, err := service.Update(context.Background(), tt.input)
			assert.Equal(t, tt.expects, err)
		})
	}
}

func TestFindAllEmployees(t *testing.T) {
	tests := []struct {
		name    string
		mock    func(mockEmployeeRepo *mocks.MockEmployeeRepository)
		expects []web.EmployeeResponse
		err     error
	}{
		{
			name: "Success",
			mock: func(mockEmployeeRepo *mocks.MockEmployeeRepository) {
				mockEmployeeRepo.EXPECT().FindAll(gomock.Any()).Return([]domain.Employee{employeeModelTpl}, nil)
			},
			expects: []web.EmployeeResponse{employeeResponseTpl},
			err:     nil,
		},
		{
			name: "Database Error",
			mock: func(mockEmployeeRepo *mocks.MockEmployeeRepository) {
				mockEmployeeRepo.EXPECT().FindAll(gomock.Any()).Return(nil, errors.New("database error"))
			},
			expects: nil,
			err:     errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockEmployeeRepo := mocks.NewMockEmployeeRepository(ctrl)
			tt.mock(mockEmployeeRepo)

			service := NewEmployeeService(mockEmployeeRepo, validator.New())
			result, err := service.FindAll(context.Background())
			assert.Equal(t, tt.expects, result)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestFindByIdEmployee(t *testing.T) {
	tests := []struct {
		name    string
		mock    func(mockEmployeeRepo *mocks.MockEmployeeRepository)
		input   uint64
		expects web.EmployeeResponse
		err     error
	}{
		{
			name: "Success",
			mock: func(mockEmployeeRepo *mocks.MockEmployeeRepository) {
				mockEmployeeRepo.EXPECT().FindById(gomock.Any(), uint64(1)).Return(employeeModelTpl, nil)
			},
			input:   employeeModelTpl.EmployeeID,
			expects: employeeResponseTpl,
			err:     nil,
		},
		{
			name: "Not Found",
			mock: func(mockEmployeeRepo *mocks.MockEmployeeRepository) {
				mockEmployeeRepo.EXPECT().FindById(gomock.Any(), uint64(1)).Return(domain.Employee{}, errors.New("not found"))
			},
			input:   1,
			expects: web.EmployeeResponse{},
			err:     errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockEmployeeRepo := mocks.NewMockEmployeeRepository(ctrl)
			tt.mock(mockEmployeeRepo)

			service := NewEmployeeService(mockEmployeeRepo, validator.New())
			result, err := service.FindById(context.Background(), tt.input)
			assert.Equal(t, tt.expects, result)
			assert.Equal(t, tt.err, err)
		})
	}
}
