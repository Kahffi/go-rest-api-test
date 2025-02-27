package service

import (
	"context"
	"errors"
	"github.com/Kahffi/go-rest-api-test/exception"
	"github.com/Kahffi/go-rest-api-test/helper"
	"github.com/Kahffi/go-rest-api-test/model/domain"
	"github.com/Kahffi/go-rest-api-test/model/web"
	"github.com/Kahffi/go-rest-api-test/repository"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	Validate          *validator.Validate
}

func NewProductService(productRepository repository.ProductRepository, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		Validate:          validate,
	}
}

// Create Product
func (service *ProductServiceImpl) Create(ctx context.Context, request web.ProductCreateRequest) (web.ProductResponse, error) {
	if err := service.Validate.Struct(request); err != nil {
		return web.ProductResponse{}, err
	}

	product := domain.Product{Name: request.Name}
	savedProduct, err := service.ProductRepository.Save(ctx, product)
	if err != nil {
		return web.ProductResponse{}, err
	}

	return helper.ToProductResponse(savedProduct), nil
}

// Update Product
func (service *ProductServiceImpl) Update(ctx context.Context, request web.ProductUpdateRequest) (web.ProductResponse, error) {
	if err := service.Validate.Struct(request); err != nil {
		return web.ProductResponse{}, err
	}

	product, err := service.ProductRepository.FindById(ctx, request.Id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return web.ProductResponse{}, exception.NewNotFoundError("Product not found")
	} else if err != nil {
		return web.ProductResponse{}, err
	}

	product.Name = request.Name
	updatedProduct, err := service.ProductRepository.Update(ctx, product)
	if err != nil {
		return web.ProductResponse{}, err
	}

	return helper.ToProductResponse(updatedProduct), nil
}

// Delete Product
func (service *ProductServiceImpl) Delete(ctx context.Context, productId uint64) error {
	product, err := service.ProductRepository.FindById(ctx, productId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return exception.NewNotFoundError("Product not found")
	} else if err != nil {
		return err
	}

	return service.ProductRepository.Delete(ctx, product)
}

// Find Product By ID
func (service *ProductServiceImpl) FindById(ctx context.Context, productId uint64) (web.ProductResponse, error) {
	product, err := service.ProductRepository.FindById(ctx, productId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return web.ProductResponse{}, exception.NewNotFoundError("Product not found")
	} else if err != nil {
		return web.ProductResponse{}, err
	}

	return helper.ToProductResponse(product), nil
}

// Find All Categories
func (service *ProductServiceImpl) FindAll(ctx context.Context) ([]web.ProductResponse, error) {
	categories, err := service.ProductRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return helper.ToProductResponses(categories), nil
}
