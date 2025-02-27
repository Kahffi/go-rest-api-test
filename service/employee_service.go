package service

import (
	"context"
	"github.com/Kahffi/go-rest-api-test/model/web"
)

type EmployeeService interface {
	Create(ctx context.Context, request web.EmployeeCreateRequest) (web.EmployeeResponse, error)
	Update(ctx context.Context, request web.EmployeeUpdateRequest) (web.EmployeeResponse, error)
	Delete(ctx context.Context, employeeId uint64) error
	FindById(ctx context.Context, employeeId uint64) (web.EmployeeResponse, error)
	FindAll(ctx context.Context) ([]web.EmployeeResponse, error)
}
