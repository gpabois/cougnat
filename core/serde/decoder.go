package serde

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/gpabois/cougnat/core/iter"
	"github.com/gpabois/cougnat/core/option"
	"github.com/gpabois/cougnat/core/result"
)

type Element interface {
	Key() string
	Value() any
}

type Decoder interface {
	// Init the decoder and return data to be decoded along the way
	Init() result.Result[any]
	// Decode as a primary type
	DecodePrimaryType(data any, typ reflect.Type) result.Result[reflect.Value]
	// Iter over encoded slice element
	IterSlice(data any) result.Result[iter.Iterator[any]]
	// Iter over encoded map element (key/value)
	IterMap(data any) result.Result[iter.Iterator[Element]]
}

func searchElement(decoder Decoder, node any, key string) option.Option[Element] {
	return iter.Find(
		decoder.IterMap(node),
		func(el Element) bool { return el.Key() == key },
	)
}

func decodeSlice(decoder Decoder, encoded any, typ reflect.Type) result.Result[reflect.Value] {
	valTyp := typ.Elem()

	iterRes := decoder.IterSlice(encoded)

	if iterRes.HasFailed() {
		return result.Result[reflect.Value]{}.Failed(iterRes.UnwrapError())
	}

	res := iter.Result_FromIter[[]reflect.Value](
		iter.Map(
			iterRes.Expect(),
			func(encoded any) result.Result[reflect.Value] {
				return decode(decoder, encoded, valTyp)
			},
		),
	)

	arr := reflect.New(typ)
	for _, el := range res.Expect() {
		arr.Elem().Set(reflect.Append(arr.Elem(), el))
	}
	return result.Success(arr.Elem())
}

type reflectElement struct {
	Key   string
	Value reflect.Value
}

func decodeMapElements(decoder Decoder, encoded any, typ reflect.Type) result.Result[[]reflectElement] {
	iterRes := decoder.IterMap(encoded)

	if iterRes.HasFailed() {
		return result.Result[[]reflectElement]{}.Failed(iterRes.UnwrapError())
	}

	return iter.Result_FromIter[[]reflectElement](
		iter.Map(
			iterRes.Expect(),
			func(element Element) result.Result[reflectElement] {
				return result.Map[reflect.Value, reflectElement](
					decode(decoder, element.Value(), typ.Elem()),
					func(decoded reflect.Value) reflectElement {
						return reflectElement{
							Key:   element.Key(),
							Value: decoded,
						}
					},
				)
			},
		),
	)
}

func decodeMap(decoder Decoder, encoded any, typ reflect.Type) result.Result[reflect.Value] {
	val := reflect.New(typ)

	res := decodeMapElements(decoder, encoded, typ)

	if res.HasFailed() {
		return result.Result[reflect.Value]{}.Failed(res.UnwrapError())
	}

	for _, el := range res.Expect() {
		val.SetMapIndex(reflect.ValueOf(el.Key), el.Value)
	}

	return result.Success(val.Elem())
}

func decodeStruct(decoder Decoder, encoded any, typ reflect.Type) result.Result[reflect.Value] {
	val := reflect.New(typ)
	for i := 0; i < typ.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Name
		marshalName, ok := typ.Field(i).Tag.Lookup("serde")
		if ok {
			fieldName = marshalName
		}

		cOpt := searchElement(node, fieldName)

		if cOpt.IsNone() || !field.IsValid() {
			continue
		}

		// Decode option
		if field.CanAddr() && option.IsMutableOption(field.Addr().Interface()) {
			innerType := field.Interface().(option.IOption).TypeOf()
			res := decode(cOpt.Expect(), innerType)
			if res.HasFailed() {
				return result.Failed[reflect.Value](res.UnwrapError())
			}
			resSet := field.Addr().Interface().(option.IMutableOption).TrySet(res.Expect())
			if resSet.HasFailed() {
				return result.Failed[reflect.Value](res.UnwrapError())
			}
		} else { // Decode normally
			res := decode(cOpt.Expect(), field.Type())
			if res.HasFailed() {
				return result.Failed[reflect.Value](res.UnwrapError())
			}
			field.Set(res.Expect())
		}
	}

	return result.Success(val.Elem())
}

func decode(decoder Decoder, encoded any, typ reflect.Type) result.Result[reflect.Value] {
	switch typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Bool, reflect.Float32, reflect.Float64, reflect.String:
		return decoder.DecodePrimaryType(encoded, typ)
	case reflect.Array, reflect.Slice:
		return decodeSlice(encoded, typ)
	case reflect.Map:
		return decodeMap(encoded, typ)
	case reflect.Struct:
		return decodeStruct(encoded, typ)
	default:
		return result.Result[reflect.Value]{}.Failed(errors.New(fmt.Sprintf("type %v cannot be denormalised", typ.Kind())))
	}
}
