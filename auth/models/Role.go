package auth_models

import "github.com/gpabois/cougnat/core/option"

type RoleID = string

type Role struct {
	ID          option.Option[RoleID]
	ObjectID    ObjectID
	Name        string
	permissions []string
	subjects    []ActorID
}
