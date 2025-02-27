// Code generated by mockery v2.16.0. DO NOT EDIT.

package proxy

import (
	context "context"

	rootcoordpb "github.com/milvus-io/milvus/internal/proto/rootcoordpb"
	mock "github.com/stretchr/testify/mock"
)

// mockTimestampAllocator is an autogenerated mock type for the timestampAllocatorInterface type
type mockTimestampAllocator struct {
	mock.Mock
}

type mockTimestampAllocator_Expecter struct {
	mock *mock.Mock
}

func (_m *mockTimestampAllocator) EXPECT() *mockTimestampAllocator_Expecter {
	return &mockTimestampAllocator_Expecter{mock: &_m.Mock}
}

// AllocTimestamp provides a mock function with given fields: ctx, req
func (_m *mockTimestampAllocator) AllocTimestamp(ctx context.Context, req *rootcoordpb.AllocTimestampRequest) (*rootcoordpb.AllocTimestampResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *rootcoordpb.AllocTimestampResponse
	if rf, ok := ret.Get(0).(func(context.Context, *rootcoordpb.AllocTimestampRequest) *rootcoordpb.AllocTimestampResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rootcoordpb.AllocTimestampResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *rootcoordpb.AllocTimestampRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockTimestampAllocator_AllocTimestamp_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AllocTimestamp'
type mockTimestampAllocator_AllocTimestamp_Call struct {
	*mock.Call
}

// AllocTimestamp is a helper method to define mock.On call
//  - ctx context.Context
//  - req *rootcoordpb.AllocTimestampRequest
func (_e *mockTimestampAllocator_Expecter) AllocTimestamp(ctx interface{}, req interface{}) *mockTimestampAllocator_AllocTimestamp_Call {
	return &mockTimestampAllocator_AllocTimestamp_Call{Call: _e.mock.On("AllocTimestamp", ctx, req)}
}

func (_c *mockTimestampAllocator_AllocTimestamp_Call) Run(run func(ctx context.Context, req *rootcoordpb.AllocTimestampRequest)) *mockTimestampAllocator_AllocTimestamp_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*rootcoordpb.AllocTimestampRequest))
	})
	return _c
}

func (_c *mockTimestampAllocator_AllocTimestamp_Call) Return(_a0 *rootcoordpb.AllocTimestampResponse, _a1 error) *mockTimestampAllocator_AllocTimestamp_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTnewMockTimestampAllocator interface {
	mock.TestingT
	Cleanup(func())
}

// newMockTimestampAllocator creates a new instance of mockTimestampAllocator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockTimestampAllocator(t mockConstructorTestingTnewMockTimestampAllocator) *mockTimestampAllocator {
	mock := &mockTimestampAllocator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
