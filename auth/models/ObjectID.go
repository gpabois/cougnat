package auth_models

type ObjectID struct {
	nature string
	id     string
}

func NewObjectID(nature string, id string) ObjectID {
	return ObjectID{
		nature,
		id,
	}
}
