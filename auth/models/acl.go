package models

import "github.com/gpabois/cougnat/core/collection/tuple"

type AccessControl tuple.Triplet[ActorID, string, ObjectID]

func NewAccessControl(actorID ActorID, right string, objID ObjectID) AccessControl {
	return AccessControl{actorID, right, objID}
}
