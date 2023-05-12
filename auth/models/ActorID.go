package models

import "github.com/gpabois/cougnat/core/option"

type ActorID struct {
	nature string
	id     string
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
	return id.id != ""
}

func AnonymousID(id string) ActorID {
	return ActorID{
		nature: "anonymous",
		id:     id,
	}
}

func UserID(id string) ActorID {
	return ActorID{
		nature: "user",
		id:     id,
	}
}

func GroupID(id string) ActorID {
	return ActorID{
		nature: "group",
		id:     id,
	}
}

func ServiceID(id string) ActorID {
	return ActorID{
		nature: "service",
		id:     id,
	}
}
