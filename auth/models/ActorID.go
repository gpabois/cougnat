package models

import "github.com/gpabois/cougnat/core/option"

type ActorID struct {
	nature string
	id     option.Option[string]
}

func ActorID_TryFromAny(val any) option.Option[ActorID] {
	if val == nil {
		return option.None[ActorID]()
	}

	actorID, ok := val.(ActorID)
	if ok == false {
		return option.None[ActorID]()
	} else {
		return option.Some(actorID)
	}
}

func (id ActorID) IsBound() bool {
	return id.id.IsSome()
}

func (id ActorID) IsUser() bool {
	return id.nature == "user"
}

func AnonymousID(id option.Option[string]) ActorID {
	return ActorID{
		nature: "anonymous",
		id:     id,
	}
}

func UserID(id string) ActorID {
	return ActorID{
		nature: "user",
		id:     option.Some(id),
	}
}

func GroupID(id string) ActorID {
	return ActorID{
		nature: "group",
		id:     option.Some(id),
	}
}

func OrganisationID(id string) ActorID {
	return ActorID{
		nature: "service",
		id:     option.Some(id),
	}
}
