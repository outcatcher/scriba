// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "github.com/outcatcher/scriba/internal/entities"
	mock "github.com/stretchr/testify/mock"

	time "time"

	uuid "github.com/google/uuid"
)

// Mockrepository is an autogenerated mock type for the repository type
type Mockrepository struct {
	mock.Mock
}

type Mockrepository_Expecter struct {
	mock *mock.Mock
}

func (_m *Mockrepository) EXPECT() *Mockrepository_Expecter {
	return &Mockrepository_Expecter{mock: &_m.Mock}
}

// CreateUserFromTG provides a mock function with given fields: ctx, telegramID
func (_m *Mockrepository) CreateUserFromTG(ctx context.Context, telegramID int64) (uuid.UUID, error) {
	ret := _m.Called(ctx, telegramID)

	var r0 uuid.UUID
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (uuid.UUID, error)); ok {
		return rf(ctx, telegramID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) uuid.UUID); ok {
		r0 = rf(ctx, telegramID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, telegramID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Mockrepository_CreateUserFromTG_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUserFromTG'
type Mockrepository_CreateUserFromTG_Call struct {
	*mock.Call
}

// CreateUserFromTG is a helper method to define mock.On call
//   - ctx context.Context
//   - telegramID int64
func (_e *Mockrepository_Expecter) CreateUserFromTG(ctx interface{}, telegramID interface{}) *Mockrepository_CreateUserFromTG_Call {
	return &Mockrepository_CreateUserFromTG_Call{Call: _e.mock.On("CreateUserFromTG", ctx, telegramID)}
}

func (_c *Mockrepository_CreateUserFromTG_Call) Run(run func(ctx context.Context, telegramID int64)) *Mockrepository_CreateUserFromTG_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *Mockrepository_CreateUserFromTG_Call) Return(_a0 uuid.UUID, _a1 error) *Mockrepository_CreateUserFromTG_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Mockrepository_CreateUserFromTG_Call) RunAndReturn(run func(context.Context, int64) (uuid.UUID, error)) *Mockrepository_CreateUserFromTG_Call {
	_c.Call.Return(run)
	return _c
}

// FindUserByTelegramID provides a mock function with given fields: ctx, telegramID
func (_m *Mockrepository) FindUserByTelegramID(ctx context.Context, telegramID int64) (*entities.Player, error) {
	ret := _m.Called(ctx, telegramID)

	var r0 *entities.Player
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*entities.Player, error)); ok {
		return rf(ctx, telegramID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *entities.Player); ok {
		r0 = rf(ctx, telegramID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Player)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, telegramID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Mockrepository_FindUserByTelegramID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindUserByTelegramID'
type Mockrepository_FindUserByTelegramID_Call struct {
	*mock.Call
}

// FindUserByTelegramID is a helper method to define mock.On call
//   - ctx context.Context
//   - telegramID int64
func (_e *Mockrepository_Expecter) FindUserByTelegramID(ctx interface{}, telegramID interface{}) *Mockrepository_FindUserByTelegramID_Call {
	return &Mockrepository_FindUserByTelegramID_Call{Call: _e.mock.On("FindUserByTelegramID", ctx, telegramID)}
}

func (_c *Mockrepository_FindUserByTelegramID_Call) Run(run func(ctx context.Context, telegramID int64)) *Mockrepository_FindUserByTelegramID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *Mockrepository_FindUserByTelegramID_Call) Return(_a0 *entities.Player, _a1 error) *Mockrepository_FindUserByTelegramID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Mockrepository_FindUserByTelegramID_Call) RunAndReturn(run func(context.Context, int64) (*entities.Player, error)) *Mockrepository_FindUserByTelegramID_Call {
	_c.Call.Return(run)
	return _c
}

// GetCountHistoryForPeriod provides a mock function with given fields: ctx, id, startDate, endDate
func (_m *Mockrepository) GetCountHistoryForPeriod(ctx context.Context, id uuid.UUID, startDate time.Time, endDate time.Time) ([]entities.CountHistoryEvent, error) {
	ret := _m.Called(ctx, id, startDate, endDate)

	var r0 []entities.CountHistoryEvent
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, time.Time, time.Time) ([]entities.CountHistoryEvent, error)); ok {
		return rf(ctx, id, startDate, endDate)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, time.Time, time.Time) []entities.CountHistoryEvent); ok {
		r0 = rf(ctx, id, startDate, endDate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.CountHistoryEvent)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, time.Time, time.Time) error); ok {
		r1 = rf(ctx, id, startDate, endDate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Mockrepository_GetCountHistoryForPeriod_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCountHistoryForPeriod'
type Mockrepository_GetCountHistoryForPeriod_Call struct {
	*mock.Call
}

// GetCountHistoryForPeriod is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
//   - startDate time.Time
//   - endDate time.Time
func (_e *Mockrepository_Expecter) GetCountHistoryForPeriod(ctx interface{}, id interface{}, startDate interface{}, endDate interface{}) *Mockrepository_GetCountHistoryForPeriod_Call {
	return &Mockrepository_GetCountHistoryForPeriod_Call{Call: _e.mock.On("GetCountHistoryForPeriod", ctx, id, startDate, endDate)}
}

func (_c *Mockrepository_GetCountHistoryForPeriod_Call) Run(run func(ctx context.Context, id uuid.UUID, startDate time.Time, endDate time.Time)) *Mockrepository_GetCountHistoryForPeriod_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(time.Time), args[3].(time.Time))
	})
	return _c
}

func (_c *Mockrepository_GetCountHistoryForPeriod_Call) Return(_a0 []entities.CountHistoryEvent, _a1 error) *Mockrepository_GetCountHistoryForPeriod_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Mockrepository_GetCountHistoryForPeriod_Call) RunAndReturn(run func(context.Context, uuid.UUID, time.Time, time.Time) ([]entities.CountHistoryEvent, error)) *Mockrepository_GetCountHistoryForPeriod_Call {
	_c.Call.Return(run)
	return _c
}

// GetPlayerCount provides a mock function with given fields: ctx, id
func (_m *Mockrepository) GetPlayerCount(ctx context.Context, id uuid.UUID) (int32, error) {
	ret := _m.Called(ctx, id)

	var r0 int32
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (int32, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) int32); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(int32)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Mockrepository_GetPlayerCount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPlayerCount'
type Mockrepository_GetPlayerCount_Call struct {
	*mock.Call
}

// GetPlayerCount is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
func (_e *Mockrepository_Expecter) GetPlayerCount(ctx interface{}, id interface{}) *Mockrepository_GetPlayerCount_Call {
	return &Mockrepository_GetPlayerCount_Call{Call: _e.mock.On("GetPlayerCount", ctx, id)}
}

func (_c *Mockrepository_GetPlayerCount_Call) Run(run func(ctx context.Context, id uuid.UUID)) *Mockrepository_GetPlayerCount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *Mockrepository_GetPlayerCount_Call) Return(_a0 int32, _a1 error) *Mockrepository_GetPlayerCount_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Mockrepository_GetPlayerCount_Call) RunAndReturn(run func(context.Context, uuid.UUID) (int32, error)) *Mockrepository_GetPlayerCount_Call {
	_c.Call.Return(run)
	return _c
}

// InsertPlayerCountChange provides a mock function with given fields: ctx, playerID, delta
func (_m *Mockrepository) InsertPlayerCountChange(ctx context.Context, playerID uuid.UUID, delta int16) error {
	ret := _m.Called(ctx, playerID, delta)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, int16) error); ok {
		r0 = rf(ctx, playerID, delta)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Mockrepository_InsertPlayerCountChange_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InsertPlayerCountChange'
type Mockrepository_InsertPlayerCountChange_Call struct {
	*mock.Call
}

// InsertPlayerCountChange is a helper method to define mock.On call
//   - ctx context.Context
//   - playerID uuid.UUID
//   - delta int16
func (_e *Mockrepository_Expecter) InsertPlayerCountChange(ctx interface{}, playerID interface{}, delta interface{}) *Mockrepository_InsertPlayerCountChange_Call {
	return &Mockrepository_InsertPlayerCountChange_Call{Call: _e.mock.On("InsertPlayerCountChange", ctx, playerID, delta)}
}

func (_c *Mockrepository_InsertPlayerCountChange_Call) Run(run func(ctx context.Context, playerID uuid.UUID, delta int16)) *Mockrepository_InsertPlayerCountChange_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(int16))
	})
	return _c
}

func (_c *Mockrepository_InsertPlayerCountChange_Call) Return(_a0 error) *Mockrepository_InsertPlayerCountChange_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Mockrepository_InsertPlayerCountChange_Call) RunAndReturn(run func(context.Context, uuid.UUID, int16) error) *Mockrepository_InsertPlayerCountChange_Call {
	_c.Call.Return(run)
	return _c
}

// ListPlayers provides a mock function with given fields: ctx
func (_m *Mockrepository) ListPlayers(ctx context.Context) ([]entities.Player, error) {
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

// Mockrepository_ListPlayers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListPlayers'
type Mockrepository_ListPlayers_Call struct {
	*mock.Call
}

// ListPlayers is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Mockrepository_Expecter) ListPlayers(ctx interface{}) *Mockrepository_ListPlayers_Call {
	return &Mockrepository_ListPlayers_Call{Call: _e.mock.On("ListPlayers", ctx)}
}

func (_c *Mockrepository_ListPlayers_Call) Run(run func(ctx context.Context)) *Mockrepository_ListPlayers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Mockrepository_ListPlayers_Call) Return(_a0 []entities.Player, _a1 error) *Mockrepository_ListPlayers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Mockrepository_ListPlayers_Call) RunAndReturn(run func(context.Context) ([]entities.Player, error)) *Mockrepository_ListPlayers_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockrepository creates a new instance of Mockrepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockrepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Mockrepository {
	mock := &Mockrepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
