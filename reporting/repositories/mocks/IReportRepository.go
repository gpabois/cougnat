// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	option "github.com/gpabois/gostd/option"
	mock "github.com/stretchr/testify/mock"

	reporting_models "github.com/gpabois/cougnat/reporting/models"

	reporting_repositories "github.com/gpabois/cougnat/reporting/repositories"

	result "github.com/gpabois/gostd/result"

	transaction "github.com/gpabois/cougnat/core/transaction"
)

// IReportRepository is an autogenerated mock type for the IReportRepository type
type IReportRepository struct {
	mock.Mock
}

type IReportRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *IReportRepository) EXPECT() *IReportRepository_Expecter {
	return &IReportRepository_Expecter{mock: &_m.Mock}
}

// Begin provides a mock function with given fields: root
func (_m *IReportRepository) Begin(root option.Option[transaction.Transaction]) result.Result[reporting_repositories.IReportRepositoryTx] {
	ret := _m.Called(root)

	var r0 result.Result[reporting_repositories.IReportRepositoryTx]
	if rf, ok := ret.Get(0).(func(option.Option[transaction.Transaction]) result.Result[reporting_repositories.IReportRepositoryTx]); ok {
		r0 = rf(root)
	} else {
		r0 = ret.Get(0).(result.Result[reporting_repositories.IReportRepositoryTx])
	}

	return r0
}

// IReportRepository_Begin_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Begin'
type IReportRepository_Begin_Call struct {
	*mock.Call
}

// Begin is a helper method to define mock.On call
//   - root option.Option[transaction.Transaction]
func (_e *IReportRepository_Expecter) Begin(root interface{}) *IReportRepository_Begin_Call {
	return &IReportRepository_Begin_Call{Call: _e.mock.On("Begin", root)}
}

func (_c *IReportRepository_Begin_Call) Run(run func(root option.Option[transaction.Transaction])) *IReportRepository_Begin_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(option.Option[transaction.Transaction]))
	})
	return _c
}

func (_c *IReportRepository_Begin_Call) Return(_a0 result.Result[reporting_repositories.IReportRepositoryTx]) *IReportRepository_Begin_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IReportRepository_Begin_Call) RunAndReturn(run func(option.Option[transaction.Transaction]) result.Result[reporting_repositories.IReportRepositoryTx]) *IReportRepository_Begin_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: report
func (_m *IReportRepository) Create(report reporting_models.NewReport) result.Result[int] {
	ret := _m.Called(report)

	var r0 result.Result[int]
	if rf, ok := ret.Get(0).(func(reporting_models.NewReport) result.Result[int]); ok {
		r0 = rf(report)
	} else {
		r0 = ret.Get(0).(result.Result[int])
	}

	return r0
}

// IReportRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type IReportRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - report reporting_models.NewReport
func (_e *IReportRepository_Expecter) Create(report interface{}) *IReportRepository_Create_Call {
	return &IReportRepository_Create_Call{Call: _e.mock.On("Create", report)}
}

func (_c *IReportRepository_Create_Call) Run(run func(report reporting_models.NewReport)) *IReportRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(reporting_models.NewReport))
	})
	return _c
}

func (_c *IReportRepository_Create_Call) Return(_a0 result.Result[int]) *IReportRepository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IReportRepository_Create_Call) RunAndReturn(run func(reporting_models.NewReport) result.Result[int]) *IReportRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: reportID
func (_m *IReportRepository) Delete(reportID int) result.Result[bool] {
	ret := _m.Called(reportID)

	var r0 result.Result[bool]
	if rf, ok := ret.Get(0).(func(int) result.Result[bool]); ok {
		r0 = rf(reportID)
	} else {
		r0 = ret.Get(0).(result.Result[bool])
	}

	return r0
}

// IReportRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type IReportRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - reportID int
func (_e *IReportRepository_Expecter) Delete(reportID interface{}) *IReportRepository_Delete_Call {
	return &IReportRepository_Delete_Call{Call: _e.mock.On("Delete", reportID)}
}

func (_c *IReportRepository_Delete_Call) Run(run func(reportID int)) *IReportRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *IReportRepository_Delete_Call) Return(_a0 result.Result[bool]) *IReportRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IReportRepository_Delete_Call) RunAndReturn(run func(int) result.Result[bool]) *IReportRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// GetById provides a mock function with given fields: id
func (_m *IReportRepository) GetById(id int) result.Result[option.Option[reporting_models.Report]] {
	ret := _m.Called(id)

	var r0 result.Result[option.Option[reporting_models.Report]]
	if rf, ok := ret.Get(0).(func(int) result.Result[option.Option[reporting_models.Report]]); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(result.Result[option.Option[reporting_models.Report]])
	}

	return r0
}

// IReportRepository_GetById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetById'
type IReportRepository_GetById_Call struct {
	*mock.Call
}

// GetById is a helper method to define mock.On call
//   - id int
func (_e *IReportRepository_Expecter) GetById(id interface{}) *IReportRepository_GetById_Call {
	return &IReportRepository_GetById_Call{Call: _e.mock.On("GetById", id)}
}

func (_c *IReportRepository_GetById_Call) Run(run func(id int)) *IReportRepository_GetById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *IReportRepository_GetById_Call) Return(_a0 result.Result[option.Option[reporting_models.Report]]) *IReportRepository_GetById_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IReportRepository_GetById_Call) RunAndReturn(run func(int) result.Result[option.Option[reporting_models.Report]]) *IReportRepository_GetById_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewIReportRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewIReportRepository creates a new instance of IReportRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIReportRepository(t mockConstructorTestingTNewIReportRepository) *IReportRepository {
	mock := &IReportRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
