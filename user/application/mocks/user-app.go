// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	domain "api-your-accounts/user/domain"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// IUserApp is an autogenerated mock type for the IUserApp type
type IUserApp struct {
	mock.Mock
}

// Exists provides a mock function with given fields: ctx, uuid, email
func (_m *IUserApp) Exists(ctx context.Context, uuid string, email string) (bool, error) {
	ret := _m.Called(ctx, uuid, email)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (bool, error)); ok {
		return rf(ctx, uuid, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) bool); ok {
		r0 = rf(ctx, uuid, email)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, uuid, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: ctx, uuid, email
func (_m *IUserApp) Login(ctx context.Context, uuid string, email string) (string, error) {
	ret := _m.Called(ctx, uuid, email)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (string, error)); ok {
		return rf(ctx, uuid, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) string); ok {
		r0 = rf(ctx, uuid, email)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, uuid, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignUp provides a mock function with given fields: ctx, user
func (_m *IUserApp) SignUp(ctx context.Context, user *domain.User) (*domain.User, error) {
	ret := _m.Called(ctx, user)

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.User) (*domain.User, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *domain.User) *domain.User); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *domain.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIUserApp interface {
	mock.TestingT
	Cleanup(func())
}

// NewIUserApp creates a new instance of IUserApp. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIUserApp(t mockConstructorTestingTNewIUserApp) *IUserApp {
	mock := &IUserApp{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
