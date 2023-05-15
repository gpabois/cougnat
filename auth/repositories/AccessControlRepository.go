package repositories

import (
	"github.com/gpabois/cougnat/auth/models"
	"github.com/gpabois/cougnat/core/option"
	"github.com/gpabois/cougnat/core/result"
)

type RoleRepository interface {
	// Get a role by its ID
	GetByID(roleID models.RoleID) result.Result[option.Option[models.Role]]
	// Get a role by its name and ObjectID
	GetByNameAndObjectID(roleName string, objectID models.ObjectID) result.Result[option.Option[models.Role]]
	// Return true if any role with the triplet SubjectID / Permission / ObjectID exist
	ExistWithSPO(subject models.ActorID, perm string, objectID models.ObjectID) result.Result[bool]
}
