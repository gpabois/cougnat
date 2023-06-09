package services

import (
	"github.com/gpabois/cougnat/auth/models"
	"github.com/gpabois/cougnat/auth/repositories"
	"github.com/gpabois/cougnat/core/option"
	"github.com/gpabois/cougnat/core/result"
)

//go:generate mockery
type IAuthorizationService interface {
	// Create a role and add it to the subject.
	// Equivalent to CreateRole + AddRoleTo
	CreateAndAddRoleTo(subject models.ActorID, roleName string, object option.Option[models.ObjectID], permissions []string) result.Result[bool]
	// Quick operations
	RemoveByObjectID(objectID models.ObjectID) result.Result[bool]

	// Manage roles
	CreateRole(roleName string, object option.Option[models.ObjectID], permissions []string) result.Result[models.RoleID]
	DeleteRole(roleID models.RoleID) result.Result[bool]
	AddRoleTo(subject models.ActorID, roleID models.RoleID) result.Result[bool]
	RemoveRoleFrom(subject models.ActorID, roleID models.RoleID) result.Result[bool]

	// Manage permissions
	AssignPermission(role models.RoleID, perm string) result.Result[bool]
	AssignPermissions(role models.RoleID, perms []string, object models.ObjectID) result.Result[bool]
	RemovePermission(role models.ActorID, perm string, object models.ObjectID) result.Result[bool]
	RemovePermissions(role models.ActorID, perms []string, object models.ObjectID) result.Result[bool]

	// Check if the subject has the permission on the object
	HasPermission(subject models.ActorID, perm string, object option.Option[models.ObjectID]) result.Result[bool]
}

type AuthorizationService struct {
	roleRepo repositories.RoleRepository
}

func (authz *AuthorizationService) HasPermission(subject models.ActorID, perm string, object models.ObjectID) result.Result[bool] {
	authz.roleRepo.ExistWithSPO(subject, perm, object)
	// If it's an user, we can check if any group has the permission
	if subject.IsUser() {
		return authz.roleRepo.ExistWithSPO(subject, perm, object)
	}

	return result.Success(true)
}
