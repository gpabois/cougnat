// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	models "github.com/gpabois/cougnat/auth/models"
	mock "github.com/stretchr/testify/mock"

	option "github.com/gpabois/cougnat/core/option"

	result "github.com/gpabois/cougnat/core/result"
)

// IAuthorizationService is an autogenerated mock type for the IAuthorizationService type
type IAuthorizationService struct {
	mock.Mock
}

type IAuthorizationService_Expecter struct {
	mock *mock.Mock
}

func (_m *IAuthorizationService) EXPECT() *IAuthorizationService_Expecter {
	return &IAuthorizationService_Expecter{mock: &_m.Mock}
}

// AddRoleTo provides a mock function with given fields: subject, roleID
func (_m *IAuthorizationService) AddRoleTo(subject models.ActorID, roleID string) result.Result[bool] {
	ret := _m.Called(subject, roleID)

	var r0 result.Result[bool]
	if rf, ok := ret.Get(0).(func(models.ActorID, string) result.Result[bool]); ok {
		r0 = rf(subject, roleID)
	} else {
		r0 = ret.Get(0).(result.Result[bool])
	}

	return r0
}

// IAuthorizationService_AddRoleTo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddRoleTo'
type IAuthorizationService_AddRoleTo_Call struct {
	*mock.Call
}

// AddRoleTo is a helper method to define mock.On call
//   - subject models.ActorID
//   - roleID string
func (_e *IAuthorizationService_Expecter) AddRoleTo(subject interface{}, roleID interface{}) *IAuthorizationService_AddRoleTo_Call {
	return &IAuthorizationService_AddRoleTo_Call{Call: _e.mock.On("AddRoleTo", subject, roleID)}
}

func (_c *IAuthorizationService_AddRoleTo_Call) Run(run func(subject models.ActorID, roleID string)) *IAuthorizationService_AddRoleTo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(models.ActorID), args[1].(string))
	})
	return _c
}

func (_c *IAuthorizationService_AddRoleTo_Call) Return(_a0 result.Result[bool]) *IAuthorizationService_AddRoleTo_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IAuthorizationService_AddRoleTo_Call) RunAndReturn(run func(models.ActorID, string) result.Result[bool]) *IAuthorizationService_AddRoleTo_Call {
	_c.Call.Return(run)
	return _c
}

// AssignPermission provides a mock function with given fields: role, perm
func (_m *IAuthorizationService) AssignPermission(role string, perm string) result.Result[bool] {
	ret := _m.Called(role, perm)

	var r0 result.Result[bool]
	if rf, ok := ret.Get(0).(func(string, string) result.Result[bool]); ok {
		r0 = rf(role, perm)
	} else {
		r0 = ret.Get(0).(result.Result[bool])
	}

	return r0
}

// IAuthorizationService_AssignPermission_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AssignPermission'
type IAuthorizationService_AssignPermission_Call struct {
	*mock.Call
}

// AssignPermission is a helper method to define mock.On call
//   - role string
//   - perm string
func (_e *IAuthorizationService_Expecter) AssignPermission(role interface{}, perm interface{}) *IAuthorizationService_AssignPermission_Call {
	return &IAuthorizationService_AssignPermission_Call{Call: _e.mock.On("AssignPermission", role, perm)}
}

func (_c *IAuthorizationService_AssignPermission_Call) Run(run func(role string, perm string)) *IAuthorizationService_AssignPermission_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *IAuthorizationService_AssignPermission_Call) Return(_a0 result.Result[bool]) *IAuthorizationService_AssignPermission_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IAuthorizationService_AssignPermission_Call) RunAndReturn(run func(string, string) result.Result[bool]) *IAuthorizationService_AssignPermission_Call {
	_c.Call.Return(run)
	return _c
}

// AssignPermissions provides a mock function with given fields: role, perms, object
func (_m *IAuthorizationService) AssignPermissions(role string, perms []string, object models.ObjectID) result.Result[bool] {
	ret := _m.Called(role, perms, object)

	var r0 result.Result[bool]
	if rf, ok := ret.Get(0).(func(string, []string, models.ObjectID) result.Result[bool]); ok {
		r0 = rf(role, perms, object)
	} else {
		r0 = ret.Get(0).(result.Result[bool])
	}

	return r0
}

// IAuthorizationService_AssignPermissions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AssignPermissions'
type IAuthorizationService_AssignPermissions_Call struct {
	*mock.Call
}

// AssignPermissions is a helper method to define mock.On call
//   - role string
//   - perms []string
//   - object models.ObjectID
func (_e *IAuthorizationService_Expecter) AssignPermissions(role interface{}, perms interface{}, object interface{}) *IAuthorizationService_AssignPermissions_Call {
	return &IAuthorizationService_AssignPermissions_Call{Call: _e.mock.On("AssignPermissions", role, perms, object)}
}

func (_c *IAuthorizationService_AssignPermissions_Call) Run(run func(role string, perms []string, object models.ObjectID)) *IAuthorizationService_AssignPermissions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].([]string), args[2].(models.ObjectID))
	})
	return _c
}

func (_c *IAuthorizationService_AssignPermissions_Call) Return(_a0 result.Result[bool]) *IAuthorizationService_AssignPermissions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IAuthorizationService_AssignPermissions_Call) RunAndReturn(run func(string, []string, models.ObjectID) result.Result[bool]) *IAuthorizationService_AssignPermissions_Call {
	_c.Call.Return(run)
	return _c
}

// CreateAndAddRoleTo provides a mock function with given fields: subject, roleName, object, permissions
func (_m *IAuthorizationService) CreateAndAddRoleTo(subject models.ActorID, roleName string, object option.Option[models.ObjectID], permissions []string) result.Result[bool] {
	ret := _m.Called(subject, roleName, object, permissions)

	var r0 result.Result[bool]
	if rf, ok := ret.Get(0).(func(models.ActorID, string, option.Option[models.ObjectID], []string) result.Result[bool]); ok {
		r0 = rf(subject, roleName, object, permissions)
	} else {
		r0 = ret.Get(0).(result.Result[bool])
	}

	return r0
}

// IAuthorizationService_CreateAndAddRoleTo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateAndAddRoleTo'
type IAuthorizationService_CreateAndAddRoleTo_Call struct {
	*mock.Call
}

// CreateAndAddRoleTo is a helper method to define mock.On call
//   - subject models.ActorID
//   - roleName string
//   - object option.Option[models.ObjectID]
//   - permissions []string
func (_e *IAuthorizationService_Expecter) CreateAndAddRoleTo(subject interface{}, roleName interface{}, object interface{}, permissions interface{}) *IAuthorizationService_CreateAndAddRoleTo_Call {
	return &IAuthorizationService_CreateAndAddRoleTo_Call{Call: _e.mock.On("CreateAndAddRoleTo", subject, roleName, object, permissions)}
}

func (_c *IAuthorizationService_CreateAndAddRoleTo_Call) Run(run func(subject models.ActorID, roleName string, object option.Option[models.ObjectID], permissions []string)) *IAuthorizationService_CreateAndAddRoleTo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(models.ActorID), args[1].(string), args[2].(option.Option[models.ObjectID]), args[3].([]string))
	})
	return _c
}

func (_c *IAuthorizationService_CreateAndAddRoleTo_Call) Return(_a0 result.Result[bool]) *IAuthorizationService_CreateAndAddRoleTo_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IAuthorizationService_CreateAndAddRoleTo_Call) RunAndReturn(run func(models.ActorID, string, option.Option[models.ObjectID], []string) result.Result[bool]) *IAuthorizationService_CreateAndAddRoleTo_Call {
	_c.Call.Return(run)
	return _c
}

// CreateRole provides a mock function with given fields: roleName, object, permissions
func (_m *IAuthorizationService) CreateRole(roleName string, object option.Option[models.ObjectID], permissions []string) result.Result[string] {
	ret := _m.Called(roleName, object, permissions)

	var r0 result.Result[string]
	if rf, ok := ret.Get(0).(func(string, option.Option[models.ObjectID], []string) result.Result[string]); ok {
		r0 = rf(roleName, object, permissions)
	} else {
		r0 = ret.Get(0).(result.Result[string])
	}

	return r0
}

// IAuthorizationService_CreateRole_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateRole'
type IAuthorizationService_CreateRole_Call struct {
	*mock.Call
}

// CreateRole is a helper method to define mock.On call
//   - roleName string
//   - object option.Option[models.ObjectID]
//   - permissions []string
func (_e *IAuthorizationService_Expecter) CreateRole(roleName interface{}, object interface{}, permissions interface{}) *IAuthorizationService_CreateRole_Call {
	return &IAuthorizationService_CreateRole_Call{Call: _e.mock.On("CreateRole", roleName, object, permissions)}
}

func (_c *IAuthorizationService_CreateRole_Call) Run(run func(roleName string, object option.Option[models.ObjectID], permissions []string)) *IAuthorizationService_CreateRole_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(option.Option[models.ObjectID]), args[2].([]string))
	})
	return _c
}

func (_c *IAuthorizationService_CreateRole_Call) Return(_a0 result.Result[string]) *IAuthorizationService_CreateRole_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IAuthorizationService_CreateRole_Call) RunAndReturn(run func(string, option.Option[models.ObjectID], []string) result.Result[string]) *IAuthorizationService_CreateRole_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteRole provides a mock function with given fields: roleID
func (_m *IAuthorizationService) DeleteRole(roleID string) result.Result[bool] {
	ret := _m.Called(roleID)

	var r0 result.Result[bool]
	if rf, ok := ret.Get(0).(func(string) result.Result[bool]); ok {
		r0 = rf(roleID)
	} else {
		r0 = ret.Get(0).(result.Result[bool])
	}

	return r0
}

// IAuthorizationService_DeleteRole_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteRole'
type IAuthorizationService_DeleteRole_Call struct {
	*mock.Call
}

// DeleteRole is a helper method to define mock.On call
//   - roleID string
func (_e *IAuthorizationService_Expecter) DeleteRole(roleID interface{}) *IAuthorizationService_DeleteRole_Call {
	return &IAuthorizationService_DeleteRole_Call{Call: _e.mock.On("DeleteRole", roleID)}
}

func (_c *IAuthorizationService_DeleteRole_Call) Run(run func(roleID string)) *IAuthorizationService_DeleteRole_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *IAuthorizationService_DeleteRole_Call) Return(_a0 result.Result[bool]) *IAuthorizationService_DeleteRole_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IAuthorizationService_DeleteRole_Call) RunAndReturn(run func(string) result.Result[bool]) *IAuthorizationService_DeleteRole_Call {
	_c.Call.Return(run)
	return _c
}

// HasPermission provides a mock function with given fields: subject, perm, object
func (_m *IAuthorizationService) HasPermission(subject models.ActorID, perm string, object option.Option[models.ObjectID]) result.Result[bool] {
	ret := _m.Called(subject, perm, object)

	var r0 result.Result[bool]
	if rf, ok := ret.Get(0).(func(models.ActorID, string, option.Option[models.ObjectID]) result.Result[bool]); ok {
		r0 = rf(subject, perm, object)
	} else {
		r0 = ret.Get(0).(result.Result[bool])
	}

	return r0
}

// IAuthorizationService_HasPermission_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HasPermission'
type IAuthorizationService_HasPermission_Call struct {
	*mock.Call
}

// HasPermission is a helper method to define mock.On call
//   - subject models.ActorID
//   - perm string
//   - object option.Option[models.ObjectID]
func (_e *IAuthorizationService_Expecter) HasPermission(subject interface{}, perm interface{}, object interface{}) *IAuthorizationService_HasPermission_Call {
	return &IAuthorizationService_HasPermission_Call{Call: _e.mock.On("HasPermission", subject, perm, object)}
}

func (_c *IAuthorizationService_HasPermission_Call) Run(run func(subject models.ActorID, perm string, object option.Option[models.ObjectID])) *IAuthorizationService_HasPermission_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(models.ActorID), args[1].(string), args[2].(option.Option[models.ObjectID]))
	})
	return _c
}

func (_c *IAuthorizationService_HasPermission_Call) Return(_a0 result.Result[bool]) *IAuthorizationService_HasPermission_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IAuthorizationService_HasPermission_Call) RunAndReturn(run func(models.ActorID, string, option.Option[models.ObjectID]) result.Result[bool]) *IAuthorizationService_HasPermission_Call {
	_c.Call.Return(run)
	return _c
}

// RemovePermission provides a mock function with given fields: role, perm, object
func (_m *IAuthorizationService) RemovePermission(role models.ActorID, perm string, object models.ObjectID) result.Result[bool] {
	ret := _m.Called(role, perm, object)

	var r0 result.Result[bool]
	if rf, ok := ret.Get(0).(func(models.ActorID, string, models.ObjectID) result.Result[bool]); ok {
		r0 = rf(role, perm, object)
	} else {
		r0 = ret.Get(0).(result.Result[bool])
	}

	return r0
}

// IAuthorizationService_RemovePermission_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemovePermission'
type IAuthorizationService_RemovePermission_Call struct {
	*mock.Call
}

// RemovePermission is a helper method to define mock.On call
//   - role models.ActorID
//   - perm string
//   - object models.ObjectID
func (_e *IAuthorizationService_Expecter) RemovePermission(role interface{}, perm interface{}, object interface{}) *IAuthorizationService_RemovePermission_Call {
	return &IAuthorizationService_RemovePermission_Call{Call: _e.mock.On("RemovePermission", role, perm, object)}
}

func (_c *IAuthorizationService_RemovePermission_Call) Run(run func(role models.ActorID, perm string, object models.ObjectID)) *IAuthorizationService_RemovePermission_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(models.ActorID), args[1].(string), args[2].(models.ObjectID))
	})
	return _c
}

func (_c *IAuthorizationService_RemovePermission_Call) Return(_a0 result.Result[bool]) *IAuthorizationService_RemovePermission_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IAuthorizationService_RemovePermission_Call) RunAndReturn(run func(models.ActorID, string, models.ObjectID) result.Result[bool]) *IAuthorizationService_RemovePermission_Call {
	_c.Call.Return(run)
	return _c
}

// RemovePermissions provides a mock function with given fields: role, perms, object
func (_m *IAuthorizationService) RemovePermissions(role models.ActorID, perms []string, object models.ObjectID) result.Result[bool] {
	ret := _m.Called(role, perms, object)

	var r0 result.Result[bool]
	if rf, ok := ret.Get(0).(func(models.ActorID, []string, models.ObjectID) result.Result[bool]); ok {
		r0 = rf(role, perms, object)
	} else {
		r0 = ret.Get(0).(result.Result[bool])
	}

	return r0
}

// IAuthorizationService_RemovePermissions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemovePermissions'
type IAuthorizationService_RemovePermissions_Call struct {
	*mock.Call
}

// RemovePermissions is a helper method to define mock.On call
//   - role models.ActorID
//   - perms []string
//   - object models.ObjectID
func (_e *IAuthorizationService_Expecter) RemovePermissions(role interface{}, perms interface{}, object interface{}) *IAuthorizationService_RemovePermissions_Call {
	return &IAuthorizationService_RemovePermissions_Call{Call: _e.mock.On("RemovePermissions", role, perms, object)}
}

func (_c *IAuthorizationService_RemovePermissions_Call) Run(run func(role models.ActorID, perms []string, object models.ObjectID)) *IAuthorizationService_RemovePermissions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(models.ActorID), args[1].([]string), args[2].(models.ObjectID))
	})
	return _c
}

func (_c *IAuthorizationService_RemovePermissions_Call) Return(_a0 result.Result[bool]) *IAuthorizationService_RemovePermissions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IAuthorizationService_RemovePermissions_Call) RunAndReturn(run func(models.ActorID, []string, models.ObjectID) result.Result[bool]) *IAuthorizationService_RemovePermissions_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveRoleFrom provides a mock function with given fields: subject, roleID
func (_m *IAuthorizationService) RemoveRoleFrom(subject models.ActorID, roleID string) result.Result[bool] {
	ret := _m.Called(subject, roleID)

	var r0 result.Result[bool]
	if rf, ok := ret.Get(0).(func(models.ActorID, string) result.Result[bool]); ok {
		r0 = rf(subject, roleID)
	} else {
		r0 = ret.Get(0).(result.Result[bool])
	}

	return r0
}

// IAuthorizationService_RemoveRoleFrom_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveRoleFrom'
type IAuthorizationService_RemoveRoleFrom_Call struct {
	*mock.Call
}

// RemoveRoleFrom is a helper method to define mock.On call
//   - subject models.ActorID
//   - roleID string
func (_e *IAuthorizationService_Expecter) RemoveRoleFrom(subject interface{}, roleID interface{}) *IAuthorizationService_RemoveRoleFrom_Call {
	return &IAuthorizationService_RemoveRoleFrom_Call{Call: _e.mock.On("RemoveRoleFrom", subject, roleID)}
}

func (_c *IAuthorizationService_RemoveRoleFrom_Call) Run(run func(subject models.ActorID, roleID string)) *IAuthorizationService_RemoveRoleFrom_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(models.ActorID), args[1].(string))
	})
	return _c
}

func (_c *IAuthorizationService_RemoveRoleFrom_Call) Return(_a0 result.Result[bool]) *IAuthorizationService_RemoveRoleFrom_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IAuthorizationService_RemoveRoleFrom_Call) RunAndReturn(run func(models.ActorID, string) result.Result[bool]) *IAuthorizationService_RemoveRoleFrom_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewIAuthorizationService interface {
	mock.TestingT
	Cleanup(func())
}

// NewIAuthorizationService creates a new instance of IAuthorizationService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIAuthorizationService(t mockConstructorTestingTNewIAuthorizationService) *IAuthorizationService {
	mock := &IAuthorizationService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}