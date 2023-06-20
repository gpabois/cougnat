package json

import (
	"errors"
	"fmt"
	"io"
	"reflect"

	"github.com/gpabois/cougnat/core/iter"
	"github.com/gpabois/cougnat/core/option"
	"github.com/gpabois/cougnat/core/result"
)

func Decode[T any](r io.Reader) result.Result[T] {
	var val T
	parser := NewParser(r)
	jsonRes := parser.Parse()

	if jsonRes.HasFailed() {
		return result.Result[T]{}.Failed(jsonRes.UnwrapError())
	}

	decRes := decode(jsonRes.Expect(), reflect.ValueOf(val).Type())

	if decRes.HasFailed() {
		return result.Result[T]{}.Failed(decRes.UnwrapError())
	}

	return result.Success(decRes.Expect().Interface().(T))
}

func searchElement(node Document, key string) option.Option[Element] {
	return iter.Find(
		iter.IterSlice(&node.Pairs),
		func(el Element) bool { return el.Key == key },
	)
}

func decodeStruct(ast any, typ reflect.Type) result.Result[reflect.Value] {
	switch node := ast.(type) {
	case Json:
		if !node.IsDocument() {
			return result.Result[reflect.Value]{}.Failed(errors.New("not a document"))
		}
		return decodeMap(node.ExpectDocument(), typ)
	case Value:
		if !node.IsDocument() {
			return result.Result[reflect.Value]{}.Failed(errors.New("not a document"))
		}
		return decodeMap(node.ExpectDocument(), typ)
	case Document:
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
	default:
		return result.Result[reflect.Value]{}.Failed(errors.New("not a document"))
	}
}

func decodeMap(ast any, typ reflect.Type) result.Result[reflect.Value] {
	switch node := ast.(type) {
	case Json:
		if !node.IsDocument() {
			return result.Result[reflect.Value]{}.Failed(errors.New("not a document"))
		}
		return decodeMap(node.ExpectDocument(), typ)
	case Value:
		if !node.IsDocument() {
			return result.Result[reflect.Value]{}.Failed(errors.New("not a document"))
		}
		return decodeMap(node.ExpectDocument(), typ)
	case Document:
		val := reflect.New(reflect.MapOf(typ.Key(), typ.Elem()))
		for _, el := range node.Pairs {
			valRes := decode(el.Value, typ.Elem())
			if valRes.HasFailed() {
				return result.Result[reflect.Value]{}.Failed(valRes.UnwrapError())
			}

			val.SetMapIndex(reflect.ValueOf(el.Key), valRes.Expect())
		}

		return result.Success(val.Elem())
	default:
		return result.Result[reflect.Value]{}.Failed(errors.New("not a document"))
	}
}

func decodeSlice(ast any, typ reflect.Type) result.Result[reflect.Value] {
	elTyp := typ.Elem()

	switch node := ast.(type) {
	case Json:
		if !node.IsArray() {
			return result.Result[reflect.Value]{}.Failed(errors.New("expecting an array"))
		}
		return decodeSlice(node.ExpectArray(), typ)
	case Value:
		if !node.IsArray() {
			return result.Result[reflect.Value]{}.Failed(errors.New("expecting an array"))
		}
		return decodeSlice(node.ExpectArray(), typ)

	case Array:
		arr := reflect.New(typ)
		for _, el := range node.Elements {
			res := decode(el, elTyp)
			if res.HasFailed() {
				return result.Result[reflect.Value]{}.Failed(res.UnwrapError())
			}

			arr.Elem().Set(reflect.Append(arr.Elem(), res.Expect()))
		}

		return result.Success(arr.Elem())
	default:
		return result.Result[reflect.Value]{}.Failed(errors.New("expecting an array"))
	}
}

func decodePrimaryTypes(ast any, typ reflect.Type) result.Result[reflect.Value] {
	val, ok := ast.(Value)

	if !ok {
		return result.Result[reflect.Value]{}.Failed(errors.New("not a value"))
	}

	switch typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if !val.IsInteger() {
			return result.Result[reflect.Value]{}.Failed(errors.New("not an integer"))
		}

		return result.Success(reflect.ValueOf(val.ExpectInteger()))
	default:
		return result.Result[reflect.Value]{}.Failed(errors.New("not a primary type"))
	}
}

func decode(ast any, typ reflect.Type) result.Result[reflect.Value] {
	switch typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Bool, reflect.Float32, reflect.Float64, reflect.String:
		return decodePrimaryTypes(ast, typ)
	case reflect.Array, reflect.Slice:
		return decodeSlice(ast, typ)
	case reflect.Map:
		return decodeMap(ast, typ)
	case reflect.Struct:
		return decodeStruct(ast, typ)
	default:
		return result.Result[reflect.Value]{}.Failed(errors.New(fmt.Sprintf("type %v cannot be denormalised", typ.Kind())))
	}
}
