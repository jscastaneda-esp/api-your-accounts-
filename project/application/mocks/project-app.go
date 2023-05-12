// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	domain "api-your-accounts/project/domain"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// IProjectApp is an autogenerated mock type for the IProjectApp type
type IProjectApp struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, project
func (_m *IProjectApp) Create(ctx context.Context, project *domain.Project) (*domain.Project, error) {
	ret := _m.Called(ctx, project)

	var r0 *domain.Project
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Project) (*domain.Project, error)); ok {
		return rf(ctx, project)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Project) *domain.Project); ok {
		r0 = rf(ctx, project)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Project)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *domain.Project) error); ok {
		r1 = rf(ctx, project)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *IProjectApp) Delete(ctx context.Context, id uint) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReadByUser provides a mock function with given fields: ctx, userId
func (_m *IProjectApp) ReadByUser(ctx context.Context, userId uint) ([]*domain.Project, error) {
	ret := _m.Called(ctx, userId)

	var r0 []*domain.Project
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) ([]*domain.Project, error)); ok {
		return rf(ctx, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint) []*domain.Project); ok {
		r0 = rf(ctx, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Project)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIProjectApp interface {
	mock.TestingT
	Cleanup(func())
}

// NewIProjectApp creates a new instance of IProjectApp. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIProjectApp(t mockConstructorTestingTNewIProjectApp) *IProjectApp {
	mock := &IProjectApp{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
