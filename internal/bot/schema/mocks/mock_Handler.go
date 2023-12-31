// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	telebot "gopkg.in/telebot.v3"
)

// MockHandler is an autogenerated mock type for the Handler type
type MockHandler struct {
	mock.Mock
}

type MockHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *MockHandler) EXPECT() *MockHandler_Expecter {
	return &MockHandler_Expecter{mock: &_m.Mock}
}

// Handle provides a mock function with given fields: endpoint, h, m
func (_m *MockHandler) Handle(endpoint interface{}, h telebot.HandlerFunc, m ...telebot.MiddlewareFunc) {
	_va := make([]interface{}, len(m))
	for _i := range m {
		_va[_i] = m[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, endpoint, h)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// MockHandler_Handle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Handle'
type MockHandler_Handle_Call struct {
	*mock.Call
}

// Handle is a helper method to define mock.On call
//   - endpoint interface{}
//   - h telebot.HandlerFunc
//   - m ...telebot.MiddlewareFunc
func (_e *MockHandler_Expecter) Handle(endpoint interface{}, h interface{}, m ...interface{}) *MockHandler_Handle_Call {
	return &MockHandler_Handle_Call{Call: _e.mock.On("Handle",
		append([]interface{}{endpoint, h}, m...)...)}
}

func (_c *MockHandler_Handle_Call) Run(run func(endpoint interface{}, h telebot.HandlerFunc, m ...telebot.MiddlewareFunc)) *MockHandler_Handle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]telebot.MiddlewareFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(telebot.MiddlewareFunc)
			}
		}
		run(args[0].(interface{}), args[1].(telebot.HandlerFunc), variadicArgs...)
	})
	return _c
}

func (_c *MockHandler_Handle_Call) Return() *MockHandler_Handle_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockHandler_Handle_Call) RunAndReturn(run func(interface{}, telebot.HandlerFunc, ...telebot.MiddlewareFunc)) *MockHandler_Handle_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockHandler creates a new instance of MockHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockHandler {
	mock := &MockHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
