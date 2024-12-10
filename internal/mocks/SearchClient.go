// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	context "context"

	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"

	search "github.com/bangumi/server/internal/search"
)

// SearchClient is an autogenerated mock type for the Client type
type SearchClient struct {
	mock.Mock
}

type SearchClient_Expecter struct {
	mock *mock.Mock
}

func (_m *SearchClient) EXPECT() *SearchClient_Expecter {
	return &SearchClient_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *SearchClient) Close() {
	_m.Called()
}

// SearchClient_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type SearchClient_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *SearchClient_Expecter) Close() *SearchClient_Close_Call {
	return &SearchClient_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *SearchClient_Close_Call) Run(run func()) *SearchClient_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *SearchClient_Close_Call) Return() *SearchClient_Close_Call {
	_c.Call.Return()
	return _c
}

func (_c *SearchClient_Close_Call) RunAndReturn(run func()) *SearchClient_Close_Call {
	_c.Call.Return(run)
	return _c
}

// EventAdded provides a mock function with given fields: ctx, id, target
func (_m *SearchClient) EventAdded(ctx context.Context, id uint32, target search.SearchTarget) error {
	ret := _m.Called(ctx, id, target)

	if len(ret) == 0 {
		panic("no return value specified for EventAdded")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, search.SearchTarget) error); ok {
		r0 = rf(ctx, id, target)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SearchClient_EventAdded_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EventAdded'
type SearchClient_EventAdded_Call struct {
	*mock.Call
}

// EventAdded is a helper method to define mock.On call
//   - ctx context.Context
//   - id uint32
//   - target search.SearchTarget
func (_e *SearchClient_Expecter) EventAdded(ctx interface{}, id interface{}, target interface{}) *SearchClient_EventAdded_Call {
	return &SearchClient_EventAdded_Call{Call: _e.mock.On("EventAdded", ctx, id, target)}
}

func (_c *SearchClient_EventAdded_Call) Run(run func(ctx context.Context, id uint32, target search.SearchTarget)) *SearchClient_EventAdded_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(search.SearchTarget))
	})
	return _c
}

func (_c *SearchClient_EventAdded_Call) Return(_a0 error) *SearchClient_EventAdded_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SearchClient_EventAdded_Call) RunAndReturn(run func(context.Context, uint32, search.SearchTarget) error) *SearchClient_EventAdded_Call {
	_c.Call.Return(run)
	return _c
}

// EventDelete provides a mock function with given fields: ctx, id, target
func (_m *SearchClient) EventDelete(ctx context.Context, id uint32, target search.SearchTarget) error {
	ret := _m.Called(ctx, id, target)

	if len(ret) == 0 {
		panic("no return value specified for EventDelete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, search.SearchTarget) error); ok {
		r0 = rf(ctx, id, target)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SearchClient_EventDelete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EventDelete'
type SearchClient_EventDelete_Call struct {
	*mock.Call
}

// EventDelete is a helper method to define mock.On call
//   - ctx context.Context
//   - id uint32
//   - target search.SearchTarget
func (_e *SearchClient_Expecter) EventDelete(ctx interface{}, id interface{}, target interface{}) *SearchClient_EventDelete_Call {
	return &SearchClient_EventDelete_Call{Call: _e.mock.On("EventDelete", ctx, id, target)}
}

func (_c *SearchClient_EventDelete_Call) Run(run func(ctx context.Context, id uint32, target search.SearchTarget)) *SearchClient_EventDelete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(search.SearchTarget))
	})
	return _c
}

func (_c *SearchClient_EventDelete_Call) Return(_a0 error) *SearchClient_EventDelete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SearchClient_EventDelete_Call) RunAndReturn(run func(context.Context, uint32, search.SearchTarget) error) *SearchClient_EventDelete_Call {
	_c.Call.Return(run)
	return _c
}

// EventUpdate provides a mock function with given fields: ctx, id, target
func (_m *SearchClient) EventUpdate(ctx context.Context, id uint32, target search.SearchTarget) error {
	ret := _m.Called(ctx, id, target)

	if len(ret) == 0 {
		panic("no return value specified for EventUpdate")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, search.SearchTarget) error); ok {
		r0 = rf(ctx, id, target)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SearchClient_EventUpdate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EventUpdate'
type SearchClient_EventUpdate_Call struct {
	*mock.Call
}

// EventUpdate is a helper method to define mock.On call
//   - ctx context.Context
//   - id uint32
//   - target search.SearchTarget
func (_e *SearchClient_Expecter) EventUpdate(ctx interface{}, id interface{}, target interface{}) *SearchClient_EventUpdate_Call {
	return &SearchClient_EventUpdate_Call{Call: _e.mock.On("EventUpdate", ctx, id, target)}
}

func (_c *SearchClient_EventUpdate_Call) Run(run func(ctx context.Context, id uint32, target search.SearchTarget)) *SearchClient_EventUpdate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(search.SearchTarget))
	})
	return _c
}

func (_c *SearchClient_EventUpdate_Call) Return(_a0 error) *SearchClient_EventUpdate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SearchClient_EventUpdate_Call) RunAndReturn(run func(context.Context, uint32, search.SearchTarget) error) *SearchClient_EventUpdate_Call {
	_c.Call.Return(run)
	return _c
}

// Handle provides a mock function with given fields: c, target
func (_m *SearchClient) Handle(c echo.Context, target search.SearchTarget) error {
	ret := _m.Called(c, target)

	if len(ret) == 0 {
		panic("no return value specified for Handle")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context, search.SearchTarget) error); ok {
		r0 = rf(c, target)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SearchClient_Handle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Handle'
type SearchClient_Handle_Call struct {
	*mock.Call
}

// Handle is a helper method to define mock.On call
//   - c echo.Context
//   - target search.SearchTarget
func (_e *SearchClient_Expecter) Handle(c interface{}, target interface{}) *SearchClient_Handle_Call {
	return &SearchClient_Handle_Call{Call: _e.mock.On("Handle", c, target)}
}

func (_c *SearchClient_Handle_Call) Run(run func(c echo.Context, target search.SearchTarget)) *SearchClient_Handle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(echo.Context), args[1].(search.SearchTarget))
	})
	return _c
}

func (_c *SearchClient_Handle_Call) Return(_a0 error) *SearchClient_Handle_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SearchClient_Handle_Call) RunAndReturn(run func(echo.Context, search.SearchTarget) error) *SearchClient_Handle_Call {
	_c.Call.Return(run)
	return _c
}

// NewSearchClient creates a new instance of SearchClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSearchClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *SearchClient {
	mock := &SearchClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
