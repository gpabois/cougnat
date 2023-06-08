package models

import (
	"github.com/gpabois/cougnat/core/option"
)

type AccessControl struct {
	Actor      ActorID
	Permission string
	Object     option.Option[ObjectID]
}

func NewAccessControl(actorID ActorID, right string, objID option.Option[ObjectID]) AccessControl {
	return AccessControl{Actor: actorID, Permission: right, Object: objID}
}
