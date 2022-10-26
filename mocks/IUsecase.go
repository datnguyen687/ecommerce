// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"
	models "ecommerce/internal/models"

	mock "github.com/stretchr/testify/mock"
)

// IUsecase is an autogenerated mock type for the IUsecase type
type IUsecase struct {
	mock.Mock
}

// FilterProducts provides a mock function with given fields: ctx, filter, orderBy
func (_m *IUsecase) FilterProducts(ctx context.Context, filter models.ProductFilter, orderBy string) ([]models.Product, error) {
	ret := _m.Called(ctx, filter, orderBy)

	var r0 []models.Product
	if rf, ok := ret.Get(0).(func(context.Context, models.ProductFilter, string) []models.Product); ok {
		r0 = rf(ctx, filter, orderBy)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.ProductFilter, string) error); ok {
		r1 = rf(ctx, filter, orderBy)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewIUsecase creates a new instance of IUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIUsecase(t mockConstructorTestingTNewIUsecase) *IUsecase {
	mock := &IUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
