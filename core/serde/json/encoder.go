package json

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	"github.com/gpabois/cougnat/core/collection"
	"github.com/gpabois/cougnat/core/result"
)

type EncoderState struct {
	typ     int
	counter int
	buffer  bytes.Buffer
}

const ROOT_STATE = 0
const MAP_STATE = 1
const MAP_KEY_STATE = 2
const MAP_VALUE_STATE = 3
const ARRAY_STATE = 4
const ARRAY_VALUE_STATE = 5

type Encoder struct {
	states collection.Stack[EncoderState]
	writer io.Writer
}

func (enc *Encoder) EncodeInt64(value int64) result.Result[bool] {
	return enc.WriteString(fmt.Sprintf("%d", value))
}
func (enc *Encoder) EncodeFloat64(value float64) result.Result[bool] {
	return enc.WriteString(fmt.Sprintf("%f", value))
}
func (enc *Encoder) EncodeBool(value bool) result.Result[bool] {
	if value {
		return enc.WriteString("true")
	} else {
		return enc.WriteString("false")
	}
}
func (enc *Encoder) EncodeString(value string) result.Result[bool] {
	return enc.WriteString(fmt.Sprintf("\"%s\"", value))
}

func (enc *Encoder) WriteString(s string) result.Result[bool] {
	_, err := enc.writer.Write([]byte(s))

	if err != nil {
		return result.Failed[bool](err)
	}

	return result.Success(true)
}

func (enc *Encoder) PushArray() result.Result[bool] {
	enc.states.Push(EncoderState{typ: ARRAY_STATE})
	return enc.WriteString("[")
}

func (enc *Encoder) PushArrayValue() result.Result[bool] {
	currentStateRes := enc.states.Last().IntoResult(errors.New("no state"))
	if currentStateRes.HasFailed() {
		return result.Result[bool]{}.Failed(currentStateRes.UnwrapError())
	}

	currentState := currentStateRes.Expect()

	if currentState.typ != ARRAY_STATE {
		return result.Result[bool]{}.Failed(errors.New("expecting the encoder to be in a map state"))
	}

	if currentState.counter > 0 {
		enc.WriteString(",")
	}
	enc.states.Push(EncoderState{typ: ARRAY_VALUE_STATE})

	return result.Success(true)
}

func (enc *Encoder) PushMap() result.Result[bool] {
	enc.states.Push(EncoderState{typ: MAP_STATE})
	enc.WriteString("{")
	return result.Success(true)
}

func (enc *Encoder) PushMapKey() result.Result[bool] {
	currentStateRes := enc.states.Last().IntoResult(errors.New("no state"))

	if currentStateRes.HasFailed() {
		return result.Result[bool]{}.Failed(currentStateRes.UnwrapError())
	}

	currentState := currentStateRes.Expect()

	if currentState.typ != MAP_STATE {
		return result.Result[bool]{}.Failed(errors.New("expecting the encoder to be in a map state"))
	}

	if currentState.counter > 0 {
		if res := enc.WriteString(","); res.HasFailed() {
			return result.Result[bool]{}.Failed(res.UnwrapError())
		}
	}

	currentState.counter++
	enc.states.Push(EncoderState{typ: MAP_KEY_STATE})
	return result.Success(true)
}

func (enc *Encoder) PushMapValue() result.Result[bool] {
	enc.states.Push(EncoderState{typ: MAP_VALUE_STATE})
	enc.WriteString(":")
	return result.Success(true)
}

func (enc *Encoder) Pop() result.Result[bool] {
	state := enc.states.Pop().IntoResult(errors.New("no state was popped"))

	if state.HasFailed() {
		return result.Result[bool]{}.Failed(state.UnwrapError())
	}

	switch state.Expect().typ {
	case MAP_STATE:
		return enc.WriteString("}")
	case ARRAY_STATE:
		return enc.WriteString("]")
	case ARRAY_VALUE_STATE:
	case MAP_VALUE_STATE:
	case MAP_KEY_STATE:
		return result.Success(true)
	default:
		return result.Failed[bool](errors.New("invalid encoder state"))
	}

	return result.Success(true)
}
