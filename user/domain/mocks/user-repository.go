// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	domain "api-your-accounts/user/domain"
	context "context"

	mock "github.com/stretchr/testify/mock"

	transaction "api-your-accounts/shared/domain/transaction"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, user
func (_m *UserRepository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
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

// ExistsByEmail provides a mock function with given fields: ctx, email
func (_m *UserRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	ret := _m.Called(ctx, email)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (bool, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExistsByUUID provides a mock function with given fields: ctx, uuid
func (_m *UserRepository) ExistsByUUID(ctx context.Context, uuid string) (bool, error) {
	ret := _m.Called(ctx, uuid)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (bool, error)); ok {
		return rf(ctx, uuid)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, uuid)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, uuid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByUUIDAndEmail provides a mock function with given fields: ctx, uuid, email
func (_m *UserRepository) FindByUUIDAndEmail(ctx context.Context, uuid string, email string) (*domain.User, error) {
	ret := _m.Called(ctx, uuid, email)

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*domain.User, error)); ok {
		return rf(ctx, uuid, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *domain.User); ok {
		r0 = rf(ctx, uuid, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, uuid, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WithTransaction provides a mock function with given fields: tx
func (_m *UserRepository) WithTransaction(tx transaction.Transaction) domain.UserRepository {
	ret := _m.Called(tx)

	var r0 domain.UserRepository
	if rf, ok := ret.Get(0).(func(transaction.Transaction) domain.UserRepository); ok {
		r0 = rf(tx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.UserRepository)
		}
	}

	return r0
}

type mockConstructorTestingTNewUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepository(t mockConstructorTestingTNewUserRepository) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
