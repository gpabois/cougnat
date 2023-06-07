package serde

import (
	"errors"
	"fmt"

	"github.com/gpabois/cougnat/core/result"
)

func Marshal[T any](value T, contentType string) result.Result[[]byte] {
	switch contentType {
	case "application/bson":
		return MarshalBson(value)
	case "application/json":
		return MarshalJson(value)
	default:
		return result.Failed[[]byte](errors.New(fmt.Sprintf("unhandled content-type %s", contentType)))
	}
}

func UnMarshal[T any](value []byte, contentType string) result.Result[T] {
	switch contentType {
	case "application/bson":
		return UnMarshalBson[T](value)
	case "application/json":
		return UnMarshalJson[T](value)
	default:
		return result.Failed[T](errors.New(fmt.Sprintf("unhandled content-type %s", contentType)))
	}
}
