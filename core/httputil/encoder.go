package httputil

import (
	"context"
	"net/http"

	"github.com/gpabois/cougnat/core/serde"
)

func EncodeResponse[T any](ctx context.Context, w http.ResponseWriter, response any) error {
	contentType := "application/json"
	encodeResult := serde.MarshalStream(
		w,
		HttpResult[T]{}.Success(response.(T)),
		contentType,
	)
	return encodeResult.UnwrapError()
}

func EncodeError[T any](_ context.Context, err error, w http.ResponseWriter) {
	contentType := "application/json"
	w.Header().Set("Content-Type", contentType)
	serde.MarshalStream(
		w,
		HttpResult[T]{}.Failed(err),
		contentType,
	)
}
