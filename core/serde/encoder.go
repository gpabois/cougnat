package serde

import (
	"reflect"

	"github.com/gpabois/cougnat/core/option"
	"github.com/gpabois/cougnat/core/result"
)

type Encoder interface {
	EncodeInt64(value int64) result.Result[bool]
	EncodeFloat64(value float64) result.Result[bool]
	EncodeBool(value bool) result.Result[bool]
	EncodeString(value string) result.Result[bool]
	PushArray() result.Result[bool]
	PushArrayValue() result.Result[bool]
	PushMap() result.Result[bool]
	PushMapKey() result.Result[bool]
	PushMapValue() result.Result[bool]
	Pop() result.Result[bool]
}

func Encode[T any](enc Encoder, value T) {
	encode(enc, reflect.ValueOf(value))
}

func encode(enc Encoder, value reflect.Value) {
	switch value.Type().Kind() {
	case reflect.Int64:
	case reflect.Int32:
	case reflect.Int16:
	case reflect.Int8:
	case reflect.Int:
		enc.EncodeInt64(value.Int())
	case reflect.String:
		enc.EncodeString(value.String())
	case reflect.Float64:
	case reflect.Float32:
		enc.EncodeFloat64(value.Float())
	case reflect.Bool:
		enc.EncodeBool(value.Bool())
	case reflect.Array:
	case reflect.Slice:
		encodeSlice(enc, value)
	case reflect.Map:
		encodeMap(enc, value)
	case reflect.Struct:
		encodeStruct(enc, value)
	}
}

func encodeStruct(enc Encoder, value reflect.Value) {
	typ := value.Type()
	for i := 0; i < typ.NumField(); i++ {
		fieldValue := value.Field(i)
		fieldName := typ.Field(i).Name

		marshalName, ok := typ.Field(i).Tag.Lookup("serde")
		if ok {
			fieldName = marshalName
		}

		if option.IsOption(fieldValue) {
			optVal := fieldValue.Interface().(option.IOption)

			// Don't set nil value
			if optVal.IsNone() {
				continue
			}

			enc.PushMapKey()
			Encode(enc, reflect.ValueOf(fieldName))
			enc.Pop()

			enc.PushMapValue()
			Encode(enc, reflect.ValueOf(optVal.Get()))
			enc.Pop()
		} else {
			enc.PushMapKey()
			Encode(enc, reflect.ValueOf(fieldName))
			enc.Pop()

			enc.PushMapValue()
			Encode(enc, reflect.ValueOf(fieldValue))
			enc.Pop()
		}
	}
	enc.Pop()
}

func encodeMap(enc Encoder, value reflect.Value) {
	enc.PushMap()
	for _, mapKey := range value.MapKeys() {
		enc.PushMapKey()
		mapValue := value.MapIndex(mapKey)
		Encode(enc, mapKey)
		enc.Pop()

		enc.PushMapValue()
		Encode(enc, mapValue)
		enc.Pop()
	}
	enc.Pop()
}
func encodeSlice(enc Encoder, value reflect.Value) {
	enc.PushArray()
	for i := 0; i < value.Len(); i++ {
		enc.PushArrayValue()
		Encode(enc, value.Index(i))
		enc.Pop()
	}
	enc.Pop()
}
