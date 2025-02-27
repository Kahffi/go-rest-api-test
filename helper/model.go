package helper

import (
	"github.com/Kahffi/go-rest-api-test/model/domain"
	"github.com/Kahffi/go-rest-api-test/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}

func ToCustomerResponse(customer domain.Customer) web.CustomerResponse {
	return web.CustomerResponse{
		Id:         customer.CustomerID,
		Name:       customer.Name,
		Email:      customer.Email,
		Phone:      customer.Phone,
		Address:    customer.Address,
		LoyaltyPts: customer.LoyaltyPts,
	}
}

func ToCustomerResponses(customers []domain.Customer) []web.CustomerResponse {
	var customerResponses []web.CustomerResponse
	for _, customer := range customers {
		customerResponses = append(customerResponses, ToCustomerResponse(customer))
	}
	return customerResponses
}

func ToEmployeeResponse(employee domain.Employee) web.EmployeeResponse {
	return web.EmployeeResponse{
		Id:        employee.EmployeeID,
		Name:      employee.Name,
		Role:      employee.Role,
		Email:     employee.Email,
		Phone:     employee.Phone,
		DateHired: employee.DateHired,
	}
}

func ToEmployeeResponses(emails []domain.Employee) []web.EmployeeResponse {
	var employeeResponses []web.EmployeeResponse
	for _, email := range emails {
		employeeResponses = append(employeeResponses, ToEmployeeResponse(email))
	}
	return employeeResponses
}

func ToProductResponse(product domain.Product) web.ProductResponse {
	return web.ProductResponse{
		Id:          product.ProductID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		StockQty:    product.StockQty,
		CategoryID:  int(product.CategoryId),
		SKU:         product.SKU,
		TaxRate:     product.TaxRate,
	}
}

func ToProductResponses(products []domain.Product) []web.ProductResponse {
	var productResponses []web.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, ToProductResponse(product))
	}
	return productResponses
}
