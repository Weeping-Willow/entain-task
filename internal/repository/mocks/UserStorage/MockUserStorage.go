// Code generated by mockery. DO NOT EDIT.

package mockUserStorage

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockUserStorage is an autogenerated mock type for the UserStorage type
type MockUserStorage struct {
	mock.Mock
}

type MockUserStorage_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUserStorage) EXPECT() *MockUserStorage_Expecter {
	return &MockUserStorage_Expecter{mock: &_m.Mock}
}

// GetUserBalance provides a mock function with given fields: ctx, userID
func (_m *MockUserStorage) GetUserBalance(ctx context.Context, userID uint64) (float64, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetUserBalance")
	}

	var r0 float64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (float64, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) float64); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserStorage_GetUserBalance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserBalance'
type MockUserStorage_GetUserBalance_Call struct {
	*mock.Call
}

// GetUserBalance is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uint64
func (_e *MockUserStorage_Expecter) GetUserBalance(ctx interface{}, userID interface{}) *MockUserStorage_GetUserBalance_Call {
	return &MockUserStorage_GetUserBalance_Call{Call: _e.mock.On("GetUserBalance", ctx, userID)}
}

func (_c *MockUserStorage_GetUserBalance_Call) Run(run func(ctx context.Context, userID uint64)) *MockUserStorage_GetUserBalance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *MockUserStorage_GetUserBalance_Call) Return(_a0 float64, _a1 error) *MockUserStorage_GetUserBalance_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserStorage_GetUserBalance_Call) RunAndReturn(run func(context.Context, uint64) (float64, error)) *MockUserStorage_GetUserBalance_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUserStorage creates a new instance of MockUserStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUserStorage(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUserStorage {
	mock := &MockUserStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}