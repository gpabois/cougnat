package auth_models

import "github.com/gpabois/gostd/option"

type ActorID struct {
	ID     option.Option[string]
	Nature string
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
	return id.ID.IsSome()
}

func (id ActorID) IsUser() bool {
	return id.Nature == "user"
}

func AnonymousID(id option.Option[string]) ActorID {
	return ActorID{
		Nature: "anonymous",
		ID:     id,
	}
}

func UserID(id string) ActorID {
	return ActorID{
		Nature: "user",
		ID:     option.Some(id),
	}
}

func GroupID(id string) ActorID {
	return ActorID{
		Nature: "group",
		ID:     option.Some(id),
	}
}

func OrganisationID(id string) ActorID {
	return ActorID{
		Nature: "service",
		ID:     option.Some(id),
	}
}
