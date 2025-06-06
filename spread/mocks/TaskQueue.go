// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// TaskQueue is an autogenerated mock type for the TaskQueue type
type TaskQueue struct {
	mock.Mock
}

// DistributeTask provides a mock function with given fields: taskName, priority, taskpayload
func (_m *TaskQueue) DistributeTask(taskName string, priority string, taskpayload interface{}) error {
	ret := _m.Called(taskName, priority, taskpayload)

	if len(ret) == 0 {
		panic("no return value specified for DistributeTask")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, interface{}) error); ok {
		r0 = rf(taskName, priority, taskpayload)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewTaskQueue creates a new instance of TaskQueue. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTaskQueue(t interface {
	mock.TestingT
	Cleanup(func())
}) *TaskQueue {
	mock := &TaskQueue{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
