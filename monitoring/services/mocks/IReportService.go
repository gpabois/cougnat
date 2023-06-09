// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	models "github.com/gpabois/cougnat/reporting/models"
	mock "github.com/stretchr/testify/mock"

	result "github.com/gpabois/cougnat/core/result"
)

// IReportService is an autogenerated mock type for the IReportService type
type IReportService struct {
	mock.Mock
}

type IReportService_Expecter struct {
	mock *mock.Mock
}

func (_m *IReportService) EXPECT() *IReportService_Expecter {
	return &IReportService_Expecter{mock: &_m.Mock}
}

// HandleDeletedReport provides a mock function with given fields: deletedReport
func (_m *IReportService) HandleDeletedReport(deletedReport string) result.Result[bool] {
	ret := _m.Called(deletedReport)

	var r0 result.Result[bool]
	if rf, ok := ret.Get(0).(func(string) result.Result[bool]); ok {
		r0 = rf(deletedReport)
	} else {
		r0 = ret.Get(0).(result.Result[bool])
	}

	return r0
}

// IReportService_HandleDeletedReport_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HandleDeletedReport'
type IReportService_HandleDeletedReport_Call struct {
	*mock.Call
}

// HandleDeletedReport is a helper method to define mock.On call
//   - deletedReport string
func (_e *IReportService_Expecter) HandleDeletedReport(deletedReport interface{}) *IReportService_HandleDeletedReport_Call {
	return &IReportService_HandleDeletedReport_Call{Call: _e.mock.On("HandleDeletedReport", deletedReport)}
}

func (_c *IReportService_HandleDeletedReport_Call) Run(run func(deletedReport string)) *IReportService_HandleDeletedReport_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *IReportService_HandleDeletedReport_Call) Return(_a0 result.Result[bool]) *IReportService_HandleDeletedReport_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IReportService_HandleDeletedReport_Call) RunAndReturn(run func(string) result.Result[bool]) *IReportService_HandleDeletedReport_Call {
	_c.Call.Return(run)
	return _c
}

// HandleNewReport provides a mock function with given fields: newReport
func (_m *IReportService) HandleNewReport(newReport models.Report) result.Result[bool] {
	ret := _m.Called(newReport)

	var r0 result.Result[bool]
	if rf, ok := ret.Get(0).(func(models.Report) result.Result[bool]); ok {
		r0 = rf(newReport)
	} else {
		r0 = ret.Get(0).(result.Result[bool])
	}

	return r0
}

// IReportService_HandleNewReport_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HandleNewReport'
type IReportService_HandleNewReport_Call struct {
	*mock.Call
}

// HandleNewReport is a helper method to define mock.On call
//   - newReport models.Report
func (_e *IReportService_Expecter) HandleNewReport(newReport interface{}) *IReportService_HandleNewReport_Call {
	return &IReportService_HandleNewReport_Call{Call: _e.mock.On("HandleNewReport", newReport)}
}

func (_c *IReportService_HandleNewReport_Call) Run(run func(newReport models.Report)) *IReportService_HandleNewReport_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(models.Report))
	})
	return _c
}

func (_c *IReportService_HandleNewReport_Call) Return(_a0 result.Result[bool]) *IReportService_HandleNewReport_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IReportService_HandleNewReport_Call) RunAndReturn(run func(models.Report) result.Result[bool]) *IReportService_HandleNewReport_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewIReportService interface {
	mock.TestingT
	Cleanup(func())
}

// NewIReportService creates a new instance of IReportService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIReportService(t mockConstructorTestingTNewIReportService) *IReportService {
	mock := &IReportService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}