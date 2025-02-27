package repository

import (
	"context"
	"errors"
	"github.com/Kahffi/go-rest-api-test/model/domain"
	"github.com/Kahffi/go-rest-api-test/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockProductRepository(ctrl)
	ctx := context.Background()

	completeProduct := domain.Product{
		ProductID:   0,
		Name:        "Product ggagal",
		Description: "duskripsi",
		Price:       100,
		StockQty:    2,
		CategoryId:  3,
		SKU:         "WKWKWK",
		TaxRate:     10,
		Category:    domain.Category{},
	}

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
				product := completeProduct
				repo.EXPECT().Save(ctx, product).Return(product, nil)
			},
			method: func() (interface{}, error) {
				return repo.Save(ctx, completeProduct)
			},
			expect:    completeProduct,
			expectErr: false,
		},
		{
			name: "Save Failure",
			mock: func() {
				repo.EXPECT().Save(ctx, gomock.Any()).Return(domain.Product{}, errors.New("error saving"))
			},
			method: func() (interface{}, error) {
				return repo.Save(ctx, domain.Product{Name: "Invalid"})
			},
			expect:    domain.Product{},
			expectErr: true,
		},
		{
			name: "Update Success",
			mock: func() {
				product := completeProduct
				repo.EXPECT().Update(ctx, product).Return(product, nil)
			},
			method: func() (interface{}, error) {
				return repo.Update(ctx, completeProduct)
			},
			expect:    completeProduct,
			expectErr: false,
		},
		{
			name: "FindById Success",
			mock: func() {
				repo.EXPECT().FindById(ctx, completeProduct.ProductID).Return(completeProduct, nil)
			},
			method: func() (interface{}, error) {
				return repo.FindById(ctx, completeProduct.ProductID)
			},
			expect:    completeProduct,
			expectErr: false,
		},
		{
			name: "FindById Not Found",
			mock: func() {
				repo.EXPECT().FindById(ctx, uint64(999)).Return(domain.Product{}, errors.New("not found"))
			},
			method: func() (interface{}, error) {
				return repo.FindById(ctx, 999)
			},
			expect:    domain.Product{},
			expectErr: true,
		},
		{
			name: "FindAll Success",
			mock: func() {
				repo.EXPECT().FindAll(ctx).Return([]domain.Product{completeProduct}, nil)
			},
			method: func() (interface{}, error) {
				return repo.FindAll(ctx)
			},
			expect:    []domain.Product{completeProduct},
			expectErr: false,
		},
		{
			name: "Delete Success",
			mock: func() {
				repo.EXPECT().Delete(ctx, domain.Product{ProductID: 1}).Return(nil)
			},
			method: func() (interface{}, error) {
				return nil, repo.Delete(ctx, domain.Product{ProductID: 1})
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
