package serde

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	"github.com/gpabois/cougnat/core/result"
	"github.com/gpabois/cougnat/core/serde/bson"
	"github.com/gpabois/cougnat/core/serde/decoder"
	"github.com/gpabois/cougnat/core/serde/json"
)

func getDecoderFromReader(r io.Reader, contentType string) result.Result[decoder.Decoder] {
	switch contentType {
	case "application/bson":
		return result.Success[decoder.Decoder](bson.NewDecoder(r))
	case "application/json":
		return result.Success[decoder.Decoder](json.NewDecoder(r))
	default:
		return result.Failed[decoder.Decoder](errors.New(fmt.Sprintf("cannot decode content-type '%s'", contentType)))
	}
}

func getDecoderFromBytes(b []byte, contentType string) result.Result[decoder.Decoder] {
	buf := bytes.NewBuffer(b)
	return getDecoderFromReader(buf, contentType)
}

func Deserialize[T any](value []byte, contentType string) result.Result[T] {
	res := getDecoderFromBytes(value, contentType)
	if res.HasFailed() {
		return result.Result[T]{}.Failed(res.UnwrapError())
	}

	return decoder.Decode[T](res.Expect())
}
