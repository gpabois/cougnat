// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	models "github.com/gpabois/cougnat/reporting/models"
	mock "github.com/stretchr/testify/mock"

	result "github.com/gpabois/cougnat/core/result"
)

// IReportEventEmitter is an autogenerated mock type for the IReportEventEmitter type
type IReportEventEmitter struct {
	mock.Mock
}

type IReportEventEmitter_Expecter struct {
	mock *mock.Mock
}

func (_m *IReportEventEmitter) EXPECT() *IReportEventEmitter_Expecter {
	return &IReportEventEmitter_Expecter{mock: &_m.Mock}
}

// OnDeletedReport provides a mock function with given fields: report
func (_m *IReportEventEmitter) OnDeletedReport(report string) result.Result[bool] {
	ret := _m.Called(report)

	var r0 result.Result[bool]
	if rf, ok := ret.Get(0).(func(string) result.Result[bool]); ok {
		r0 = rf(report)
	} else {
		r0 = ret.Get(0).(result.Result[bool])
	}

	return r0
}

// IReportEventEmitter_OnDeletedReport_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OnDeletedReport'
type IReportEventEmitter_OnDeletedReport_Call struct {
	*mock.Call
}

// OnDeletedReport is a helper method to define mock.On call
//   - report string
func (_e *IReportEventEmitter_Expecter) OnDeletedReport(report interface{}) *IReportEventEmitter_OnDeletedReport_Call {
	return &IReportEventEmitter_OnDeletedReport_Call{Call: _e.mock.On("OnDeletedReport", report)}
}

func (_c *IReportEventEmitter_OnDeletedReport_Call) Run(run func(report string)) *IReportEventEmitter_OnDeletedReport_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *IReportEventEmitter_OnDeletedReport_Call) Return(_a0 result.Result[bool]) *IReportEventEmitter_OnDeletedReport_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IReportEventEmitter_OnDeletedReport_Call) RunAndReturn(run func(string) result.Result[bool]) *IReportEventEmitter_OnDeletedReport_Call {
	_c.Call.Return(run)
	return _c
}

// OnNewReport provides a mock function with given fields: report
func (_m *IReportEventEmitter) OnNewReport(report models.Report) result.Result[bool] {
	ret := _m.Called(report)

	var r0 result.Result[bool]
	if rf, ok := ret.Get(0).(func(models.Report) result.Result[bool]); ok {
		r0 = rf(report)
	} else {
		r0 = ret.Get(0).(result.Result[bool])
	}

	return r0
}

// IReportEventEmitter_OnNewReport_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OnNewReport'
type IReportEventEmitter_OnNewReport_Call struct {
	*mock.Call
}

// OnNewReport is a helper method to define mock.On call
//   - report models.Report
func (_e *IReportEventEmitter_Expecter) OnNewReport(report interface{}) *IReportEventEmitter_OnNewReport_Call {
	return &IReportEventEmitter_OnNewReport_Call{Call: _e.mock.On("OnNewReport", report)}
}

func (_c *IReportEventEmitter_OnNewReport_Call) Run(run func(report models.Report)) *IReportEventEmitter_OnNewReport_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(models.Report))
	})
	return _c
}

func (_c *IReportEventEmitter_OnNewReport_Call) Return(_a0 result.Result[bool]) *IReportEventEmitter_OnNewReport_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IReportEventEmitter_OnNewReport_Call) RunAndReturn(run func(models.Report) result.Result[bool]) *IReportEventEmitter_OnNewReport_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewIReportEventEmitter interface {
	mock.TestingT
	Cleanup(func())
}

// NewIReportEventEmitter creates a new instance of IReportEventEmitter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIReportEventEmitter(t mockConstructorTestingTNewIReportEventEmitter) *IReportEventEmitter {
	mock := &IReportEventEmitter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
