// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/kelvin950/spread/internals/core/domain"

	mock "github.com/stretchr/testify/mock"

	s3 "github.com/aws/aws-sdk-go-v2/service/s3"

	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// IS3 is an autogenerated mock type for the IS3 type
type IS3 struct {
	mock.Mock
}

// CompleteMultiPart provides a mock function with given fields: ctx, input
func (_m *IS3) CompleteMultiPart(ctx context.Context, input domain.CompleteMultiPart) (*s3.CompleteMultipartUploadOutput, error) {
	ret := _m.Called(ctx, input)

	if len(ret) == 0 {
		panic("no return value specified for CompleteMultiPart")
	}

	var r0 *s3.CompleteMultipartUploadOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.CompleteMultiPart) (*s3.CompleteMultipartUploadOutput, error)); ok {
		return rf(ctx, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.CompleteMultiPart) *s3.CompleteMultipartUploadOutput); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.CompleteMultipartUploadOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.CompleteMultiPart) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateMultiPartUpload provides a mock function with given fields: ctx, input
func (_m *IS3) CreateMultiPartUpload(ctx context.Context, input domain.CreateMultiPartUpload) (*s3.CreateMultipartUploadOutput, error) {
	ret := _m.Called(ctx, input)

	if len(ret) == 0 {
		panic("no return value specified for CreateMultiPartUpload")
	}

	var r0 *s3.CreateMultipartUploadOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.CreateMultiPartUpload) (*s3.CreateMultipartUploadOutput, error)); ok {
		return rf(ctx, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.CreateMultiPartUpload) *s3.CreateMultipartUploadOutput); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.CreateMultipartUploadOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.CreateMultiPartUpload) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePresignMultiPart provides a mock function with given fields: ctx, input
func (_m *IS3) CreatePresignMultiPart(ctx context.Context, input domain.UplaodMultiPart) (*v4.PresignedHTTPRequest, error) {
	ret := _m.Called(ctx, input)

	if len(ret) == 0 {
		panic("no return value specified for CreatePresignMultiPart")
	}

	var r0 *v4.PresignedHTTPRequest
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.UplaodMultiPart) (*v4.PresignedHTTPRequest, error)); ok {
		return rf(ctx, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.UplaodMultiPart) *v4.PresignedHTTPRequest); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v4.PresignedHTTPRequest)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.UplaodMultiPart) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIS3 creates a new instance of IS3. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIS3(t interface {
	mock.TestingT
	Cleanup(func())
}) *IS3 {
	mock := &IS3{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
