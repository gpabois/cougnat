package json

import (
	"errors"
	"io"
	"reflect"

	"github.com/gpabois/cougnat/core/iter"
	"github.com/gpabois/cougnat/core/result"
	"github.com/gpabois/cougnat/core/serde"
)

type Decoder struct {
	parser Parser
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		parser: *NewParser(r),
	}
}

func (decoder *Decoder) Init() result.Result[any] {
	return decoder.parser.Parse().ToAny()
}

func (decoder *Decoder) DecodePrimaryType(data any, typ reflect.Type) result.Result[reflect.Value] {
	return decodePrimaryTypes(data, typ)
}

func (decoder *Decoder) IterMap(ast any) result.Result[iter.Iterator[serde.Element]] {
	switch node := ast.(type) {
	case Json:
		if !node.IsDocument() {
			return result.Result[iter.Iterator[serde.Element]]{}.Failed(errors.New("expecting a map"))
		}
		return decoder.IterMap(node.ExpectDocument())
	case Value:
		if !node.IsDocument() {
			return result.Result[iter.Iterator[serde.Element]]{}.Failed(errors.New("expecting an array"))
		}
		return decoder.IterMap(node.ExpectArray())
	case Document:
		return result.Success(iter.Map(
			iter.IterSlice(&node.Pairs),
			func(pair Element) serde.Element {
				return pair
			},
		))
	default:
		return result.Result[iter.Iterator[serde.Element]]{}.Failed(errors.New("not a map"))
	}
}

func (decoder *Decoder) IterSlice(ast any) result.Result[iter.Iterator[any]] {
	switch node := ast.(type) {
	case Json:
		if !node.IsArray() {
			return result.Result[iter.Iterator[any]]{}.Failed(errors.New("expecting an array"))
		}
		return decoder.IterSlice(node.ExpectArray())
	case Value:
		if !node.IsArray() {
			return result.Result[iter.Iterator[any]]{}.Failed(errors.New("expecting an array"))
		}
		return decoder.IterSlice(node.ExpectArray())
	case Array:
		return result.Success(iter.Map(iter.IterSlice(&node.Elements), func(el Value) any { return any(el) }))
	default:
		return result.Result[iter.Iterator[any]]{}.Failed(errors.New("not an array"))
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
	case reflect.Float32, reflect.Float64:
		if !val.IsFloat() {
			return result.Result[reflect.Value]{}.Failed(errors.New("not a float"))
		}
		return result.Success(reflect.ValueOf(val.ExpectFloat()))
	case reflect.Bool:
		if !val.IsBoolean() {
			return result.Result[reflect.Value]{}.Failed(errors.New("not a boolean"))
		}
		return result.Success(reflect.ValueOf(val.ExpectBoolean()))
	case reflect.String:
		if !val.IsString() {
			return result.Result[reflect.Value]{}.Failed(errors.New("not a string"))
		}
		return result.Success(reflect.ValueOf(val.ExpectString()))

	default:
		return result.Result[reflect.Value]{}.Failed(errors.New("not a primary type"))
	}
}
