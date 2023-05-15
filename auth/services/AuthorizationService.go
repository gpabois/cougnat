package services

import (
	"github.com/gpabois/cougnat/auth/models"
	"github.com/gpabois/cougnat/auth/repositories"
	"github.com/gpabois/cougnat/core/result"
)

// Core Authorization Service
type AuthorizationService interface {
	// Create a role and add it to the subject.
	// Equivalent to CreateRole + AddRoleTo
	CreateAndAddRoleTo(subject models.ActorID, roleName string, object models.ObjectID, permissions []string)

	// Manage roles
	CreateRole(roleName string, object models.ObjectID, permissions []string) result.Result[models.RoleID]
	DeleteRole(roleID models.RoleID) result.Result[bool]
	AddRoleTo(subject models.ActorID, roleID models.RoleID) result.Result[bool]
	RemoveRoleFrom(subject models.ActorID, roleID models.RoleID) result.Result[bool]

	// Manage permissions
	AssignPermission(role models.RoleID, perm string) result.Result[bool]
	AssignPermissions(role models.RoleID, perms []string, object models.ObjectID) result.Result[bool]
	RemovePermission(role models.ActorID, perm string, object models.ObjectID) result.Result[bool]
	RemovePermissions(role models.ActorID, perms []string, object models.ObjectID) result.Result[bool]

	// Check if the subject has the permission on the object
	HasPermission(subject models.ActorID, perm string, object models.ObjectID) result.Result[bool]
}

type ImplAuthorizationService struct {
	roleRepo repositories.RoleRepository
}

func (authz ImplAuthorizationService) HasPermission(subject models.ActorID, perm string, object models.ObjectID) result.Result[bool] {
	authz.roleRepo.ExistWithSPO(subject, perm, object)
	// If it's an user, we can check if any group has the permission
	if subject.IsUser() {

	}
}
