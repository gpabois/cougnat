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

var primaryTypes = []reflect.Kind{
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
}

func Normalise[T any](value T) any {
	val := reflect.ValueOf(value)

	// Deref pointer
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if slices.Contains(primaryTypes, val.Kind()) {
		return value
	}

	norm := make(map[string]any)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := val.Type().Field(i).Name

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

			norm[fieldName] = Normalise(optVal.Get())
		} else {
			norm[fieldName] = Normalise(fieldVal)
		}
	}

	return norm
}

func denormaliseByReflectType(typ reflect.Type, val any) result.Result[any] {

	// Primary type, easy !
	if slices.Contains(primaryTypes, typ.Kind()) {
		switch typ.Kind() {
		case reflect.Int64:
		case reflect.Int32:
		case reflect.Int16:
		case reflect.Int8:
		case reflect.Int:
			integer := reflect.ValueOf(val).Int()
			return result.Success[any](int(integer))
		default:
			return result.Success(val)
		}

	}

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
			field.Set(reflect.ValueOf(res.Expect()))
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
