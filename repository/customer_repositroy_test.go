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

func TestCustomerRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockCustomerRepository(ctrl)
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
				customer := domain.Customer{
					CustomerID: 0,
					Name:       "Electruc",
					Email:      "electruc@gmail.com",
					Phone:      "093762876",
					Address:    "hybrid",
					LoyaltyPts: 0,
				}
				repo.EXPECT().Save(ctx, customer).Return(customer, nil)
			},
			method: func() (interface{}, error) {
				return repo.Save(ctx, domain.Customer{CustomerID: 0,
					Name:       "Electruc",
					Email:      "electruc@gmail.com",
					Phone:      "093762876",
					Address:    "hybrid",
					LoyaltyPts: 0})
			},
			expect: domain.Customer{CustomerID: 0,
				Name:       "Electruc",
				Email:      "electruc@gmail.com",
				Phone:      "093762876",
				Address:    "hybrid",
				LoyaltyPts: 0,
			},
			expectErr: false,
		},
		{
			name: "Save Failure",
			mock: func() {
				repo.EXPECT().Save(ctx, gomock.Any()).Return(domain.Customer{}, errors.New("error saving"))
			},
			method: func() (interface{}, error) {
				return repo.Save(ctx, domain.Customer{
					Name: "Invalid",
				})
			},
			expect:    domain.Customer{},
			expectErr: true,
		},
		{
			name: "Update Success",
			mock: func() {
				customer := domain.Customer{
					CustomerID: 1,
					Name:       "Updated Name",
					Email:      "ppp@gmail.com",
					Phone:      "098373467836",
					Address:    "mac",
					LoyaltyPts: 0,
				}
				repo.EXPECT().Update(ctx, customer).Return(customer, nil)
			},
			method: func() (interface{}, error) {
				return repo.Update(ctx, domain.Customer{
					CustomerID: 1,
					Name:       "Updated Name",
					Email:      "ppp@gmail.com",
					Phone:      "098373467836",
					Address:    "mac",
					LoyaltyPts: 0,
				})
			},
			expect: domain.Customer{CustomerID: 1,
				Name:       "Updated Name",
				Email:      "ppp@gmail.com",
				Phone:      "098373467836",
				Address:    "mac",
				LoyaltyPts: 0,
			},
			expectErr: false,
		},
		{
			name: "FindById Success",
			mock: func() {
				repo.EXPECT().FindById(ctx, uint64(0)).Return(domain.Customer{
					CustomerID: 0,
					Name:       "Electruc",
					Email:      "electruc@gmail.com",
					Phone:      "093762876",
					Address:    "hybrid",
					LoyaltyPts: 0,
				}, nil)
			},
			method: func() (interface{}, error) {
				return repo.FindById(ctx, 0)
			},
			expect: domain.Customer{
				CustomerID: 0,
				Name:       "Electruc",
				Email:      "electruc@gmail.com",
				Phone:      "093762876",
				Address:    "hybrid",
				LoyaltyPts: 0,
			},
			expectErr: false,
		},
		{
			name: "FindById Not Found",
			mock: func() {
				repo.EXPECT().FindById(ctx, uint64(999)).Return(domain.Customer{}, errors.New("not found"))
			},
			method: func() (interface{}, error) {
				return repo.FindById(ctx, 999)
			},
			expect:    domain.Customer{},
			expectErr: true,
		},
		{
			name: "FindAll Success",
			mock: func() {
				repo.EXPECT().FindAll(ctx).Return([]domain.Customer{{
					CustomerID: 0,
					Name:       "Electruc",
					Email:      "electruc@gmail.com",
					Phone:      "093762876",
					Address:    "hybrid",
					LoyaltyPts: 0,
				}}, nil)
			},
			method: func() (interface{}, error) {
				return repo.FindAll(ctx)
			},
			expect: []domain.Customer{{
				CustomerID: 0,
				Name:       "Electruc",
				Email:      "electruc@gmail.com",
				Phone:      "093762876",
				Address:    "hybrid",
				LoyaltyPts: 0,
			}},
			expectErr: false,
		},
		{
			name: "Delete Success",
			mock: func() {
				repo.EXPECT().Delete(ctx, domain.Customer{
					CustomerID: 0,
					Name:       "Electruc",
					Email:      "electruc@gmail.com",
					Phone:      "093762876",
					Address:    "hybrid",
					LoyaltyPts: 0,
				}).Return(nil)
			},
			method: func() (interface{}, error) {
				return nil, repo.Delete(ctx, domain.Customer{
					CustomerID: 0,
					Name:       "Electruc",
					Email:      "electruc@gmail.com",
					Phone:      "093762876",
					Address:    "hybrid",
					LoyaltyPts: 0,
				})
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
