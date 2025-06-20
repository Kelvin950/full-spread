// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/kelvin950/spread/internals/core/domain"
	mock "github.com/stretchr/testify/mock"
)

// NewMockIS3 creates a new instance of MockIS3. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockIS3(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockIS3 {
	mock := &MockIS3{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// MockIS3 is an autogenerated mock type for the IS3 type
type MockIS3 struct {
	mock.Mock
}

type MockIS3_Expecter struct {
	mock *mock.Mock
}

func (_m *MockIS3) EXPECT() *MockIS3_Expecter {
	return &MockIS3_Expecter{mock: &_m.Mock}
}

// CompleteMultiPart provides a mock function for the type MockIS3
func (_mock *MockIS3) CompleteMultiPart(ctx context.Context, input domain.CompleteMultiPart) (*s3.CompleteMultipartUploadOutput, error) {
	ret := _mock.Called(ctx, input)

	if len(ret) == 0 {
		panic("no return value specified for CompleteMultiPart")
	}

	var r0 *s3.CompleteMultipartUploadOutput
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, domain.CompleteMultiPart) (*s3.CompleteMultipartUploadOutput, error)); ok {
		return returnFunc(ctx, input)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, domain.CompleteMultiPart) *s3.CompleteMultipartUploadOutput); ok {
		r0 = returnFunc(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.CompleteMultipartUploadOutput)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, domain.CompleteMultiPart) error); ok {
		r1 = returnFunc(ctx, input)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockIS3_CompleteMultiPart_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CompleteMultiPart'
type MockIS3_CompleteMultiPart_Call struct {
	*mock.Call
}

// CompleteMultiPart is a helper method to define mock.On call
//   - ctx context.Context
//   - input domain.CompleteMultiPart
func (_e *MockIS3_Expecter) CompleteMultiPart(ctx interface{}, input interface{}) *MockIS3_CompleteMultiPart_Call {
	return &MockIS3_CompleteMultiPart_Call{Call: _e.mock.On("CompleteMultiPart", ctx, input)}
}

func (_c *MockIS3_CompleteMultiPart_Call) Run(run func(ctx context.Context, input domain.CompleteMultiPart)) *MockIS3_CompleteMultiPart_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 context.Context
		if args[0] != nil {
			arg0 = args[0].(context.Context)
		}
		var arg1 domain.CompleteMultiPart
		if args[1] != nil {
			arg1 = args[1].(domain.CompleteMultiPart)
		}
		run(
			arg0,
			arg1,
		)
	})
	return _c
}

func (_c *MockIS3_CompleteMultiPart_Call) Return(completeMultipartUploadOutput *s3.CompleteMultipartUploadOutput, err error) *MockIS3_CompleteMultiPart_Call {
	_c.Call.Return(completeMultipartUploadOutput, err)
	return _c
}

func (_c *MockIS3_CompleteMultiPart_Call) RunAndReturn(run func(ctx context.Context, input domain.CompleteMultiPart) (*s3.CompleteMultipartUploadOutput, error)) *MockIS3_CompleteMultiPart_Call {
	_c.Call.Return(run)
	return _c
}

// CreateMultiPartUpload provides a mock function for the type MockIS3
func (_mock *MockIS3) CreateMultiPartUpload(ctx context.Context, input domain.CreateMultiPartUpload) (*s3.CreateMultipartUploadOutput, error) {
	ret := _mock.Called(ctx, input)

	if len(ret) == 0 {
		panic("no return value specified for CreateMultiPartUpload")
	}

	var r0 *s3.CreateMultipartUploadOutput
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, domain.CreateMultiPartUpload) (*s3.CreateMultipartUploadOutput, error)); ok {
		return returnFunc(ctx, input)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, domain.CreateMultiPartUpload) *s3.CreateMultipartUploadOutput); ok {
		r0 = returnFunc(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.CreateMultipartUploadOutput)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, domain.CreateMultiPartUpload) error); ok {
		r1 = returnFunc(ctx, input)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockIS3_CreateMultiPartUpload_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateMultiPartUpload'
type MockIS3_CreateMultiPartUpload_Call struct {
	*mock.Call
}

// CreateMultiPartUpload is a helper method to define mock.On call
//   - ctx context.Context
//   - input domain.CreateMultiPartUpload
func (_e *MockIS3_Expecter) CreateMultiPartUpload(ctx interface{}, input interface{}) *MockIS3_CreateMultiPartUpload_Call {
	return &MockIS3_CreateMultiPartUpload_Call{Call: _e.mock.On("CreateMultiPartUpload", ctx, input)}
}

func (_c *MockIS3_CreateMultiPartUpload_Call) Run(run func(ctx context.Context, input domain.CreateMultiPartUpload)) *MockIS3_CreateMultiPartUpload_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 context.Context
		if args[0] != nil {
			arg0 = args[0].(context.Context)
		}
		var arg1 domain.CreateMultiPartUpload
		if args[1] != nil {
			arg1 = args[1].(domain.CreateMultiPartUpload)
		}
		run(
			arg0,
			arg1,
		)
	})
	return _c
}

func (_c *MockIS3_CreateMultiPartUpload_Call) Return(createMultipartUploadOutput *s3.CreateMultipartUploadOutput, err error) *MockIS3_CreateMultiPartUpload_Call {
	_c.Call.Return(createMultipartUploadOutput, err)
	return _c
}

func (_c *MockIS3_CreateMultiPartUpload_Call) RunAndReturn(run func(ctx context.Context, input domain.CreateMultiPartUpload) (*s3.CreateMultipartUploadOutput, error)) *MockIS3_CreateMultiPartUpload_Call {
	_c.Call.Return(run)
	return _c
}

// CreatePresignMultiPart provides a mock function for the type MockIS3
func (_mock *MockIS3) CreatePresignMultiPart(ctx context.Context, input domain.UplaodMultiPart) (*v4.PresignedHTTPRequest, error) {
	ret := _mock.Called(ctx, input)

	if len(ret) == 0 {
		panic("no return value specified for CreatePresignMultiPart")
	}

	var r0 *v4.PresignedHTTPRequest
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, domain.UplaodMultiPart) (*v4.PresignedHTTPRequest, error)); ok {
		return returnFunc(ctx, input)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, domain.UplaodMultiPart) *v4.PresignedHTTPRequest); ok {
		r0 = returnFunc(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v4.PresignedHTTPRequest)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, domain.UplaodMultiPart) error); ok {
		r1 = returnFunc(ctx, input)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// MockIS3_CreatePresignMultiPart_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreatePresignMultiPart'
type MockIS3_CreatePresignMultiPart_Call struct {
	*mock.Call
}

// CreatePresignMultiPart is a helper method to define mock.On call
//   - ctx context.Context
//   - input domain.UplaodMultiPart
func (_e *MockIS3_Expecter) CreatePresignMultiPart(ctx interface{}, input interface{}) *MockIS3_CreatePresignMultiPart_Call {
	return &MockIS3_CreatePresignMultiPart_Call{Call: _e.mock.On("CreatePresignMultiPart", ctx, input)}
}

func (_c *MockIS3_CreatePresignMultiPart_Call) Run(run func(ctx context.Context, input domain.UplaodMultiPart)) *MockIS3_CreatePresignMultiPart_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 context.Context
		if args[0] != nil {
			arg0 = args[0].(context.Context)
		}
		var arg1 domain.UplaodMultiPart
		if args[1] != nil {
			arg1 = args[1].(domain.UplaodMultiPart)
		}
		run(
			arg0,
			arg1,
		)
	})
	return _c
}

func (_c *MockIS3_CreatePresignMultiPart_Call) Return(presignedHTTPRequest *v4.PresignedHTTPRequest, err error) *MockIS3_CreatePresignMultiPart_Call {
	_c.Call.Return(presignedHTTPRequest, err)
	return _c
}

func (_c *MockIS3_CreatePresignMultiPart_Call) RunAndReturn(run func(ctx context.Context, input domain.UplaodMultiPart) (*v4.PresignedHTTPRequest, error)) *MockIS3_CreatePresignMultiPart_Call {
	_c.Call.Return(run)
	return _c
}
