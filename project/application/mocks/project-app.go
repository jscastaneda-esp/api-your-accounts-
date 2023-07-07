// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"
	application "your-accounts-api/project/application"

	domain "your-accounts-api/project/domain"

	mock "github.com/stretchr/testify/mock"
)

// IProjectApp is an autogenerated mock type for the IProjectApp type
type IProjectApp struct {
	mock.Mock
}

// Clone provides a mock function with given fields: ctx, baseId
func (_m *IProjectApp) Clone(ctx context.Context, baseId uint) (uint, error) {
	ret := _m.Called(ctx, baseId)

	var r0 uint
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) (uint, error)); ok {
		return rf(ctx, baseId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint) uint); ok {
		r0 = rf(ctx, baseId)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, baseId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, createData
func (_m *IProjectApp) Create(ctx context.Context, createData application.CreateData) (uint, error) {
	ret := _m.Called(ctx, createData)

	var r0 uint
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, application.CreateData) (uint, error)); ok {
		return rf(ctx, createData)
	}
	if rf, ok := ret.Get(0).(func(context.Context, application.CreateData) uint); ok {
		r0 = rf(ctx, createData)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(context.Context, application.CreateData) error); ok {
		r1 = rf(ctx, createData)
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

// FindByUser provides a mock function with given fields: ctx, userId
func (_m *IProjectApp) FindByUser(ctx context.Context, userId uint) ([]*application.FindByUserRecord, error) {
	ret := _m.Called(ctx, userId)

	var r0 []*application.FindByUserRecord
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) ([]*application.FindByUserRecord, error)); ok {
		return rf(ctx, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint) []*application.FindByUserRecord); ok {
		r0 = rf(ctx, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*application.FindByUserRecord)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindLogsByProject provides a mock function with given fields: ctx, projectId
func (_m *IProjectApp) FindLogsByProject(ctx context.Context, projectId uint) ([]*domain.ProjectLog, error) {
	ret := _m.Called(ctx, projectId)

	var r0 []*domain.ProjectLog
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) ([]*domain.ProjectLog, error)); ok {
		return rf(ctx, projectId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint) []*domain.ProjectLog); ok {
		r0 = rf(ctx, projectId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.ProjectLog)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, projectId)
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
