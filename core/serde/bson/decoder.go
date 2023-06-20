package bson

import (
	"errors"
	"io"
	"reflect"

	"github.com/gpabois/cougnat/core/result"
)

type Decoder struct {
	parser *Parser
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{parser: NewParser(r)}
}

func (d *Decoder) Init() result.Result[any] {
	return d.parser.Parse().ToAny()
}

func (d *Decoder) DecodePrimaryType(ast any, typ reflect.Type) result.Result[reflect.Value] {
	return decodePrimaryTypes(ast, typ)
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
