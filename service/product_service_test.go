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

var productResponseTpl = web.ProductResponse{
	Id:          1,
	Name:        "Barang mewwah",
	Description: "mewah bingit",
	Price:       10000,
	StockQty:    100,
	CategoryID:  32,
	SKU:         "MWH",
	TaxRate:     10,
}

var productModelTpl = domain.Product{
	ProductID:   1,
	Name:        "Barang mewwah",
	Description: "mewah bingit",
	Price:       10000,
	StockQty:    100,
	CategoryId:  32,
	SKU:         "MWH",
	TaxRate:     10,
	Category:    domain.Category{},
}

func TestCreateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepository(ctrl)
	mockValidator := validator.New()
	productService := NewProductService(mockRepo, mockValidator)

	productCreateReq := web.ProductCreateRequest{
		Name:        "Barang mewwah",
		Description: "mewah bingit",
		Price:       10000,
		StockQty:    100,
		CategoryID:  32,
		SKU:         "MWH",
		TaxRate:     10,
	}

	tests := []struct {
		name      string
		input     web.ProductCreateRequest
		mock      func()
		expect    web.ProductResponse
		expectErr bool
	}{
		{
			name:  "success",
			input: productCreateReq,
			mock: func() {
				mockRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(productModelTpl, nil)
			},
			expect:    productResponseTpl,
			expectErr: false,
		},
		{
			name:      "validation error",
			input:     web.ProductCreateRequest{Name: ""},
			mock:      func() {},
			expect:    web.ProductResponse{},
			expectErr: true,
		},
		{
			name:  "repository error",
			input: productCreateReq,
			mock: func() {
				mockRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(domain.Product{}, errors.New("database error"))
			},
			expect:    web.ProductResponse{},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			resp, err := productService.Create(context.Background(), tt.input)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expect, resp)
			}
		})
	}
}

func TestDeleteProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepository(ctrl)
	productService := NewProductService(mockRepo, validator.New())

	tests := []struct {
		name      string
		productId uint64
		mock      func()
		expectErr bool
	}{
		{
			name:      "success",
			productId: 1,
			mock: func() {
				mockRepo.EXPECT().FindById(gomock.Any(), productModelTpl.ProductID).Return(productModelTpl, nil)
				mockRepo.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectErr: false,
		},
		{
			name:      "not found",
			productId: 99,
			mock: func() {
				mockRepo.EXPECT().FindById(gomock.Any(), uint64(99)).Return(domain.Product{}, errors.New("not found"))
			},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := productService.Delete(context.Background(), tt.productId)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateProduct(t *testing.T) {
	productUpdateReqTpl := web.ProductUpdateRequest{
		Id:          1,
		Name:        "Barang mewwah",
		Description: "mewah bingit",
		Price:       10000,
		StockQty:    100,
		CategoryID:  32,
		SKU:         "MWH",
		TaxRate:     10,
	}

	tests := []struct {
		name    string
		mock    func(mockProductRepo *mocks.MockProductRepository)
		input   web.ProductUpdateRequest
		expects error
	}{
		{
			name: "Success",
			mock: func(mockProductRepo *mocks.MockProductRepository) {
				mockProductRepo.EXPECT().FindById(gomock.Any(), uint64(1)).Return(productModelTpl, nil)
				mockProductRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(productModelTpl, nil)
			},
			input:   productUpdateReqTpl,
			expects: nil,
		},
		{
			name: "Product Not Found",
			mock: func(mockProductRepo *mocks.MockProductRepository) {
				mockProductRepo.EXPECT().FindById(gomock.Any(), productModelTpl.ProductID).Return(domain.Product{}, errors.New("product not found"))
			},
			input:   productUpdateReqTpl,
			expects: errors.New("product not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockProductRepo := mocks.NewMockProductRepository(ctrl)
			tt.mock(mockProductRepo)

			service := NewProductService(mockProductRepo, validator.New())
			_, err := service.Update(context.Background(), tt.input)
			assert.Equal(t, tt.expects, err)
		})
	}
}

func TestFindAllProducts(t *testing.T) {
	tests := []struct {
		name    string
		mock    func(mockProductRepo *mocks.MockProductRepository)
		expects []web.ProductResponse
		err     error
	}{
		{
			name: "Success",
			mock: func(mockProductRepo *mocks.MockProductRepository) {
				mockProductRepo.EXPECT().FindAll(gomock.Any()).Return([]domain.Product{productModelTpl}, nil)
			},
			expects: []web.ProductResponse{productResponseTpl},
			err:     nil,
		},
		{
			name: "Database Error",
			mock: func(mockProductRepo *mocks.MockProductRepository) {
				mockProductRepo.EXPECT().FindAll(gomock.Any()).Return(nil, errors.New("database error"))
			},
			expects: nil,
			err:     errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockProductRepo := mocks.NewMockProductRepository(ctrl)
			tt.mock(mockProductRepo)

			service := NewProductService(mockProductRepo, validator.New())
			result, err := service.FindAll(context.Background())
			assert.Equal(t, tt.expects, result)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestFindByIdProduct(t *testing.T) {
	tests := []struct {
		name    string
		mock    func(mockProductRepo *mocks.MockProductRepository)
		input   uint64
		expects web.ProductResponse
		err     error
	}{
		{
			name: "Success",
			mock: func(mockProductRepo *mocks.MockProductRepository) {
				mockProductRepo.EXPECT().FindById(gomock.Any(), uint64(1)).Return(productModelTpl, nil)
			},
			input:   productModelTpl.ProductID,
			expects: productResponseTpl,
			err:     nil,
		},
		{
			name: "Not Found",
			mock: func(mockProductRepo *mocks.MockProductRepository) {
				mockProductRepo.EXPECT().FindById(gomock.Any(), uint64(1)).Return(domain.Product{}, errors.New("not found"))
			},
			input:   1,
			expects: web.ProductResponse{},
			err:     errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockProductRepo := mocks.NewMockProductRepository(ctrl)
			tt.mock(mockProductRepo)

			service := NewProductService(mockProductRepo, validator.New())
			result, err := service.FindById(context.Background(), tt.input)
			assert.Equal(t, tt.expects, result)
			assert.Equal(t, tt.err, err)
		})
	}
}
