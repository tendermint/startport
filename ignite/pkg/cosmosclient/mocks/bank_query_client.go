// Code generated by mockery v2.27.1. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	types "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// BankQueryClient is an autogenerated mock type for the QueryClient type
type BankQueryClient struct {
	mock.Mock
}

type BankQueryClient_Expecter struct {
	mock *mock.Mock
}

func (_m *BankQueryClient) EXPECT() *BankQueryClient_Expecter {
	return &BankQueryClient_Expecter{mock: &_m.Mock}
}

// AllBalances provides a mock function with given fields: ctx, in, opts
func (_m *BankQueryClient) AllBalances(ctx context.Context, in *types.QueryAllBalancesRequest, opts ...grpc.CallOption) (*types.QueryAllBalancesResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryAllBalancesResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAllBalancesRequest, ...grpc.CallOption) (*types.QueryAllBalancesResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAllBalancesRequest, ...grpc.CallOption) *types.QueryAllBalancesResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryAllBalancesResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryAllBalancesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BankQueryClient_AllBalances_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AllBalances'
type BankQueryClient_AllBalances_Call struct {
	*mock.Call
}

// AllBalances is a helper method to define mock.On call
//   - ctx context.Context
//   - in *types.QueryAllBalancesRequest
//   - opts ...grpc.CallOption
func (_e *BankQueryClient_Expecter) AllBalances(ctx interface{}, in interface{}, opts ...interface{}) *BankQueryClient_AllBalances_Call {
	return &BankQueryClient_AllBalances_Call{Call: _e.mock.On("AllBalances",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *BankQueryClient_AllBalances_Call) Run(run func(ctx context.Context, in *types.QueryAllBalancesRequest, opts ...grpc.CallOption)) *BankQueryClient_AllBalances_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*types.QueryAllBalancesRequest), variadicArgs...)
	})
	return _c
}

func (_c *BankQueryClient_AllBalances_Call) Return(_a0 *types.QueryAllBalancesResponse, _a1 error) *BankQueryClient_AllBalances_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BankQueryClient_AllBalances_Call) RunAndReturn(run func(context.Context, *types.QueryAllBalancesRequest, ...grpc.CallOption) (*types.QueryAllBalancesResponse, error)) *BankQueryClient_AllBalances_Call {
	_c.Call.Return(run)
	return _c
}

// Balance provides a mock function with given fields: ctx, in, opts
func (_m *BankQueryClient) Balance(ctx context.Context, in *types.QueryBalanceRequest, opts ...grpc.CallOption) (*types.QueryBalanceResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryBalanceResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryBalanceRequest, ...grpc.CallOption) (*types.QueryBalanceResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryBalanceRequest, ...grpc.CallOption) *types.QueryBalanceResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryBalanceResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryBalanceRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BankQueryClient_Balance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Balance'
type BankQueryClient_Balance_Call struct {
	*mock.Call
}

// Balance is a helper method to define mock.On call
//   - ctx context.Context
//   - in *types.QueryBalanceRequest
//   - opts ...grpc.CallOption
func (_e *BankQueryClient_Expecter) Balance(ctx interface{}, in interface{}, opts ...interface{}) *BankQueryClient_Balance_Call {
	return &BankQueryClient_Balance_Call{Call: _e.mock.On("Balance",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *BankQueryClient_Balance_Call) Run(run func(ctx context.Context, in *types.QueryBalanceRequest, opts ...grpc.CallOption)) *BankQueryClient_Balance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*types.QueryBalanceRequest), variadicArgs...)
	})
	return _c
}

func (_c *BankQueryClient_Balance_Call) Return(_a0 *types.QueryBalanceResponse, _a1 error) *BankQueryClient_Balance_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BankQueryClient_Balance_Call) RunAndReturn(run func(context.Context, *types.QueryBalanceRequest, ...grpc.CallOption) (*types.QueryBalanceResponse, error)) *BankQueryClient_Balance_Call {
	_c.Call.Return(run)
	return _c
}

// DenomMetadata provides a mock function with given fields: ctx, in, opts
func (_m *BankQueryClient) DenomMetadata(ctx context.Context, in *types.QueryDenomMetadataRequest, opts ...grpc.CallOption) (*types.QueryDenomMetadataResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryDenomMetadataResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryDenomMetadataRequest, ...grpc.CallOption) (*types.QueryDenomMetadataResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryDenomMetadataRequest, ...grpc.CallOption) *types.QueryDenomMetadataResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryDenomMetadataResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryDenomMetadataRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BankQueryClient_DenomMetadata_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DenomMetadata'
type BankQueryClient_DenomMetadata_Call struct {
	*mock.Call
}

// DenomMetadata is a helper method to define mock.On call
//   - ctx context.Context
//   - in *types.QueryDenomMetadataRequest
//   - opts ...grpc.CallOption
func (_e *BankQueryClient_Expecter) DenomMetadata(ctx interface{}, in interface{}, opts ...interface{}) *BankQueryClient_DenomMetadata_Call {
	return &BankQueryClient_DenomMetadata_Call{Call: _e.mock.On("DenomMetadata",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *BankQueryClient_DenomMetadata_Call) Run(run func(ctx context.Context, in *types.QueryDenomMetadataRequest, opts ...grpc.CallOption)) *BankQueryClient_DenomMetadata_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*types.QueryDenomMetadataRequest), variadicArgs...)
	})
	return _c
}

func (_c *BankQueryClient_DenomMetadata_Call) Return(_a0 *types.QueryDenomMetadataResponse, _a1 error) *BankQueryClient_DenomMetadata_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BankQueryClient_DenomMetadata_Call) RunAndReturn(run func(context.Context, *types.QueryDenomMetadataRequest, ...grpc.CallOption) (*types.QueryDenomMetadataResponse, error)) *BankQueryClient_DenomMetadata_Call {
	_c.Call.Return(run)
	return _c
}

// DenomOwners provides a mock function with given fields: ctx, in, opts
func (_m *BankQueryClient) DenomOwners(ctx context.Context, in *types.QueryDenomOwnersRequest, opts ...grpc.CallOption) (*types.QueryDenomOwnersResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryDenomOwnersResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryDenomOwnersRequest, ...grpc.CallOption) (*types.QueryDenomOwnersResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryDenomOwnersRequest, ...grpc.CallOption) *types.QueryDenomOwnersResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryDenomOwnersResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryDenomOwnersRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BankQueryClient_DenomOwners_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DenomOwners'
type BankQueryClient_DenomOwners_Call struct {
	*mock.Call
}

// DenomOwners is a helper method to define mock.On call
//   - ctx context.Context
//   - in *types.QueryDenomOwnersRequest
//   - opts ...grpc.CallOption
func (_e *BankQueryClient_Expecter) DenomOwners(ctx interface{}, in interface{}, opts ...interface{}) *BankQueryClient_DenomOwners_Call {
	return &BankQueryClient_DenomOwners_Call{Call: _e.mock.On("DenomOwners",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *BankQueryClient_DenomOwners_Call) Run(run func(ctx context.Context, in *types.QueryDenomOwnersRequest, opts ...grpc.CallOption)) *BankQueryClient_DenomOwners_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*types.QueryDenomOwnersRequest), variadicArgs...)
	})
	return _c
}

func (_c *BankQueryClient_DenomOwners_Call) Return(_a0 *types.QueryDenomOwnersResponse, _a1 error) *BankQueryClient_DenomOwners_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BankQueryClient_DenomOwners_Call) RunAndReturn(run func(context.Context, *types.QueryDenomOwnersRequest, ...grpc.CallOption) (*types.QueryDenomOwnersResponse, error)) *BankQueryClient_DenomOwners_Call {
	_c.Call.Return(run)
	return _c
}

// DenomsMetadata provides a mock function with given fields: ctx, in, opts
func (_m *BankQueryClient) DenomsMetadata(ctx context.Context, in *types.QueryDenomsMetadataRequest, opts ...grpc.CallOption) (*types.QueryDenomsMetadataResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryDenomsMetadataResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryDenomsMetadataRequest, ...grpc.CallOption) (*types.QueryDenomsMetadataResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryDenomsMetadataRequest, ...grpc.CallOption) *types.QueryDenomsMetadataResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryDenomsMetadataResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryDenomsMetadataRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BankQueryClient_DenomsMetadata_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DenomsMetadata'
type BankQueryClient_DenomsMetadata_Call struct {
	*mock.Call
}

// DenomsMetadata is a helper method to define mock.On call
//   - ctx context.Context
//   - in *types.QueryDenomsMetadataRequest
//   - opts ...grpc.CallOption
func (_e *BankQueryClient_Expecter) DenomsMetadata(ctx interface{}, in interface{}, opts ...interface{}) *BankQueryClient_DenomsMetadata_Call {
	return &BankQueryClient_DenomsMetadata_Call{Call: _e.mock.On("DenomsMetadata",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *BankQueryClient_DenomsMetadata_Call) Run(run func(ctx context.Context, in *types.QueryDenomsMetadataRequest, opts ...grpc.CallOption)) *BankQueryClient_DenomsMetadata_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*types.QueryDenomsMetadataRequest), variadicArgs...)
	})
	return _c
}

func (_c *BankQueryClient_DenomsMetadata_Call) Return(_a0 *types.QueryDenomsMetadataResponse, _a1 error) *BankQueryClient_DenomsMetadata_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BankQueryClient_DenomsMetadata_Call) RunAndReturn(run func(context.Context, *types.QueryDenomsMetadataRequest, ...grpc.CallOption) (*types.QueryDenomsMetadataResponse, error)) *BankQueryClient_DenomsMetadata_Call {
	_c.Call.Return(run)
	return _c
}

// Params provides a mock function with given fields: ctx, in, opts
func (_m *BankQueryClient) Params(ctx context.Context, in *types.QueryParamsRequest, opts ...grpc.CallOption) (*types.QueryParamsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryParamsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryParamsRequest, ...grpc.CallOption) (*types.QueryParamsResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryParamsRequest, ...grpc.CallOption) *types.QueryParamsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryParamsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryParamsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BankQueryClient_Params_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Params'
type BankQueryClient_Params_Call struct {
	*mock.Call
}

// Params is a helper method to define mock.On call
//   - ctx context.Context
//   - in *types.QueryParamsRequest
//   - opts ...grpc.CallOption
func (_e *BankQueryClient_Expecter) Params(ctx interface{}, in interface{}, opts ...interface{}) *BankQueryClient_Params_Call {
	return &BankQueryClient_Params_Call{Call: _e.mock.On("Params",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *BankQueryClient_Params_Call) Run(run func(ctx context.Context, in *types.QueryParamsRequest, opts ...grpc.CallOption)) *BankQueryClient_Params_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*types.QueryParamsRequest), variadicArgs...)
	})
	return _c
}

func (_c *BankQueryClient_Params_Call) Return(_a0 *types.QueryParamsResponse, _a1 error) *BankQueryClient_Params_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BankQueryClient_Params_Call) RunAndReturn(run func(context.Context, *types.QueryParamsRequest, ...grpc.CallOption) (*types.QueryParamsResponse, error)) *BankQueryClient_Params_Call {
	_c.Call.Return(run)
	return _c
}

// SendEnabled provides a mock function with given fields: ctx, in, opts
func (_m *BankQueryClient) SendEnabled(ctx context.Context, in *types.QuerySendEnabledRequest, opts ...grpc.CallOption) (*types.QuerySendEnabledResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QuerySendEnabledResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QuerySendEnabledRequest, ...grpc.CallOption) (*types.QuerySendEnabledResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QuerySendEnabledRequest, ...grpc.CallOption) *types.QuerySendEnabledResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QuerySendEnabledResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QuerySendEnabledRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BankQueryClient_SendEnabled_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendEnabled'
type BankQueryClient_SendEnabled_Call struct {
	*mock.Call
}

// SendEnabled is a helper method to define mock.On call
//   - ctx context.Context
//   - in *types.QuerySendEnabledRequest
//   - opts ...grpc.CallOption
func (_e *BankQueryClient_Expecter) SendEnabled(ctx interface{}, in interface{}, opts ...interface{}) *BankQueryClient_SendEnabled_Call {
	return &BankQueryClient_SendEnabled_Call{Call: _e.mock.On("SendEnabled",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *BankQueryClient_SendEnabled_Call) Run(run func(ctx context.Context, in *types.QuerySendEnabledRequest, opts ...grpc.CallOption)) *BankQueryClient_SendEnabled_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*types.QuerySendEnabledRequest), variadicArgs...)
	})
	return _c
}

func (_c *BankQueryClient_SendEnabled_Call) Return(_a0 *types.QuerySendEnabledResponse, _a1 error) *BankQueryClient_SendEnabled_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BankQueryClient_SendEnabled_Call) RunAndReturn(run func(context.Context, *types.QuerySendEnabledRequest, ...grpc.CallOption) (*types.QuerySendEnabledResponse, error)) *BankQueryClient_SendEnabled_Call {
	_c.Call.Return(run)
	return _c
}

// SpendableBalanceByDenom provides a mock function with given fields: ctx, in, opts
func (_m *BankQueryClient) SpendableBalanceByDenom(ctx context.Context, in *types.QuerySpendableBalanceByDenomRequest, opts ...grpc.CallOption) (*types.QuerySpendableBalanceByDenomResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QuerySpendableBalanceByDenomResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QuerySpendableBalanceByDenomRequest, ...grpc.CallOption) (*types.QuerySpendableBalanceByDenomResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QuerySpendableBalanceByDenomRequest, ...grpc.CallOption) *types.QuerySpendableBalanceByDenomResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QuerySpendableBalanceByDenomResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QuerySpendableBalanceByDenomRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BankQueryClient_SpendableBalanceByDenom_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SpendableBalanceByDenom'
type BankQueryClient_SpendableBalanceByDenom_Call struct {
	*mock.Call
}

// SpendableBalanceByDenom is a helper method to define mock.On call
//   - ctx context.Context
//   - in *types.QuerySpendableBalanceByDenomRequest
//   - opts ...grpc.CallOption
func (_e *BankQueryClient_Expecter) SpendableBalanceByDenom(ctx interface{}, in interface{}, opts ...interface{}) *BankQueryClient_SpendableBalanceByDenom_Call {
	return &BankQueryClient_SpendableBalanceByDenom_Call{Call: _e.mock.On("SpendableBalanceByDenom",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *BankQueryClient_SpendableBalanceByDenom_Call) Run(run func(ctx context.Context, in *types.QuerySpendableBalanceByDenomRequest, opts ...grpc.CallOption)) *BankQueryClient_SpendableBalanceByDenom_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*types.QuerySpendableBalanceByDenomRequest), variadicArgs...)
	})
	return _c
}

func (_c *BankQueryClient_SpendableBalanceByDenom_Call) Return(_a0 *types.QuerySpendableBalanceByDenomResponse, _a1 error) *BankQueryClient_SpendableBalanceByDenom_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BankQueryClient_SpendableBalanceByDenom_Call) RunAndReturn(run func(context.Context, *types.QuerySpendableBalanceByDenomRequest, ...grpc.CallOption) (*types.QuerySpendableBalanceByDenomResponse, error)) *BankQueryClient_SpendableBalanceByDenom_Call {
	_c.Call.Return(run)
	return _c
}

// SpendableBalances provides a mock function with given fields: ctx, in, opts
func (_m *BankQueryClient) SpendableBalances(ctx context.Context, in *types.QuerySpendableBalancesRequest, opts ...grpc.CallOption) (*types.QuerySpendableBalancesResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QuerySpendableBalancesResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QuerySpendableBalancesRequest, ...grpc.CallOption) (*types.QuerySpendableBalancesResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QuerySpendableBalancesRequest, ...grpc.CallOption) *types.QuerySpendableBalancesResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QuerySpendableBalancesResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QuerySpendableBalancesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BankQueryClient_SpendableBalances_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SpendableBalances'
type BankQueryClient_SpendableBalances_Call struct {
	*mock.Call
}

// SpendableBalances is a helper method to define mock.On call
//   - ctx context.Context
//   - in *types.QuerySpendableBalancesRequest
//   - opts ...grpc.CallOption
func (_e *BankQueryClient_Expecter) SpendableBalances(ctx interface{}, in interface{}, opts ...interface{}) *BankQueryClient_SpendableBalances_Call {
	return &BankQueryClient_SpendableBalances_Call{Call: _e.mock.On("SpendableBalances",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *BankQueryClient_SpendableBalances_Call) Run(run func(ctx context.Context, in *types.QuerySpendableBalancesRequest, opts ...grpc.CallOption)) *BankQueryClient_SpendableBalances_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*types.QuerySpendableBalancesRequest), variadicArgs...)
	})
	return _c
}

func (_c *BankQueryClient_SpendableBalances_Call) Return(_a0 *types.QuerySpendableBalancesResponse, _a1 error) *BankQueryClient_SpendableBalances_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BankQueryClient_SpendableBalances_Call) RunAndReturn(run func(context.Context, *types.QuerySpendableBalancesRequest, ...grpc.CallOption) (*types.QuerySpendableBalancesResponse, error)) *BankQueryClient_SpendableBalances_Call {
	_c.Call.Return(run)
	return _c
}

// SupplyOf provides a mock function with given fields: ctx, in, opts
func (_m *BankQueryClient) SupplyOf(ctx context.Context, in *types.QuerySupplyOfRequest, opts ...grpc.CallOption) (*types.QuerySupplyOfResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QuerySupplyOfResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QuerySupplyOfRequest, ...grpc.CallOption) (*types.QuerySupplyOfResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QuerySupplyOfRequest, ...grpc.CallOption) *types.QuerySupplyOfResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QuerySupplyOfResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QuerySupplyOfRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BankQueryClient_SupplyOf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SupplyOf'
type BankQueryClient_SupplyOf_Call struct {
	*mock.Call
}

// SupplyOf is a helper method to define mock.On call
//   - ctx context.Context
//   - in *types.QuerySupplyOfRequest
//   - opts ...grpc.CallOption
func (_e *BankQueryClient_Expecter) SupplyOf(ctx interface{}, in interface{}, opts ...interface{}) *BankQueryClient_SupplyOf_Call {
	return &BankQueryClient_SupplyOf_Call{Call: _e.mock.On("SupplyOf",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *BankQueryClient_SupplyOf_Call) Run(run func(ctx context.Context, in *types.QuerySupplyOfRequest, opts ...grpc.CallOption)) *BankQueryClient_SupplyOf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*types.QuerySupplyOfRequest), variadicArgs...)
	})
	return _c
}

func (_c *BankQueryClient_SupplyOf_Call) Return(_a0 *types.QuerySupplyOfResponse, _a1 error) *BankQueryClient_SupplyOf_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BankQueryClient_SupplyOf_Call) RunAndReturn(run func(context.Context, *types.QuerySupplyOfRequest, ...grpc.CallOption) (*types.QuerySupplyOfResponse, error)) *BankQueryClient_SupplyOf_Call {
	_c.Call.Return(run)
	return _c
}

// TotalSupply provides a mock function with given fields: ctx, in, opts
func (_m *BankQueryClient) TotalSupply(ctx context.Context, in *types.QueryTotalSupplyRequest, opts ...grpc.CallOption) (*types.QueryTotalSupplyResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryTotalSupplyResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryTotalSupplyRequest, ...grpc.CallOption) (*types.QueryTotalSupplyResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryTotalSupplyRequest, ...grpc.CallOption) *types.QueryTotalSupplyResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryTotalSupplyResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryTotalSupplyRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BankQueryClient_TotalSupply_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TotalSupply'
type BankQueryClient_TotalSupply_Call struct {
	*mock.Call
}

// TotalSupply is a helper method to define mock.On call
//   - ctx context.Context
//   - in *types.QueryTotalSupplyRequest
//   - opts ...grpc.CallOption
func (_e *BankQueryClient_Expecter) TotalSupply(ctx interface{}, in interface{}, opts ...interface{}) *BankQueryClient_TotalSupply_Call {
	return &BankQueryClient_TotalSupply_Call{Call: _e.mock.On("TotalSupply",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *BankQueryClient_TotalSupply_Call) Run(run func(ctx context.Context, in *types.QueryTotalSupplyRequest, opts ...grpc.CallOption)) *BankQueryClient_TotalSupply_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*types.QueryTotalSupplyRequest), variadicArgs...)
	})
	return _c
}

func (_c *BankQueryClient_TotalSupply_Call) Return(_a0 *types.QueryTotalSupplyResponse, _a1 error) *BankQueryClient_TotalSupply_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BankQueryClient_TotalSupply_Call) RunAndReturn(run func(context.Context, *types.QueryTotalSupplyRequest, ...grpc.CallOption) (*types.QueryTotalSupplyResponse, error)) *BankQueryClient_TotalSupply_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewBankQueryClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewBankQueryClient creates a new instance of BankQueryClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBankQueryClient(t mockConstructorTestingTNewBankQueryClient) *BankQueryClient {
	mock := &BankQueryClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
