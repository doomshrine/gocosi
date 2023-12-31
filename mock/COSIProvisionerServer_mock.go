// Code generated by mockery v2.32.0. DO NOT EDIT.

package mock

import (
	context "context"

	cosi "sigs.k8s.io/container-object-storage-interface-spec"

	mock "github.com/stretchr/testify/mock"
)

// MockCOSIProvisionerServer is an autogenerated mock type for the COSIProvisionerServer type
type MockCOSIProvisionerServer struct {
	mock.Mock
}

type MockCOSIProvisionerServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCOSIProvisionerServer) EXPECT() *MockCOSIProvisionerServer_Expecter {
	return &MockCOSIProvisionerServer_Expecter{mock: &_m.Mock}
}

// DriverCreateBucket provides a mock function with given fields: _a0, _a1
func (_m *MockCOSIProvisionerServer) DriverCreateBucket(_a0 context.Context, _a1 *cosi.DriverCreateBucketRequest) (*cosi.DriverCreateBucketResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *cosi.DriverCreateBucketResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cosi.DriverCreateBucketRequest) (*cosi.DriverCreateBucketResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cosi.DriverCreateBucketRequest) *cosi.DriverCreateBucketResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cosi.DriverCreateBucketResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cosi.DriverCreateBucketRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCOSIProvisionerServer_DriverCreateBucket_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DriverCreateBucket'
type MockCOSIProvisionerServer_DriverCreateBucket_Call struct {
	*mock.Call
}

// DriverCreateBucket is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *cosi.DriverCreateBucketRequest
func (_e *MockCOSIProvisionerServer_Expecter) DriverCreateBucket(_a0 interface{}, _a1 interface{}) *MockCOSIProvisionerServer_DriverCreateBucket_Call {
	return &MockCOSIProvisionerServer_DriverCreateBucket_Call{Call: _e.mock.On("DriverCreateBucket", _a0, _a1)}
}

func (_c *MockCOSIProvisionerServer_DriverCreateBucket_Call) Run(run func(_a0 context.Context, _a1 *cosi.DriverCreateBucketRequest)) *MockCOSIProvisionerServer_DriverCreateBucket_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*cosi.DriverCreateBucketRequest))
	})
	return _c
}

func (_c *MockCOSIProvisionerServer_DriverCreateBucket_Call) Return(_a0 *cosi.DriverCreateBucketResponse, _a1 error) *MockCOSIProvisionerServer_DriverCreateBucket_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCOSIProvisionerServer_DriverCreateBucket_Call) RunAndReturn(run func(context.Context, *cosi.DriverCreateBucketRequest) (*cosi.DriverCreateBucketResponse, error)) *MockCOSIProvisionerServer_DriverCreateBucket_Call {
	_c.Call.Return(run)
	return _c
}

// DriverDeleteBucket provides a mock function with given fields: _a0, _a1
func (_m *MockCOSIProvisionerServer) DriverDeleteBucket(_a0 context.Context, _a1 *cosi.DriverDeleteBucketRequest) (*cosi.DriverDeleteBucketResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *cosi.DriverDeleteBucketResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cosi.DriverDeleteBucketRequest) (*cosi.DriverDeleteBucketResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cosi.DriverDeleteBucketRequest) *cosi.DriverDeleteBucketResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cosi.DriverDeleteBucketResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cosi.DriverDeleteBucketRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCOSIProvisionerServer_DriverDeleteBucket_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DriverDeleteBucket'
type MockCOSIProvisionerServer_DriverDeleteBucket_Call struct {
	*mock.Call
}

// DriverDeleteBucket is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *cosi.DriverDeleteBucketRequest
func (_e *MockCOSIProvisionerServer_Expecter) DriverDeleteBucket(_a0 interface{}, _a1 interface{}) *MockCOSIProvisionerServer_DriverDeleteBucket_Call {
	return &MockCOSIProvisionerServer_DriverDeleteBucket_Call{Call: _e.mock.On("DriverDeleteBucket", _a0, _a1)}
}

func (_c *MockCOSIProvisionerServer_DriverDeleteBucket_Call) Run(run func(_a0 context.Context, _a1 *cosi.DriverDeleteBucketRequest)) *MockCOSIProvisionerServer_DriverDeleteBucket_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*cosi.DriverDeleteBucketRequest))
	})
	return _c
}

func (_c *MockCOSIProvisionerServer_DriverDeleteBucket_Call) Return(_a0 *cosi.DriverDeleteBucketResponse, _a1 error) *MockCOSIProvisionerServer_DriverDeleteBucket_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCOSIProvisionerServer_DriverDeleteBucket_Call) RunAndReturn(run func(context.Context, *cosi.DriverDeleteBucketRequest) (*cosi.DriverDeleteBucketResponse, error)) *MockCOSIProvisionerServer_DriverDeleteBucket_Call {
	_c.Call.Return(run)
	return _c
}

// DriverGrantBucketAccess provides a mock function with given fields: _a0, _a1
func (_m *MockCOSIProvisionerServer) DriverGrantBucketAccess(_a0 context.Context, _a1 *cosi.DriverGrantBucketAccessRequest) (*cosi.DriverGrantBucketAccessResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *cosi.DriverGrantBucketAccessResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cosi.DriverGrantBucketAccessRequest) (*cosi.DriverGrantBucketAccessResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cosi.DriverGrantBucketAccessRequest) *cosi.DriverGrantBucketAccessResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cosi.DriverGrantBucketAccessResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cosi.DriverGrantBucketAccessRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCOSIProvisionerServer_DriverGrantBucketAccess_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DriverGrantBucketAccess'
type MockCOSIProvisionerServer_DriverGrantBucketAccess_Call struct {
	*mock.Call
}

// DriverGrantBucketAccess is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *cosi.DriverGrantBucketAccessRequest
func (_e *MockCOSIProvisionerServer_Expecter) DriverGrantBucketAccess(_a0 interface{}, _a1 interface{}) *MockCOSIProvisionerServer_DriverGrantBucketAccess_Call {
	return &MockCOSIProvisionerServer_DriverGrantBucketAccess_Call{Call: _e.mock.On("DriverGrantBucketAccess", _a0, _a1)}
}

func (_c *MockCOSIProvisionerServer_DriverGrantBucketAccess_Call) Run(run func(_a0 context.Context, _a1 *cosi.DriverGrantBucketAccessRequest)) *MockCOSIProvisionerServer_DriverGrantBucketAccess_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*cosi.DriverGrantBucketAccessRequest))
	})
	return _c
}

func (_c *MockCOSIProvisionerServer_DriverGrantBucketAccess_Call) Return(_a0 *cosi.DriverGrantBucketAccessResponse, _a1 error) *MockCOSIProvisionerServer_DriverGrantBucketAccess_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCOSIProvisionerServer_DriverGrantBucketAccess_Call) RunAndReturn(run func(context.Context, *cosi.DriverGrantBucketAccessRequest) (*cosi.DriverGrantBucketAccessResponse, error)) *MockCOSIProvisionerServer_DriverGrantBucketAccess_Call {
	_c.Call.Return(run)
	return _c
}

// DriverRevokeBucketAccess provides a mock function with given fields: _a0, _a1
func (_m *MockCOSIProvisionerServer) DriverRevokeBucketAccess(_a0 context.Context, _a1 *cosi.DriverRevokeBucketAccessRequest) (*cosi.DriverRevokeBucketAccessResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *cosi.DriverRevokeBucketAccessResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *cosi.DriverRevokeBucketAccessRequest) (*cosi.DriverRevokeBucketAccessResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cosi.DriverRevokeBucketAccessRequest) *cosi.DriverRevokeBucketAccessResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cosi.DriverRevokeBucketAccessResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cosi.DriverRevokeBucketAccessRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCOSIProvisionerServer_DriverRevokeBucketAccess_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DriverRevokeBucketAccess'
type MockCOSIProvisionerServer_DriverRevokeBucketAccess_Call struct {
	*mock.Call
}

// DriverRevokeBucketAccess is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *cosi.DriverRevokeBucketAccessRequest
func (_e *MockCOSIProvisionerServer_Expecter) DriverRevokeBucketAccess(_a0 interface{}, _a1 interface{}) *MockCOSIProvisionerServer_DriverRevokeBucketAccess_Call {
	return &MockCOSIProvisionerServer_DriverRevokeBucketAccess_Call{Call: _e.mock.On("DriverRevokeBucketAccess", _a0, _a1)}
}

func (_c *MockCOSIProvisionerServer_DriverRevokeBucketAccess_Call) Run(run func(_a0 context.Context, _a1 *cosi.DriverRevokeBucketAccessRequest)) *MockCOSIProvisionerServer_DriverRevokeBucketAccess_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*cosi.DriverRevokeBucketAccessRequest))
	})
	return _c
}

func (_c *MockCOSIProvisionerServer_DriverRevokeBucketAccess_Call) Return(_a0 *cosi.DriverRevokeBucketAccessResponse, _a1 error) *MockCOSIProvisionerServer_DriverRevokeBucketAccess_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCOSIProvisionerServer_DriverRevokeBucketAccess_Call) RunAndReturn(run func(context.Context, *cosi.DriverRevokeBucketAccessRequest) (*cosi.DriverRevokeBucketAccessResponse, error)) *MockCOSIProvisionerServer_DriverRevokeBucketAccess_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCOSIProvisionerServer creates a new instance of MockCOSIProvisionerServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCOSIProvisionerServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCOSIProvisionerServer {
	mock := &MockCOSIProvisionerServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
