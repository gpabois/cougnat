// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	events "github.com/gpabois/cougnat/reporting/events"
	mock "github.com/stretchr/testify/mock"

	result "github.com/gpabois/cougnat/core/result"
)

// IReportEventReceiver is an autogenerated mock type for the IReportEventReceiver type
type IReportEventReceiver struct {
	mock.Mock
}

type IReportEventReceiver_Expecter struct {
	mock *mock.Mock
}

func (_m *IReportEventReceiver) EXPECT() *IReportEventReceiver_Expecter {
	return &IReportEventReceiver_Expecter{mock: &_m.Mock}
}

// OnDeletedReport provides a mock function with given fields: queueName
func (_m *IReportEventReceiver) OnDeletedReport(queueName string) result.Result[chan events.DeletedReportEvent] {
	ret := _m.Called(queueName)

	var r0 result.Result[chan events.DeletedReportEvent]
	if rf, ok := ret.Get(0).(func(string) result.Result[chan events.DeletedReportEvent]); ok {
		r0 = rf(queueName)
	} else {
		r0 = ret.Get(0).(result.Result[chan events.DeletedReportEvent])
	}

	return r0
}

// IReportEventReceiver_OnDeletedReport_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OnDeletedReport'
type IReportEventReceiver_OnDeletedReport_Call struct {
	*mock.Call
}

// OnDeletedReport is a helper method to define mock.On call
//   - queueName string
func (_e *IReportEventReceiver_Expecter) OnDeletedReport(queueName interface{}) *IReportEventReceiver_OnDeletedReport_Call {
	return &IReportEventReceiver_OnDeletedReport_Call{Call: _e.mock.On("OnDeletedReport", queueName)}
}

func (_c *IReportEventReceiver_OnDeletedReport_Call) Run(run func(queueName string)) *IReportEventReceiver_OnDeletedReport_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *IReportEventReceiver_OnDeletedReport_Call) Return(_a0 result.Result[chan events.DeletedReportEvent]) *IReportEventReceiver_OnDeletedReport_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IReportEventReceiver_OnDeletedReport_Call) RunAndReturn(run func(string) result.Result[chan events.DeletedReportEvent]) *IReportEventReceiver_OnDeletedReport_Call {
	_c.Call.Return(run)
	return _c
}

// OnNewReport provides a mock function with given fields: queueName
func (_m *IReportEventReceiver) OnNewReport(queueName string) result.Result[chan events.NewReportEvent] {
	ret := _m.Called(queueName)

	var r0 result.Result[chan events.NewReportEvent]
	if rf, ok := ret.Get(0).(func(string) result.Result[chan events.NewReportEvent]); ok {
		r0 = rf(queueName)
	} else {
		r0 = ret.Get(0).(result.Result[chan events.NewReportEvent])
	}

	return r0
}

// IReportEventReceiver_OnNewReport_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OnNewReport'
type IReportEventReceiver_OnNewReport_Call struct {
	*mock.Call
}

// OnNewReport is a helper method to define mock.On call
//   - queueName string
func (_e *IReportEventReceiver_Expecter) OnNewReport(queueName interface{}) *IReportEventReceiver_OnNewReport_Call {
	return &IReportEventReceiver_OnNewReport_Call{Call: _e.mock.On("OnNewReport", queueName)}
}

func (_c *IReportEventReceiver_OnNewReport_Call) Run(run func(queueName string)) *IReportEventReceiver_OnNewReport_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *IReportEventReceiver_OnNewReport_Call) Return(_a0 result.Result[chan events.NewReportEvent]) *IReportEventReceiver_OnNewReport_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IReportEventReceiver_OnNewReport_Call) RunAndReturn(run func(string) result.Result[chan events.NewReportEvent]) *IReportEventReceiver_OnNewReport_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewIReportEventReceiver interface {
	mock.TestingT
	Cleanup(func())
}

// NewIReportEventReceiver creates a new instance of IReportEventReceiver. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIReportEventReceiver(t mockConstructorTestingTNewIReportEventReceiver) *IReportEventReceiver {
	mock := &IReportEventReceiver{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}