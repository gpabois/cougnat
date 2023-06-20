package json

type Json struct {
	array    Array
	document Document
	set      int
}

func (json Json) IsArray() bool {
	return json.set == 1
}

func (json Json) ExpectArray() Array {
	if !json.IsArray() {
		panic("not an array")
	}

	return json.array
}

func (json Json) Array(array Array) Json {
	json.array = array
	json.set = 1
	return json
}

func (json Json) Document(document Document) Json {
	json.document = document
	json.set = 2
	return json
}

func (json Json) IsDocument() bool {
	return json.set == 1
}

func (json Json) ExpectDocument() Document {
	if !json.IsDocument() {
		panic("not a document")
	}

	return json.document
}

type Document struct {
	Pairs []Element
}

type Element struct {
	Key   string
	Value Value
}

type Array struct {
	Elements []Value
}

type Value struct {
	documentValue Document
	arrayValue    Array
	boolValue     bool
	stringValue   string
	integerValue  int
	floatValue    float64

	// Define the value set (similarly to union)
	set int
}

func (val Value) IsDocument() bool {
	return val.set == 1
}

func (val Value) ExpectDocument() Document {
	if !val.IsDocument() {
		panic("not a document")
	}

	return val.documentValue
}

func (val Value) Document(document Document) Value {
	return Value{
		documentValue: document,
		set:           1,
	}
}

func (val Value) IsArray() bool {
	return val.set == 2
}

func (val Value) ExpectArray() Array {
	if !val.IsArray() {
		panic("not an array")
	}

	return val.arrayValue
}

func (val Value) Array(array Array) Value {
	return Value{
		arrayValue: array,
		set:        2,
	}
}

func (val Value) Bool(bval bool) Value {
	return Value{
		boolValue: bval,
		set:       3,
	}
}

func (val Value) String(sval string) Value {
	return Value{
		stringValue: sval,
		set:         3,
	}
}

func (val Value) IsInteger() bool {
	return val.set == 2
}

func (val Value) ExpectInteger() int {
	if !val.IsArray() {
		panic("not an integer")
	}

	return val.integerValue
}

func (val Value) Integer(ival int) Value {
	return Value{
		integerValue: ival,
		set:          4,
	}
}

func (val Value) Float(fval float64) Value {
	return Value{
		floatValue: fval,
		set:        5,
	}
}

func (val Value) Null() Value {
	return Value{
		set: 6,
	}
}
