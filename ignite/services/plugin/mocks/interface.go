// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	plugin "github.com/ignite/cli/ignite/services/plugin"
	mock "github.com/stretchr/testify/mock"
)

// PluginInterface is an autogenerated mock type for the Interface type
type PluginInterface struct {
	mock.Mock
}

type PluginInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *PluginInterface) EXPECT() *PluginInterface_Expecter {
	return &PluginInterface_Expecter{mock: &_m.Mock}
}

// Commands provides a mock function with given fields:
func (_m *PluginInterface) Commands() ([]plugin.Command, error) {
	ret := _m.Called()

	var r0 []plugin.Command
	if rf, ok := ret.Get(0).(func() []plugin.Command); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]plugin.Command)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PluginInterface_Commands_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Commands'
type PluginInterface_Commands_Call struct {
	*mock.Call
}

// Commands is a helper method to define mock.On call
func (_e *PluginInterface_Expecter) Commands() *PluginInterface_Commands_Call {
	return &PluginInterface_Commands_Call{Call: _e.mock.On("Commands")}
}

func (_c *PluginInterface_Commands_Call) Run(run func()) *PluginInterface_Commands_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PluginInterface_Commands_Call) Return(_a0 []plugin.Command, _a1 error) *PluginInterface_Commands_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Execute provides a mock function with given fields: cmd, args
func (_m *PluginInterface) Execute(cmd plugin.Command, args []string) error {
	ret := _m.Called(cmd, args)

	var r0 error
	if rf, ok := ret.Get(0).(func(plugin.Command, []string) error); ok {
		r0 = rf(cmd, args)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PluginInterface_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type PluginInterface_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - cmd plugin.Command
//   - args []string
func (_e *PluginInterface_Expecter) Execute(cmd interface{}, args interface{}) *PluginInterface_Execute_Call {
	return &PluginInterface_Execute_Call{Call: _e.mock.On("Execute", cmd, args)}
}

func (_c *PluginInterface_Execute_Call) Run(run func(cmd plugin.Command, args []string)) *PluginInterface_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(plugin.Command), args[1].([]string))
	})
	return _c
}

func (_c *PluginInterface_Execute_Call) Return(_a0 error) *PluginInterface_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

// ExecuteHookCleanUp provides a mock function with given fields: hook, args
func (_m *PluginInterface) ExecuteHookCleanUp(hook plugin.Hook, args []string) error {
	ret := _m.Called(hook, args)

	var r0 error
	if rf, ok := ret.Get(0).(func(plugin.Hook, []string) error); ok {
		r0 = rf(hook, args)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PluginInterface_ExecuteHookCleanUp_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExecuteHookCleanUp'
type PluginInterface_ExecuteHookCleanUp_Call struct {
	*mock.Call
}

// ExecuteHookCleanUp is a helper method to define mock.On call
//   - hook plugin.Hook
//   - args []string
func (_e *PluginInterface_Expecter) ExecuteHookCleanUp(hook interface{}, args interface{}) *PluginInterface_ExecuteHookCleanUp_Call {
	return &PluginInterface_ExecuteHookCleanUp_Call{Call: _e.mock.On("ExecuteHookCleanUp", hook, args)}
}

func (_c *PluginInterface_ExecuteHookCleanUp_Call) Run(run func(hook plugin.Hook, args []string)) *PluginInterface_ExecuteHookCleanUp_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(plugin.Hook), args[1].([]string))
	})
	return _c
}

func (_c *PluginInterface_ExecuteHookCleanUp_Call) Return(_a0 error) *PluginInterface_ExecuteHookCleanUp_Call {
	_c.Call.Return(_a0)
	return _c
}

// ExecuteHookPost provides a mock function with given fields: hook, args
func (_m *PluginInterface) ExecuteHookPost(hook plugin.Hook, args []string) error {
	ret := _m.Called(hook, args)

	var r0 error
	if rf, ok := ret.Get(0).(func(plugin.Hook, []string) error); ok {
		r0 = rf(hook, args)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PluginInterface_ExecuteHookPost_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExecuteHookPost'
type PluginInterface_ExecuteHookPost_Call struct {
	*mock.Call
}

// ExecuteHookPost is a helper method to define mock.On call
//   - hook plugin.Hook
//   - args []string
func (_e *PluginInterface_Expecter) ExecuteHookPost(hook interface{}, args interface{}) *PluginInterface_ExecuteHookPost_Call {
	return &PluginInterface_ExecuteHookPost_Call{Call: _e.mock.On("ExecuteHookPost", hook, args)}
}

func (_c *PluginInterface_ExecuteHookPost_Call) Run(run func(hook plugin.Hook, args []string)) *PluginInterface_ExecuteHookPost_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(plugin.Hook), args[1].([]string))
	})
	return _c
}

func (_c *PluginInterface_ExecuteHookPost_Call) Return(_a0 error) *PluginInterface_ExecuteHookPost_Call {
	_c.Call.Return(_a0)
	return _c
}

// ExecuteHookPre provides a mock function with given fields: hook, args
func (_m *PluginInterface) ExecuteHookPre(hook plugin.Hook, args []string) error {
	ret := _m.Called(hook, args)

	var r0 error
	if rf, ok := ret.Get(0).(func(plugin.Hook, []string) error); ok {
		r0 = rf(hook, args)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PluginInterface_ExecuteHookPre_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExecuteHookPre'
type PluginInterface_ExecuteHookPre_Call struct {
	*mock.Call
}

// ExecuteHookPre is a helper method to define mock.On call
//   - hook plugin.Hook
//   - args []string
func (_e *PluginInterface_Expecter) ExecuteHookPre(hook interface{}, args interface{}) *PluginInterface_ExecuteHookPre_Call {
	return &PluginInterface_ExecuteHookPre_Call{Call: _e.mock.On("ExecuteHookPre", hook, args)}
}

func (_c *PluginInterface_ExecuteHookPre_Call) Run(run func(hook plugin.Hook, args []string)) *PluginInterface_ExecuteHookPre_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(plugin.Hook), args[1].([]string))
	})
	return _c
}

func (_c *PluginInterface_ExecuteHookPre_Call) Return(_a0 error) *PluginInterface_ExecuteHookPre_Call {
	_c.Call.Return(_a0)
	return _c
}

// Hooks provides a mock function with given fields:
func (_m *PluginInterface) Hooks() ([]plugin.Hook, error) {
	ret := _m.Called()

	var r0 []plugin.Hook
	if rf, ok := ret.Get(0).(func() []plugin.Hook); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]plugin.Hook)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PluginInterface_Hooks_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Hooks'
type PluginInterface_Hooks_Call struct {
	*mock.Call
}

// Hooks is a helper method to define mock.On call
func (_e *PluginInterface_Expecter) Hooks() *PluginInterface_Hooks_Call {
	return &PluginInterface_Hooks_Call{Call: _e.mock.On("Hooks")}
}

func (_c *PluginInterface_Hooks_Call) Run(run func()) *PluginInterface_Hooks_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PluginInterface_Hooks_Call) Return(_a0 []plugin.Hook, _a1 error) *PluginInterface_Hooks_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewPluginInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewPluginInterface creates a new instance of PluginInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPluginInterface(t mockConstructorTestingTNewPluginInterface) *PluginInterface {
	mock := &PluginInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
