package models

type ActorID struct {
	nature string
	id     string
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
