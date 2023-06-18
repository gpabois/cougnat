package serde

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/gpabois/cougnat/core/option"
	"github.com/gpabois/cougnat/core/result"
	"golang.org/x/exp/slices"
)

type NormalisedStruct = map[string]any

var PRIMARY_TYPES = []reflect.Kind{
	reflect.Bool,
	reflect.Int,
	reflect.Int8,
	reflect.Int16,
	reflect.Int32,
	reflect.Int64,
	reflect.Uint,
	reflect.Uint8,
	reflect.Uint16,
	reflect.Uint32,
	reflect.Uint64,
	reflect.Uintptr,
	reflect.Float32,
	reflect.Float64,
	reflect.Complex64,
	reflect.Complex128,
	reflect.String,
	reflect.Interface,
}

type Encoder interface {
	WriteInt(value int)
	WriteFloat32(value float32)
	WriteBool(value bool)
	WriteString(value string)
	PushArray()
	PushMap()
	PushMapKey()
	PushMapValue()
	Pop()
}

func encode(enc Encoder, value reflect.Value) {

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
			encode(enc, reflect.ValueOf(fieldName))
			enc.Pop()

			enc.PushMapValue()
			encode(enc, reflect.ValueOf(optVal.Get()))
			enc.Pop()
		} else {
			enc.PushMapKey()
			encode(enc, reflect.ValueOf(fieldName))
			enc.Pop()

			enc.PushMapValue()
			encode(enc, reflect.ValueOf(fieldValue))
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
		encode(enc, mapKey)
		enc.Pop()

		enc.PushMapValue()
		encode(enc, mapValue)
		enc.Pop()
	}
	enc.Pop()
}
func encodeSlice(enc Encoder, value reflect.Value) {
	enc.PushArray()
	for i := 0; i < value.Len(); i++ {
		encode(enc, value.Index(i))
	}
	enc.Pop()
}

func normalise(val any) any {
	typ := reflect.ValueOf(val).Type()
	valOf := reflect.ValueOf(val)

	if slices.Contains(PRIMARY_TYPES, typ.Kind()) {
		return val
	}

	if typ.Kind() == reflect.Slice {
		slc := []any{}
		for i := 0; i < valOf.Len(); i++ {
			slc = append(slc, normalise(valOf.Index(i).Interface()))
		}
		return any(slc)
	}

	if typ.Kind() == reflect.Map {
		mapp := map[string]any{}

		rmap := reflect.ValueOf(val)
		for _, key := range rmap.MapKeys() {
			mapp[key.String()] = normalise(rmap.MapIndex(key).Elem().Interface())
		}
		return any(mapp)
	}

	// We handle struct-based types
	if typ.Kind() != reflect.Struct {
		return result.Failed[any](errors.New(fmt.Sprintf("type %v cannot be denormalised", typ.Kind())))
	}

	norm := make(map[string]any)
	anyValOf := reflect.ValueOf(val)

	for i := 0; i < typ.NumField(); i++ {
		field := anyValOf.Field(i)

		fieldName := typ.Field(i).Name

		marshalName, ok := typ.Field(i).Tag.Lookup("serde")
		if ok {
			fieldName = marshalName
		}

		if !field.CanInterface() {
			continue
		}

		fieldVal := field.Interface()

		if option.IsOption(fieldVal) {
			optVal := fieldVal.(option.IOption)

			// Don't set nil value
			if optVal.IsNone() {
				continue
			}

			norm[fieldName] = normalise(optVal.Get())
		} else {
			norm[fieldName] = normalise(fieldVal)
		}
	}

	return norm
}

func Normalise[T any](value T) any {
	return normalise(value)
}

func denormaliseStruct(typ reflect.Type, norm reflect.Value) result.Result[reflect.Value] {
	// Instantiate a struct
	value := reflect.New(typ).Elem()

	for i := 0; i < typ.NumField(); i++ {
		field := value.Field(i)
		fieldType := field.Type()
		fieldName := typ.Field(i).Name

		marshalName, ok := typ.Field(i).Tag.Lookup("serde")
		if ok {
			fieldName = marshalName
		}

		normValue := norm.MapIndex(reflect.ValueOf(fieldName))
		if normValue.IsValid() == false {
			continue
		}
		normValue = normValue.Elem()

		// Take care of optional values
		if field.CanAddr() && option.IsMutableOption(field.Addr().Interface()) && normValue.IsValid() {
			res := denormalise(field.Interface().(option.IOption).TypeOf(), normValue)

			if res.HasFailed() {
				return result.Failed[reflect.Value](res.UnwrapError())
			}

			resSet := field.Addr().Interface().(option.IMutableOption).TrySet(res.Expect())

			if resSet.HasFailed() {
				return result.Failed[reflect.Value](res.UnwrapError())
			}
		} else if normValue.IsValid() { // Manage the rest
			res := denormalise(fieldType, normValue)
			if res.HasFailed() {
				return result.Failed[reflect.Value](res.UnwrapError())
			}

			// Set the value
			if field.IsValid() {
				field.Set(res.Expect())
			}
		}
	}

	return result.Success(value)
}

func denormaliseMap(typ reflect.Type, norm reflect.Value) result.Result[reflect.Value] {
	//keyType := nval.Type().Key()
	valType := typ.Elem()
	val := reflect.New(reflect.MapOf(typ.Key(), typ.Elem()))

	for _, key := range norm.MapKeys() {
		field := norm.MapIndex(key)

		denormRes := denormalise(valType, field.Elem())
		if denormRes.HasFailed() {
			return result.Result[reflect.Value]{}.Failed(denormRes.UnwrapError())
		}

		val.SetMapIndex(key, denormRes.Expect())
	}

	return result.Success(val.Elem())
}

func denormaliseSlice(typ reflect.Type, elValues reflect.Value) result.Result[reflect.Value] {
	arrElType := typ.Elem()

	slc := reflect.New(reflect.SliceOf(arrElType))
	//
	for i := 0; i < elValues.Len(); i++ {
		// Denormalise the element
		res := denormalise(arrElType, elValues.Index(i))

		if res.HasFailed() {
			return result.Failed[reflect.Value](res.UnwrapError())
		}
		// Add
		if slc.Elem().IsValid() {
			slc.Elem().Set(reflect.Append(slc.Elem(), reflect.ValueOf(res.Expect())))
		}
	}

	// Denormalise the array
	return result.Success(slc.Elem())
}

func denormalise(typ reflect.Type, val reflect.Value) result.Result[reflect.Value] {
	// Primary type, easy !
	if slices.Contains(PRIMARY_TYPES, typ.Kind()) {
		return result.Success(val)
	}

	// Slice-based values, we need to denormalise each element.
	if typ.Kind() == reflect.Slice || typ.Kind() == reflect.Array {
		return denormaliseSlice(typ, val)
	}

	// Map-based values
	if typ.Kind() == reflect.Map {
		return denormaliseMap(typ, val)
	}

	// We handle struct-based types
	if typ.Kind() != reflect.Struct {
		return result.Failed[reflect.Value](errors.New(fmt.Sprintf("type %v cannot be denormalised", typ.Kind())))
	}

	return denormaliseStruct(typ, val)
}

func DeNormalise[T any](value any) result.Result[T] {
	typ := reflect.TypeOf((*T)(nil)).Elem()
	val := reflect.ValueOf(value)
	return result.Map(
		denormalise(typ, val),
		func(val reflect.Value) T {
			return val.Interface().(T)
		},
	)
}
