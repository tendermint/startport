// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	types "github.com/tendermint/spn/x/campaign/types"
)

// CampaignClient is an autogenerated mock type for the CampaignClient type
type CampaignClient struct {
	mock.Mock
}

// Campaign provides a mock function with given fields: ctx, in, opts
func (_m *CampaignClient) Campaign(ctx context.Context, in *types.QueryGetCampaignRequest, opts ...grpc.CallOption) (*types.QueryGetCampaignResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryGetCampaignResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryGetCampaignRequest, ...grpc.CallOption) *types.QueryGetCampaignResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryGetCampaignResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryGetCampaignRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CampaignAll provides a mock function with given fields: ctx, in, opts
func (_m *CampaignClient) CampaignAll(ctx context.Context, in *types.QueryAllCampaignRequest, opts ...grpc.CallOption) (*types.QueryAllCampaignResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryAllCampaignResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAllCampaignRequest, ...grpc.CallOption) *types.QueryAllCampaignResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryAllCampaignResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryAllCampaignRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CampaignChains provides a mock function with given fields: ctx, in, opts
func (_m *CampaignClient) CampaignChains(ctx context.Context, in *types.QueryGetCampaignChainsRequest, opts ...grpc.CallOption) (*types.QueryGetCampaignChainsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryGetCampaignChainsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryGetCampaignChainsRequest, ...grpc.CallOption) *types.QueryGetCampaignChainsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryGetCampaignChainsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryGetCampaignChainsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MainnetAccount provides a mock function with given fields: ctx, in, opts
func (_m *CampaignClient) MainnetAccount(ctx context.Context, in *types.QueryGetMainnetAccountRequest, opts ...grpc.CallOption) (*types.QueryGetMainnetAccountResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryGetMainnetAccountResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryGetMainnetAccountRequest, ...grpc.CallOption) *types.QueryGetMainnetAccountResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryGetMainnetAccountResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryGetMainnetAccountRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MainnetAccountAll provides a mock function with given fields: ctx, in, opts
func (_m *CampaignClient) MainnetAccountAll(ctx context.Context, in *types.QueryAllMainnetAccountRequest, opts ...grpc.CallOption) (*types.QueryAllMainnetAccountResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryAllMainnetAccountResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAllMainnetAccountRequest, ...grpc.CallOption) *types.QueryAllMainnetAccountResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryAllMainnetAccountResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryAllMainnetAccountRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MainnetAccountBalance provides a mock function with given fields: ctx, in, opts
func (_m *CampaignClient) MainnetAccountBalance(ctx context.Context, in *types.QueryGetMainnetAccountBalanceRequest, opts ...grpc.CallOption) (*types.QueryGetMainnetAccountBalanceResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryGetMainnetAccountBalanceResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryGetMainnetAccountBalanceRequest, ...grpc.CallOption) *types.QueryGetMainnetAccountBalanceResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryGetMainnetAccountBalanceResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryGetMainnetAccountBalanceRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MainnetAccountBalanceAll provides a mock function with given fields: ctx, in, opts
func (_m *CampaignClient) MainnetAccountBalanceAll(ctx context.Context, in *types.QueryAllMainnetAccountBalanceRequest, opts ...grpc.CallOption) (*types.QueryAllMainnetAccountBalanceResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryAllMainnetAccountBalanceResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAllMainnetAccountBalanceRequest, ...grpc.CallOption) *types.QueryAllMainnetAccountBalanceResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryAllMainnetAccountBalanceResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryAllMainnetAccountBalanceRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Params provides a mock function with given fields: ctx, in, opts
func (_m *CampaignClient) Params(ctx context.Context, in *types.QueryParamsRequest, opts ...grpc.CallOption) (*types.QueryParamsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryParamsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryParamsRequest, ...grpc.CallOption) *types.QueryParamsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryParamsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryParamsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SpecialAllocationsBalance provides a mock function with given fields: ctx, in, opts
func (_m *CampaignClient) SpecialAllocationsBalance(ctx context.Context, in *types.QuerySpecialAllocationsBalanceRequest, opts ...grpc.CallOption) (*types.QuerySpecialAllocationsBalanceResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QuerySpecialAllocationsBalanceResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QuerySpecialAllocationsBalanceRequest, ...grpc.CallOption) *types.QuerySpecialAllocationsBalanceResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QuerySpecialAllocationsBalanceResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QuerySpecialAllocationsBalanceRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TotalShares provides a mock function with given fields: ctx, in, opts
func (_m *CampaignClient) TotalShares(ctx context.Context, in *types.QueryTotalSharesRequest, opts ...grpc.CallOption) (*types.QueryTotalSharesResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryTotalSharesResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryTotalSharesRequest, ...grpc.CallOption) *types.QueryTotalSharesResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryTotalSharesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryTotalSharesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCampaignClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewCampaignClient creates a new instance of CampaignClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCampaignClient(t mockConstructorTestingTNewCampaignClient) *CampaignClient {
	mock := &CampaignClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
