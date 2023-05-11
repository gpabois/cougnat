package option

import (
	"reflect"
)

type ReflectOption interface {
	ValueOf() reflect.Value
	TypeOf() reflect.Type
	// Return false is the expected value type is not right
	Set(value any) bool
}

func IsOption(value any) bool {
	_, ok := value.(ReflectOption)
	return ok
}

type Option[T any] struct {
	value *T
}

// Returns nil, or the value if any
func (opt Option[T]) ValueOf() reflect.Value {
	if opt.value == nil {
		return reflect.ValueOf(nil)
	}

	return reflect.ValueOf(*opt.value)
}

func (opt Option[T]) TypeOf() reflect.Type {
	return reflect.TypeOf((*T)(nil))
}

func (opt Option[T]) Set(val any) bool {
	refVal := reflect.ValueOf(val)

	if refVal.Kind() == reflect.Ptr && refVal.IsNil() {
		opt.value = nil
		return true
	} else {
		inner, ok := val.(T)
		if !ok {
			return false
		}
		opt.value = &inner
		return true
	}
}

func (opt Option[T]) IsSome() bool {
	return opt.value != nil
}

func (opt Option[T]) IsNone() bool {
	return opt.value == nil
}

func (opt Option[T]) Expect() T {
	if !opt.IsSome() {
		panic("empty value")
	}

	return *opt.value
}

func Some[T any](value T) Option[T] {
	return Option[T]{value: &value}
}

func None[T any]() Option[T] {
	return Option[T]{value: nil}
}
