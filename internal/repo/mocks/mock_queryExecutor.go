// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	sql "database/sql"

	sqlx "github.com/jmoiron/sqlx"
)

// MockqueryExecutor is an autogenerated mock type for the queryExecutor type
type MockqueryExecutor struct {
	mock.Mock
}

type MockqueryExecutor_Expecter struct {
	mock *mock.Mock
}

func (_m *MockqueryExecutor) EXPECT() *MockqueryExecutor_Expecter {
	return &MockqueryExecutor_Expecter{mock: &_m.Mock}
}

// BindNamed provides a mock function with given fields: _a0, _a1
func (_m *MockqueryExecutor) BindNamed(_a0 string, _a1 interface{}) (string, []interface{}, error) {
	ret := _m.Called(_a0, _a1)

	var r0 string
	var r1 []interface{}
	var r2 error
	if rf, ok := ret.Get(0).(func(string, interface{}) (string, []interface{}, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(string, interface{}) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, interface{}) []interface{}); ok {
		r1 = rf(_a0, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]interface{})
		}
	}

	if rf, ok := ret.Get(2).(func(string, interface{}) error); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockqueryExecutor_BindNamed_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BindNamed'
type MockqueryExecutor_BindNamed_Call struct {
	*mock.Call
}

// BindNamed is a helper method to define mock.On call
//   - _a0 string
//   - _a1 interface{}
func (_e *MockqueryExecutor_Expecter) BindNamed(_a0 interface{}, _a1 interface{}) *MockqueryExecutor_BindNamed_Call {
	return &MockqueryExecutor_BindNamed_Call{Call: _e.mock.On("BindNamed", _a0, _a1)}
}

func (_c *MockqueryExecutor_BindNamed_Call) Run(run func(_a0 string, _a1 interface{})) *MockqueryExecutor_BindNamed_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(interface{}))
	})
	return _c
}

func (_c *MockqueryExecutor_BindNamed_Call) Return(_a0 string, _a1 []interface{}, _a2 error) *MockqueryExecutor_BindNamed_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockqueryExecutor_BindNamed_Call) RunAndReturn(run func(string, interface{}) (string, []interface{}, error)) *MockqueryExecutor_BindNamed_Call {
	_c.Call.Return(run)
	return _c
}

// DriverName provides a mock function with given fields:
func (_m *MockqueryExecutor) DriverName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockqueryExecutor_DriverName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DriverName'
type MockqueryExecutor_DriverName_Call struct {
	*mock.Call
}

// DriverName is a helper method to define mock.On call
func (_e *MockqueryExecutor_Expecter) DriverName() *MockqueryExecutor_DriverName_Call {
	return &MockqueryExecutor_DriverName_Call{Call: _e.mock.On("DriverName")}
}

func (_c *MockqueryExecutor_DriverName_Call) Run(run func()) *MockqueryExecutor_DriverName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockqueryExecutor_DriverName_Call) Return(_a0 string) *MockqueryExecutor_DriverName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockqueryExecutor_DriverName_Call) RunAndReturn(run func() string) *MockqueryExecutor_DriverName_Call {
	_c.Call.Return(run)
	return _c
}

// ExecContext provides a mock function with given fields: ctx, query, args
func (_m *MockqueryExecutor) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 sql.Result
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) (sql.Result, error)); ok {
		return rf(ctx, query, args...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) sql.Result); ok {
		r0 = rf(ctx, query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sql.Result)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockqueryExecutor_ExecContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExecContext'
type MockqueryExecutor_ExecContext_Call struct {
	*mock.Call
}

// ExecContext is a helper method to define mock.On call
//   - ctx context.Context
//   - query string
//   - args ...interface{}
func (_e *MockqueryExecutor_Expecter) ExecContext(ctx interface{}, query interface{}, args ...interface{}) *MockqueryExecutor_ExecContext_Call {
	return &MockqueryExecutor_ExecContext_Call{Call: _e.mock.On("ExecContext",
		append([]interface{}{ctx, query}, args...)...)}
}

func (_c *MockqueryExecutor_ExecContext_Call) Run(run func(ctx context.Context, query string, args ...interface{})) *MockqueryExecutor_ExecContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockqueryExecutor_ExecContext_Call) Return(_a0 sql.Result, _a1 error) *MockqueryExecutor_ExecContext_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockqueryExecutor_ExecContext_Call) RunAndReturn(run func(context.Context, string, ...interface{}) (sql.Result, error)) *MockqueryExecutor_ExecContext_Call {
	_c.Call.Return(run)
	return _c
}

// GetContext provides a mock function with given fields: ctx, dest, query, args
func (_m *MockqueryExecutor) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, ctx, dest, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, string, ...interface{}) error); ok {
		r0 = rf(ctx, dest, query, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockqueryExecutor_GetContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetContext'
type MockqueryExecutor_GetContext_Call struct {
	*mock.Call
}

// GetContext is a helper method to define mock.On call
//   - ctx context.Context
//   - dest interface{}
//   - query string
//   - args ...interface{}
func (_e *MockqueryExecutor_Expecter) GetContext(ctx interface{}, dest interface{}, query interface{}, args ...interface{}) *MockqueryExecutor_GetContext_Call {
	return &MockqueryExecutor_GetContext_Call{Call: _e.mock.On("GetContext",
		append([]interface{}{ctx, dest, query}, args...)...)}
}

func (_c *MockqueryExecutor_GetContext_Call) Run(run func(ctx context.Context, dest interface{}, query string, args ...interface{})) *MockqueryExecutor_GetContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(interface{}), args[2].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockqueryExecutor_GetContext_Call) Return(_a0 error) *MockqueryExecutor_GetContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockqueryExecutor_GetContext_Call) RunAndReturn(run func(context.Context, interface{}, string, ...interface{}) error) *MockqueryExecutor_GetContext_Call {
	_c.Call.Return(run)
	return _c
}

// QueryContext provides a mock function with given fields: ctx, query, args
func (_m *MockqueryExecutor) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 *sql.Rows
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) (*sql.Rows, error)); ok {
		return rf(ctx, query, args...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) *sql.Rows); ok {
		r0 = rf(ctx, query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Rows)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockqueryExecutor_QueryContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryContext'
type MockqueryExecutor_QueryContext_Call struct {
	*mock.Call
}

// QueryContext is a helper method to define mock.On call
//   - ctx context.Context
//   - query string
//   - args ...interface{}
func (_e *MockqueryExecutor_Expecter) QueryContext(ctx interface{}, query interface{}, args ...interface{}) *MockqueryExecutor_QueryContext_Call {
	return &MockqueryExecutor_QueryContext_Call{Call: _e.mock.On("QueryContext",
		append([]interface{}{ctx, query}, args...)...)}
}

func (_c *MockqueryExecutor_QueryContext_Call) Run(run func(ctx context.Context, query string, args ...interface{})) *MockqueryExecutor_QueryContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockqueryExecutor_QueryContext_Call) Return(_a0 *sql.Rows, _a1 error) *MockqueryExecutor_QueryContext_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockqueryExecutor_QueryContext_Call) RunAndReturn(run func(context.Context, string, ...interface{}) (*sql.Rows, error)) *MockqueryExecutor_QueryContext_Call {
	_c.Call.Return(run)
	return _c
}

// QueryRowxContext provides a mock function with given fields: ctx, query, args
func (_m *MockqueryExecutor) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 *sqlx.Row
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) *sqlx.Row); ok {
		r0 = rf(ctx, query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqlx.Row)
		}
	}

	return r0
}

// MockqueryExecutor_QueryRowxContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryRowxContext'
type MockqueryExecutor_QueryRowxContext_Call struct {
	*mock.Call
}

// QueryRowxContext is a helper method to define mock.On call
//   - ctx context.Context
//   - query string
//   - args ...interface{}
func (_e *MockqueryExecutor_Expecter) QueryRowxContext(ctx interface{}, query interface{}, args ...interface{}) *MockqueryExecutor_QueryRowxContext_Call {
	return &MockqueryExecutor_QueryRowxContext_Call{Call: _e.mock.On("QueryRowxContext",
		append([]interface{}{ctx, query}, args...)...)}
}

func (_c *MockqueryExecutor_QueryRowxContext_Call) Run(run func(ctx context.Context, query string, args ...interface{})) *MockqueryExecutor_QueryRowxContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockqueryExecutor_QueryRowxContext_Call) Return(_a0 *sqlx.Row) *MockqueryExecutor_QueryRowxContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockqueryExecutor_QueryRowxContext_Call) RunAndReturn(run func(context.Context, string, ...interface{}) *sqlx.Row) *MockqueryExecutor_QueryRowxContext_Call {
	_c.Call.Return(run)
	return _c
}

// QueryxContext provides a mock function with given fields: ctx, query, args
func (_m *MockqueryExecutor) QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 *sqlx.Rows
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) (*sqlx.Rows, error)); ok {
		return rf(ctx, query, args...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) *sqlx.Rows); ok {
		r0 = rf(ctx, query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqlx.Rows)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockqueryExecutor_QueryxContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryxContext'
type MockqueryExecutor_QueryxContext_Call struct {
	*mock.Call
}

// QueryxContext is a helper method to define mock.On call
//   - ctx context.Context
//   - query string
//   - args ...interface{}
func (_e *MockqueryExecutor_Expecter) QueryxContext(ctx interface{}, query interface{}, args ...interface{}) *MockqueryExecutor_QueryxContext_Call {
	return &MockqueryExecutor_QueryxContext_Call{Call: _e.mock.On("QueryxContext",
		append([]interface{}{ctx, query}, args...)...)}
}

func (_c *MockqueryExecutor_QueryxContext_Call) Run(run func(ctx context.Context, query string, args ...interface{})) *MockqueryExecutor_QueryxContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockqueryExecutor_QueryxContext_Call) Return(_a0 *sqlx.Rows, _a1 error) *MockqueryExecutor_QueryxContext_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockqueryExecutor_QueryxContext_Call) RunAndReturn(run func(context.Context, string, ...interface{}) (*sqlx.Rows, error)) *MockqueryExecutor_QueryxContext_Call {
	_c.Call.Return(run)
	return _c
}

// Rebind provides a mock function with given fields: _a0
func (_m *MockqueryExecutor) Rebind(_a0 string) string {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockqueryExecutor_Rebind_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Rebind'
type MockqueryExecutor_Rebind_Call struct {
	*mock.Call
}

// Rebind is a helper method to define mock.On call
//   - _a0 string
func (_e *MockqueryExecutor_Expecter) Rebind(_a0 interface{}) *MockqueryExecutor_Rebind_Call {
	return &MockqueryExecutor_Rebind_Call{Call: _e.mock.On("Rebind", _a0)}
}

func (_c *MockqueryExecutor_Rebind_Call) Run(run func(_a0 string)) *MockqueryExecutor_Rebind_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockqueryExecutor_Rebind_Call) Return(_a0 string) *MockqueryExecutor_Rebind_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockqueryExecutor_Rebind_Call) RunAndReturn(run func(string) string) *MockqueryExecutor_Rebind_Call {
	_c.Call.Return(run)
	return _c
}

// SelectContext provides a mock function with given fields: ctx, dest, query, args
func (_m *MockqueryExecutor) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, ctx, dest, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, string, ...interface{}) error); ok {
		r0 = rf(ctx, dest, query, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockqueryExecutor_SelectContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SelectContext'
type MockqueryExecutor_SelectContext_Call struct {
	*mock.Call
}

// SelectContext is a helper method to define mock.On call
//   - ctx context.Context
//   - dest interface{}
//   - query string
//   - args ...interface{}
func (_e *MockqueryExecutor_Expecter) SelectContext(ctx interface{}, dest interface{}, query interface{}, args ...interface{}) *MockqueryExecutor_SelectContext_Call {
	return &MockqueryExecutor_SelectContext_Call{Call: _e.mock.On("SelectContext",
		append([]interface{}{ctx, dest, query}, args...)...)}
}

func (_c *MockqueryExecutor_SelectContext_Call) Run(run func(ctx context.Context, dest interface{}, query string, args ...interface{})) *MockqueryExecutor_SelectContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(interface{}), args[2].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockqueryExecutor_SelectContext_Call) Return(_a0 error) *MockqueryExecutor_SelectContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockqueryExecutor_SelectContext_Call) RunAndReturn(run func(context.Context, interface{}, string, ...interface{}) error) *MockqueryExecutor_SelectContext_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockqueryExecutor creates a new instance of MockqueryExecutor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockqueryExecutor(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockqueryExecutor {
	mock := &MockqueryExecutor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
