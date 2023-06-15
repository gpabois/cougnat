package serde

import (
	"errors"
	"fmt"
	"io"

	"github.com/gpabois/cougnat/core/result"
)

func Marshal[T any](value T, contentType string) result.Result[[]byte] {
	switch contentType {
	case "application/bson":
		return MarshalBson(value)
	case "application/json":
		return MarshalJson(value)
	default:
		return result.Failed[[]byte](errors.New(fmt.Sprintf("cannot unmarshal content-type '%s'", contentType)))
	}
}

func MarshalStream[T any](w io.Writer, value T, contentType string) result.Result[int] {
	res := Marshal(value, contentType)

	if res.HasFailed() {
		return result.Result[int]{}.Failed(res.UnwrapError())
	}

	written, err := w.Write(res.Expect())

	if err != nil {
		return result.Result[int]{}.Failed(err)
	}

	return result.Success(written)
}

func UnMarshalStream[T any](stream io.Reader, contentType string) result.Result[T] {
	value, err := io.ReadAll(stream)

	if err != nil {
		return result.Result[T]{}.Failed(err)
	}

	switch contentType {
	case "application/bson":
		return UnMarshalBson[T](value)
	case "application/json":
		return UnMarshalJson[T](value)
	default:
		return result.Failed[T](errors.New(fmt.Sprintf("unhandled content-type %s", contentType)))
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
