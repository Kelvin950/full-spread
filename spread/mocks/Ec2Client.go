// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	ec2 "github.com/aws/aws-sdk-go-v2/service/ec2"
	mock "github.com/stretchr/testify/mock"
)

// Ec2Client is an autogenerated mock type for the Ec2Client type
type Ec2Client struct {
	mock.Mock
}

// CreateInstance provides a mock function with given fields: timestarted, taskid
func (_m *Ec2Client) CreateInstance(timestarted string, taskid string) (*ec2.RunInstancesOutput, error) {
	ret := _m.Called(timestarted, taskid)

	if len(ret) == 0 {
		panic("no return value specified for CreateInstance")
	}

	var r0 *ec2.RunInstancesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*ec2.RunInstancesOutput, error)); ok {
		return rf(timestarted, taskid)
	}
	if rf, ok := ret.Get(0).(func(string, string) *ec2.RunInstancesOutput); ok {
		r0 = rf(timestarted, taskid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.RunInstancesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(timestarted, taskid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DestroyInstance provides a mock function with given fields: instanceId
func (_m *Ec2Client) DestroyInstance(instanceId string) error {
	ret := _m.Called(instanceId)

	if len(ret) == 0 {
		panic("no return value specified for DestroyInstance")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(instanceId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewEc2Client creates a new instance of Ec2Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEc2Client(t interface {
	mock.TestingT
	Cleanup(func())
}) *Ec2Client {
	mock := &Ec2Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
