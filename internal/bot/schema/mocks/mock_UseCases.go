// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "github.com/outcatcher/scriba/internal/entities"
	mock "github.com/stretchr/testify/mock"
)

// MockUseCases is an autogenerated mock type for the UseCases type
type MockUseCases struct {
	mock.Mock
}

type MockUseCases_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUseCases) EXPECT() *MockUseCases_Expecter {
	return &MockUseCases_Expecter{mock: &_m.Mock}
}

// GetPlayerCountByTelegramID provides a mock function with given fields: ctx, telegramID
func (_m *MockUseCases) GetPlayerCountByTelegramID(ctx context.Context, telegramID int64) (int32, error) {
	ret := _m.Called(ctx, telegramID)

	var r0 int32
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (int32, error)); ok {
		return rf(ctx, telegramID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) int32); ok {
		r0 = rf(ctx, telegramID)
	} else {
		r0 = ret.Get(0).(int32)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, telegramID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUseCases_GetPlayerCountByTelegramID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPlayerCountByTelegramID'
type MockUseCases_GetPlayerCountByTelegramID_Call struct {
	*mock.Call
}

// GetPlayerCountByTelegramID is a helper method to define mock.On call
//   - ctx context.Context
//   - telegramID int64
func (_e *MockUseCases_Expecter) GetPlayerCountByTelegramID(ctx interface{}, telegramID interface{}) *MockUseCases_GetPlayerCountByTelegramID_Call {
	return &MockUseCases_GetPlayerCountByTelegramID_Call{Call: _e.mock.On("GetPlayerCountByTelegramID", ctx, telegramID)}
}

func (_c *MockUseCases_GetPlayerCountByTelegramID_Call) Run(run func(ctx context.Context, telegramID int64)) *MockUseCases_GetPlayerCountByTelegramID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockUseCases_GetPlayerCountByTelegramID_Call) Return(_a0 int32, _a1 error) *MockUseCases_GetPlayerCountByTelegramID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUseCases_GetPlayerCountByTelegramID_Call) RunAndReturn(run func(context.Context, int64) (int32, error)) *MockUseCases_GetPlayerCountByTelegramID_Call {
	_c.Call.Return(run)
	return _c
}

// ListPlayers provides a mock function with given fields: ctx
func (_m *MockUseCases) ListPlayers(ctx context.Context) ([]entities.Player, error) {
	ret := _m.Called(ctx)

	var r0 []entities.Player
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]entities.Player, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []entities.Player); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.Player)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUseCases_ListPlayers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListPlayers'
type MockUseCases_ListPlayers_Call struct {
	*mock.Call
}

// ListPlayers is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockUseCases_Expecter) ListPlayers(ctx interface{}) *MockUseCases_ListPlayers_Call {
	return &MockUseCases_ListPlayers_Call{Call: _e.mock.On("ListPlayers", ctx)}
}

func (_c *MockUseCases_ListPlayers_Call) Run(run func(ctx context.Context)) *MockUseCases_ListPlayers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockUseCases_ListPlayers_Call) Return(_a0 []entities.Player, _a1 error) *MockUseCases_ListPlayers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUseCases_ListPlayers_Call) RunAndReturn(run func(context.Context) ([]entities.Player, error)) *MockUseCases_ListPlayers_Call {
	_c.Call.Return(run)
	return _c
}

// RegisterWithTelegram provides a mock function with given fields: ctx, telegramID
func (_m *MockUseCases) RegisterWithTelegram(ctx context.Context, telegramID int64) error {
	ret := _m.Called(ctx, telegramID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, telegramID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUseCases_RegisterWithTelegram_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RegisterWithTelegram'
type MockUseCases_RegisterWithTelegram_Call struct {
	*mock.Call
}

// RegisterWithTelegram is a helper method to define mock.On call
//   - ctx context.Context
//   - telegramID int64
func (_e *MockUseCases_Expecter) RegisterWithTelegram(ctx interface{}, telegramID interface{}) *MockUseCases_RegisterWithTelegram_Call {
	return &MockUseCases_RegisterWithTelegram_Call{Call: _e.mock.On("RegisterWithTelegram", ctx, telegramID)}
}

func (_c *MockUseCases_RegisterWithTelegram_Call) Run(run func(ctx context.Context, telegramID int64)) *MockUseCases_RegisterWithTelegram_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockUseCases_RegisterWithTelegram_Call) Return(_a0 error) *MockUseCases_RegisterWithTelegram_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUseCases_RegisterWithTelegram_Call) RunAndReturn(run func(context.Context, int64) error) *MockUseCases_RegisterWithTelegram_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateCountByTelegramID provides a mock function with given fields: ctx, telegramID, delta
func (_m *MockUseCases) UpdateCountByTelegramID(ctx context.Context, telegramID int64, delta int16) error {
	ret := _m.Called(ctx, telegramID, delta)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int16) error); ok {
		r0 = rf(ctx, telegramID, delta)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUseCases_UpdateCountByTelegramID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateCountByTelegramID'
type MockUseCases_UpdateCountByTelegramID_Call struct {
	*mock.Call
}

// UpdateCountByTelegramID is a helper method to define mock.On call
//   - ctx context.Context
//   - telegramID int64
//   - delta int16
func (_e *MockUseCases_Expecter) UpdateCountByTelegramID(ctx interface{}, telegramID interface{}, delta interface{}) *MockUseCases_UpdateCountByTelegramID_Call {
	return &MockUseCases_UpdateCountByTelegramID_Call{Call: _e.mock.On("UpdateCountByTelegramID", ctx, telegramID, delta)}
}

func (_c *MockUseCases_UpdateCountByTelegramID_Call) Run(run func(ctx context.Context, telegramID int64, delta int16)) *MockUseCases_UpdateCountByTelegramID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64), args[2].(int16))
	})
	return _c
}

func (_c *MockUseCases_UpdateCountByTelegramID_Call) Return(_a0 error) *MockUseCases_UpdateCountByTelegramID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUseCases_UpdateCountByTelegramID_Call) RunAndReturn(run func(context.Context, int64, int16) error) *MockUseCases_UpdateCountByTelegramID_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUseCases creates a new instance of MockUseCases. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUseCases(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUseCases {
	mock := &MockUseCases{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}