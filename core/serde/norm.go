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

func normaliseByReflectType(val any) any {
	typ := reflect.ValueOf(val).Type()
	valOf := reflect.ValueOf(val)

	if slices.Contains(PRIMARY_TYPES, typ.Kind()) {
		return val
	}

	if typ.Kind() == reflect.Slice {
		slc := []any{}
		for i := 0; i < valOf.Len(); i++ {
			slc = append(slc, normaliseByReflectType(valOf.Index(i).Interface()))
		}
		return any(slc)
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

			norm[fieldName] = normaliseByReflectType(optVal.Get())
		} else {
			norm[fieldName] = normaliseByReflectType(fieldVal)
		}
	}

	return norm
}

func Normalise[T any](value T) any {
	return normaliseByReflectType(value)
}

func denormaliseByReflectType(typ reflect.Type, val any) result.Result[any] {
	// Primary type, easy !
	if slices.Contains(PRIMARY_TYPES, typ.Kind()) {
		switch typ.Kind() {
		case reflect.Int64:
		case reflect.Int32:
		case reflect.Int16:
		case reflect.Int8:
		case reflect.Int:
			integer := reflect.ValueOf(val).Int()
			return result.Success[any](int(integer))
		// Type is any we return it
		case reflect.Interface:
			return result.Success(val)
		default:
			return result.Success(val)
		}
	}

	// Slice-based values, we need to denormalise each element.
	if typ.Kind() == reflect.Slice || typ.Kind() == reflect.Array {
		arrElType := typ.Elem()
		elValues := reflect.ValueOf(val)

		slc := reflect.New(reflect.SliceOf(arrElType))
		//
		for i := 0; i < elValues.Len(); i++ {
			// Denormalise the element
			res := denormaliseByReflectType(arrElType, elValues.Index(i).Interface())

			if res.HasFailed() {
				return result.Failed[any](res.UnwrapError())
			}
			// Add
			if slc.Elem().IsValid() {
				slc.Elem().Set(reflect.Append(slc.Elem(), reflect.ValueOf(res.Expect())))
			}
		}

		// Denormalise the array
		return result.Success(slc.Elem().Interface())
	}

	// We handle struct-based types
	if typ.Kind() != reflect.Struct {
		return result.Failed[any](errors.New(fmt.Sprintf("type %v cannot be denormalised", typ.Kind())))
	}

	norm, ok := val.(NormalisedStruct)

	if !ok {
		return result.Failed[any](errors.New(fmt.Sprintf("expecting a normalised struct got: %v", val)))
	}

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
		normValue, ok := norm[fieldName]

		// Take care of optional values
		if field.CanAddr() && option.IsMutableOption(field.Addr().Interface()) {
			res := denormaliseByReflectType(field.Interface().(option.IOption).TypeOf(), normValue)

			if res.HasFailed() {
				return result.Failed[any](res.UnwrapError())
			}

			resSet := field.
				Addr().
				Interface().(option.IMutableOption).
				TrySet(res.Expect())

			if resSet.HasFailed() {
				return result.Failed[any](res.UnwrapError())
			}
		} else { // Manage the rest
			if !ok {
				continue
			}

			if !field.CanInterface() || !field.CanSet() {
				continue
			}

			res := denormaliseByReflectType(fieldType, normValue)

			if res.HasFailed() {
				return result.Failed[any](res.UnwrapError())
			}

			// Set the value
			if field.IsValid() {
				field.Set(reflect.ValueOf(res.Expect()))
			}
		}
	}

	return result.Success(value.Interface())
}

func DeNormalise[T any](val any) result.Result[T] {
	typ := reflect.TypeOf((*T)(nil)).Elem()
	return result.Map(
		denormaliseByReflectType(typ, val),
		func(val any) T {
			return val.(T)
		},
	)
}
