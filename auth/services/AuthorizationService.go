package services

import (
	"github.com/gpabois/cougnat/auth/models"
	"github.com/gpabois/cougnat/core/result"
)

type AuthorizationService interface {
	// Permissions
	AddPermission(actorID models.ActorID, perm string, object models.ObjectID) result.Result[bool]
	AddPermissions(actorID models.ActorID, perms []string, object models.ObjectID) result.Result[bool]
	RemovePermission(actorID models.ActorID, perm string, object models.ObjectID) result.Result[bool]
	RemovePermissions(actorID models.ActorID, perms []string, object models.ObjectID) result.Result[bool]
	HasPermission(actorID models.ActorID, perm string, object models.ObjectID) result.Result[bool]
}
