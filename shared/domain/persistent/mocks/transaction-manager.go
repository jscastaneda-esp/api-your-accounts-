// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	persistent "your-accounts-api/shared/domain/persistent"

	mock "github.com/stretchr/testify/mock"
)

// TransactionManager is an autogenerated mock type for the TransactionManager type
type TransactionManager struct {
	mock.Mock
}

// Transaction provides a mock function with given fields: fc
func (_m *TransactionManager) Transaction(fc func(persistent.Transaction) error) error {
	ret := _m.Called(fc)

	if len(ret) == 0 {
		panic("no return value specified for Transaction")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(func(persistent.Transaction) error) error); ok {
		r0 = rf(fc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewTransactionManager creates a new instance of TransactionManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTransactionManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *TransactionManager {
	mock := &TransactionManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
