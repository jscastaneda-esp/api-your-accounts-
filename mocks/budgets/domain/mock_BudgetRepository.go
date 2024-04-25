// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks_domain

import (
	context "context"
	domain "your-accounts-api/budgets/domain"

	mock "github.com/stretchr/testify/mock"

	persistent "your-accounts-api/shared/domain/persistent"
)

// MockBudgetRepository is an autogenerated mock type for the BudgetRepository type
type MockBudgetRepository struct {
	mock.Mock
}

type MockBudgetRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockBudgetRepository) EXPECT() *MockBudgetRepository_Expecter {
	return &MockBudgetRepository_Expecter{mock: &_m.Mock}
}

// Delete provides a mock function with given fields: ctx, id
func (_m *MockBudgetRepository) Delete(ctx context.Context, id uint) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBudgetRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockBudgetRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - id uint
func (_e *MockBudgetRepository_Expecter) Delete(ctx interface{}, id interface{}) *MockBudgetRepository_Delete_Call {
	return &MockBudgetRepository_Delete_Call{Call: _e.mock.On("Delete", ctx, id)}
}

func (_c *MockBudgetRepository_Delete_Call) Run(run func(ctx context.Context, id uint)) *MockBudgetRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint))
	})
	return _c
}

func (_c *MockBudgetRepository_Delete_Call) Return(_a0 error) *MockBudgetRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBudgetRepository_Delete_Call) RunAndReturn(run func(context.Context, uint) error) *MockBudgetRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Save provides a mock function with given fields: ctx, _a1
func (_m *MockBudgetRepository) Save(ctx context.Context, _a1 domain.Budget) (uint, error) {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 uint
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Budget) (uint, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.Budget) uint); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.Budget) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBudgetRepository_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type MockBudgetRepository_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 domain.Budget
func (_e *MockBudgetRepository_Expecter) Save(ctx interface{}, _a1 interface{}) *MockBudgetRepository_Save_Call {
	return &MockBudgetRepository_Save_Call{Call: _e.mock.On("Save", ctx, _a1)}
}

func (_c *MockBudgetRepository_Save_Call) Run(run func(ctx context.Context, _a1 domain.Budget)) *MockBudgetRepository_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.Budget))
	})
	return _c
}

func (_c *MockBudgetRepository_Save_Call) Return(_a0 uint, _a1 error) *MockBudgetRepository_Save_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBudgetRepository_Save_Call) RunAndReturn(run func(context.Context, domain.Budget) (uint, error)) *MockBudgetRepository_Save_Call {
	_c.Call.Return(run)
	return _c
}

// Search provides a mock function with given fields: ctx, id
func (_m *MockBudgetRepository) Search(ctx context.Context, id uint) (domain.Budget, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Search")
	}

	var r0 domain.Budget
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) (domain.Budget, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint) domain.Budget); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.Budget)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBudgetRepository_Search_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Search'
type MockBudgetRepository_Search_Call struct {
	*mock.Call
}

// Search is a helper method to define mock.On call
//   - ctx context.Context
//   - id uint
func (_e *MockBudgetRepository_Expecter) Search(ctx interface{}, id interface{}) *MockBudgetRepository_Search_Call {
	return &MockBudgetRepository_Search_Call{Call: _e.mock.On("Search", ctx, id)}
}

func (_c *MockBudgetRepository_Search_Call) Run(run func(ctx context.Context, id uint)) *MockBudgetRepository_Search_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint))
	})
	return _c
}

func (_c *MockBudgetRepository_Search_Call) Return(_a0 domain.Budget, _a1 error) *MockBudgetRepository_Search_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBudgetRepository_Search_Call) RunAndReturn(run func(context.Context, uint) (domain.Budget, error)) *MockBudgetRepository_Search_Call {
	_c.Call.Return(run)
	return _c
}

// SearchAllByExample provides a mock function with given fields: ctx, example
func (_m *MockBudgetRepository) SearchAllByExample(ctx context.Context, example domain.Budget) ([]domain.Budget, error) {
	ret := _m.Called(ctx, example)

	if len(ret) == 0 {
		panic("no return value specified for SearchAllByExample")
	}

	var r0 []domain.Budget
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Budget) ([]domain.Budget, error)); ok {
		return rf(ctx, example)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.Budget) []domain.Budget); ok {
		r0 = rf(ctx, example)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Budget)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.Budget) error); ok {
		r1 = rf(ctx, example)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBudgetRepository_SearchAllByExample_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SearchAllByExample'
type MockBudgetRepository_SearchAllByExample_Call struct {
	*mock.Call
}

// SearchAllByExample is a helper method to define mock.On call
//   - ctx context.Context
//   - example domain.Budget
func (_e *MockBudgetRepository_Expecter) SearchAllByExample(ctx interface{}, example interface{}) *MockBudgetRepository_SearchAllByExample_Call {
	return &MockBudgetRepository_SearchAllByExample_Call{Call: _e.mock.On("SearchAllByExample", ctx, example)}
}

func (_c *MockBudgetRepository_SearchAllByExample_Call) Run(run func(ctx context.Context, example domain.Budget)) *MockBudgetRepository_SearchAllByExample_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.Budget))
	})
	return _c
}

func (_c *MockBudgetRepository_SearchAllByExample_Call) Return(_a0 []domain.Budget, _a1 error) *MockBudgetRepository_SearchAllByExample_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBudgetRepository_SearchAllByExample_Call) RunAndReturn(run func(context.Context, domain.Budget) ([]domain.Budget, error)) *MockBudgetRepository_SearchAllByExample_Call {
	_c.Call.Return(run)
	return _c
}

// WithTransaction provides a mock function with given fields: tx
func (_m *MockBudgetRepository) WithTransaction(tx persistent.Transaction) domain.BudgetRepository {
	ret := _m.Called(tx)

	if len(ret) == 0 {
		panic("no return value specified for WithTransaction")
	}

	var r0 domain.BudgetRepository
	if rf, ok := ret.Get(0).(func(persistent.Transaction) domain.BudgetRepository); ok {
		r0 = rf(tx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.BudgetRepository)
		}
	}

	return r0
}

// MockBudgetRepository_WithTransaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithTransaction'
type MockBudgetRepository_WithTransaction_Call struct {
	*mock.Call
}

// WithTransaction is a helper method to define mock.On call
//   - tx persistent.Transaction
func (_e *MockBudgetRepository_Expecter) WithTransaction(tx interface{}) *MockBudgetRepository_WithTransaction_Call {
	return &MockBudgetRepository_WithTransaction_Call{Call: _e.mock.On("WithTransaction", tx)}
}

func (_c *MockBudgetRepository_WithTransaction_Call) Run(run func(tx persistent.Transaction)) *MockBudgetRepository_WithTransaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(persistent.Transaction))
	})
	return _c
}

func (_c *MockBudgetRepository_WithTransaction_Call) Return(_a0 domain.BudgetRepository) *MockBudgetRepository_WithTransaction_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBudgetRepository_WithTransaction_Call) RunAndReturn(run func(persistent.Transaction) domain.BudgetRepository) *MockBudgetRepository_WithTransaction_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockBudgetRepository creates a new instance of MockBudgetRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockBudgetRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockBudgetRepository {
	mock := &MockBudgetRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
