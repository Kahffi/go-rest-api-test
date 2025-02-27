package repository

import (
	"context"
	"errors"
	"github.com/aronipurwanto/go-restful-api/model/domain"
	"github.com/aronipurwanto/go-restful-api/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmployeeRepository(t *testing.T) {
	CompleteEmployee := domain.Employee{
		EmployeeID: 1,
		Name:       "Kahpi",
		Role:       "COY",
		Email:      "kahffi@emial.com",
		Phone:      "9838783",
		DateHired:  "12/34/2005",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockEmployeeRepository(ctrl)
	ctx := context.Background()

	tests := []struct {
		name      string
		mock      func()
		method    func() (interface{}, error)
		expect    interface{}
		expectErr bool
	}{
		{
			name: "Save Success",
			mock: func() {
				employee := CompleteEmployee
				repo.EXPECT().Save(ctx, employee).Return(employee, nil)
			},
			method: func() (interface{}, error) {
				return repo.Save(ctx, CompleteEmployee)
			},
			expect:    CompleteEmployee,
			expectErr: false,
		},
		{
			name: "Save Failure",
			mock: func() {
				repo.EXPECT().Save(ctx, gomock.Any()).Return(domain.Employee{}, errors.New("error saving"))
			},
			method: func() (interface{}, error) {
				return repo.Save(ctx, domain.Employee{Name: "Invalid"})
			},
			expect:    domain.Employee{},
			expectErr: true,
		},
		{
			name: "Update Success",
			mock: func() {
				employee := CompleteEmployee
				repo.EXPECT().Update(ctx, employee).Return(employee, nil)
			},
			method: func() (interface{}, error) {
				return repo.Update(ctx, CompleteEmployee)
			},
			expect:    CompleteEmployee,
			expectErr: false,
		},
		{
			name: "FindById Success",
			mock: func() {
				repo.EXPECT().FindById(ctx, uint64(1)).Return(domain.Employee{EmployeeID: 1, Name: "Electronics"}, nil)
			},
			method: func() (interface{}, error) {
				return repo.FindById(ctx, 1)
			},
			expect:    domain.Employee{EmployeeID: 1, Name: "Electronics"},
			expectErr: false,
		},
		{
			name: "FindById Not Found",
			mock: func() {
				repo.EXPECT().FindById(ctx, uint64(999)).Return(domain.Employee{}, errors.New("not found"))
			},
			method: func() (interface{}, error) {
				return repo.FindById(ctx, 999)
			},
			expect:    domain.Employee{},
			expectErr: true,
		},
		{
			name: "FindAll Success",
			mock: func() {
				repo.EXPECT().FindAll(ctx).Return([]domain.Employee{CompleteEmployee}, nil)
			},
			method: func() (interface{}, error) {
				return repo.FindAll(ctx)
			},
			expect:    []domain.Employee{CompleteEmployee},
			expectErr: false,
		},
		{
			name: "Delete Success",
			mock: func() {
				repo.EXPECT().Delete(ctx, domain.Employee{EmployeeID: 1}).Return(nil)
			},
			method: func() (interface{}, error) {
				return nil, repo.Delete(ctx, domain.Employee{EmployeeID: 1})
			},
			expect:    nil,
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			result, err := tt.method()

			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expect, result)
			}
		})
	}
}
