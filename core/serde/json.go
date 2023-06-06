package serde

import (
	"github.com/gpabois/cougnat/core/result"
	"github.com/mongodb/mongo-tools-common/json"
)

// Marshal a BSON Document
func MarshalJson[T any](value T) result.Result[[]byte] {
	normRes := Normalise(value)
	data, err := json.Marshal(normRes)
	if err != nil {
		return result.Failed[[]byte](err)
	}
	return result.Success(data)
}

func UnMarshalJson[T any](raw []byte) result.Result[T] {
	var norm map[string]any = make(map[string]any)
	err := json.Unmarshal(raw, &norm)

	if err != nil {
		return result.Failed[T](err)
	}

	return DeNormalise[T](norm)
}
